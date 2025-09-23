package dtos

type PaginationMeta struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalRows  int64 `json:"totalRows"`
	TotalPages int   `json:"totalPages"`
}

type PaginatedUsersResponse struct {
	Meta  PaginationMeta    `json:"meta"`
	Users []UserResponseDTO `json:"users"`
}
