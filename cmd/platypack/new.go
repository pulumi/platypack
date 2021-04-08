package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pulumi/platypack/cmd/generator"
	"github.com/spf13/cobra"
)

func newNewCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:        "new <language> <name>",
		Aliases:    []string{},
		SuggestFor: []string{},
		Args:       cobra.ExactArgs(2),
		Short:      "create a new pulumi package",
		Long: `create a new pulumi package in the current directory
		- supported languages: 'typescript'
		`,
		Run: func(cmd *cobra.Command, args []string) {
			language := args[0]
			supportedLanguages := []string{"typescript"}
			isSupported := false
			for _, s := range supportedLanguages {
				if s == language {
					isSupported = true
				}
			}
			if !isSupported {
				fmt.Printf("unsupported language received: %s\nplease input one of: \n\t- %v",
					language,
					strings.Join(supportedLanguages, "\n\t- "),
				)
				os.Exit(1)
			}
			name := args[1]
			dir, err := os.Getwd()
			if err != nil {
				fmt.Printf("failed to get working dir: %v\n", err)
				os.Exit(1)
			}

			g, err := generator.NewGenerator(language, name, dir)
			if err != nil {
				fmt.Printf("failed to create package generator: %v\n", err)
				os.Exit(1)
			}

			fmt.Printf("generating new pulumi package '%s' in language '%s' in directory '%s'\n", name, language, dir)
			fmt.Println("...")

			err = g.Generate()
			if err != nil {
				fmt.Printf("failed to generate package: %v\n", err)
			}
			fmt.Println("done!")
		},
	}

	return cmd
}
