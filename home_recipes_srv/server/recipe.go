package server

/**
** RECIPE DEFINITION
 */

// Recipe DTO
type Recipe struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Steps       []string     `json:"steps"`
	Ingredients []Ingredient `json:"ingredients"`
}

// GetName function
func (r *Recipe) GetName() string {
	return r.Name
}

// SetName function
func (r *Recipe) SetName(name string) {
	r.Name = name
}

// GetDescription function
func (r *Recipe) GetDescription() string {
	return r.Description
}

// SetDescription function
func (r *Recipe) SetDescription(description string) {
	r.Description = description
}

// GetSteps function
func (r *Recipe) GetSteps() []string {
	return r.Steps
}

// SetSteps function
func (r *Recipe) SetSteps(steps []string) {
	r.Steps = steps
}

// SetIngredients function
func (r *Recipe) SetIngredients(ingredients []Ingredient) {
	r.Ingredients = ingredients
}

// GetIngredients function
func (r *Recipe) GetIngredients() []Ingredient {
	return r.Ingredients
}

// Interface ResponseObject Implementation
func (r *Recipe) getObjectInfo() string {
	return r.GetName()
}
