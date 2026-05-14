package response

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

// Response adalah envelope standar untuk semua API response.
type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

// PaginationMeta menyimpan informasi paginasi.
type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	TotalItems  int64 `json:"total_items"`
	TotalPages  int   `json:"total_pages"`
}

// PaginatedResponse adalah envelope untuk response dengan paginasi.
type PaginatedResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message,omitempty"`
	Data    any            `json:"data,omitempty"`
	Meta    PaginationMeta `json:"meta"`
}

// --- Success Responses ---

// Success mengirim response sukses dengan data.
func Success(c *echo.Context, code int, data any) error {
	return c.JSON(code, Response{
		Success: true,
		Data:    data,
	})
}

// SuccessWithMessage mengirim response sukses dengan pesan dan data.
func SuccessWithMessage(c *echo.Context, code int, message string, data any) error {
	return c.JSON(code, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Created mengirim response 201 Created.
func Created(c *echo.Context, data any) error {
	return Success(c, http.StatusCreated, data)
}

// NoContent mengirim response 204 No Content.
func NoContent(c *echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

// Paginated mengirim response sukses dengan data dan informasi paginasi.
func Paginated(c *echo.Context, data any, meta PaginationMeta) error {
	return c.JSON(http.StatusOK, PaginatedResponse{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

// --- Error Responses ---

// Error mengirim response error dengan pesan.
func Error(c *echo.Context, code int, message string) error {
	return c.JSON(code, Response{
		Success: false,
		Message: message,
	})
}

// ErrorWithDetails mengirim response error dengan detail validasi.
func ErrorWithDetails(c *echo.Context, code int, message string, errors any) error {
	return c.JSON(code, Response{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}

// --- Shortcut Error Responses ---

// BadRequest mengirim response 400.
func BadRequest(c *echo.Context, message string) error {
	return Error(c, http.StatusBadRequest, message)
}

// Unauthorized mengirim response 401.
func Unauthorized(c *echo.Context, message string) error {
	return Error(c, http.StatusUnauthorized, message)
}

// Forbidden mengirim response 403.
func Forbidden(c *echo.Context, message string) error {
	return Error(c, http.StatusForbidden, message)
}

// NotFound mengirim response 404.
func NotFound(c *echo.Context, message string) error {
	return Error(c, http.StatusNotFound, message)
}

// InternalError mengirim response 500.
func InternalError(c *echo.Context, message string) error {
	return Error(c, http.StatusInternalServerError, message)
}
