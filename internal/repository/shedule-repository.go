package repository

import (
	"errors"
	"time"
)

type SheduleRepository struct {
	sheduleConfig *Shedule
}

var SheduleConfig = Shedule{
	Enabled:          true,
	Interval_seconds: 30,
	Last_update:      time.Now(),
	Next_update:      time.Now().Add(30 * time.Second),
}

func NewSheduleRepository(sheduleConfig *Shedule) *SheduleRepository {
	return &SheduleRepository{
		sheduleConfig: sheduleConfig,
	}
}

func (r *SheduleRepository) GetSheduleConfig() *Shedule {
	return r.sheduleConfig
}

func (r *SheduleRepository) ChangeSheduleConfig(enabled bool, interval_seconds int) (*Shedule,error) {

	if r.sheduleConfig == nil {
		return nil, errors.New("shedule config not initialization")
	}

	r.sheduleConfig.Enabled = enabled
	r.sheduleConfig.Interval_seconds = interval_seconds
	r.sheduleConfig.Last_update = time.Now()
	r.sheduleConfig.Next_update = time.Now().Add(time.Duration(interval_seconds) * time.Second)

	return r.sheduleConfig,nil
}
