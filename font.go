package main

import (
	"log"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func GetFont() font.Face {
	fontBytes, err := os.ReadFile("./assets/Roboto-Regular.ttf")
	if err != nil {
		log.Fatal(err)
	}

	parsedFont, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	face, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return face
}
