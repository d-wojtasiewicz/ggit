/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"ggit/cmd/repoInit"
	"ggit/internal/factory"
	"ggit/internal/filesystem"
	"ggit/internal/repository"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ggit",
	Short: "Go Global information tracker",
	Long: `Go Global information tracker 
GGit is a basic command line version control software. 
It attempt to implement very simplified and basic Git core commands.
	
For example:`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	fs := factory.NewFactory()
	cwd, err := filesystem.GetCWD()
	if err != nil {
		fmt.Printf("%e", err)
		os.Exit(1)
	}
	r, err := repository.NewRepository(fs, cwd)
	if err != nil {
		fmt.Printf("%e", err)
		os.Exit(1)
	}
	rootCmd.AddCommand(repoInit.NewCommandInit(r))
}
