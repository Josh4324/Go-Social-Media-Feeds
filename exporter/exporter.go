package exporter

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

type Socialmedia interface {
	Feed() []string
	Fame() int
}

func ExportTxt(u Socialmedia, filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return errors.New("an error occurred opening the file: " + err.Error())
	}
	for _, fd := range u.Feed() {
		n, err := f.Write([]byte(fd + "\n"))
		if err != nil {
			return errors.New("an error occurred writing to file: " + err.Error())
		}
		fmt.Printf("wrote %d bytes \n", n)
	}
	return nil

}

func ExportJson(u Socialmedia, filename string) error {

	feed := make(map[int]string)
	for index, fd := range u.Feed() {
		feed[index+1] = fd
	}
	json, err := json.Marshal(feed)
	if err != nil {
		return errors.New("an error occurred opening the file: " + err.Error())
	}
	ioutil.WriteFile(filename, json, os.ModePerm)
	return nil
}

func ExportXml(u Socialmedia, filename string) error {

	out, err := xml.MarshalIndent(u.Feed(), " ", "  ")
	if err != nil {
		return errors.New("an error occurred opening the file: " + err.Error())
	}
	ioutil.WriteFile(filename, out, os.ModePerm)
	return nil
}
