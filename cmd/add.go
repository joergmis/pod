package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	url string

	addCmd = &cobra.Command{
		Use:   "add",
		Short: "Add a podcast feed",
		Run: func(cmd *cobra.Command, args []string) {
			if url == "" {
				fmt.Println("no URL given")
				return
			}

			feeds := viper.GetStringSlice("feeds")

			alreadyTracked := false
			for _, feed := range feeds {
				if feed == url {
					alreadyTracked = true
				}
			}

			if alreadyTracked {
				fmt.Println("feed is already tracked")
				return
			}

			feeds = append(feeds, url)
			viper.Set("feeds", feeds)

			if err := viper.WriteConfig(); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("tracking feed %s\n", url)
		},
	}
)

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.PersistentFlags().StringVar(&url, "url", "", "Add a url to track")
}
