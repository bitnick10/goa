package vc

import (
	"crypto/md5"
	"encoding/hex"
	"image"
	"image/draw"
	"io"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/Bitnick2002/freetype-go/freetype"
)

type VerifyCode map[string]*Code

type Code struct {
	Code      string
	BirthDate time.Time
}

func New(sweepInterval time.Duration) VerifyCode {
	verifyCode := make(VerifyCode)
	ticker := time.NewTicker(sweepInterval)
	go func() {
		for {
			<-ticker.C
			verifyCode.Sweep(sweepInterval)
		}
	}()
	return verifyCode
}
func (vc VerifyCode) Sweep(duration time.Duration) {
	for key, value := range vc {
		if value.BirthDate.Add(duration).Before(time.Now()) {
			delete(vc, key)
		}
	}
}

func (vc VerifyCode) Add() (key string) {
	h := md5.New()
	io.WriteString(h, randString())
	key = hex.EncodeToString(h.Sum(nil))

	for vc[key] != nil {
		io.WriteString(h, strconv.Itoa(rd()))
		key = hex.EncodeToString(h.Sum(nil))
	}
	vc[key] = &Code{randCode(), time.Now()}
	return key
}

func (vc VerifyCode) Delete(key string) {
	delete(vc, key)
}

func (code Code) ToImage(ttfFileName string) (*image.RGBA, error) {
	fontStyle, err := ioutil.ReadFile(ttfFileName)
	if err != nil {
		return nil, err
	}
	font, err := freetype.ParseFont(fontStyle)
	if err != nil {
		return nil, err
	}
	fg, bg := image.Black, image.White
	img := image.NewRGBA(image.Rect(0, 0, 80, 30))
	draw.Draw(img, img.Bounds(), bg, image.ZP, draw.Src)

	context := freetype.NewContext()
	context.SetDPI(72)
	context.SetFont(font)
	context.SetFontSize(24)
	context.SetClip(img.Bounds())
	context.SetDst(img)
	context.SetSrc(fg)

	pt := freetype.Pt(10, 25)
	context.DrawString(code.Code, pt)

	return img, nil
}
