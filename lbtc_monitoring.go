package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
	"github.com/engenegr/localbitcoins_monitoring/api"
)

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "currency",
			Usage: "The currency you want to sell your BTCs for.",
		},
	}
	app.Action = func(c *cli.Context) error {
		currency := c.String("currency")
		if currency == "" {
			panic("--currency is a required argument")
		}
		var keywords []string
		for _, kw := range c.Args() {
			keywords = append(keywords, strings.ToLower(kw))
		}
		fmt.Println("Currency: " + currency)
		fmt.Println(fmt.Sprintf("Keywords: %v", keywords))
		m := api.NewMonitor()
		_, err := m.GatherBuyers(currency, keywords)
		m.ListCurrencies()
		return err
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

