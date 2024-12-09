package dto

// Sort represents a sorting option for a column.
type Sort struct {
	ColId string `json:"col_id"` // ColId specifies the column to sort.
	Sort  string `json:"sort"`   // Sort specifies the sorting direction (e.g., "asc" or "desc").
}

// Filter represents a filtering option applied to data.
type Filter struct {
	Type       string `json:"type"`       // Type specifies the type of filter (e.g., "range").
	From       string `json:"from"`       // From specifies the starting value for the filter range.
	To         string `json:"to"`         // To specifies the ending value for the filter range.
	FilterType string `json:"filterType"` // FilterType specifies the category of filter (e.g., "date", "number").
}

// DynaimcFilter represents dynamic sorting and filtering options.
type DynaimcFilter struct {
	Sort   *[]Sort           `json:"sort"`   // Sort contains a list of sorting options.
	Filter map[string]Filter `json:"filter"` // Filter contains a map of column IDs to their filter configurations.
}

// PagedList represents a paginated list of items.
type PagedList[T any] struct {
	PageNumber      int   `json:"pageNumber"`      // PageNumber is the current page number.
	TotalRows       int64 `json:"totalRows"`       // TotalRows is the total number of rows available.
	TotalPages      int   `json:"totalPages"`      // TotalPages is the total number of pages available.
	HasPerviousPage bool  `json:"hasPerviousPage"` // HasPerviousPage indicates whether there is a previous page.
	HasNextPage     bool  `json:"hasNextPage"`     // HasNextPage indicates whether there is a next page.
	Items           *[]T  `json:"items"`           // Items contains the list of items for the current page.
}

// PaginationInput represents the basic pagination parameters.
type PaginationInput struct {
	PageSize   int `json:"pageSize"`   // PageSize specifies the number of items per page.
	PageNumber int `json:"pageNumber"` // PageNumber specifies the current page number.
}

// PaginationInputWithFilter extends PaginationInput with dynamic filtering and sorting options.
type PaginationInputWithFilter struct {
	PaginationInput // Embeds PaginationInput for basic pagination parameters.
	DynaimcFilter   // Embeds DynaimcFilter for filtering and sorting options.
}

// GetOffset calculates the starting index for database queries based on the current page and page size.
func (p *PaginationInputWithFilter) GetOffset() int {
	return (p.PageNumber - 1) * p.PageSize
}

// GetPageNumber ensures the page number is valid and returns it.
func (p *PaginationInputWithFilter) GetPageNumber() int {
	if p.PageNumber == 0 {
		p.PageNumber = 1 // Default to the first page if PageNumber is not set.
	}
	return p.PageNumber
}

// GetPageSize ensures the page size is valid and returns it.
func (p *PaginationInputWithFilter) GetPageSize() int {
	if p.PageSize == 0 {
		p.PageSize = 10 // Default to 10 items per page if PageSize is not set.
	}
	return p.PageSize
}
