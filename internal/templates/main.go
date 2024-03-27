package templates

import _ "embed"

//go:embed files/main.go.tmpl
var mainTemplate string

func GetMainTemplate() string {
	return mainTemplate
}
