package middlewares

import (
	"context"

	"github.com/labstack/echo/v5"

	"backend/pkg/response"
)

// tenantKeyType adalah custom type untuk key context.
type tenantKeyType string

const (
	// TenantContextKey adalah key untuk menyimpan schema name tenant di dalam context.
	TenantContextKey tenantKeyType = "tenant_schema"
)

// TenantResolver mengekstrak identitas institusi (tenant) dari request
// dan menyimpannya di context untuk digunakan oleh layer selanjutnya.
// Tenant bisa didapat dari header 'X-Tenant-ID' atau subdomain.
func TenantResolver() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			// Prioritas 1: Ambil dari header
			tenantID := c.Request().Header.Get("X-Tenant-ID")

			// Prioritas 2: Jika tidak ada di header, ambil dari origin/domain (opsional)
			// ... logika untuk parse domain bisa ditambahkan di sini

			if tenantID == "" {
				return response.BadRequest(c, "X-Tenant-ID header is required for this endpoint")
			}

			// Validasi tenantID format jika perlu
			// ...

			// Format schema name berdasarkan tenant ID (misal: tenant_uuid)
			schemaName := "tenant_" + tenantID

			// Set ke dalam request context
			ctx := context.WithValue(c.Request().Context(), TenantContextKey, schemaName)
			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}

// GetTenantSchema mengambil nama skema tenant dari context.
func GetTenantSchema(ctx context.Context) string {
	if schema, ok := ctx.Value(TenantContextKey).(string); ok {
		return schema
	}
	// Fallback ke skema public jika tidak ada context tenant
	return "public"
}
