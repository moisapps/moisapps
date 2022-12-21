package cmd

import (
	"github.com/moisapps/moisapps/internal/pkg/infrastructure"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "moisapps",
	Short: "Moisapps is an automation for application creation and release",
	Long:  "Moisapps cli a tool for application creation",
}

var Name string

func Execute() error {
	infrastructure.SetupDatabase()
	rootCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "Nome da aplicacao")

	return rootCmd.Execute()
}
