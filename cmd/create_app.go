package cmd

import (
	"github.com/moisapps/moisapps/internal/pkg/entity"
	"github.com/moisapps/moisapps/internal/pkg/infrastructure"
	"github.com/spf13/cobra"
)

func AppCmd() *cobra.Command {
	var technology, version, path string
	createAppCmd := &cobra.Command{
		Use:   "app",
		Short: "desenha uma app",
		RunE: func(cmd *cobra.Command, args []string) error {
			newApplication, err := entity.NewApplication(Name, technology, version, path)
			if err != nil {
				return err
			}
			err = newApplication.Create(infrastructure.DB)
			return err
		},
	}
	createAppCmd.Flags().StringVarP(&technology, "technology", "t", "", "Tipo da tecnologia da aplicacao a ser criada")
	createAppCmd.Flags().StringVarP(&version, "version", "v", "", "Versão da tecnologia da aplicacao a ser criada")
	createAppCmd.Flags().StringVarP(&path, "path", "p", "", "Caminho para criar a aplicacao. Caso não seja informado, será utilizado o caminho: /tmp/moisapps/<Nome informado>")
	return createAppCmd
}
