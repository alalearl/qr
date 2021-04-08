package main

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
	"log"

	qrcode "github.com/skip2/go-qrcode"
)

func showImage(m image.Image) string {
	var buf bytes.Buffer
	err := png.Encode(&buf, m)
	if err != nil {
		panic(err)
	}
	enc := base64.StdEncoding.EncodeToString(buf.Bytes())
	return enc;
}

func generate(input string) string {
	code, err := qrcode.New(input, qrcode.Low)
	if err != nil {
		log.Fatal(err)
	}
	return showImage(code.Image(256))
}
