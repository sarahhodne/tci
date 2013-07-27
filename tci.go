package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"sort"
)

var (
	showHelp    = flag.Bool("h", false, "show help")
	showVersion = flag.Bool("v", false, "print version string")
)

type cmd struct {
	f interface{}
	a string // args
	d string // short description
}

var (
	selfName = os.Args[0]
	cmds     = map[string]cmd{}
	cmdHelp  = map[string]string{}
)

const (
	usage1 = `
Each command takes zero or more options and zero or more arguments.
In addition, there are some global options that can be used with any command.
The exit status is 0 on success, 1 otherwise.

Global Options:
`

	usage2 = `Commands:
`
)

func usage() {
	fmt.Fprintf(os.Stderr, "Use: %s [options] <command> [options] [args]\n", selfName)
	fmt.Fprint(os.Stderr, usage1)
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr)

	fmt.Fprint(os.Stderr, usage2)
	var max int
	var names []string
	us := make(map[string]string)
	for k := range cmds {
		u := k + " " + cmds[k].a
		if len(u) > max {
			max = len(u)
		}
		names = append(names, k)
		us[k] = u
	}
	sort.Strings(names)
	for _, k := range names {
		fmt.Fprintf(os.Stderr, "  %-*s - %s\n", max, us[k], cmds[k].d)
	}
}

func bail(e error) {
	fmt.Fprintln(os.Stderr, "Error:", e)
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		usage()
		return
	}

	if *showVersion {
		fmt.Println("tci", version)
		return
	}

	if flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "%s: missing command\n", os.Args[0])
		usage()
		os.Exit(127)
	}

	cmd := flag.Arg(0)

	c, ok := cmds[cmd]
	if !ok {
		fmt.Fprintln(os.Stderr, "Unknown command:", cmd)
		usage()
		os.Exit(127)
	}

	os.Args = flag.Args()
	flag.Parse()

	if *showHelp {
		help(cmd)
		return
	}

	args := flag.Args()
	ft := reflect.TypeOf(c.f)
	if len(args) != ft.NumIn() {
		fmt.Fprintf(os.Stderr, "%s: wrong number of arguments\n", cmd)
		help(cmd)
		os.Exit(127)
	}

	vals := make([]reflect.Value, len(args))
	for i, s := range args {
		vals[i] = reflect.ValueOf(s)
	}
	fv := reflect.ValueOf(c.f)
	fv.Call(vals)
}
