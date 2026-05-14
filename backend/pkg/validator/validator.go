package validator

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

// ValidationError merepresentasikan satu kesalahan validasi pada field tertentu.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrors adalah kumpulan error validasi yang mengimplementasikan interface error.
type ValidationErrors []ValidationError

// Error mengimplementasikan interface error.
func (ve ValidationErrors) Error() string {
	var msgs []string
	for _, e := range ve {
		msgs = append(msgs, fmt.Sprintf("%s: %s", e.Field, e.Message))
	}
	return strings.Join(msgs, "; ")
}

// CustomValidator mengimplementasikan interface echo.Validator
// menggunakan go-playground/validator sebagai engine.
type CustomValidator struct {
	validate *validator.Validate
}

// New membuat instance CustomValidator baru.
func New() *CustomValidator {
	v := validator.New(validator.WithRequiredStructEnabled())
	return &CustomValidator{validate: v}
}

// Validate memvalidasi struct dan mengembalikan ValidationErrors jika gagal.
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validate.Struct(i); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return formatErrors(validationErrors)
		}
		return err
	}
	return nil
}

// formatErrors mengkonversi go-playground validation errors ke format yang user-friendly.
func formatErrors(errs validator.ValidationErrors) ValidationErrors {
	result := make(ValidationErrors, 0, len(errs))
	for _, fe := range errs {
		result = append(result, ValidationError{
			Field:   toSnakeCase(fe.Field()),
			Message: messageForTag(fe),
		})
	}
	return result
}

// messageForTag mengembalikan pesan error yang mudah dipahami berdasarkan tag validasi.
func messageForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "field ini wajib diisi"
	case "email":
		return "format email tidak valid"
	case "min":
		return fmt.Sprintf("minimal %s karakter", fe.Param())
	case "max":
		return fmt.Sprintf("maksimal %s karakter", fe.Param())
	case "len":
		return fmt.Sprintf("harus tepat %s karakter", fe.Param())
	case "uuid":
		return "format UUID tidak valid"
	case "url":
		return "format URL tidak valid"
	case "oneof":
		return fmt.Sprintf("harus salah satu dari: %s", fe.Param())
	case "gte":
		return fmt.Sprintf("harus lebih besar atau sama dengan %s", fe.Param())
	case "lte":
		return fmt.Sprintf("harus lebih kecil atau sama dengan %s", fe.Param())
	case "e164":
		return "format nomor telepon tidak valid (gunakan format E.164)"
	default:
		return fmt.Sprintf("validasi '%s' gagal", fe.Tag())
	}
}

// toSnakeCase mengubah PascalCase/camelCase menjadi snake_case.
func toSnakeCase(s string) string {
	var result strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(r))
		} else {
			result.WriteRune(r)
		}
	}
	return result.String()
}
