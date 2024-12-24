package dto

// CreateColorRequest for creating a new color.
type CreateColorRequest struct {
	Name    string `json:"name" binding:"required,max=15"`
	HexCode string `json:"hexCode" binding:"required,max=15"`
}

// UpdateColorRequest for updating color details.
type UpdateColorRequest struct {
	Name    string `json:"name,omitempty"`
	HexCode string `json:"hexCode,omitempty"`
}

// ColorResponse for returning color details.
type ColorResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	HexCode string `json:"hexCode,omitempty"`
}
