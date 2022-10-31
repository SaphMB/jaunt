package training

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/SaphMB/jaunt/lib/retriever"
	"github.com/SaphMB/jaunt/lib/swagger"
)

type TrainingPeriod struct {
	StartDate time.Time
	EndDate   time.Time
	Duration  time.Duration
	RaceDate  time.Time
}

type TrainingLogger struct {
	ctx            context.Context
	TrainingPeriod TrainingPeriod
	Retriever      retriever.Retriever
}

func (t TrainingLogger) Activities() (activities []swagger.SummaryActivity, err error) {
	activities, resp, err := t.Retriever.GetLoggedInAthleteActivities(t.ctx, nil)
	if err != nil {
		return []swagger.SummaryActivity{}, fmt.Errorf("error retrieving activities: %s", err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		content, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return []swagger.SummaryActivity{}, fmt.Errorf("unexpected status code %d when retrieving activities: %s", resp.StatusCode, err.Error())
		}
		return []swagger.SummaryActivity{}, fmt.Errorf("unexpected status code %d when retrieving activities: %s", resp.StatusCode, content)
	}

	defer resp.Body.Close()
	return activities, nil
}
