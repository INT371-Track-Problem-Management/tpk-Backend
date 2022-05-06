package repositories

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CheckHealthy(ctx echo.Context, conn *gorm.DB) (*string, error) {
	var data string
	ctx.Logger()
	err := conn.Raw(`SELECT "Healthy" FROM DUAL`).Scan(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
