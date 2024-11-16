package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(db)
}

var db = &cobra.Command{
	Use:   "db",
	Short: "the db you want to use in the project",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		comand := exec.Command("echo", "hello world")
		output, err := comand.Output()
		if err != nil {
			fmt.Print(string(output))
		}
		fmt.Println("Next-Gen v0.0.1")
	},
}
