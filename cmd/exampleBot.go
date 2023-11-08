/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	TeleToken = os.Getenv("TELE_TOKEN")
)

// exampleBotCmd represents the exampleBot command
var exampleBotCmd = &cobra.Command{
	Use:     "exampleBot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("exampleBot %s started", appVersion)
		pref := telebot.Settings{
			Token:  os.Getenv("TELE_TOKEN"),
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		}

		exampleBot, err := telebot.NewBot(pref)
		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
			return
		}

		exampleBot.Handle(telebot.OnText, func(c telebot.Context) error {
			log.Print(c.Message().Payload, c.Text())
			payload := c.Message().Payload
			switch payload {
			case "hello":
				err = c.Send(fmt.Sprintf("My name is ExampleBot %s", appVersion))

			}
			return err
		})
		exampleBot.Start()
	},
}

func init() {
	rootCmd.AddCommand(exampleBotCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exampleBotCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exampleBotCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
