package service

import (
	"cryptoserver/internal/repository"
	"cryptoserver/internal/transport"
)

type SheduleService struct {
	sheduleRepository *repository.SheduleRepository
}

func NewSheduleService(sheduleRepository *repository.SheduleRepository) *SheduleService {
	return &SheduleService{
		sheduleRepository: sheduleRepository,
	}
}

func (s *SheduleService) GetSheduleInfo() *repository.Shedule{
	sheduleInfo := s.sheduleRepository.GetSheduleConfig()
	return sheduleInfo
}

func(s *SheduleService) ChangeShedule(changeSheduleRequest *transport.ChangeSheduleRequest) (*repository.Shedule,error) {


	enabled := changeSheduleRequest.Enabled
	interval_seconds := changeSheduleRequest.Interval_seconds

	shedule,err :=  s.sheduleRepository.ChangeSheduleConfig(enabled,interval_seconds)
	if err != nil {
		return nil,err
	}

	return shedule,nil

}