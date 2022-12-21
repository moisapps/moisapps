package application

import (
	cmd2 "github.com/moisapps/moisapps/cmd"
	"github.com/moisapps/moisapps/internal/pkg/entity"
	"github.com/moisapps/moisapps/internal/pkg/infrastructure"
	"github.com/spf13/cobra"
)

func appCmd() *cobra.Command {
	var technology, version, path string
	createAppCmd := &cobra.Command{
		Use:   "app",
		Short: "desenha uma app",
		RunE: func(cmd *cobra.Command, args []string) error {
			nodeApp := entity.NewNodeApp(cmd2.Name, version, path, infrastructure.DB)
			err := nodeApp.Create()
			return err
		},
	}
	createAppCmd.Flags().StringVarP(&technology, "technology", "t", "", "Tipo da tecnologia da aplicacao a ser criada")
	createAppCmd.Flags().StringVarP(&version, "version", "v", "", "Versão da tecnologia da aplicacao a ser criada")
	createAppCmd.Flags().StringVarP(&path, "path", "p", "", "Caminho para criar a aplicacao")
	return createAppCmd
}
