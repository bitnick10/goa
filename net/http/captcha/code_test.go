package captcha

import (
	"fmt"
	"image/png"
	"net/http"
	"strconv"
	"testing"
	"time"
)

func Test_(t *testing.T) {
	http.HandleFunc("/pic", pic)
	http.HandleFunc("/", index)

	for i := 9000; i < 65535; i++ {
		s := &http.Server{
			Addr:           ":" + strconv.Itoa(i),
			ReadTimeout:    30 * time.Second,
			WriteTimeout:   30 * time.Second,
			MaxHeaderBytes: 1 << 20,
		}
		fmt.Println("opened at port ", i)
		err := s.ListenAndServe()
		if err != nil {
			continue
		} else {
			fmt.Println("else")
			break
		}
	}
}

func pic(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "image/png")
	captchaCode := New(time.Second)
	key := captchaCode.Add()
	image, _ := captchaCode.Get(key).ToImage("C:\\luxi-fonts\\luxisr.ttf")
	png.Encode(res, image)
}

func index(w http.ResponseWriter, req *http.Request) {
	str := "<meta charset=\"utf-8\"><h3>golang 图片验证码例子</h3><img border=\"1\" src=\"/pic\" alt=\"图片验证码\" onclick=\"this.src='/pic'\" />"
	for i := 0; i < 100; i++ {
		str += fmt.Sprintf("<img border=\"1\" src=\"/pic?v=%d\" alt=\"图片验证码\" onclick=\"this.src='/pic'\" />", i)
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(str))
}
