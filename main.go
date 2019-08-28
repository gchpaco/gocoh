package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Profile struct {
	Name   string `xml:",chardata"`
	Exec   string `xml:"exec,attr"`
	Order  int    `xml:"order,attr"`
	Params string `xml:"params,attr"`
}

type File struct {
	Name string   `xml:"name,attr"`
	Size int      `xml:"size,attr"`
	MD5  string   `xml:"md5,attr"`
	URLs []string `xml:"url"`
}

type Manifest struct {
	Label    string    `xml:"label"`
	Profiles []Profile `xml:"profiles>launch"`
	Files    []File    `xml:"filelist>file"`
}

func main() {
	var manifest Manifest
	if err := xml.NewDecoder(bufio.NewReader(os.Stdin)).Decode(&manifest); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Label: %s\n", manifest.Label)
	for _, profile := range manifest.Profiles {
		fmt.Printf("Profile %s:\n\tExec: %s\n\tOrder: %d\n\tParams: %s\n",
			profile.Name, profile.Exec, profile.Order, profile.Params)
	}
	for _, file := range manifest.Files {
		fmt.Printf("File %s (%d sources)\n", file.Name, len(file.URLs))
	}
}
