package manifest

import (
	"encoding/xml"
	"io"
)

// Read loads an XML wire document describing a CoH manifest and
// either returns the manifest or returns an error.
func Read(r io.Reader) (manifest *Manifest, err error) {
	manifest = new(Manifest)
	err = xml.NewDecoder(r).Decode(&manifest)
	return
}
