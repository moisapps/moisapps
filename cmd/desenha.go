package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(desenhaCmd())
}

func desenhaCmd() *cobra.Command {
	listTechsCmd := listTechnologiesCmd()
	desenha := cobra.Command{
		Use:   "desenha",
		Short: "desenha um(a) app, namespace, etc",
	}
	desenha.AddCommand(listTechsCmd)

	return &desenha
}
