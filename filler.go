package main

import (
	"fmt"
	"sort"

	"github.com/shopspring/decimal"
	mgmt "github.com/stakemachine/graph-indexer-cli/graphql"
	"github.com/stakemachine/graph-indexer-cli/utils"
)

type SubgraphsPool []SubgraphsPoolEntity

type SubgraphsPoolEntity struct {
	SubgraphHash      string
	SignalledTokens   decimal.Decimal
	StakedTokens      decimal.Decimal
	Capacity          decimal.Decimal
	AvailableCapacity decimal.Decimal
	Amount            decimal.Decimal
	CurrentRatio      decimal.Decimal
	PotentialRatio    decimal.Decimal
	Enabled           bool
}

func (sp SubgraphsPool) Staked() decimal.Decimal {
	stakedSum := decimal.NewFromInt(0)
	for _, s := range sp {
		if s.Enabled {
			stakedSum = stakedSum.Add(s.StakedTokens)
		}
	}
	return stakedSum
}

func (sp SubgraphsPool) AvailableCapacity() decimal.Decimal {
	availableCapacity := decimal.NewFromInt(0)
	for _, s := range sp {
		if s.Enabled {
			availableCapacity = availableCapacity.Add(s.AvailableCapacity)
		}
	}
	return availableCapacity
}

func (sp SubgraphsPool) TotalCapacity() decimal.Decimal {
	capacity := decimal.NewFromInt(0)
	for _, s := range sp {
		if s.Enabled {
			capacity = capacity.Add(s.Capacity)
		}
	}
	return capacity
}

func (sp SubgraphsPool) Ratio(newStake decimal.Decimal) decimal.Decimal {
	ratio := newStake.Add(sp.Staked()).Div(sp.TotalCapacity())
	return ratio
}

type SubgraphDeployments struct {
	deployments  []mgmt.SubgraphDeployment
	graphNetwork *mgmt.GraphNetwork // Assuming GraphNetwork is the type of graphNetwork
}

func (s SubgraphDeployments) Len() int {
	return len(s.deployments)
}

func (s SubgraphDeployments) Less(i, j int) bool {
	// Call the AvailableCapacity method for each SubgraphDeployment
	capacityI, errI := s.deployments[i].AvailableCapacity(s.graphNetwork)
	capacityJ, errJ := s.deployments[j].AvailableCapacity(s.graphNetwork)

	// If there's an error getting capacity, treat the deployment as having less capacity
	if errI != nil {
		return false
	}
	if errJ != nil {
		return true
	}

	// Compare the capacities
	return capacityI.Cmp(capacityJ) > 0
}

func (s SubgraphDeployments) Swap(i, j int) {
	s.deployments[i], s.deployments[j] = s.deployments[j], s.deployments[i]
}

func CreateSubgraphsPool(subgraphs []mgmt.SubgraphDeployment, indexingStatuses []mgmt.IndexingStatus, graphNetwork mgmt.GraphNetwork, _ string) (SubgraphsPool, error) {
	subgraphDeployments := SubgraphDeployments{
		deployments:  subgraphs,
		graphNetwork: &graphNetwork,
	}
	subgraphsPool := SubgraphsPool{}
	sort.Sort(subgraphDeployments)

	indexingStatusMap := make(map[string]mgmt.IndexingStatus)
	for _, status := range indexingStatuses {
		indexingStatusMap[status.Subgraph] = status
	}

	var errors []error

	for _, s := range subgraphDeployments.deployments {
		availableCapacity, err := s.AvailableCapacity(&graphNetwork)
		if err != nil {
			errors = append(errors, err)
			continue
		}

		subgraphHash, err := s.Hash()
		if err != nil {
			errors = append(errors, err)
			continue
		}

		status, exists := indexingStatusMap[subgraphHash]
		if !exists || !status.Synced || status.Health != "healthy" {
			continue
		}

		capacity, err := s.Capacity(&graphNetwork)
		if err != nil {
			errors = append(errors, err)
			continue
		}

		currentRatio, err := s.CurrentRatio(&graphNetwork)
		if err != nil {
			errors = append(errors, err)
			continue
		}

		stakedTokens, err := utils.ToDecimal(s.StakedTokens, 18)
		if err != nil {
			errors = append(errors, fmt.Errorf("subgraphStakedTokensDecimal error: %w", err))
			continue
		}

		signalledTokens, err := utils.ToDecimal(s.SignalledTokens, 18)
		if err != nil {
			errors = append(errors, fmt.Errorf("subgraphSignalledTokensDecimal error: %w", err))
			continue
		}

		subgraphsPool = append(subgraphsPool, SubgraphsPoolEntity{
			SubgraphHash:      subgraphHash,
			StakedTokens:      stakedTokens,
			SignalledTokens:   signalledTokens,
			Capacity:          capacity,
			AvailableCapacity: availableCapacity,
			Amount:            decimal.Decimal{},
			CurrentRatio:      currentRatio,
			PotentialRatio:    decimal.Decimal{},
			Enabled:           true,
		})
	}

	if len(errors) > 0 {
		return subgraphsPool, fmt.Errorf("encountered errors: %v", errors)
	}

	return subgraphsPool, nil
}
