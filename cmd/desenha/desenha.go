package desenha

import (
	"github.com/moisapps/moisapps/cmd"
	"github.com/moisapps/moisapps/cmd/application"
	"github.com/spf13/cobra"
)

func init() {
	cmd.rootCmd.AddCommand(desenhaCmd())
}

func desenhaCmd() *cobra.Command {
	listTechsCmd := cmd.listTechnologiesCmd()
	createAppCmd := application.appCmd()
	desenha := cobra.Command{
		Use:   "desenha",
		Short: "desenha um(a) app, namespace, etc",
	}
	desenha.AddCommand(listTechsCmd, createAppCmd)

	return &desenha
}
