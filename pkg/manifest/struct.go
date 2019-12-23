package manifest

// Profile collects launch data for City of Heroes; executable
// location, parameters to use, etc.
type Profile struct {
	Name   string `xml:",chardata"`
	Exec   string `xml:"exec,attr"`
	Order  int    `xml:"order,attr"`
	Params string `xml:"params,attr"`
}

// File specifies files that should exist on disk in relative
// pathnames, what metadata they should have, and how to download new
// versions if needed.
type File struct {
	Name string   `xml:"name,attr"`
	Size int      `xml:"size,attr"`
	MD5  string   `xml:"md5,attr"`
	URLs []string `xml:"url"`
}

// Manifest is a collection of Files and Profiles, as respresented in
// the XML manifest.
type Manifest struct {
	Label    string    `xml:"label"`
	Profiles []Profile `xml:"profiles>launch"`
	Files    []File    `xml:"filelist>file"`
}
