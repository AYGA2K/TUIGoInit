package templates

import _ "embed"

type CobraTemplates struct{}

//go:embed files/cmd/root.go.tmpl
var rootCmdTemplate string

func (c CobraTemplates) Root() string {
	return rootCmdTemplate
}
