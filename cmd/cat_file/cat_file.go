package cat_file

import (
	"fmt"
	"ggit/internal/repository"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func NewCommandInit(r *repository.Repository) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Provide contents or details of repository objects",
		Long:  "Provide contents or details of repository objects",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCatFile(r, args)
		},
	}
	return cmd
}

func validArgs(args []string) error {
	obj := repository.GitObjects()
	if !lo.Contains(obj, args[0]) {
		return fmt.Errorf("unsuported object of type: %s, supported types %v", args[0], obj)
	}
	return nil
}

func runCatFile(r *repository.Repository, args []string) error {
	if err := validArgs(args); err != nil {
		return err
	}
	data, err := r.CatObject(args[1])
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}
