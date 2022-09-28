package authentication

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Logout(ctx echo.Context, conn *gorm.DB) error {
	token := GetTokenFromHeadler(ctx)
	err := DisbleToken(conn, &token)
	if err != nil {
		return err
	}
	return nil
}

func DisbleToken(conn *gorm.DB, token *string) error {
	sql := fmt.Sprintf(
		`
		UPDATE tokenApp
		SET status = 'I'
		WHERE token = '%v';
		`, *token,
	)
	err := conn.Table("tokenApp").Exec(sql).Error
	if err != nil {
		return err
	}
	return nil
}
