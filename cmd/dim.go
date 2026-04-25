package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

var brightness float64
var dimming float64

var dimCmd = &cobra.Command{
	Use:   "dim input output",
	Short: "Dim an imge using ImageMagick",
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

		modulate := strconv.FormatFloat(brightness, 'f', -1, 64)
		dim := strconv.FormatFloat(dimming, 'f', -1, 64)

		brightnessCommand := exec.Command(
			"magick",
			input,
			"-modulate",
			modulate,
			output,
		)

		if err := brightnessCommand.Run(); err != nil {
			return fmt.Errorf("failed to adjust brightness of image: %w", err)
		}

		dimmingCommand := exec.Command(
			"magick",
			output,
			"-fill",
			fmt.Sprintf("rgba(0,0,0,%s)", dim),
			"-draw" ,
			"rectangle 0,0,10000,10000",
			output,
		)

		if err := dimmingCommand.Run(); err != nil {
			return fmt.Errorf("failed to adjust brightness of image: %w", err)
		}

		fmt.Println("Wrote:", output)
		return nil
	},
}

func init() {
	mediaCmd.AddCommand(dimCmd)
	dimCmd.Flags().Float64VarP(&brightness, "brightness", "b", 70.0, "Brigthness")
	dimCmd.Flags().Float64VarP(&dimming, "dimming", "d", 0.8, "Dimming")
}

