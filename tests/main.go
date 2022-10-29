package main

import (
	"fmt"

	goargs "github.com/budimanlai/go-args"
)

func main() {
	args := goargs.NewArgs()
	args.Parse()

	fmt.Println(`script name:`, args.ScriptName)
	fmt.Println(`Command:`, args.Command)
	fmt.Println(`Port:`, args.GetInt(`port`))
	fmt.Println(`TLS:`, args.GetIntOr(`tls`, 443))
	fmt.Println(`Token:`, args.GetString(`token`))
}
