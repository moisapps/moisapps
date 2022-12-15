package cmd

import (
	"github.com/moisapps/moisapps/internal"
	"github.com/spf13/cobra"
)

func appCmd() *cobra.Command {
	var technology, version, path string
	createAppCmd := &cobra.Command{
		Use:   "app",
		Short: "desenha uma app",
		RunE: func(cmd *cobra.Command, args []string) error {
			nodeApp := internal.NewNodeApp(Name, version, path, internal.DB)
			err := nodeApp.Create()
			return err
		},
	}
	createAppCmd.Flags().StringVarP(&technology, "technology", "t", "", "Tipo da tecnologia da aplicacao a ser criada")
	createAppCmd.Flags().StringVarP(&version, "version", "v", "", "Versão da tecnologia da aplicacao a ser criada")
	createAppCmd.Flags().StringVarP(&path, "path", "p", "", "Caminho para criar a aplicacao")
	return createAppCmd
}
