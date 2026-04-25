package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

var blurRadius float64

var blurCmd = &cobra.Command{
	Use:   "blur input output",
	Short: "Blur an imge using ImageMagick",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		input := args[0]
		output := args[1]

		if _, err := os.Stat(input); err != nil {
			return fmt.Errorf("input file does not exist: %s", input)
		}

		if _, err := exec.LookPath("magick"); err != nil {
			return fmt.Errorf("missing dependency: magick")
		}

		radius := strconv.FormatFloat(blurRadius, 'f', -1, 64)

		command := exec.Command(
			"magick",
			input,
			"-blur",
			"0x"+radius,
			output,
		)

		if err := command.Run(); err != nil {
			return fmt.Errorf("failed to blur image: %w", err)
		}

		fmt.Println("Wrote:", output)
		return nil
	},
}

func init() {
	mediaCmd.AddCommand(blurCmd)
	blurCmd.Flags().Float64VarP(&blurRadius, "radius", "r", 8.0, "Blur radius")
}
