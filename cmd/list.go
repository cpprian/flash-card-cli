/*
Copyright © 2022 Cyprian Szczepański <cpprian456@gmail.com>

*/
package cmd

import (
	"fmt"
	"github.com/cpprian/flash-card-cli/card"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"text/tabwriter"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all flash cards or last score",
	Long: 
	`Display all flashcards, each of them contains
question and answer separated by a comma.
	
Or display your last score using --score flag.
This option includes displaying:
	- questions with your answers
	- if your answer was incorrect, also show the correct answer
	- points for good anwears
	- number of all questions
	- how much time do you need to answer each question`,
	Run: listRun,
}

var score bool

func listRun(cmd *cobra.Command, args []string) {
	w := tabwriter.NewWriter(os.Stdout, 2, 0, 4, ' ', 0)
	// TODO: implementation for --score option
	if score {

		w.Flush()
		return
	}

	c := card.NewCardContainer()
	card.ReadData(viper.GetString("datafile"), *c)
	for key, value := range c.Question {
		fmt.Fprintln(w, key+"\t"+value)
	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&score, "score", false, "Show score from score.json")
}
