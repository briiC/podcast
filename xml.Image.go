package podcast

// Image ..
type Image struct {
	// Text  string `xml:",chardata"`
	Href  string `xml:"href,attr,omitempty"`
	URL   string `xml:"url,omitempty"`
	Title string `xml:"title,omitempty"`
	Link  string `xml:"link,omitempty"`
}

// IsEmpty ..
func (image *Image) IsEmpty() bool {
	return image == nil || image.URL == "" || image.Title == "" || image.Link == ""
}

// UnmarshalYAML ..
func (image *Image) UnmarshalYAML(unmarshal func(interface{}) error) error {
	unmarshal(&image.URL)
	return nil
}

// String return URL or Href
func (image *Image) String() string {
	if image.URL != "" {
		return image.URL
	}
	if image.Href != "" {
		return image.Href
	}
	return image.Link
}
