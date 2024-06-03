package main

import (
	"fmt"
	"log"
	"os"
	"time"

	newsapi "daysling.com/news/internal/news-api"
	"github.com/urfave/cli/v2"
)

func main() {
	// This is the main function for the news command
	app := &cli.App{
		Name:    "news-scrap",
		Usage:   "Scrap news from various sources of NewsAPI.org",
		Version: "1.0.1",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "Path to the configuration file",
				Value:   "./config.yaml",
			},
		},
		Action: func(c *cli.Context) error {
			cfg, err := newsapi.NewConfig(c.Path("config"))
			if err != nil {
				log.Fatal(err)
				return err
			}
			db, err := newsapi.NewConnection(cfg)
			if err != nil {
				log.Fatal(err)
				return err
			}
			// loop through 1-5 pages
			fmt.Printf("[INFO] Ran time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
			for i := 1; i <= 5; i++ {
				news, err := newsapi.FetchNews(cfg.ApiKey, cfg.Sources, i)
				if err != nil {
					log.Fatal(err)
					return err
				}
				// save the news to the database
				rows, err := newsapi.SaveNews(db, news, cfg.Database.Table)
				if err != nil {
					log.Fatal(err)
					return err
				}
				fmt.Printf("[INFO] Saved %d rows from page %d.\n", rows, i)
			}
			fmt.Printf("[INFO] Done!\n\n") // another new line since i use stdout to redirect to a file.
			defer db.Close()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
