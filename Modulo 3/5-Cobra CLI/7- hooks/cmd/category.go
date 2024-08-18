/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
		fmt.Println("Category called with name: ", category) //go run main.go category -n=Ingrid
	},
	PreRun: func(cmd *cobra.Command, args []string) { //executa antes do comando ser executado
		fmt.Println("Chamado antes do run")
	},
	PostRun: func(cmd *cobra.Command, args []string) { //executa depois do comando ser executado
		fmt.Println("Chamado depois do run")
	},
	//RunE: func(cmd *cobra.Command, args []string) error { //executa o comando e retorna um erro
	//	fmt.Println("Chamado runE")
	//	//return nil
	//	return fmt.Errorf("Ocorreu um erro")
	//},
	//PreRunE: func(cmd *cobra.Command, args []string) error {
	//	return fmt.Errorf("Ocorreu um erro antes do RunE")
	//},
	//PostRunE: func(cmd *cobra.Command, args []string) error {
	//	return fmt.Errorf("Ocorreu um erro pos do RunE")
	//},
}

var category string //variavel vai receber o valor passado na flag -n

func init() {
	rootCmd.AddCommand(categoryCmd)
	categoryCmd.PersistentFlags().StringVarP(&category, "name", "n", "X", "Name of the category") //salva o valor passado na variavel category
}
