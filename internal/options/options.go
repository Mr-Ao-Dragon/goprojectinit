package options

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	. "github.com/tangzixiang/goprojectinit/pkg/utils"
)

var (
	opts    HelpOptions
	version = "1.0.3"
)

// Parse 开始解析命令行参数
func Parse() *HelpOptions {

	parse := parseOptions()

	if parse == nil {
		return nil
	}

	if opts.Version {
		fmt.Printf("%v \n", version)
		return nil
	}

	if len(opts.Args.ProjectName) < 1 {
		fmt.Printf("the required argument `projectname (at least 1 argument)` was not provided\n\n")
		parse.WriteHelp(os.Stderr)
		return nil
	}

	Verbose = opts.Verbose

	return &opts
}

func parseOptions() *flags.Parser {

	parse := flags.NewParser(&opts, flags.Default)
	_, err := parse.Parse()
	if err == nil {
		return parse
	}

	// 直接输入 -h 或则 --help
	if err.(*flags.Error).Type == flags.ErrHelp {
		return nil
	}

	fmt.Println()
	parse.WriteHelp(os.Stderr)
	return nil
}

type HelpOptions struct {
	Version    bool    `short:"v" long:"version" description:"show this tool version"`
	Verbose    bool    `short:"b" long:"verbose" description:"Show verbose debug information"`
	Cover      bool    `short:"c" long:"cover" description:"if the project path exists ,cover the directory and init the project"`
	IsTool     bool    `short:"t" long:"istool" description:"istool mean this project is a tool project,so the main-file will be placed in project root directory"`
	TargetPath *string `short:"p" long:"targetpath" description:"Project should init in the which directory,default is current path,if target directory not exists will be created"`
	ConfigPath *string `short:"f" long:"configfile" description:"which init-config file should be use,if not set, default file will be download"`
	Args       struct {
		ProjectName []string `positional-arg-name:"projectname" description:"init the project with this name, the first name will be named for project,then all remaining names will be sub service name in cmd directory"`
	} `positional-args:"yes" required:"yes"`
}
