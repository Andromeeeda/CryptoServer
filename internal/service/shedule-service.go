package service

import (
	"cryptoserver/internal/repository"
	"cryptoserver/internal/transport"
	"log"
	"time"
)

type SheduleService struct {
	sheduleRepository *repository.SheduleRepository
	cryptoService     *CryptoService
}

func NewSheduleService(sheduleRepository *repository.SheduleRepository, cryptoService *CryptoService) *SheduleService {
	return &SheduleService{
		sheduleRepository: sheduleRepository,
		cryptoService:     cryptoService,
	}
}

//фоновая горутина

func (s *SheduleService) SheduleStart() {
	go func() {
		for {
			sheduleConfig := s.sheduleRepository.GetSheduleConfig()

			time.Sleep(time.Duration(sheduleConfig.Interval_seconds) * time.Second)

			if sheduleConfig.Enabled {

				updated_count, err := s.cryptoService.RefreshAllPrices()
				if err != nil {
					log.Printf("Auto-refresh error: %v", err)
					continue 
				}

				log.Printf("AutoRefresh updated %d cryptos", updated_count)
			}
		}
	}()
}

func (s *SheduleService) GetSheduleInfo() *repository.Shedule {
	sheduleInfo := s.sheduleRepository.GetSheduleConfig()
	return sheduleInfo
}

func (s *SheduleService) ChangeShedule(changeSheduleRequest *transport.ChangeSheduleRequest) (*repository.Shedule, error) {

	enabled := changeSheduleRequest.Enabled
	interval_seconds := changeSheduleRequest.Interval_seconds

	shedule, err := s.sheduleRepository.ChangeSheduleConfig(enabled, interval_seconds)
	if err != nil {
		return nil, err
	}

	return shedule, nil

}

func (s *SheduleService) TriggerUpdate() (int, time.Time, error) {
	update_count, err := s.cryptoService.RefreshAllPrices()
	if err != nil {
		return update_count, time.Now(), err
	}

	return update_count, time.Now(), nil
}
