package Json

import (
	"TON/internal/dto"

	"github.com/labstack/echo/v4"
)

func JSONError(c echo.Context, status int, errMsg string, details string) error {
	return c.JSON(status, dto.ErrorResponseDTO{
		Error:   errMsg,
		Details: details,
	})
}
