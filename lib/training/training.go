package training

import "time"

type TrainingPeriod struct {
	StartDate time.Time
	EndDate time.Time
	Duration time.Duration
	RaceDate time.Time
}

type TrainingLogger struct {
	Retriever Retriever
}