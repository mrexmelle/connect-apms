package opts

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "connect-apms",
	Short:   "Connect APMS",
	Long:    "Connect APMS - An approval management system for Connect",
	Version: "0.1.0",
}
