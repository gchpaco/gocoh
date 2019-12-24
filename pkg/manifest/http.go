package manifest

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type badStatusCode struct {
	StatusCode int
}

func (e badStatusCode) Error() string {
	return fmt.Sprintf("Bad status code seen: %d", e.StatusCode)
}

// Download retrieves a file from its mirrors, trying each in turn
// until success or ultimate failure.
func (file *File) Download(base string, verbose bool) error {
	localized, err := file.localizeTo(base)
	if err != nil {
		return err
	}
	if verbose {
		log.Printf("Localized file is %s\n", localized)
	}
	dir := filepath.Dir(localized)
	if verbose {
		log.Printf("Localized file directory is %s\n", dir)
	}
	if verbose {
		log.Printf("Making parent directories\n")
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	if verbose {
		log.Printf("Creating file\n")
	}
	f, err := os.Create(localized)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, url := range file.URLs {
		if verbose {
			log.Printf("Trying URL %s\n", url)
		}
		shouldContinue, err := copyTo(url, f, verbose)
		if verbose {
			log.Printf("Got %s, %s\n", shouldContinue, err)
		}
		if err == nil {
			return nil
		} else if shouldContinue {
			log.Printf("Saw non-fatal error %v on URL %s; continuing", err, url)
			continue
		} else {
			return err
		}
	}
	return fmt.Errorf("Unable to download %s; all mirrors errored", localized)
}

func copyTo(url string, writer io.Writer, verbose bool) (bool, error) {
	resp, err := http.Get(url)
	if err != nil {
		return true, err
	}
	defer resp.Body.Close()
	if verbose {
		log.Printf("Successfully got %d status code", resp.StatusCode)
	}
	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		if verbose {
			log.Printf("Didn't like status code, aborting")
		}
		return true, badStatusCode{resp.StatusCode}
	}
	if _, err := io.Copy(writer, resp.Body); err != nil {
		// Don't want to continue to try if this fails; the mirror's
		// fine, we probably filled the disk or something dumb.
		return false, err
	}
	return true, nil
}
