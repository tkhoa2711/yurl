package cmd

import (
	"fmt"

	"github.com/tkhoa2711/yurl/yurllib"

	"github.com/spf13/cobra"
)

// checkAASACmd represents the check command for Apple App Site Association file locally
var checkAASACmd = &cobra.Command{
	Use:   "check <path>",
	Short: "check the given file against Apple's requirements",
	Run: func(cmd *cobra.Command, args []string) {
		output, err := yurllib.CheckAASAFile(args[0])

		if err != nil {
			fmt.Println("Got error: %w", err)
		}

		for _, item := range output {
			fmt.Print(item)
		}

	},
}

func init() {
	aasaCmd.AddCommand(checkAASACmd)
}
