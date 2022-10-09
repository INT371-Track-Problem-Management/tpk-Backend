package repositories

import (
	"fmt"
	"tpk-backend/app/model/entity"
	"tpk-backend/app/model/request"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func FetcStatDashBoard(ctx echo.Context, conn *gorm.DB, req request.Stat) (*entity.Stat, error) {

	count_all := 
	`
	SELECT COUNT(r.reportId) as total_report
	FROM reports r
	WHERE 1=1
	`

	count_electric := 
	`
	SELECT COUNT(r.reportId) as electric
	FROM reports r
	WHERE r.categoriesReport = 'electric'
	`

	count_water := 
	`
	SELECT COUNT(r.reportId) as water
	FROM reports r
	WHERE r.categoriesReport = 'water'
	`

	count_electric_device := 
	`
	SELECT COUNT(r.reportId) as electric_device
	FROM reports r
	WHERE r.categoriesReport = 'electric_device'
	`

	count_water_machine := 
	`
	SELECT COUNT(r.reportId) as water_machine
	FROM reports r
	WHERE r.categoriesReport = 'water_machine'
	`

	count_furniture := 
	`
	SELECT COUNT(r.reportId) as furniture
	FROM reports r
	WHERE r.categoriesReport = 'furniture'
	`

	count_building := 
	`
	SELECT COUNT(r.reportId) as building
	FROM reports r
	WHERE r.categoriesReport = 'building'
	`

	count_other := 
	`
	SELECT COUNT(r.reportId) as other
	FROM reports r
	WHERE r.categoriesReport = 'other'
	`

	var cond string
	if req.Month != 0 {
		cond += fmt.Sprintf(`AND MONTH(r.createAt) = %v`, req.Month)
	}
	if req.Year != 0 {
		cond += fmt.Sprintf(` AND YEAR(r.createAt) = %v`, req.Year)
	}

	fmt.Println("condition: " + cond)

	if cond != "" {
		count_all += cond
		count_electric += cond
		count_water += cond
		count_electric_device += cond
		count_water_machine += cond
		count_furniture += cond
		count_building += cond
		count_other += cond
	}

	sql := fmt.Sprintf(
		`
		SELECT total_report, electric, water, electric_device, water_machine, furniture, building, other
		FROM 
		(
		%v
		) as total_report,
		(
		%v
		) as electric,
		(
		%v
		) as water,
		(
		%v
		) as electric_device,
		(
		%v
		) as water_machine,
		(
		%v
		) as furniture,
		(
		%v
		) as building,
		(
		%v
		) as other
		`,
		count_all,
		count_electric,
		count_water,
		count_electric_device,
		count_water_machine,
		count_furniture,
		count_building,
		count_other,
	)

	var stat entity.Stat
	err := conn.Raw(sql).Scan(&stat).Error
	if err != nil {
		return nil, err
	}

	return &stat, nil
}
