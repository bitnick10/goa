package captcha

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

type CaptchaCode struct {
	codemap map[string]*Code
	// OnBeforeSweep func()
	OnAfterSweep func()
}

type Code struct {
	Code      string
	BirthDate time.Time
}

func New(sweepInterval time.Duration) *CaptchaCode {
	cc := &CaptchaCode{codemap: make(CaptchaCode)}
	ticker := time.NewTicker(sweepInterval)
	go func() {
		for {
			<-ticker.C
			cc.Sweep(sweepInterval)
		}
	}()
	return cc
}
func (cc *CaptchaCode) Get(key string) (code *Code) {
	return cc.codemap[key]
}
func (cc *CaptchaCode) Delete(key string) {
	delete(cc.codemap, key)
}
func (cc *CaptchaCode) Len(key string) int {
	return len(cc.codemap)
}
func (cc *CaptchaCode) Sweep(duration time.Duration) {
	for key, value := range cc.codemap {
		if value.BirthDate.Add(duration).Before(time.Now()) {
			delete(cc, key)
		}
	}
	OnAfterSweep()
}

func (cc *CaptchaCode) Add() (key string) {
	h := md5.New()
	io.WriteString(h, randString())
	key = hex.EncodeToString(h.Sum(nil))

	for cc.codemap[key] != nil {
		io.WriteString(h, strconv.Itoa(rd()))
		key = hex.EncodeToString(h.Sum(nil))
	}
	cc.codemap[key] = &Code{randCode(), time.Now()}
	return key
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
