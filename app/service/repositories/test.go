package repositories

import (
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TestRepository(ctx echo.Context, req request.Test, conn *gorm.DB) (*entity.Test, error) {
	var data entity.Test
	ctx.Logger()
	err := conn.Table("test").Select("USERID", "NAME").Where("USERID = ?", req.UserId).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}
