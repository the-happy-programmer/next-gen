package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ngen",
	Short: "nextjs starter kit generator",
	Long:  `A starter kit to be able to save time on creating a boilerplate website https://github.com/the-happy-programmer/next-gen`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello from there")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
