package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	mgmt "github.com/stakemachine/graph-indexer-cli/graphql"
)

type mockHTTPError struct {
	message string
}

func (e *mockHTTPError) Error() string {
	return e.message
}

type mockHTTPClient struct {
	mockDoFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) RoundTrip(req *http.Request) (*http.Response, error) {
	if m.mockDoFunc != nil {
		return m.mockDoFunc(req)
	}
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(`{"data": {}}`)),
	}, nil
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.mockDoFunc != nil {
		return m.mockDoFunc(req)
	}
	// Default mock response
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(`{"data": {}}`)),
	}, nil
}

// Helper function to create a test client
func newTestClient(mockDoFunc func(_ *http.Request) (*http.Response, error)) *http.Client {
	return &http.Client{
		Transport: &mockHTTPClient{mockDoFunc: mockDoFunc},
	}
}

func TestGetIndexingRule(t *testing.T) {
	testCases := []struct {
		name         string
		deploymentID string
		args         []string
		shouldError  bool
	}{
		{
			name:         "Valid hex ID",
			deploymentID: "0x1234567890123456789012345678901234567890123456789012345678901234",
			args:         []string{"0x1234567890123456789012345678901234567890123456789012345678901234"},
			shouldError:  false,
		},
		{
			name:         "Valid IPFS hash",
			deploymentID: "QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF",
			args:         []string{"QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF"},
			shouldError:  false,
		},
		{
			name:         "Invalid ID",
			deploymentID: "invalid-id",
			args:         []string{"invalid-id"},
			shouldError:  true,
		},
		{
			name:         "No args provided",
			deploymentID: "",
			args:         []string{},
			shouldError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := newTestClient(func(req *http.Request) (*http.Response, error) {
				if tc.shouldError {
					return nil, &mockHTTPError{message: "test error"}
				}
				return &http.Response{
					StatusCode: 200,
					Body:       io.NopCloser(bytes.NewBufferString(`{"data": {"indexingRule": {}}}`)),
				}, nil
			})

			err := getIndexingRule(context.Background(), "http://localhost:8000", tc.args, *client)

			if tc.shouldError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestSetIndexingRule(t *testing.T) {
	testCases := []struct {
		name         string
		deploymentID string
		args         []string
		shouldError  bool
	}{
		{
			name:         "Valid rule set",
			deploymentID: "QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF",
			args:         []string{"allocationAmount", "1000", "decisionBasis", "always"},
			shouldError:  false,
		},
		{
			name:         "Invalid parameters",
			deploymentID: "invalid-id",
			args:         []string{"invalid", "params"},
			shouldError:  true,
		},
		{
			name:         "Empty parameters",
			deploymentID: "QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF",
			args:         []string{},
			shouldError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := newTestClient(func(req *http.Request) (*http.Response, error) {
				if tc.shouldError {
					return nil, &mockHTTPError{message: "test error"}
				}

				body := req.Body
				defer body.Close()
				bodyBytes, err := io.ReadAll(body)
				if err != nil {
					return nil, err
				}
				bodyStr := string(bodyBytes)

				// If this is the mutation
				if strings.Contains(bodyStr, "mutation") {
					return &http.Response{
						StatusCode: 200,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"data": {
								"indexingRule": {
									"identifier": "0x31edcacc9a53bc8ab4be2eeb0d873409da4c4228cb2d60e4243bd3b4e8af7500",
									"allocationAmount": "1000000000000000000000",
									"decisionBasis": "always"
								}
							}
						}`)),
					}, nil
				}

				// If this is the subsequent query
				return &http.Response{
					StatusCode: 200,
					Body: io.NopCloser(bytes.NewBufferString(`{
						"data": {
							"indexingRule": {
								"identifier": "0x31edcacc9a53bc8ab4be2eeb0d873409da4c4228cb2d60e4243bd3b4e8af7500",
								"allocationAmount": "1000000000000000000000",
								"decisionBasis": "always"
							}
						}
					}`)),
				}, nil
			})

			err := setIndexingRule(context.Background(), "http://localhost:8000", tc.deploymentID, tc.args, *client)

			if tc.shouldError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestSetCostModel(t *testing.T) {
	testCases := []struct {
		name         string
		deploymentID string
		args         []string
		shouldError  bool
	}{
		{
			name:         "Valid cost model",
			deploymentID: "QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF",
			args:         []string{"costModel", "default => 0.1;"},
			shouldError:  false,
		},
		{
			name:         "Invalid deployment ID",
			deploymentID: "invalid-id",
			args:         []string{"costModel", "default => 0.1;"},
			shouldError:  true,
		},
		{
			name:         "Empty args",
			deploymentID: "QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF",
			args:         []string{},
			shouldError:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := newTestClient(func(req *http.Request) (*http.Response, error) {
				if tc.shouldError {
					return nil, &mockHTTPError{message: "test error"}
				}

				body := req.Body
				defer body.Close()
				bodyBytes, err := io.ReadAll(body)
				if err != nil {
					return nil, err
				}
				bodyStr := string(bodyBytes)

				// If this is the mutation
				if strings.Contains(bodyStr, "mutation") {
					return &http.Response{
						StatusCode: 200,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"data": {
								"setCostModel": {
									"deployment": "0x31edcacc9a53bc8ab4be2eeb0d873409da4c4228cb2d60e4243bd3b4e8af7500",
									"model": "default => 0.1;",
									"variables": ""
								}
							}
						}`)),
					}, nil
				}

				// If this is a query for existing models
				return &http.Response{
					StatusCode: 200,
					Body: io.NopCloser(bytes.NewBufferString(`{
						"data": {
							"costModels": [{
								"deployment": "0x31edcacc9a53bc8ab4be2eeb0d873409da4c4228cb2d60e4243bd3b4e8af7500",
								"model": "default => 0.1;",
								"variables": ""
							}]
						}
					}`)),
				}, nil
			})

			err := setCostModel("http://localhost:8000", tc.deploymentID, tc.args, *client)

			if tc.shouldError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestSignals(t *testing.T) {
	testCases := []struct {
		name            string
		networkSubgraph string
		indexNode       string
		shouldError     bool
	}{
		{
			name:            "Valid endpoints",
			networkSubgraph: "https://api.thegraph.com/subgraphs/name/graphprotocol/graph-network-mainnet",
			indexNode:       "http://localhost:8030/graphql",
			shouldError:     false,
		},
		{
			name:            "Invalid network subgraph",
			networkSubgraph: "invalid-url",
			indexNode:       "http://localhost:8030/graphql",
			shouldError:     true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := newTestClient(func(req *http.Request) (*http.Response, error) {
				if tc.shouldError {
					return nil, &mockHTTPError{message: "test error"}
				}

				// Network subgraph endpoints
				if strings.Contains(req.URL.String(), "mainnet") {
					// GraphNetwork query
					if strings.Contains(req.Header.Get("X-GraphQL-Query"), "graphNetwork") {
						return &http.Response{
							StatusCode: 200,
							Body: io.NopCloser(bytes.NewBufferString(`{
								"data": {
									"graphNetwork": {
										"currentEpoch": 420,
										"totalTokensSignalled": "1000",
										"totalTokensAllocated": "500"
									}
								}
							}`)),
						}, nil
					}

					// SubgraphDeployments query
					return &http.Response{
						StatusCode: 200,
						Body: io.NopCloser(bytes.NewBufferString(`{
							"data": {
								"subgraphDeployments": [
									{
										"id": "0x1234",
										"signalAmount": "100",
										"signalledTokens": "100",
										"stakedTokens": "50",
										"originalName": "test-graph"
									}
								]
							}
						}`)),
					}, nil
				}

				// Index node endpoint
				return &http.Response{
					StatusCode: 200,
					Body: io.NopCloser(bytes.NewBufferString(`{
						"data": {
							"indexingStatuses": [
								{
									"subgraph": "0x1234",
									"synced": true,
									"health": "healthy",
									"chains": [
										{
											"network": "mainnet",
											"chainHeadBlock": {"number": "100"},
											"latestBlock": {"number": "90"}
										}
									]
								}
							]
						}
					}`)),
				}, nil
			})

			err := signals(context.Background(), tc.networkSubgraph, tc.indexNode, *client)

			if tc.shouldError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestCreateSubgraphsPool(t *testing.T) {
	testCases := []struct {
		name                string
		subgraphDeployments []mgmt.SubgraphDeployment
		indexingStatuses    []mgmt.IndexingStatus
		minSignal           string
		expectedPoolSize    int
	}{
		{
			name: "Valid pool creation",
			subgraphDeployments: []mgmt.SubgraphDeployment{
				{
					ID:              "0x1234",
					SignalAmount:    "1000",
					SignalledTokens: "1000",
					StakedTokens:    "500",
					OriginalName:    "test-graph",
				},
			},
			indexingStatuses: []mgmt.IndexingStatus{
				{
					Subgraph: "0x1234",
					Synced:   true,
					Health:   "healthy",
					Chains: []mgmt.Chains{
						{
							Network: "mainnet",
							ChainHeadBlock: mgmt.ChainHeadBlock{
								Number: "100",
							},
							LatestBlock: mgmt.LatestBlock{
								Number: "90",
							},
						},
					},
				},
			},
			minSignal:        "0",
			expectedPoolSize: 1,
		},
		{
			name:                "Empty deployments",
			subgraphDeployments: []mgmt.SubgraphDeployment{},
			indexingStatuses:    []mgmt.IndexingStatus{},
			minSignal:           "0",
			expectedPoolSize:    0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			pool, _ := CreateSubgraphsPool(tc.subgraphDeployments, tc.indexingStatuses, mgmt.GraphNetwork{
				TotalTokensSignalled: "1000",
				TotalTokensAllocated: "500",
			}, tc.minSignal)
			if len(pool) != tc.expectedPoolSize {
				t.Errorf("Expected pool size %d, got %d", tc.expectedPoolSize, len(pool))
			}
		})
	}
}
