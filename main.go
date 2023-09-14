package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type status int

const divisor = 4;

const (
	todo status = iota
	inProgress
	done
)

/* STYLING */
var (
	columnStyle = lipgloss.NewStyle().Padding(1, 2)
	focusedStyle = lipgloss.NewStyle().Padding(1, 2).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("62"))
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

/* CUSTOM ITEM */
type Task struct {
	status status
	title string
	description string
}

// implement the list.Item interface
func (t Task) FilterValue() string {
	return t.title
}

func (t Task) Title() string {
	return t.title
}

func (t Task) Description() string {
	return t.description
}

/* MAIN MODEL */
type Model struct {
	focused status
	lists []list.Model
	loaded bool
	err error
}

func New() *Model {
	return &Model{}
}

func (m *Model) initLists(width, height int) {
	defaultList := list.New([]list.Item{}, list.NewDefaultDelegate(), width/divisor, height)
	defaultList.SetShowHelp(false)
	m.lists = []list.Model{defaultList, defaultList, defaultList}

	// init to do
	m.lists[todo].Title = "To Do"
	m.lists[todo].SetItems([]list.Item{
		Task{status: todo, title: "buy milk", description: "strawberry milk"},
		Task{status: todo, title: "eat sushi", description: "negitoro roll, miso soup, something else"},
		Task{status: todo, title: "fold laundry", description: "or wear wrinkly t-shirts lol"},
	})
	// init in progress
	m.lists[inProgress].Title = "In Progress"
	m.lists[inProgress].SetItems([]list.Item{
		Task{status: todo, title: "write code", description: "don't worry, it's Go"},
	})
	// init done
	m.lists[done].Title = "Done"
	m.lists[done].SetItems([]list.Item{
		Task{status: todo, title: "stay cool", description: "as a cucumber"},
	})
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.WindowSizeMsg:
			if !m.loaded {
				m.initLists(msg.Width, msg.Height)
				m.loaded = true
			}
	}

	var cmd tea.Cmd
	m.lists[m.focused], cmd = m.lists[m.focused].Update(msg)

	return m, cmd 
}

func (m Model) View() string {
	if m.loaded {
		todoView := m.lists[todo].View()
		inProgView := m.lists[inProgress].View()
		doneView := m.lists[done].View()

		return lipgloss.JoinHorizontal(lipgloss.Left, todoView, inProgView, doneView)
	} else {
		return "loading..."
	}
}

func main() {
	m := New()
	p := tea.NewProgram(m)

	if err := p.Start(); err != nil {
		fmt.Println(err)
		os.Exit(1);
	}
}