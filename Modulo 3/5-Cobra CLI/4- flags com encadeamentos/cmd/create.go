/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// comando que criou esse arquivo : cobra-cli add create -p 'categoryCmd' -> cria o comando create sendo filho de category

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help() //mostra as flags do comando
	},
}

func init() {
	categoryCmd.AddCommand(createCmd) //create nao é filho de root e sim de category, casso queira mudar o pai, basta mudar o comando aqui ex: rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.
	//Temos dois tipos de flags: persistent flags e local flags
	//Persistent flags são globais e podem ser usadas por todos os comandos(pai) e subcomandos(filhos)
	createCmd.PersistentFlags().String("name1", "", "Name of the category 1") /*
		- se rodar o comando go run main.go category create a flag name aparece
		- se rodar o comando go run main.go category  a flag name não aparece
	*/

	//Local flags são específicas de um comando ou subcomando e só podem ser usadas por ele
	createCmd.Flags().String("name3", "", "Name of the category 3") //vai funcionar so em create
}
