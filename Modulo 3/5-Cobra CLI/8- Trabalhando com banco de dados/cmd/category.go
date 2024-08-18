/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// Hooks são funções que são executadas automaticamente em determinados eventos do ciclo de vida do comando.
// Hooks são úteis para executar tarefas comuns, como inicializar flags, configurar loggers, etc.
// Hooks permite fazer execucoes antes e depois do comando ser executado

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
		cmd.Help()
	},
}

var category string

func init() {
	rootCmd.AddCommand(categoryCmd)
}
