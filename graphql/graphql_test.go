package graphql

import (
	"context"
	"testing"
)

// GraphQLQuerier defines the interface we need to mock
type GraphQLQuerier interface {
	Query(ctx context.Context, q interface{}, variables map[string]interface{}) error
}

// GraphService can now use either the real client or our mock
type testGraphService struct {
	querier GraphQLQuerier
}

// Custom mock implementation
type mockQuerier struct {
	response interface{}
	err      error
}

func (m *mockQuerier) Query(_ context.Context, q interface{}, _ map[string]interface{}) error {
	if m.err != nil {
		return m.err
	}

	switch v := q.(type) {
	case *struct{ IndexingStatuses []IndexingStatus }:
		if m.response == nil {
			*v = struct{ IndexingStatuses []IndexingStatus }{
				IndexingStatuses: []IndexingStatus{},
			}
			return nil
		}
		*v = struct{ IndexingStatuses []IndexingStatus }{
			IndexingStatuses: []IndexingStatus{
				{
					Subgraph: "Qm123",
					Node:     "test-node",
					Health:   "healthy",
					Synced:   true,
					Chains: []Chains{
						{
							Network: "mainnet",
							ChainHeadBlock: ChainHeadBlock{
								Number: "100",
							},
							LatestBlock: LatestBlock{
								Number: "90",
							},
						},
					},
				},
			},
		}
	case *struct{ Status Status }:
		if m.response == struct{}{} {
			*v = struct{ Status Status }{
				Status: Status{
					IndexerRegistration: IndexerRegistration{
						URL:        "",
						Address:    "",
						Registered: false,
					},
				},
			}
			return nil
		}
		*v = struct{ Status Status }{
			Status: Status{
				IndexerRegistration: IndexerRegistration{
					URL:        "http://test.com",
					Address:    "0x123",
					Registered: true,
				},
			},
		}
	case *struct{ CostModels []CostModel }:
		if m.response == nil {
			*v = struct{ CostModels []CostModel }{
				CostModels: []CostModel{},
			}
			return nil
		}
		*v = struct{ CostModels []CostModel }{
			CostModels: []CostModel{
				{
					Deployment: "test-deployment",
					Model:      "default => 0.1;",
					Variables:  `{"test": "value"}`,
				},
			},
		}
	}
	return nil
}

type queryError struct {
	Message string
}

func (e *queryError) Error() string {
	return e.Message
}

// Mock versions of the service methods for testing
func (s *testGraphService) GetIndexingStatuses() ([]IndexingStatus, error) {
	var q struct {
		IndexingStatuses []IndexingStatus
	}
	err := s.querier.Query(context.Background(), &q, nil)
	return q.IndexingStatuses, err
}

func (s *testGraphService) GetStatus() (Status, error) {
	var q struct {
		Status Status
	}
	err := s.querier.Query(context.Background(), &q, nil)
	return q.Status, err
}

func (s *testGraphService) GetCostModelsWithVariables() ([]CostModel, error) {
	var q struct {
		CostModels []CostModel
	}
	err := s.querier.Query(context.Background(), &q, nil)
	return q.CostModels, err
}

func TestGetIndexingStatuses(t *testing.T) {
	t.Run("With data", func(t *testing.T) {
		mockService := &testGraphService{
			querier: &mockQuerier{
				response: struct{}{}, // trigger non-nil response
			},
		}
		statuses, err := mockService.GetIndexingStatuses()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(statuses) != 1 {
			t.Errorf("Expected 1 status, got %d", len(statuses))
			return
		}
		if statuses[0].Subgraph != "Qm123" {
			t.Errorf("Expected subgraph Qm123, got %s", statuses[0].Subgraph)
		}
		if !statuses[0].Synced {
			t.Error("Expected synced to be true")
		}
	})

	t.Run("Empty response", func(t *testing.T) {
		mockService := &testGraphService{
			querier: &mockQuerier{
				response: nil,
			},
		}
		statuses, err := mockService.GetIndexingStatuses()
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if len(statuses) != 0 {
			t.Errorf("Expected 0 statuses, got %d", len(statuses))
		}
	})
}

func TestGetStatus(t *testing.T) {
	mockService := &testGraphService{
		querier: &mockQuerier{},
	}

	status, err := mockService.GetStatus()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if status.IndexerRegistration.URL != "http://test.com" {
		t.Errorf("Expected URL http://test.com, got %s", status.IndexerRegistration.URL)
	}

	if status.IndexerRegistration.Address != "0x123" {
		t.Errorf("Expected address 0x123, got %s", status.IndexerRegistration.Address)
	}

	if !status.IndexerRegistration.Registered {
		t.Error("Expected registered to be true")
	}
}

func TestGraphServiceWithError(t *testing.T) {
	mockService := &testGraphService{
		querier: &mockQuerier{
			err: &queryError{Message: "Test error"},
		},
	}

	_, err := mockService.GetStatus()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestGetCostModels(t *testing.T) {
	mockService := &testGraphService{
		querier: &mockQuerier{},
	}

	models, err := mockService.GetCostModelsWithVariables()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Add assertions for cost models
	if models == nil {
		t.Error("Expected non-nil cost models")
	}
}

func TestGetStatusWithEmptyResponse(t *testing.T) {
	emptyMock := &mockQuerier{
		response: struct{}{},
	}
	mockService := &testGraphService{
		querier: emptyMock,
	}

	status, err := mockService.GetStatus()
	if err != nil {
		t.Errorf("Expected no error on empty response, got %v", err)
	}

	// Verify empty response handling
	if status.IndexerRegistration.URL != "" {
		t.Error("Expected empty URL for empty response")
	}
}

func TestGetIndexingStatusesWithNilResponse(t *testing.T) {
	nilMock := &mockQuerier{
		response: nil,
	}
	mockService := &testGraphService{
		querier: nilMock,
	}

	statuses, err := mockService.GetIndexingStatuses()
	if err != nil {
		t.Errorf("Expected no error on nil response, got %v", err)
	}

	if len(statuses) != 0 {
		t.Errorf("Expected empty statuses array, got %d items", len(statuses))
	}
}

func TestGraphServiceWithNetworkError(t *testing.T) {
	mockService := &testGraphService{
		querier: &mockQuerier{
			err: &queryError{Message: "Network connection failed"},
		},
	}

	_, err := mockService.GetStatus()
	if err == nil {
		t.Error("Expected network error, got nil")
	}
	if err.Error() != "Network connection failed" {
		t.Errorf("Expected 'Network connection failed' error, got '%s'", err.Error())
	}
}
