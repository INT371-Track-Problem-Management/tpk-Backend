package service

import "tpk-backend/app/models/model"

func (s serviceTPK) FetchReviewByReportId(reportId int) (*model.ReviewReports, error) {
	review, err := s.repo.FetchReviewByReportId(reportId)
	if err != nil {
		return nil, err
	}
	return review, nil
}
