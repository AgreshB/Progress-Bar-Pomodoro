package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	mcobra "github.com/muesli/mango-cobra"
	"github.com/muesli/roff"
	"github.com/spf13/cobra"
)

type model struct {
	name         string
	altscreen    bool
	duration     time.Duration
	passed       time.Duration
	start        time.Time
	timer        timer.Model
	progress     progress.Model
	quitting     bool
	interrupting bool
}

func (m model) Init() tea.Cmd {
	
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
}

func (m model) View() string {
	
}

var (
	name                string
	altscreen           bool
	winHeight, winWidth int
	version             = "dev"
	quitKeys            = key.NewBinding(key.WithKeys("esc", "q"))
	intKeys             = key.NewBinding(key.WithKeys("ctrl+c"))
	boldStyle           = lipgloss.NewStyle().Bold(true)
	italicStyle         = lipgloss.NewStyle().Italic(true)
)

const (
	padding  = 2
	maxWidth = 80
)

var rootCmd = &cobra.Command{
	Use:          "timer",
	Short:        "timer is like sleep, but with progress report",
	Version:      version,
	SilenceUsage: true,
	Args:         cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		addSuffixIfArgIsNumber(&(args[0]), "s")
		duration, err := time.ParseDuration(args[0])
		if err != nil {
			return err
		}
		var opts []tea.ProgramOption
		if altscreen {
			opts = append(opts, tea.WithAltScreen())
		}
		return nil
	},
}



func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

