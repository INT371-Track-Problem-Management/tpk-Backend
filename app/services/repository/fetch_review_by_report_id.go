package repository

import "tpk-backend/app/models/model"

func (r mysqlRepository) FetchReviewByReportId(reportId int) (*model.ReviewReports, error) {
	var review model.ReviewReports
	if err := r.conn.Table("reviewReports").Where("reportId = ?", reportId).Scan(&review).Error; err != nil {
		return nil, err
	}
	return &review, nil
}
