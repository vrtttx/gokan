package main

import "github.com/charmbracelet/bubbles/list"

type status int

const (
	todo status = iota
	inProgress
	done
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
	list list.Model
	err error
}

// TODO: call this on tea.WindowSizeMsg
func (m *Model) initList(width, height int) {
	m.list = list.New([]list.Item{}, list.NewDefaultDelegate(), width, height)
	m.list.Title = "To Do"
	m.list.SetItems([]list.Item{
		Task{status: todo, title: "buy milk", description: "strawberry milk"},
		Task{status: todo, title: "eat sushi", description: "negitoro roll, miso soup, something else"},
		Task{status: todo, title: "fold laundry", description: "or wear wrinkly t-shirts lol"},
	})
}