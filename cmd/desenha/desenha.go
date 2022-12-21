package desenha

import (
	"github.com/moisapps/moisapps/cmd"
	"github.com/moisapps/moisapps/cmd/application"
	"github.com/moisapps/moisapps/cmd/technologies"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(desenhaCmd())
}

func desenhaCmd() *cobra.Command {
	listTechsCmd := technologies.ListTechnologiesCmd()
	createAppCmd := application.AppCmd()
	desenha := cobra.Command{
		Use:   "desenha",
		Short: "desenha um(a) app, namespace, etc",
	}
	desenha.AddCommand(listTechsCmd, createAppCmd)

	return &desenha
}
