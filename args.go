package goargs

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

type Args struct {
	ScriptName string
	Command    string
	val        map[string]string
}

const (
	strLine    = `(?m)(?Ui)\s*([a-z0-9_.]+)\s*=\s*(.*)(\s+(?:#|/{2,}).*|)\s*$`
	strCommand = `\w+`
)

func NewArgs() *Args {
	return &Args{}
}

func (a *Args) Parse() {
	ar := os.Args
	base_fname := filepath.Base(ar[0])

	a.val = make(map[string]string)
	a.ScriptName = base_fname

	regexLine := regexp.MustCompile(strLine)
	regexCommand := regexp.MustCompile(strCommand)

	for _, obj := range ar[1:] {
		if matches := regexLine.FindStringSubmatch(obj); len(matches) > 0 {
			key := matches[1]
			value := matches[2]
			a.val[key] = value
		} else if matches := regexCommand.FindStringSubmatch(obj); len(matches) > 0 {
			a.Command = matches[0]
		}
	}
}

func (a *Args) GetString(name string) string {
	return a.GetStringOr(name, "")
}

// Read string property or retun defValue if property is not exists or empty
func (a *Args) GetStringOr(name string, defValue string) string {
	if val, ok := a.val[name]; ok {
		return val
	}
	return defValue
}

// Read integer property. If property is not exists or empty will return 0
func (a *Args) GetInt(name string) int {
	return a.GetIntOr(name, 0)
}

// Read integer property or return defValue if property is not exists or empty
func (a *Args) GetIntOr(name string, defValue int) int {
	if val, ok := a.val[name]; ok {
		r, e := strconv.Atoi(val)
		if e != nil {
			return defValue
		}

		return r
	}
	return defValue
}
