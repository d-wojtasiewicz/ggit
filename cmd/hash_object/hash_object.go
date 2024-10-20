package hashobject

import (
	"fmt"
	"ggit/internal/filesystem"
	"ggit/internal/repository"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func NewCommandHashObject(r *repository.Repository) *cobra.Command {
	ops := &repository.HashObject{
		Type:  "blob",
		Write: false,
	}
	var cmd = &cobra.Command{
		Use:   "hash-object",
		Short: "Compute object ID and optionally creates a blob from a file",
		Long:  "Compute object ID and optionally creates a blob from a file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ops.File = args[0]
			return runHashObject(r, ops)
		},
	}
	cmd.Flags().StringP("type", "t", ops.Type, "Specify the type")
	cmd.Flags().BoolP("write", "w", ops.Write, "Actually write the object into the database")
	return cmd
}

func validArgs(r *repository.Repository, opts *repository.HashObject) error {
	obj := repository.GitObjects()
	if !lo.Contains(obj, opts.Type) {
		return fmt.Errorf("unsuported object of type: %s, supported types %v", opts.Type, obj)
	}
	if !filesystem.Exists(r.FS, opts.File) {
		return fmt.Errorf("file %s not found", opts.File)
	}
	if r.IsInitiated() && opts.Write {
		return repository.ErrorUninitiate
	}
	return nil
}

func runHashObject(r *repository.Repository, opts *repository.HashObject) error {
	if err := validArgs(r, opts); err != nil {
		return err
	}
	data, err := r.HashObject(opts)
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}
