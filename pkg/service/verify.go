package service

import "github.com/rob-bender/meetsite-backend/pkg/repository"

type VerifyService struct {
	repo repository.TodoVerify
}

func NewVerifyService(r repository.TodoVerify) *VerifyService {
	return &VerifyService{
		repo: r,
	}
}

func (s *VerifyService) EmailVerify(uid string) (bool, int, error) {
	return s.repo.EmailVerify(uid)
}