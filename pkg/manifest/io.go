package manifest

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Read loads an XML wire document describing a CoH manifest and
// either returns the manifest or returns an error.
func Read(r io.Reader) (manifest *Manifest, err error) {
	manifest = new(Manifest)
	err = xml.NewDecoder(r).Decode(&manifest)
	return
}

func (file *File) localizeTo(base string) (string, error) {
	cleaned := filepath.Clean(file.Name)
	if filepath.Base(cleaned) == ".." ||
		filepath.Base(cleaned) == "." ||
		filepath.Base(cleaned) == "/" {
		return "", fmt.Errorf("Something weird going on with manifest; saw request for %s which is unsafe", cleaned)
	}
	return filepath.Join(base, cleaned), nil
}

// Verify checks to see if a file on disk matches the indicated File
// in the manifest.  If false, the file should be redownloaded.
func (file *File) Verify(base string) (bool, error) {
	localized, err := file.localizeTo(base)
	if err != nil {
		return false, err
	}
	if stat, err := os.Lstat(localized); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	} else if !stat.Mode().IsRegular() {
		return false, fmt.Errorf("Tried to read a file with a weird mode in %s, aborting", localized)
	} else if stat.Size() != int64(file.Size) {
		return false, nil
	} else if hash, err := md5file(localized); err != nil {
		return false, err
	} else if fmt.Sprintf("%x", hash) != file.MD5 {
		return false, nil
	} else {
		return true, nil
	}
}

func md5file(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
