package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var feedsCmd = &cobra.Command{
	Use:   "feeds",
	Short: "List the feeds that are tracked",
	Run: func(cmd *cobra.Command, args []string) {
		feeds := viper.GetStringSlice("feeds")

		for _, feed := range feeds {
			fmt.Println(feed)
		}
	},
}

func init() {
	rootCmd.AddCommand(feedsCmd)
}
