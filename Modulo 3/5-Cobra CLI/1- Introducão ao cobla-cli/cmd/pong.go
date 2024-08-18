/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// pongCmd represents the pong command
var pongCmd = &cobra.Command{
	Use:   "pong",
	Short: "Esse é o comando pong",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) { //go run main.go pong
		fmt.Println("pong called") //retorna oq esta aqui quando o comando é chamado
	},
}

func init() {
	rootCmd.AddCommand(pongCmd) //pong é adcionado no root

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pongCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pongCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
