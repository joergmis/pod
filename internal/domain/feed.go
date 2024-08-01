package domain

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Feed struct {
	Channel struct {
		Title string `xml:"title"`
		Link  string `xml:"link"`
		Items []struct {
			Title    string `xml:"title"`
			Download struct {
				Type   string `xml:"type,attr"`
				Length string `xml:"length,attr"`
				Url    string `xml:"url,attr"`
			} `xml:"enclosure"`
		} `xml:"item"`
	} `xml:"channel"`
}

func (f *Feed) Update(folder string) error {
	items := len(f.Channel.Items)

	fmt.Printf("%s: %d episodes\n", f.Channel.Title, len(f.Channel.Items))
	for idx, item := range f.Channel.Items {
		filename := fmt.Sprintf("%s_ep_%04d_%s.mp3", simplifyName(f.Channel.Title), items-idx, simplifyName(item.Title))
		path := filepath.Join(folder, filename)

		if _, err := os.Stat(path); os.IsNotExist(err) {
			fmt.Printf("downloading '%s'\n", filename)
			data, err := http.Get(item.Download.Url)
			if err != nil {
				return err
			}

			podcast, err := io.ReadAll(data.Body)
			if err != nil {
				return err
			}

			if err := os.WriteFile(path, podcast, 0644); err != nil {
				return err
			}
		} else {
			fmt.Printf("skipping '%s' - already downloaded\n", filename)
		}
	}

	return nil
}

func simplifyName(name string) string {
	r := strings.NewReplacer(":", "_", ",", "_", " ", "_", "!", "", "/", "-", "'", "", "?", "", " ", "_")
	return strings.ToLower(r.Replace(name))
}
