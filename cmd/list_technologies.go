package cmd

import (
	"fmt"
	"github.com/moisapps/moisapps/internal/pkg/entity"
	"github.com/moisapps/moisapps/internal/pkg/infrastructure"
	"github.com/spf13/cobra"
)

func listTechnologiesCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "technology",
		Short: "desenha as tecnologias e versoes suportadas",
		RunE: func(cmd *cobra.Command, args []string) error {
			repository := entity.NewRepository(infrastructure.DB)
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
