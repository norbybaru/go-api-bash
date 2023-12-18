package response

import (
	"dancing-pony/internal/common/jwt"
	"dancing-pony/internal/platform/paginator"
	"dancing-pony/internal/platform/validator"
)

type JsonResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Failed to process request"`
}

type UnauthenticatedResponse struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Unauthenticated"`
}

type ValidationErrorResponse struct {
	Success bool   `json:"success" example:"false"`
	Error   string `json:"error" example:"Failed validation"`
	Errors  interface{}
}

type PaginatedResponse struct {
	Success  bool        `json:"success"`
	Data     interface{} `json:"data"`
	Metadata interface{} `json:"meta"`
	Links    interface{} `json:"links"`
}

// Return unauthenticated response
func NewUnauthenticatedResponse() *UnauthenticatedResponse {
	return &UnauthenticatedResponse{
		Success: false,
		Error:   jwt.ErrorUnAuthenticated.Error(),
	}
}

// Return paginated results with metadata and page links
func NewPaginatedResponse(p *paginator.PaginatorResult, requestUri string, perPage int) *PaginatedResponse {
	return &PaginatedResponse{
		Success:  true,
		Data:     p.Records,
		Metadata: p.Paginator,
		Links:    p.Paginator.BuildLinks(requestUri, perPage),
	}
}

// Json response format for success
func NewJsonResponse(data interface{}) *JsonResponse {
	return &JsonResponse{
		Success: true,
		Data:    data,
	}
}

// Json response format for error
func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		Success: false,
		Error:   err.Error(),
	}
}

func NewValidationErrorResponse(v validator.ValidationMessage) *ValidationErrorResponse {
	return &ValidationErrorResponse{
		Success: false,
		Error:   v.Message,
		Errors:  v.Errors,
	}
}
