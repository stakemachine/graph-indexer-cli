package graphql

import "github.com/shurcooL/graphql"

// GraphService stores graphql client
type GraphService struct {
	Client *graphql.Client
}

type IndexingRule struct {
	Identifier              string
	IdentifierType          string
	AllocationAmount        string
	AllocationLifetime      int
	AutoRenewal             bool
	ParallelAllocations     int
	MaxAllocationPercentage string
	MinSignal               string
	MaxSignal               string
	MinStake                string
	MinAverageQueryFees     string
	Custom                  string
	DecisionBasis           string
	RequireSupported        bool
}

type BlockPointer struct {
	Number int
	Hash   string
}

type ChainIndexingStatus struct {
	Network        string
	LatestBlock    BlockPointer
	ChainHeadBlock BlockPointer
	EarliestBlock  BlockPointer
}

type IndexingError struct {
	Handler string
	Message string
}

type IndexerDeployment struct {
	SubgraphDeployment string
	Synced             bool
	Health             string
	FatalError         IndexingError
	Node               string
	Chains             []ChainIndexingStatus
}

type IndexerAllocation struct {
	ID                 string
	AllocatedTokens    string
	CreatedAtEpoch     int
	ClosedAtEpoch      int
	SubgraphDeployment string
	SignalledTokens    string
	StakedTokens       string
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
	IndexerDeployments []IndexerDeployment
	IndexingRules      []IndexingRule
	IndexerAllocations []IndexerAllocation
}

type Allocation struct {
	ID      string
	Indexer struct {
		ID string
	}
	SubgraphDeployment struct {
		ID           string
		OriginalName string
	}
	Poi             string
	CreatedAtEpoch  int
	AllocatedTokens string
}

type Indexer struct {
	StakedTokens    string
	AllocatedTokens string
	AvailableStake  string
}

type GraphNetwork struct {
	CurrentEpoch int
}

type SubgraphDeployment struct {
	ID              string `json:"id"`
	OriginalName    string `json:"originalName"`
	SignalAmount    string `json:"signalAmount"`
	SignalledTokens string `json:"signalledTokens"`
}

type ChainHeadBlock struct {
	Number string `json:"number"`
}

type LatestBlock struct {
	Number string `json:"number"`
}

type Chains struct {
	ChainHeadBlock ChainHeadBlock `json:"chainHeadBlock"`
	LatestBlock    LatestBlock    `json:"latestBlock"`
	Network        string         `json:"network"`
}

type FatalError struct {
	Handler interface{} `json:"handler"`
}

type IndexingStatus struct {
	Chains     []Chains   `json:"chains"`
	FatalError FatalError `json:"fatalError"`
	Node       string     `json:"node"`
	Subgraph   string     `json:"subgraph"`
	Health     string     `json:"health"`
}

type ProofOfIndexing string

type Epoch struct {
	EndBlock   int    `json:"endBlock"`
	ID         string `json:"id"`
	StartBlock int    `json:"startBlock"`
}
