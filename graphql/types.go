package graphql

import "github.com/shurcooL/graphql"

// GraphService stores graphql client
type GraphService struct {
	Client *graphql.Client
}

type IndexingRule struct {
	Deployment              string
	AllocationAmount        string
	ParallelAllocations     int
	MaxAllocationPercentage string
	MinSignal               string
	MaxSignal               string
	MinStake                string
	MinAverageQueryFees     string
	Custom                  string
	DecisionBasis           string
}

type GeoLocation struct {
	Latitude  string
	Longitude string
}

type IndexerRegistration struct {
	URL        string
	Address    string
	Registered bool
	Location   GeoLocation
}

type IndexerEndpointTest struct {
	Test            string
	Error           string
	PossibleActions []string
}

type IndexerEndpoint struct {
	URL     string
	Healthy bool
	Tests   []IndexerEndpointTest
}

type CostModel struct {
	Deployment string
	Model      string
	Variables  string
}

type Status struct {
	IndexerRegistration IndexerRegistration
	IndexerEndpoints    struct {
		Service  IndexerEndpoint
		Channels IndexerEndpoint
		Status   IndexerEndpoint
	}
	IndexingRules []IndexingRule
}
