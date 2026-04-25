package cmd

import "github.com/spf13/cobra"

var mediaCmd = &cobra.Command{
	Use:   "media",
	Short: "Media utilities",
}

func init() {
	rootCmd.AddCommand(mediaCmd)
}
