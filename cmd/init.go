package cmd

import (
	"log"
	"sync"

	"github.com/AYGA2K/TUIGoInit/internal/actions"
	"github.com/AYGA2K/TUIGoInit/internal/ui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)


func init() {
	rootCmd.AddCommand(cmdInit)
}

var cmdInit = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new application",
	Args:  cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		if AppName == "" {
			form := ui.NewInitForm()
			program := tea.NewProgram(form)
			if _, err := program.Run(); err != nil {
				log.Fatal(err)
			}
			AppName = *form.AppName
			InitGit = *form.InitGit
		}

		spinner := tea.NewProgram(ui.NewSpinner())
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := spinner.Run(); err != nil {
				cobra.CheckErr(err)
			}
		}()
		actions.Init(AppName, InitGit)
		actions.Setup(AppName)
		spinner.ReleaseTerminal()
	},
}
