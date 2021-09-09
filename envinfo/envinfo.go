package envinfo

import (
	"fmt"
	"os"
	"strings"
)

func BuildInfoString() string {
	info := "Args:\n"
	info += "\t" + strings.Join(os.Args, "\n\t")
	info += "\n\n#########################\n\n"
	dir, err := os.Getwd()
	if err != nil {
		info += "Couldn't get Working Directory from os.Getwd()\n"
	} else {
		info += fmt.Sprintf("Getwd: %s\n", dir)
	}
	info += fmt.Sprintf("CD: %s\n", os.Getenv("cd"))
	info += fmt.Sprintf("PWD: %s\n\n", os.Getenv("PWD"))
	info += "Environ All:\n"
	for _, e := range os.Environ() {
		info += fmt.Sprintf("\t%s\n", e)
		//pair := strings.SplitN(e, "=", 2)
		//info +=fmt.Sprintf("\t%s = \"%s\"", pair[0], pair[1])
	}
	return info
}

type Env struct {
	lines []string
	ents
}

func NewRealEnviroment() *Env {
	return &Env{lines: os.Environ()}
}
func NewMockEnviroment(s []string) *Env {
	return &Env{lines: s}
}

func (e *Env) Sprint() string {
	msg := ""
	for _, e := range e.lines {
		pair := strings.SplitN(e, "=", 2)
		for j, x := range pair {
			x = `"` + x + `"`
			pair[j] = x
		}
		msg += strings.Join(pair, `=`) + "\n"
	}
	return msg
}
