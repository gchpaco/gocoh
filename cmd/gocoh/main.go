package main

import (
	"flag"
	"github.com/gchpaco/gocoh/pkg/manifest"
	"log"
	"net/http"
)

var url = flag.String("manifest", "http://patch.savecoh.com/manifest.xml", "Manifest location to use")
var dir = flag.String("base", ".", "Base directory to check for updates against")
var verbose = flag.Bool("v", false, "Be more verbose")

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
	log.Printf("Label: %s\n", manifest.Label)
	//for _, profile := range manifest.Profiles {
	//	fmt.Printf("Profile %s:\n\tExec: %s\n\tOrder: %d\n\tParams: %s\n",
	//		profile.Name, profile.Exec, profile.Order, profile.Params)
	//}
	for _, file := range manifest.Files {
		if *verbose {
			log.Printf("Operating on file %s", file.Name)
		}
		fileOK, err := file.Verify(*dir)
		if err != nil {
			log.Fatal(err)
		}
		if !fileOK {
			log.Printf("Need to redownload %s\n", file.Name)
			if err := file.Download(*dir, *verbose); err != nil {
				log.Fatal(err)
			}
		}
	}
}
