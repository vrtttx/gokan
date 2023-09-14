package main

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