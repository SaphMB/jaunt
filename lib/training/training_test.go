package training_test

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/SaphMB/jaunt/lib/mocks"
	"github.com/SaphMB/jaunt/lib/swagger"
	"github.com/SaphMB/jaunt/lib/training"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Training", func() {
	var (
		ctrl           *gomock.Controller
		mock_retriever *mocks.MockRetriever
		resp           *http.Response
		period         training.TrainingPeriod
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mock_retriever = mocks.NewMockRetriever(ctrl)
		resp = &http.Response{}
	})

	AfterEach(func() {
		defer ctrl.Finish()
	})

	Describe("Activities", func() {
		Context("when the retriever returns an error", func() {
			It("returns the error", func() {
				logger := training.TrainingLogger{
					TrainingPeriod: period,
					Retriever:      mock_retriever,
				}

				mock_retriever.
					EXPECT().
					GetLoggedInAthleteActivities(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]swagger.SummaryActivity{}, resp, errors.New("an error"))

				activities, err := logger.Activities()
				Expect(err).To(MatchError("error retrieving activities: an error"))
				Expect(activities).To(BeEmpty())
			})
		})

		Context("when the retriever's response is not OK", func() {
			It("returns an error with the status code and response body", func() {
				mock_resp := &http.Response{
					StatusCode: http.StatusForbidden,
					Status:     "403 Forbidden",
					Body:       ioutil.NopCloser(bytes.NewBufferString("halt!")),
				}

				period := training.TrainingPeriod{}
				logger := training.TrainingLogger{
					TrainingPeriod: period,
					Retriever:      mock_retriever,
				}

				mock_retriever.
					EXPECT().
					GetLoggedInAthleteActivities(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]swagger.SummaryActivity{}, mock_resp, nil)

				activities, err := logger.Activities()
				Expect(err).To(MatchError("unexpected status code 403 when retrieving activities: halt!"))
				Expect(activities).To(BeEmpty())
			})
		})

		Context("when the retriever's response is not OK and cannot be read", func() {
			It("returns an error with the status code and the reader error", func() {
				var body mocks.ErrReader
				resp = &http.Response{
					StatusCode: http.StatusForbidden,
					Status:     "403 Forbidden",
					Body:       body,
				}
				defer ctrl.Finish()

				logger := training.TrainingLogger{
					TrainingPeriod: period,
					Retriever:      mock_retriever,
				}

				mock_retriever.
					EXPECT().
					GetLoggedInAthleteActivities(gomock.Any(), gomock.Any()).
					Times(1).
					Return([]swagger.SummaryActivity{}, resp, nil)

				activities, err := logger.Activities()
				Expect(err).To(MatchError("unexpected status code 403 when retrieving activities: an error"))
				Expect(activities).To(BeEmpty())
			})
		})

		It("returns an error with the status code and the reader error", func() {
			ctrl := gomock.NewController(GinkgoT())
			mock_retriever := mocks.NewMockRetriever(ctrl)
			mock_resp := &http.Response{
				StatusCode: http.StatusOK,
				Status:     "OK",
				Body:       ioutil.NopCloser(bytes.NewBufferString("")),
			}
			defer ctrl.Finish()

			period := training.TrainingPeriod{}
			logger := training.TrainingLogger{
				TrainingPeriod: period,
				Retriever:      mock_retriever,
			}

			expected_activities := []swagger.SummaryActivity{
				{
					Id: 1,
				},
				{
					Id: 2,
				},
			}

			mock_retriever.
				EXPECT().
				GetLoggedInAthleteActivities(gomock.Any(), gomock.Any()).
				Times(1).
				Return(expected_activities, mock_resp, nil)

			activities, err := logger.Activities()
			Expect(err).ToNot(HaveOccurred())
			Expect(activities).To(Equal(expected_activities))
		})
	})
})
