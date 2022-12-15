package cmd

import (
	"fmt"
	"github.com/moisapps/moisapps/internal"
	"github.com/spf13/cobra"
)

func listTechnologiesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "technology",
		Short: "desenha as tecnologias e versoes suportadas",
		RunE: func(cmd *cobra.Command, args []string) error {
			repository := internal.NewRepository(internal.DB)
			technologies, err := repository.FindAll()
			if err != nil {
				return err
			}
			for _, t := range technologies {
				fmt.Printf("%s\t%s\n", t.Name(), t.Version())
			}
			return nil
		},
	}
}
