package actions

import (
	"github.com/AYGA2K/TUIGoInit/internal/util"
)

func Init(appName string, gitInit bool) {
	util.MakeDirectory(appName)
	util.ChangeDirectory(appName)
	util.GoModInit(appName)
	if gitInit {
		util.GitInit()
	}
}
