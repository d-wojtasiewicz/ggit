package repoInit

import (
	"ggit/internal/repository"

	"github.com/spf13/cobra"
)

func NewCommandInit(r *repository.Repository) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "init",
		Short: "Create an empty Git repository",
		Long: `This command creates an empty Git repository - basically a .git directory with subdirectories for objects, refs/heads, refs/tags, and template files. 
	An initial branch without any commits will be created (see the --initial-branch option below for its name).`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return r.Create(true)
		},
	}
	return cmd
}
