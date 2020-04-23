package main

import (
	"gosmf/exporter"
	"gosmf/media"
)

func main() {
	fb := new(media.Facebook)
	twtr := new(media.Twitter)
	lnkdin := new(media.Linkedin)
	err := exporter.ExportJson(fb, "fb.json")
	nerr := exporter.ExportTxt(twtr, "twitter.txt")
	ne := exporter.ExportXml(lnkdin, "linkedin.xml")
	if err != nil {
		panic(err)
	}

	if nerr != nil {
		panic(nerr)
	}

	if ne != nil {
		panic(ne)
	}

}
