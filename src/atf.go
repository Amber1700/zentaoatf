package main

import (
	"flag"
	"github.com/easysoft/zentaoatf/src/action"
	configUtils "github.com/easysoft/zentaoatf/src/utils/config"
	fileUtils "github.com/easysoft/zentaoatf/src/utils/file"
	logUtils "github.com/easysoft/zentaoatf/src/utils/log"
	stdinUtils "github.com/easysoft/zentaoatf/src/utils/stdin"
	"github.com/easysoft/zentaoatf/src/utils/vari"
	"os"
)

func main() {
	var language string
	var independentFile bool
	var keywords string

	var productId string
	var moduleId string
	var taskId string
	var suiteId string
	//var caseId string

	flagSet := flag.NewFlagSet("atf", flag.ContinueOnError)

	//flagSet.Var(commonUtils.NewSliceValue([]string{}, &files), "file", "")

	flagSet.StringVar(&productId, "p", "", "")
	flagSet.StringVar(&productId, "product", "", "")

	flagSet.StringVar(&moduleId, "m", "", "")
	flagSet.StringVar(&moduleId, "module", "", "")

	flagSet.StringVar(&suiteId, "s", "", "")
	flagSet.StringVar(&suiteId, "suiteId", "", "")

	flagSet.StringVar(&taskId, "t", "", "")
	flagSet.StringVar(&taskId, "taskId", "", "")

	flagSet.StringVar(&language, "l", "", "")
	flagSet.StringVar(&language, "language", "", "")

	flagSet.BoolVar(&independentFile, "i", false, "")
	flagSet.BoolVar(&independentFile, "independent", false, "")

	flagSet.StringVar(&keywords, "k", "", "")
	flagSet.StringVar(&keywords, "keywords", "", "")

	//flagSet.StringVar(&caseId, "c", "", "")
	//flagSet.StringVar(&caseId, "case", "", "")

	var placeholder string
	flagSet.StringVar(&placeholder, "h", "", "")
	flagSet.StringVar(&placeholder, "r", "", "")
	flagSet.StringVar(&placeholder, "v", "", "")

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "run", ".")
	}

	switch os.Args[1] {
	case "run", "-r":
		files := fileUtils.GetFilesFromParams(os.Args[2:])
		if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
			action.Run(files, suiteId, taskId)
		}

	case "checkout", "co":
		if err := flagSet.Parse(os.Args[2:]); err == nil {
			action.Generate(productId, moduleId, suiteId, taskId, independentFile, language)
		}

	case "update", "up":
		if err := flagSet.Parse(os.Args[2:]); err == nil {
			action.Generate(productId, moduleId, suiteId, taskId, independentFile, language)
		}

	case "ci":
		files := fileUtils.GetFilesFromParams(os.Args[2:])
		if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
			action.CommitCases(files)
		}

	case "cr":
		files := fileUtils.GetFilesFromParams(os.Args[2:])
		if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
			action.CommitResult(files)
		}

	case "cb":
		files := fileUtils.GetFilesFromParams(os.Args[2:])
		if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
			action.CommitBug(files)
		}

	case "list", "ls", "-l":
		files := fileUtils.GetFilesFromParams(os.Args[2:])
		if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
			action.List(files, keywords)
		}

	case "view", "-v":
		files := fileUtils.GetFilesFromParams(os.Args[2:])
		if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
			action.View(files, keywords)
		}

	case "set", "-s":
		stdinUtils.InputForSet()

	case "help", "-h":
		logUtils.PrintUsage()

	default:
		if len(os.Args) > 1 { // ignore run param, like atf suite.cs, atf -task 1
			args := []string{os.Args[0]}
			args = append(args, "run")
			args = append(args, os.Args[1:]...)

			os.Args = args
			files := fileUtils.GetFilesFromParams(os.Args[2:])
			if err := flagSet.Parse(os.Args[len(files)+2:]); err == nil {
				action.Run(files, suiteId, taskId)
			}
		} else {
			logUtils.PrintUsage()
		}
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
