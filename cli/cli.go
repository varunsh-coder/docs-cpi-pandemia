package cli

import (
	"github.com/cuducos/docs-cpi-pandemia/downloader"
	"github.com/cuducos/docs-cpi-pandemia/filesystem"
	"github.com/cuducos/docs-cpi-pandemia/unzip"
	"github.com/spf13/cobra"
)

const help = "Downloads all public documents received by the CPI da Pandemia"

var workers uint
var retries uint
var directory string
var cleanUp bool

var cmd = &cobra.Command{
	Use:   "docs-cpi-pandemia",
	Short: help,
	RunE: func(_ *cobra.Command, _ []string) error {
		if cleanUp {
			filesystem.CleanDir(directory)
		}

		if err := downloader.Download(directory, workers, retries); err != nil {
			return err
		}

		return unzip.UnzipAll(directory)
	},
}

func CLI() *cobra.Command {
	cmd.Flags().BoolVarP(
		&cleanUp,
		"cleanup",
		"c",
		false,
		"Cleans up the target directory, including resetting the cache",
	)
	cmd.Flags().StringVarP(
		&directory,
		"directory",
		"d",
		"data",
		"Target directory to download the files",
	)
	cmd.Flags().UintVarP(
		&workers,
		"workers",
		"w",
		8,
		"Maximum parallels downloads allowed",
	)
	cmd.Flags().UintVarP(
		&retries,
		"retries",
		"r",
		16,
		"Maximum retries for the same URL",
	)
	return cmd
}
