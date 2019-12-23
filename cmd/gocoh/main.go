package main

import (
	"bufio"
	"fmt"
	"github.com/gchpaco/gocoh/pkg/manifest"
	"log"
)

func main() {
	manifest, err := manifest.Read(bufio.NewReader(os.Stdin))
	if err != nil {
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
