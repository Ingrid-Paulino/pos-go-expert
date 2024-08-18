/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// categoryCmd represents the category command
var categoryCmd = &cobra.Command{
	Use:   "category",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Category called with name: ", category) //go run main.go category -n=Ingrid
	},
}

var category string //variavel vai receber o valor passado na flag -n

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "X", "Name of the category") //salva o valor passado na variavel category
}
