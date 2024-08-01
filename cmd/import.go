package cmd

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"

	"github.com/joergmis/pod/internal/domain"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	opmlImport string

	importCmd = &cobra.Command{
		Use:   "import",
		Short: "Import feeds from an opml file",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			if opmlImport == "" {
				fmt.Println("no file to import")
				return
			}

			raw, err := os.ReadFile(opmlImport)
			if err != nil {
				log.Fatal(err)
			}

			data := &domain.Opml{}
			if err := xml.Unmarshal(raw, data); err != nil {
				log.Fatal(err)
			}

			feeds := viper.GetStringSlice("feeds")

			for _, item := range data.Body.Outline.Items {
				alreadyTracked := false
				for _, feed := range feeds {
					if feed == item.XmlUrl {
						alreadyTracked = true
					}
				}

				if alreadyTracked {
					fmt.Printf("feed '%s' already tracked - skipping it\n", item.Text)
					continue
				}

				fmt.Printf("start tracking '%s'\n", item.Text)
				feeds = append(feeds, item.XmlUrl)
			}

			viper.Set("feeds", feeds)
			if err := viper.WriteConfig(); err != nil {
				log.Fatal(err)
			}
		},
	}
)

func init() {
	rootCmd.AddCommand(importCmd)
	importCmd.PersistentFlags().StringVar(&opmlImport, "file", "", "Path to an opml podcast file")
}
