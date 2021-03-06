package cmd

import (
	"io"
	"os"

	"github.com/Shopify/vouch4cluster/listers/reader"
	"github.com/Shopify/vouch4cluster/process"
	"github.com/spf13/cobra"
)

var inputReader io.Reader = os.Stdin

// readerCmd represents the reader command
var readerCmd = &cobra.Command{
	Use:   "reader",
	Short: "Reads list of images to attest from a file",
	Long: `Reads list of images to attest from a file, which is 
separated by newlines.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		cfg := getVoucherCfg()

		lister := reader.NewImageLister(inputReader)

		err = process.LookupAndAttest(cfg, lister, os.Stdout)
		if nil != err {
			errorf("attesting images failed: %s", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(readerCmd)
	// readerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
