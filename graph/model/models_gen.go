// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Todo struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	Completed bool   `json:"completed"`
}

type TodoInput struct {
	Text      string `json:"text"`
	Completed *bool  `json:"completed"`
}
