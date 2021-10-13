package domain

var (
	phonyStats = &BasicStatistics{
		Avg: 3.12,
		Sum: 345346,
	}
	StatisticsDao statisticsDaoInterface
)

func init() {
	StatisticsDao = &statisticsDao{}
}

type statisticsDaoInterface interface {
	GetBasicStatistics(int64, int64) (*BasicStatistics, error)
}

type statisticsDao struct{}

func (s *statisticsDao) GetBasicStatistics(startTs int64, endTs int64) (*BasicStatistics, error) {
	// Get data from DB
	return phonyStats, nil
	// return err
}
