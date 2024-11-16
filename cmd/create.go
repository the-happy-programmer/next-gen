package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var qs = []*survey.Question{
	{
		Name: "Auth",
		Prompt: &survey.Select{
			Message: "What Auth library you want to use",
			Options: []string{"next-auth", "better-auth", "none"},
		},
	},
	{
		Name: "Database",
		Prompt: &survey.Select{
			Message: "which database you are using",
			Options: []string{"supabase", "neon", "PostgreSQL", "none"},
		},
	},
	{
		Name: "UI",
		Prompt: &survey.Select{
			Message: "Which UI library you want to use",
			Options: []string{"next-ui", "shadcn/ui"},
		},
	},
}

func init() {
	rootCmd.AddCommand(create)
}

var create = &cobra.Command{
	Use:   "create",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		answers := struct {
			Auth     string `survey:"auth"`
			Database string `survey:"db"`
			Ui       string `survey:"ui"`
		}{}

		err := survey.Ask(qs, &answers)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	},
}
