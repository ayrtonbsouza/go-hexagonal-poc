/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/ayrtonbsouza/hexagonal-architecture-poc/adapters/cli"
	"github.com/spf13/cobra"
)

var (
	action       string
	productId    string
	productName  string
	productPrice float64
)

// cliCmd represents the cli command
var cliCmd = &cobra.Command{
	Use:   "cli",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		result, err := cli.Run(&productService, action, productId, productName, productPrice)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(cliCmd)
	cliCmd.Flags().StringVarP(&action, "action", "a", "enable", "enable or disable a product")
	cliCmd.Flags().StringVarP(&productId, "id", "i", "", "product id")
	cliCmd.Flags().StringVarP(&productName, "name", "n", "", "product name")
	cliCmd.Flags().Float64VarP(&productPrice, "price", "p", 0, "product price")
}
