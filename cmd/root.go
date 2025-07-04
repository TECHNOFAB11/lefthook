package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"

	"github.com/evilmartians/lefthook/internal/lefthook"
	"github.com/evilmartians/lefthook/internal/log"
)

func newRootCmd() *cobra.Command {
	options := lefthook.Options{
		Fs: afero.NewOsFs(),
	}

	rootCmd := &cobra.Command{
		Use:   "lefthook",
		Short: "CLI tool to manage Git hooks",
		Long: heredoc.Doc(`
				After installation go to your project directory
				and execute the following command:
				lefthook install
		`),
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	rootCmd.PersistentFlags().BoolVarP(
		&options.Verbose, "verbose", "v", false, "verbose output",
	)

	rootCmd.PersistentFlags().StringVar(
		&options.Colors, "colors", "auto", "'auto', 'on', or 'off'",
	)

	rootCmd.PersistentFlags().BoolVar(
		&options.NoColors, "no-colors", false, "disable colored output",
	)

	rootCmd.PersistentFlags().StringVar(
		&options.ConfigPath, "config", "", "custom config path",
	)

	// To be dropped in next releases.
	rootCmd.Flags().BoolVarP(
		&options.Force, "force", "f", false,
		"use command-specific --force option",
	)
	rootCmd.Flags().BoolVarP(
		&options.Aggressive, "aggressive", "a", false,
		"use --force flag instead",
	)
	err := rootCmd.PersistentFlags().MarkDeprecated("no-colors", "use --colors")
	if err != nil {
		log.Warn("Unexpected error:", err)
	}
	err = rootCmd.Flags().MarkDeprecated("aggressive", "use command-specific --force option")
	if err != nil {
		log.Warn("Unexpected error:", err)
	}
	err = rootCmd.Flags().MarkDeprecated("force", "use command-specific --force option")
	if err != nil {
		log.Warn("Unexpected error:", err)
	}

	for _, subcommand := range commands {
		rootCmd.AddCommand(subcommand.New(&options))
	}

	return rootCmd
}
