package templates

import _ "embed"

type CharmTemplates struct{}

//go:embed files/internal/ui/charm.go.tmpl
var uiTemplate string

func (c CharmTemplates) UI() string {
	return uiTemplate
}
