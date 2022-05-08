package reporting

import (
	"errors"
	"testing"

	"github.com/diegojunges/reporting/reportstore"
	mockreportstore "github.com/diegojunges/reporting/reportstore/mock"
	mockuuid "github.com/diegojunges/reporting/uuid/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ReportManagerSuite struct {
	suite.Suite
	*require.Assertions

	ctrl              *gomock.Controller
	mockReportStore   *mockreportstore.MockStore
	mockUUIDGenerator *mockuuid.MockGenerator

	manager *ReportManager
}

func TestReportManagerSuite(t *testing.T) {
	suite.Run(t, new(ReportManagerSuite))
}

func (s *ReportManagerSuite) SetupTest() {
	s.Assertions = require.New(s.T())

	s.ctrl = gomock.NewController(s.T())
	s.mockReportStore = mockreportstore.NewMockStore(s.ctrl)
	s.mockUUIDGenerator = mockuuid.NewMockGenerator(s.ctrl)

	s.manager = NewReportManager(s.mockUUIDGenerator, s.mockReportStore)
}

func (s *ReportManagerSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *ReportManagerSuite) TestCreateReport() {
	reportID := "reportid"
	userID := "userid"
	title := "title"

	s.mockUUIDGenerator.EXPECT().Generate().Return(reportID).Times(1)

	s.mockReportStore.EXPECT().CreateReport(gomock.Eq(reportstore.CreateReportRequest{
		ReportID: reportID,
		UserID:   userID,
		Status:   reportstore.ReportStatusPending.String(),
		Title:    title,
	})).Return(nil).Times(1)

	actualResponse, err := s.manager.CreateReport(CreateReportRequest{
		UserID: userID,
		Title:  title,
	})
	s.NoError(err)

	expectedResponse := CreateReportResponse{
		ReportID: reportID,
	}
	s.Equal(expectedResponse, actualResponse)
}

func (s *ReportManagerSuite) TestCreateReportError() {
	s.mockUUIDGenerator.EXPECT().Generate().Return("reportid").Times(1)

	createError := errors.New("create report error")
	s.mockReportStore.EXPECT().CreateReport(gomock.Any()).Return(createError).Times(1)

	_, err := s.manager.CreateReport(CreateReportRequest{})
	s.Equal(createError, err)
}
