package main

import (
    "flag"
    "fmt"
    "os"
)

/*
type Add struct {
    FlagSet *flag.FlagSet
    Verbose bool
    Force   bool
}

type Commit struct {
    FlagSet *flag.FlagSet
    Verbose bool
    Signoff bool
}
 */

var (
    bool_flag   bool
    string_flag string
    int_flag    int
)

/*
func main_comp() {
    command := flag.NewFlagSet("command name", flag.ExitOnError)
    version := command.Bool("version", false, "Output version information and exit")

    command.Parse(os.Args[1:])

    if *version {
        fmt.Fprintf(os.Stderr, "%s v0.0.1", command.Name())
    }
}
 */

func init() {
    flag.CommandLine.Init("command name", flag.ContinueOnError)
    flag.CommandLine.Usage = func() {
        o := flag.CommandLine.Output()
        fmt.Fprintf(o, "\nUsage: %s\n", flag.CommandLine.Name())
        fmt.Fprintf(o, "\nDescription: Command description etc.\n\nOptions:\n")
        flag.PrintDefaults()
        fmt.Fprintf(o, "\nCopyright 2018 xxx.\n\n")
    }

    flag.BoolVar(&bool_flag, "bool_flag", false, "flag of boolean")
    flag.StringVar(&string_flag, "string_flag", "hello", "flag of string")
    flag.IntVar(&int_flag, "int_flag", 0, "flag of integer")

    // bool_flag := flag.Bool("bool_flag", flase, "flag of boolean")
}

func main() {
    if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
        if err != flag.ErrHelp {
            fmt.Fprintf(os.Stderr, "error: %s\n", err)
        }
        os.Exit(2)
    }

    fmt.Println("bool_flag:", bool_flag)
    fmt.Println("string_flag:", string_flag)
    fmt.Println("int_flag:", int_flag)
    fmt.Println("Args:", flag.Args())
}
