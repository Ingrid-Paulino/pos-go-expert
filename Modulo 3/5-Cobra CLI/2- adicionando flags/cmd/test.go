/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// para rodar : go run main.go test
// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		comando, _ := cmd.Flags().GetString("comando")
		if comando == "ping" {
			cmd.Println("ping")
		} else {
			cmd.Println("pong")
		}
	},
}

// go run main.go
// go run main.go test
// go run main.go test --comando=ping
// go run main.go test -c=ping

func init() {
	rootCmd.AddCommand(testCmd)
	//Adicionando flags
	testCmd.Flags().StringP("comando", "c", "", "Escolha ping ou pong") //go run main.go test --comando=ping //--comando ou -c -> São a mesma coisa, treceiro parametro é um valor default e o ultio é a descrição
	testCmd.MarkFlagRequired("comando")                                 //Obriga o usuario a passar a flag comando
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
