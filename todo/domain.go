package todo

import "github.com/pascencio/gotodo/domain"

type Todo struct {
	domain.Domain
	Title       string `json:"title"`
	Description string `json:"description"`
}
