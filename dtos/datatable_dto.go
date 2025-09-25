package dtos

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// DataTableRequest - Reusable struct for pagination, sorting, filtering, and searching
type DataTableRequest struct {
	Page          int                    `json:"page" query:"page"`
	Limit         int                    `json:"limit" query:"limit"`
	Search        string                 `json:"search" query:"search"`
	Sort          string                 `json:"sort" query:"sort"`
	SortDirection string                 `json:"sortDirection" query:"sortDirection"`
	Filter        map[string]interface{} `json:"filter" query:"filter"`
}

// NewDataTableRequest creates a new DataTableRequest with default values
func NewDataTableRequest() *DataTableRequest {
	return &DataTableRequest{
		Page:          1,
		Limit:         10,
		Search:        "",
		Sort:          "created_at",
		SortDirection: "desc",
		Filter:        make(map[string]interface{}),
	}
}

// ParseFromFiberContext parses DataTableRequest from fiber context
func (dt *DataTableRequest) ParseFromFiberContext(c *fiber.Ctx) {
	// Parse basic pagination
	if page, err := strconv.Atoi(c.Query("page", "1")); err == nil && page > 0 {
		dt.Page = page
	} else {
		dt.Page = 1
	}

	if limit, err := strconv.Atoi(c.Query("limit", "10")); err == nil && limit > 0 {
		dt.Limit = limit
	} else {
		dt.Limit = 10
	}

	// Parse search
	dt.Search = strings.TrimSpace(c.Query("search", ""))

	// Parse sort
	sort := c.Query("sort", "created_at")
	if sort != "" {
		dt.Sort = sort
	}

	// Parse sort direction
	sortDirection := strings.ToLower(c.Query("sortDirection", "desc"))
	if sortDirection == "asc" || sortDirection == "desc" {
		dt.SortDirection = sortDirection
	} else {
		dt.SortDirection = "desc"
	}

	// Parse filters - you can extend this based on your needs
	// Example implementations for common filter types:
	if status := c.Query("filter.status"); status != "" {
		switch status {
		case "true":
			dt.Filter["status"] = true
		case "false":
			dt.Filter["status"] = false
		}
	}

	if startDate := c.Query("filter.startDate"); startDate != "" {
		dt.Filter["startDate"] = startDate
	}

	if endDate := c.Query("filter.endDate"); endDate != "" {
		dt.Filter["endDate"] = endDate
	}

	// Parse array filters (comma-separated values)
	if arrayFilter := c.Query("filter.array"); arrayFilter != "" {
		values := strings.Split(arrayFilter, ",")
		var intValues []int
		for _, v := range values {
			if intVal, err := strconv.Atoi(strings.TrimSpace(v)); err == nil {
				intValues = append(intValues, intVal)
			}
		}
		if len(intValues) > 0 {
			dt.Filter["array"] = intValues
		}
	}

	// Add more filter parsing as needed
}

// Validate validates the DataTableRequest
func (dt *DataTableRequest) Validate() {
	if dt.Page < 1 {
		dt.Page = 1
	}
	if dt.Limit < 1 {
		dt.Limit = 10
	}
	if dt.Sort == "" {
		dt.Sort = "created_at"
	}
	if dt.SortDirection != "asc" && dt.SortDirection != "desc" {
		dt.SortDirection = "desc"
	}
}

// GetOffset calculates the offset for database queries
func (dt *DataTableRequest) GetOffset() int {
	return (dt.Page - 1) * dt.Limit
}

// GetOrderClause returns the order clause for database queries
func (dt *DataTableRequest) GetOrderClause() string {
	return dt.Sort + " " + dt.SortDirection
}
