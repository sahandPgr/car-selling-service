package dto

// Request to create a property category.
type CreatePropertyCategoryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"` // Category name.
	Icon string `json:"icon" binding:"max=1000"`                    // Category icon.
}

// Request to update a property category.
type UpdatePropertyCategoryRequest struct {
	Name string `json:"name,omitempty" binding:"alpha,min=3,max=15"` // Updated name.
	Icon string `json:"icon,omitempty" binding:"max=1000"`           // Updated icon.
}

// Response for a property category.
type PropertyCategoryResponse struct {
	Id         int                `json:"id"`                   // Category ID.
	Name       string             `json:"name"`                 // Category name.
	Icon       string             `json:"icon"`                 // Category icon.
	Properties []PropertyResponse `json:"properties,omitempty"` // List of properties.
}

// Request to create a property.
type CreatePropertyRequest struct {
	CategoryId  int    `json:"categoryId" binding:"required"`        // Associated category ID.
	Name        string `json:"name" binding:"required,min=3,max=15"` // Property name.
	Icon        string `json:"icon" binding:"max=1000"`              // Property icon.
	Description string `json:"description" binding:"max=1000"`       // Property description.
	DataType    string `json:"dataType" binding:"max=15"`            // Data type.
	Unt         string `json:"unit" binding:"max=15"`                // Unit.
}

// Request to update a property.
type UpdatePropertyRequest struct {
	CategoryId  int    `json:"categoryId,omitempty"`                     // Updated category ID.
	Name        string `json:"name,omitempty" binding:"min=3,max=15"`    // Updated name.
	Icon        string `json:"icon,omitempty" binding:"max=1000"`        // Updated icon.
	Description string `json:"description,omitempty" binding:"max=1000"` // Updated description.
	DataType    string `json:"dataType,omitempty" binding:"max=15"`      // Updated data type.
	Unt         string `json:"unit,omitempty" binding:"max=15"`          // Updated unit.
}

// Response for a property.
type PropertyResponse struct {
	Id          int                      `json:"id"`          // Property ID.
	Name        string                   `json:"name"`        // Property name.
	Icon        string                   `json:"icon"`        // Property icon.
	Description string                   `json:"description"` // Property description.
	DataType    string                   `json:"DataType"`    // Data type.
	Unit        string                   `json:"unit"`        // Unit.
	Category    PropertyCategoryResponse `json:"category"`    // Associated category.
}
