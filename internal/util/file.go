package util

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AYGA2K/TUIGoInit/internal/constants"
)

func MakeDirectory(name string) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		log.Fatalf(constants.ErrorMakingDirectory + ": " + name + err.Error())
	}
}

func ChangeDirectory(dir string) {
	err := os.Chdir(dir)
	if err != nil {
		log.Fatalf(constants.ErrorChangingDirectory + ": " + dir + err.Error())
	}
}

func GetCwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf(constants.ErrorGettingCurrentDirectory + err.Error())
	}
	return cwd
}

func GitInit() {
	cmd := constants.CommandGitInit
	err := RunCommand(cmd)
	if err != nil {
		log.Fatalf(constants.ErrorInitializingGit + err.Error())
	}
}

func GoModInit(name string) {
	cmd := constants.CommandGoModInit + name
	err := RunCommand(cmd)
	if err != nil {
		log.Fatalf(constants.ErrorInitializingGoMod + " " + err.Error())
	}
}

func ReadFile(name string) string {
	content, err := os.ReadFile("internal/templates/" + name)
	if err != nil {
		log.Fatalf(constants.ErrorReadingFile + ": " + name + " " + err.Error())
	}
	return string(content)
}

func WriteFile(name, content string) {
	file, err := os.Create(name)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func FillPlaceholder(content, placeholder, value string) string {
	return strings.ReplaceAll(content, fmt.Sprintf("{{%s}}", placeholder), value)
}
