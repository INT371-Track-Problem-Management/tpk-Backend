package service

import (
	"mime/multipart"
	"tpk-backend/app/models/model"
	"tpk-backend/app/models/request"
	"tpk-backend/app/pkg"
)

func (s serviceTPK) EditProfile(req request.EditProfile, email string, role string) error {
	timenow := pkg.GetDatetime()
	model := model.EditProfile{
		Fname:       req.Fname,
		Lname:       req.Lname,
		Sex:         req.Sex,
		DateOfBirth: req.DateOfBirth,
		Age:         req.Age,
		Phone:       req.Phone,
		Address:     req.Address,
		UpdateAt:    timenow,
		UpdateBy:    req.UpdateBy,
	}
	err := s.repo.EditProfile(model, email, role)
	if err != nil {
		return err
	}
	return nil
}

func (s serviceTPK) CreateProfileMedia(image *multipart.FileHeader, email string) error {
	model, err := pkg.UploadProfileFile(image, email)
	if err != nil {
		return err
	}
	session := s.database.Begin()
	if err := s.repo.CreateProfileMedia(*model, session); err != nil {
		return err
	}
	session.Commit()
	return nil
}

func (s serviceTPK) UpdateProfileMedia(image *multipart.FileHeader, email string) error {
	model, err := pkg.UploadProfileFile(image, email)
	if err != nil {
		return err
	}
	session := s.database.Begin()
	if err := s.repo.DeleteProfileMedia(email, session); err != nil {
		return err
	}
	if err := s.repo.CreateProfileMedia(*model, session); err != nil {
		return err
	}
	session.Commit()
	return nil
}

func (s serviceTPK) DeleteProfileMedia(email string) error {
	session := s.database.Begin()
	if err := s.repo.DeleteProfileMedia(email, session); err != nil {
		return err
	}
	session.Commit()
	return nil
}
