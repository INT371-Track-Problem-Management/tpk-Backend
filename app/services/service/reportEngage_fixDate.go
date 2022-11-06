package service

import "tpk-backend/app/models/response"

func (s serviceTPK) ReportEnagegeFixDateDetail(reportId string) (*response.ReportEngageFixDate, error) {
	engage, err := s.repo.ReportEnagegeByReportId(reportId)
	if err != nil {
		return nil, err
	}
	fixDate, err := s.repo.FixdateByEngageId(engage.EngageId)
	if err != nil {
		return nil, err
	}

	engageDetail := response.ReportEngageFixDate{
		EngageId:     engage.EngageId,
		Step:         engage.Step,
		SelectedDate: engage.SelectedDate,
		ReportId:     engage.ReportId,
		BuildingId:   engage.BuildingId,
		CreateBy:     engage.CreateBy,
		CreateAt:     engage.CreateAt,
		UpdateAt:     engage.UpdateAt,
		UpdateBy:     engage.UpdateBy,
		MaintainerId: engage.MaintainerId,
		Fixdate:      *fixDate,
	}

	return &engageDetail, nil
}
