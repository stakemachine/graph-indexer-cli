package graphql

import (
	"context"
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/shurcooL/graphql"
	"github.com/stakemachine/graph-indexer-cli/utils"
)

// GetStatus queries status
func (gs *GraphService) GetStatus() (Status, error) {
	var q struct {
		IndexerRegistration IndexerRegistration `graphql:"indexerRegistration"`
		IndexerEndpoints    struct {
			Service  IndexerEndpoint
			Channels IndexerEndpoint
			Status   IndexerEndpoint
		} `graphql:"indexerEndpoints"`
		IndexerDeployments []IndexerDeployment `graphql:"indexerDeployments"`
		IndexingRules      []IndexingRule      `graphql:"indexingRules(merged: true)"`
		IndexerAllocations []IndexerAllocation `graphql:"indexerAllocations"`
	}
	err := gs.Client.Query(context.Background(), &q, nil)
	if err != nil {
		return Status{}, err
	}
	status := Status{
		q.IndexerRegistration,
		q.IndexerEndpoints,
		q.IndexerDeployments,
		q.IndexingRules,
		q.IndexerAllocations,
	}
	return status, nil
}

func (gs *GraphService) GetIndexingRule(id string) (IndexingRule, error) {
	variables := make(map[string]interface{})
	if id == "all" {
		return IndexingRule{}, errors.New("use status command to get all rules")
	}
	variables["identifier"] = graphql.String(id)
	var q struct {
		IndexingRule IndexingRule `graphql:"indexingRule(identifier:$identifier, merged: true)"`
	}
	err := gs.Client.Query(context.Background(), &q, variables)
	if err != nil {
		return IndexingRule{}, err
	}
	return q.IndexingRule, nil
}

func (gs *GraphService) SetIndexingRule(deploymentID string, args []string) error {
	if len(args)%2 != 0 {
		return errors.New("an uneven number of key/value pairs was passed")
	}

	rulesMap := make(map[string]string)
	for i := 0; i < len(args); i += 2 {
		rulesMap[args[i]] = args[i+1]
	}

	var m struct {
		IndexingRule IndexingRule `graphql:"setIndexingRule(rule:{identifier:$identifier,identifierType:$identifierType,allocationAmount:$allocationAmount,allocationLifetime:$allocationLifetime,parallelAllocations:$parallelAllocations,maxAllocationPercentage:$maxAllocationPercentage,minSignal:$minSignal,maxSignal:$maxSignal,minStake:$minStake,minAverageQueryFees:$minAverageQueryFees,custom:$custom,decisionBasis:$decisionBasis})"`
	}
	var deployment string
	var err error
	if deploymentID != "global" && strings.HasPrefix(deploymentID, "Qm") {
		deployment, err = utils.SubgraphHashToHex(deploymentID)
		if err != nil {
			return err
		}
	} else {
		deployment = deploymentID
	}

	variables := make(map[string]interface{})
	variables["identifier"] = graphql.String(deployment)
	for i, p := range rulesMap {
		switch i {
		case "parallelAllocations":
			var parallelAllocations int
			parallelAllocations, err = strconv.Atoi(rulesMap["parallelAllocations"])
			if err != nil {
				return err
			}
			if parallelAllocations > 0 && parallelAllocations <= math.MaxInt32 {
				variables[i] = graphql.Int(parallelAllocations)
			}
		case "allocationAmount":
			allocationAmount, e := utils.ToWei(p, 18)
			if e != nil {
				return e
			}
			variables[i] = graphql.String(allocationAmount.String())
		default:
			variables[i] = graphql.String(p)
		}
	}

	err = gs.Client.Mutate(context.Background(), &m, variables)
	if err != nil {
		return err
	}
	return nil
}

func (gs *GraphService) DeleteIndexingRule(deploymentID string) (bool, error) {
	var m struct {
		DeleteIndexingRule graphql.Boolean `graphql:"deleteIndexingRule(deployment:$deployment)"`
	}
	var deployment string
	var err error
	if deploymentID != "global" && strings.HasPrefix(deploymentID, "Qm") {
		deployment, err = utils.SubgraphHashToHex(deploymentID)
		if err != nil {
			return false, err
		}
	} else {
		return false, errors.New("cannot delete global rule")
	}

	variables := map[string]interface{}{
		"deployment": graphql.String(deployment),
	}

	err = gs.Client.Mutate(context.Background(), &m, variables)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (gs *GraphService) GetModelsWithVariables() ([]CostModel, error) {
	var q struct {
		CostModels []CostModel `graphql:"costModels"`
	}
	err := gs.Client.Query(context.Background(), &q, nil)
	if err != nil {
		return []CostModel{}, err
	}
	return q.CostModels, nil
}

func (gs *GraphService) SetModel(model CostModel) (CostModel, error) {
	var m struct {
		CostModel CostModel `graphql:"setCostModel(costModel:{deployment:$deployment, model:$model, variables:$variables})"`
	}
	variables := make(map[string]interface{})
	variables["deployment"] = graphql.String(model.Deployment)
	if model.Model != "" {
		variables["model"] = graphql.String(model.Model)
	}
	if model.Variables != "" {
		variables["variables"] = graphql.String(model.Variables)
	}
	err := gs.Client.Mutate(context.Background(), &m, variables)
	if err != nil {
		return CostModel{}, err
	}
	return m.CostModel, nil
}

func (gs *GraphService) GetActiveAllocations(indexer string) ([]Allocation, error) {
	variables := map[string]interface{}{
		"indexer": graphql.String(strings.ToLower(indexer)),
	}

	var q struct {
		Allocations []Allocation `graphql:"allocations(first: 1000, where: {indexer:$indexer, status: Active})"`
	}

	err := gs.Client.Query(context.Background(), &q, variables)
	if err != nil {
		return []Allocation{}, err
	}
	return q.Allocations, nil
}

func (gs *GraphService) GetClosedAllocations(subgraph string, epoch int) ([]Allocation, error) {
	variables := map[string]interface{}{
		"subgraph": graphql.String(strings.ToLower(subgraph)),
		"epoch":    graphql.Int(epoch),
	}
	var q struct {
		Allocations []Allocation `graphql:"allocations(first: 1000, where: {subgraphDeployment:$subgraph, status: Closed, closedAtEpoch:$epoch})"`
	}
	err := gs.Client.Query(context.Background(), &q, variables)
	if err != nil {
		return []Allocation{}, err
	}
	return q.Allocations, nil
}

func (gs *GraphService) GetIndexerInfo(indexer string) (Indexer, error) {
	variables := map[string]interface{}{
		"indexer": graphql.String(strings.ToLower(indexer)),
	}
	var q struct {
		Indexer Indexer `graphql:"indexer(id:$indexer)"`
	}
	err := gs.Client.Query(context.Background(), &q, variables)
	if err != nil {
		return Indexer{}, err
	}
	return q.Indexer, nil
}

func (gs *GraphService) GetCurrentEpoch() (GraphNetwork, error) {
	var q struct {
		GraphNetwork GraphNetwork `graphql:"graphNetwork(id: 1)"`
	}
	err := gs.Client.Query(context.Background(), &q, nil)
	if err != nil {
		return GraphNetwork{}, err
	}
	return q.GraphNetwork, nil
}

func (gs *GraphService) GetSubgraphDeploymentsSignalled() ([]SubgraphDeployment, error) {
	var q struct {
		SubgraphDeployments []SubgraphDeployment `graphql:"subgraphDeployments(first:1000, where:{signalAmount_gte: 1})"`
	}
	err := gs.Client.Query(context.Background(), &q, nil)
	if err != nil {
		return []SubgraphDeployment{}, err
	}
	return q.SubgraphDeployments, nil
}

func (gs *GraphService) GetIndexingStatuses() ([]IndexingStatus, error) {
	var q struct {
		IndexingStatuses []IndexingStatus `graphql:"indexingStatuses"`
	}
	err := gs.Client.Query(context.Background(), &q, nil)
	if err != nil {
		return []IndexingStatus{}, err
	}
	return q.IndexingStatuses, nil
}

func (gs *GraphService) GetEpochInfo(epochNumber int) (Epoch, error) {
	variables := map[string]interface{}{
		"epochNumber": graphql.Int(epochNumber),
	}
	var q struct {
		Epoch Epoch `graphql:"epoch(id:$epochNumber)"`
	}
	err := gs.Client.Query(context.Background(), &q, variables)
	if err != nil {
		return Epoch{}, err
	}
	return q.Epoch, nil
}

func (gs *GraphService) GetProofOfIndexing(blockNumber int, blockHash, indexerAddress, subgraph string) (ProofOfIndexing, error) {
	variables := map[string]interface{}{
		"indexer":     graphql.String(strings.ToLower(indexerAddress)),
		"blockNumber": graphql.Int(blockNumber),
		"blockHash":   graphql.String(blockHash),
		"subgraph":    graphql.String(subgraph),
	}
	var q struct {
		ProofOfIndexing ProofOfIndexing `graphql:"proofOfIndexing(indexer:$indexer, blockNumber:$blockNumber, blockHash:$blockHash, subgraph:$subgraph)"`
	}
	err := gs.Client.Query(context.Background(), &q, variables)
	if err != nil {
		return "", err
	}
	return q.ProofOfIndexing, nil
}
