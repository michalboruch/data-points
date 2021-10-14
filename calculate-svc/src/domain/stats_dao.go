package domain

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const DB_DNS = "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"

var StatisticsDao statisticsDaoInterface

func init() {
	StatisticsDao = &statisticsDao{}
}

type statisticsDaoInterface interface {
	GetBasicStatistics(string, int64, int64) (*BasicStatistics, error)
}

type statisticsDao struct{}

func (s *statisticsDao) GetBasicStatistics(name string, startTs int64, endTs int64) (*BasicStatistics, error) {
	// Get data from DB
	db, err := sql.Open("postgres", DB_DNS)
	if err != nil {
		log.Fatal("Failed to open connection to the DB. err:", err)
		return nil, err
	}
	defer db.Close()
	var stats BasicStatistics
	statSql := "SELECT AVG(value) AS avg, SUM(value) AS sum FROM data_point_datapoint WHERE name = $1 AND timestamp BETWEEN $2 AND $3"
	err = db.QueryRow(statSql, name, startTs, endTs).Scan(&stats.Avg, &stats.Sum)
	if err != nil {
		log.Fatal("Failed to execute query. err: ", err)
		return nil, err
	}
	return &stats, nil
}
