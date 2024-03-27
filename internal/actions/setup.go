package actions

import (
	"github.com/AYGA2K/TUIGoInit/internal/constants"
	"github.com/AYGA2K/TUIGoInit/internal/templates"
	"github.com/AYGA2K/TUIGoInit/internal/util"
)

func Setup(appName string) {
	content := templates.GetMainTemplate()
	content = util.FillPlaceholder(content, constants.PlaceHolderAppName, appName)
	util.WriteFile("main.go", content)
	util.MakeDirectory("cmd")
	util.MakeDirectory("internal")
	util.MakeDirectory("internal/ui")
	cwd := util.GetCwd()
	util.ChangeDirectory("cmd")
	cobra := templates.CobraTemplates{}
	content = cobra.Root()
	content = util.FillPlaceholder(content, constants.PlaceHolderAppName, appName)
	util.WriteFile("root.go", content)
	util.ChangeDirectory(cwd)
	util.ChangeDirectory("internal/ui")
	charm := templates.CharmTemplates{}
	content = charm.UI()
	util.WriteFile("charm.go", content)
	util.ChangeDirectory(cwd)
	util.RunCommand(constants.CommandGoModTidy)
}
