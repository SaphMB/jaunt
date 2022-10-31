package retriever

import (
	"context"
	"net/http"

	"github.com/SaphMB/jaunt/lib/swagger"
)

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/mock_retriever.go -package=mocks github.com/SaphMB/jaunt/lib/retriever Retriever
type Retriever interface {
	GetLoggedInAthleteActivities(ctx context.Context, localVarOptionals *swagger.ActivitiesApiGetLoggedInAthleteActivitiesOpts) ([]swagger.SummaryActivity, *http.Response, error)
}
