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
		name, _ := cmd.Flags().GetString("name") //go run main.go category --name=X
		fmt.Println("Category called with name: ", name)

		name2, _ := cmd.Flags().GetString("name2") //go run main.go category -n=X   oh go run main.go category (tras valor default)
		fmt.Println("Category called with name: ", name2)

		exists, _ := cmd.Flags().GetBool("exists") //go run main.go category --exists=true /go run main.go category -e=true / go run main.go category -e
		fmt.Println("Category called with exists: ", exists)

		id, _ := cmd.Flags().GetInt16("id")          //go run main.go category --age=10 /go run main.go category -a=10 / go run main.go category -a
		fmt.Println("Category called with id: ", id) //go run main.go category -i=10 / go run main.go category i / go run main.go category --id=11
	},
}

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.PersistentFlags().String("name", "", "Name of the category ")
	categoryCmd.PersistentFlags().StringP("name2", "n", "Y", "Name of the category")      //posso ter o comando abreviado
	categoryCmd.PersistentFlags().BoolP("exists", "e", false, "Check if category exists") //go run main.go category -e=true
	categoryCmd.PersistentFlags().Int16P("id", "i", 0, "ID of the category")              //go run main.go category -i=10 /
}
