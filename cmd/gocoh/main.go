package main

import (
	"flag"
	"fmt"
	"github.com/gchpaco/gocoh/pkg/manifest"
	"log"
	"net/http"
)

var url = flag.String("manifest", "http://patch.savecoh.com/manifest.xml", "Manifest location to use")
var dir = flag.String("base", ".", "Base directory to check for updates against")

func main() {
	flag.Parse()
	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	manifest, err := manifest.Read(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Label: %s\n", manifest.Label)
	//for _, profile := range manifest.Profiles {
	//	fmt.Printf("Profile %s:\n\tExec: %s\n\tOrder: %d\n\tParams: %s\n",
	//		profile.Name, profile.Exec, profile.Order, profile.Params)
	//}
	for _, file := range manifest.Files {
		fileOK, err := file.Verify(*dir)
		if err != nil {
			log.Fatal(err)
		}
		if !fileOK {
			fmt.Printf("Need to redownload %s\n", file.Name)
		}
	}
}
