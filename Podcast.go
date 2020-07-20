package podcast

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

// Podcast ..
type Podcast struct {
	configFilepath string `yaml:"-"`
	contentPath    string `yaml:"-"`
	contentURL     string `yaml:"-"`

	Title string `yaml:"Title"`

	Feed *XMLRoot `yaml:"-"`
}

// New ..
func New(configPath, contentPath, contentURL string) (*Podcast, error) {
	log.Println("[podcast] New ")

	podcast := &Podcast{
		configFilepath: configPath,
		contentPath:    contentPath,
		contentURL:     contentURL,

		Feed: &XMLRoot{
			Itunes:  "http://www.itunes.com/dtds/podcast-1.0.dtd",
			Content: "http://purl.org/rss/1.0/modules/content/",
			Atom:    "http://www.w3.org/2005/Atom",
			Version: "2.0",
			Channel: &Channel{},
		},
	}

	// load content from given directory
	if err := podcast.Load(); err != nil {
		return podcast, err
	}

	return podcast, nil
}

// LoadFromPath ..
func (podcast *Podcast) Load() error {

	// check if podcast YAML file exists
	if _, err := os.Stat(podcast.configFilepath); os.IsNotExist(err) {
		return err
	}

	// Read YAML contents
	buf, err := ioutil.ReadFile(podcast.configFilepath)
	if err != nil {
		return err
	}

	// Parse YAML into struct
	if err := yaml.Unmarshal(buf, &podcast.Feed.Channel); err != nil {
		return err
	}

	// podcast.Feed.Channel.Title = podcast.Title

	color.Magenta("%+v", podcast.Feed.Channel)
	return nil
}

// Fix misconfigs and populate empty values with defaults  before saving ..
func (podcast *Podcast) Fix() {
	log.Println("[podcast] Fix ")
	podcast.Feed.Channel.Fix()
}

// Validate before saving ..
func (podcast *Podcast) Validate() error {
	log.Println("[podcast] Validate ")

	if err := podcast.Feed.Channel.Validate(); err != nil {
		return err
	}

	return nil
}

// SaveToFile ..
func (podcast *Podcast) SaveToFile() error {
	// fix some values
	podcast.Fix()

	// validate feed before saving to file
	if err := podcast.Validate(); err != nil {
		return err
	}

	// generate XML and save to file
	log.Println("[podcast] SaveToFile ")

	buf, err := podcast.Feed.ToXML("")
	if err != nil {
		return err
	}

	color.Yellow("%s", buf)
	return nil
}
