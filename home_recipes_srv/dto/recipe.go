package dto

import (
	"strconv"
	"strings"
)

// ModelObject interface
type ModelObject interface {
	getObjectInfo() string
}

// Recipe DTO
type Recipe struct {
	name        string       `json:"name"`
	description string       `json:"description"`
	steps       []string     `json:"steps"`
	ingredients []Ingredient `json:ingredients`
}

// Ingredient DTO
type Ingredient struct {
	Name        string
	Description string
	Quantity    int
}

// Status DTO
type Status struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

// Interface ModelObject Implementations
func (r *Recipe) getObjectInfo() string {
	return r.getName()
}

func (r *Status) getObjectInfo() string {
	info := []string{
		r.Code,
		r.Description,
	}
	resp := strings.Join(info, ": ")
	return resp
}

func (i *Ingredient) getObjectInfo() string {
	info := []string{
		i.Name,
		i.Description,
	}
	resp := strings.Join(info, ": ")
	quantity := "\nQuantity: " + strconv.Itoa(i.Quantity)
	resp += quantity
	return resp
}

// GETTERS AND SETTERS

func (r *Recipe) getName() string {
	return r.name
}
func (r *Recipe) setName(s string) {
	r.name = s
}
