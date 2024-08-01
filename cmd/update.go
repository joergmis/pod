package cmd

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/joergmis/pod/internal/domain"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	nrOfPodcastsToDownload int

	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "check all podcast feeds and download new episodes",
		Run: func(cmd *cobra.Command, args []string) {
			feeds := viper.GetStringSlice("feeds")
			if len(feeds) == 0 {
				fmt.Println("no feeds configured")
				return
			}

			downloadFolder := viper.GetString("downloads")
			if downloadFolder == "" {
				fmt.Println("no download folder configured")
				return
			}

			for _, url := range feeds {
				feed := domain.Feed{}

				resp, err := http.Get(url)
				if err != nil {
					log.Fatal(err)
				}

				data, err := io.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}

				if err := xml.Unmarshal(data, &feed); err != nil {
					log.Fatal(err)
				}

				feed.Update(downloadFolder, nrOfPodcastsToDownload)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.PersistentFlags().IntVar(&nrOfPodcastsToDownload, "count", 5, "Number of podcasts to download per feed")
}
