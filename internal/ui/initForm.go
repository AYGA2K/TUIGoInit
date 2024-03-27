package ui

import (
	"errors"

	"github.com/AYGA2K/TUIGoInit/internal/constants"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type initForm struct {
	form           *huh.Form
	AppName        *string
	InitGit        *bool
	appNameKey     string
	initGitRepoKey string
}

var catppuccin *huh.Theme = huh.ThemeCatppuccin()

func NewInitForm() initForm {
	appNameKey := constants.KeyAppName
	confirmKey := constants.KeyGitInit

	appNameInput := huh.NewInput().
		Key(appNameKey).
		Title(constants.PromptAppName).
		Validate(func(s string) error {
			if s == "" {
				return errors.New(constants.ErrorAppNameEmpty)
			}
			return nil
		})

	gitInitConfirm := huh.NewConfirm().
		Key(confirmKey).
		Title(constants.PromptGitInit)

	form := huh.NewForm(
		huh.NewGroup(appNameInput, gitInitConfirm),
	).WithTheme(catppuccin)

	return initForm{
		form:           form,
		appNameKey:     appNameKey,
		initGitRepoKey: confirmKey,
		AppName:        new(string),
		InitGit:        new(bool),
	}
}

func (f initForm) Init() tea.Cmd {
	return f.form.Init()
}

func (f initForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return f, tea.Quit
		}
	}

	var cmd tea.Cmd
	form, cmd := f.form.Update(msg)
	if form, ok := form.(*huh.Form); ok {
		f.form = form
	}

	if f.form.State == huh.StateCompleted {
		*f.AppName = f.form.GetString(f.appNameKey)
		*f.InitGit = f.form.GetBool(f.initGitRepoKey)
		return f, tea.Quit
	}

	return f, cmd
}

func (f initForm) View() string {
	return f.form.View()
}
