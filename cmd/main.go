package main

import (
	"github.com/mrexmelle/connect-apms/cmd/opts"
)

func main() {
	opts.RootCmd.CompletionOptions.DisableDefaultCmd = true
	opts.RootCmd.AddCommand(opts.ServeCmd)
	opts.RootCmd.Execute()
}
