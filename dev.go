package flogodevtool

import (
	"os"
	"path/filepath"

	"github.com/project-flogo/cli/common"

	"github.com/spf13/cobra"
)

var descCmd = &cobra.Command{
	Use:              "dev",
	Short:            "Developer tool for basic work ",
	Long:             `This command helps you to work flogo contributions `,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
	},
}
var GOPATH string
var COREPATH string

func init() {
	common.RegisterPlugin(descCmd)

	GOPATH = os.Getenv("GOPATH")

	COREPATH = filepath.Join(GOPATH, "pkg", "mod", "github.com", "project-flogo", "core@v0.9.0", "examples")
}
