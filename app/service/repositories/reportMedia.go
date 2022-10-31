package repositories

import (
	"fmt"
	"tpk-backend/app/model/entity"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UploadFile(ctx echo.Context, conn *gorm.DB, media entity.ReportMediaInsert) error {
	fmt.Println("file: ", media)
	err := conn.Table("reportMedia").Create(media).Error
	if err != nil {
		return err
	}
	return nil
}

func FileByReportId(ctx echo.Context, conn *gorm.DB, reportId int) (*[]entity.ReportMedia, error) {
	var medias []entity.ReportMedia
	err := conn.Table("reportMedia").Where("reportId = ?", reportId).Scan(&medias).Error
	if err != nil {
		return nil, err
	}
	return &medias, nil
}

func DeleteFileByReportId(ctx echo.Context, conn *gorm.DB, reportId string) error {
	sql := fmt.Sprintf(
		`
		DELETE FROM 
			reportMedia
		WHERE
			reportMedia = %v
		`, reportId)
	err := conn.Exec(sql).Error
	if err != nil {
		return err
	}
	return nil
}
