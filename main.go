package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/skip2/go-qrcode"
)

func main() {
	angga, err := os.Open("angga.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer angga.Close()

	angga2, err := ioutil.ReadAll(angga)
	if err != nil {
		log.Fatal(err)
	}

	QR := string(angga2)

	newQR := "angga-qr.png"

	err = qrcode.WriteFile(QR, qrcode.Medium, 512, newQR)

	if err != nil {
		log.Fatal(err)
	}
}
