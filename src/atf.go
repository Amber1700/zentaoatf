package main

import (
	"flag"
	"github.com/easysoft/zentaoatf/src/action"
	"github.com/easysoft/zentaoatf/src/utils/common"
	configUtils "github.com/easysoft/zentaoatf/src/utils/config"
	logUtils "github.com/easysoft/zentaoatf/src/utils/log"
	"github.com/easysoft/zentaoatf/src/utils/vari"
	"os"
)

func main() {
	var language string
	var independentFile bool

	var dir string
	var files []string
	var productId string
	var moduleId string
	var taskId string
	var suite string

	flagSet := flag.NewFlagSet("atf", flag.ContinueOnError)

	flagSet.Var(commonUtils.NewSliceValue([]string{}, &files), "f", "")
	flagSet.Var(commonUtils.NewSliceValue([]string{}, &files), "file", "")

	flagSet.StringVar(&dir, "d", "", "")
	flagSet.StringVar(&dir, "dir", "", "")

	flagSet.StringVar(&productId, "p", "", "")
	flagSet.StringVar(&productId, "product", "", "")

	flagSet.StringVar(&moduleId, "m", "", "")
	flagSet.StringVar(&moduleId, "module", "", "")

	flagSet.StringVar(&taskId, "t", "", "")
	flagSet.StringVar(&taskId, "task", "", "")

	flagSet.StringVar(&suite, "s", "", "")
	flagSet.StringVar(&suite, "suite", "", "")

	flagSet.StringVar(&language, "l", "", "")
	flagSet.StringVar(&language, "language", "", "")

	flagSet.BoolVar(&independentFile, "i", false, "")
	flagSet.BoolVar(&independentFile, "independent", false, "")

	switch os.Args[1] {
	case "run":

	case "co":
		if err := flagSet.Parse(os.Args[2:]); err == nil {
			action.GenerateScript(productId, moduleId, suite, taskId, independentFile, language)
		}

	case "update":
		if err := flagSet.Parse(os.Args[2:]); err == nil {
			action.GenerateScript(productId, moduleId, suite, taskId, independentFile, language)
		}

	case "ci":

	case "list":

	case "view":

	case "set":
		configUtils.ConfigForSet()

	case "help":

	default:
		logUtils.PrintUsage()
		os.Exit(1)
	}
}

func init() {
	if len(os.Args) > 1 {
		if os.Args[1] == "cui" {
			vari.RunFromCui = true
		} else {
			vari.RunFromCui = false
		}
	}

	configUtils.InitConfig()
}
