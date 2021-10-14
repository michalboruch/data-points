package services

import "calc/domain"

type statsService struct{}

var StatisticsService statsService

func (s *statsService) GetBasicStatistics(name string, startTs int64, endTs int64) (*domain.BasicStatistics, error) {
	return domain.StatisticsDao.GetBasicStatistics(name, startTs, endTs)
}
