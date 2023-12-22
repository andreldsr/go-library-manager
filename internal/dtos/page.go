package dtos

import "math"

type Page[T any] struct {
	Content          []T `json:"content,omitempty"`
	TotalPages       int `json:"totalPages,omitempty"`
	TotalElements    int `json:"totalElements,omitempty"`
	NumberOfElements int `json:"NumberOfElements"`
	Number           int `json:"number"`
	Size             int `json:"size,omitempty"`
}

func BuildPage[T any](content []T, count, pageNumber, pageSize int) (page Page[T]) {
	page.Content = content
	page.TotalPages = int(math.Ceil(float64(count) / float64(pageSize)))
	page.TotalElements = count
	page.NumberOfElements = len(content)
	page.Number = pageNumber
	page.Size = pageSize
	return
}
