package dtos

type OffsetBasedResponseMeta struct {
	Total      int  `json:"total,omitempty"`
	PerPage    int  `json:"perPage,omitempty"`
	Page       int  `json:"page,omitempty"`
	TotalPages int  `json:"totalPages,omitempty"`
	HasNext    bool `json:"hasNext"`
	HasPrev    bool `json:"hasPrev"`
}

type CursorBasedResponseMeta struct {
	NextCursor string `json:"nextCursor,omitempty"`
	// PrevCursor    string `json:"prevCursor,omitempty"`
	HasNext bool `json:"hasNext"`
	// HasPrev       bool   `json:"hasPrev"`
	Limit         int `json:"limit,omitempty"`
	ReturnedCount int `json:"returnedCount,omitempty"`
}

type ApiResponseList[T any, U any] struct {
	Success bool   `json:"success"`
	Data    T      `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Meta    *U     `json:"meta,omitempty"`
}

type ApiResponseSingle[T any] struct {
	Success bool   `json:"success"`
	Data    *T     `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}
