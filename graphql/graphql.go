package graphql

import (
	"context"
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/shopspring/decimal"
	"github.com/shurcooL/graphql"
	"github.com/stakemachine/graph-indexer-cli/utils"
)

// GetStatus queries status
func (gs *GraphService) GetStatus(protocolNetwork string) (Status, error) {
	variables := map[string]interface{}{
		"protocolNetwork": graphql.String(protocolNetwork),
	}
	var q struct {
		IndexerRegistration IndexerRegistration `graphql:"indexerRegistration(protocolNetwork: $protocolNetwork)"`
		IndexerEndpoints    struct {
			Service  IndexerEndpoint
			Channels IndexerEndpoint
			Status   IndexerEndpoint
		} `graphql:"indexerEndpoints(protocolNetwork: $protocolNetwork)"`
		IndexerDeployments []IndexerDeployment `graphql:"indexerDeployments(protocolNetwork: $protocolNetwork)"`
		IndexingRules      []IndexingRule      `graphql:"indexingRules(merged: true)"`
		IndexerAllocations []IndexerAllocation `graphql:"indexerAllocations(protocolNetwork: $protocolNetwork)"`
	}
	err := gs.Client.Query(context.Background(), &q, variables)
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
	spew.Dump(status)
	return status, nil
}

func (gs *GraphService) GetIndexingRule(id string) (IndexingRule, error) {
	variables := make(map[string]interface{})
	if id == "all" {
		return IndexingRule{}, errors.New("use status command to get all rules")
	}

	identifier, err := utils.ReturnIdentifier(id)
	if err != nil {
		return IndexingRule{}, err
	}
	variables["identifier"] = graphql.String(identifier)
	var q struct {
		IndexingRule IndexingRule `graphql:"indexingRule(identifier:$identifier, merged: true)"`
	}
	err = gs.Client.Query(context.Background(), &q, variables)
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
	/* 	var deployment string
	   	var err error
	   	if deploymentID != "global" && strings.HasPrefix(deploymentID, "Qm") {
	   		deployment, err = utils.SubgraphHashToHex(deploymentID)
	   		if err != nil {
	   			return err
	   		}
	   	} else {
	   		deployment = deploymentID
	   	} */

	deployment, err := utils.ReturnIdentifier(deploymentID)
	if err != nil {
		return err
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
		DeleteIndexingRule graphql.Boolean `graphql:"deleteIndexingRule(identifier:$identifier)"`
	}

	deployment, err := utils.ReturnIdentifier(deploymentID)
	if err != nil {
		return false, err
	}

	variables := map[string]interface{}{
		"identifier": graphql.String(deployment),
	}

	err = gs.Client.Mutate(context.Background(), &m, variables)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (gs *GraphService) GetCostModelsWithVariables() ([]CostModel, error) {
	var q struct {
		CostModels []CostModel `graphql:"costModels"`
	}
	err := gs.Client.Query(context.Background(), &q, nil)
	if err != nil {
		return []CostModel{}, err
	}
	return q.CostModels, nil
}

func (gs *GraphService) SetCostModel(model CostModel) (CostModel, error) {
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
	if epoch > math.MaxInt32 || epoch < math.MinInt32 {
		return nil, errors.New("epoch number out of range for int32")
	}
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

func (gs *GraphService) GetGraphNetwork() (GraphNetwork, error) {
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
		SubgraphDeployments []SubgraphDeployment `graphql:"subgraphDeployments(first:1000, where:{signalledTokens_gte: 100},orderBy: signalledTokens, orderDirection: desc)"`
	}
	err := gs.Client.Query(context.Background(), &q, nil)
	if err != nil {
		return []SubgraphDeployment{}, err
	}
	return q.SubgraphDeployments, nil
}

func (gs *GraphService) GetSubgraphDeploymentByID(deployment string) (SubgraphDeployment, error) {
	variables := map[string]interface{}{
		"deployment": graphql.String(deployment),
	}
	var q struct {
		SubgraphDeployment SubgraphDeployment `graphql:"subgraphDeployment(id:$deployment)"`
	}
	err := gs.Client.Query(context.Background(), &q, variables)
	if err != nil {
		return SubgraphDeployment{}, err
	}
	return q.SubgraphDeployment, nil
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
	if epochNumber > math.MaxInt32 || epochNumber < math.MinInt32 {
		return Epoch{}, errors.New("epoch number out of range for int32")
	}
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
	if blockNumber > math.MaxInt32 || blockNumber < math.MinInt32 {
		return "", errors.New("block number out of range for int32")
	}

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

func (s *SubgraphDeployment) Hash() (string, error) {
	subgraphHash, err := utils.SubgraphHexToHash(s.ID)
	if err != nil {
		return "", err
	}
	return subgraphHash, nil
}

func (s *SubgraphDeployment) StakedRatio(graphNetwork *GraphNetwork) (decimal.Decimal, error) {
	subgraphStakedTokens, err := decimal.NewFromString(s.StakedTokens)
	if err != nil {
		return decimal.Decimal{}, errors.New("failed to convert signal amount to big.Int")
	}
	totalTokensAllocated, err := decimal.NewFromString(graphNetwork.TotalTokensAllocated)
	if err != nil {
		return decimal.Decimal{}, errors.New("failed to convert signal amount to big.Float")
	}
	stakedRatio := subgraphStakedTokens.Div(totalTokensAllocated)
	return stakedRatio, nil
}

func (s *SubgraphDeployment) SignalledRatio(graphNetwork *GraphNetwork) (decimal.Decimal, error) {
	subgraphSignalledTokens, err := decimal.NewFromString(s.SignalledTokens)
	if err != nil {
		return decimal.Decimal{}, errors.New("failed to convert signal amount to big.Int")
	}
	totalTokensSignalled, err := decimal.NewFromString(graphNetwork.TotalTokensSignalled)
	if err != nil {
		return decimal.Decimal{}, errors.New("failed to convert signal amount to big.Float")
	}
	if totalTokensSignalled.IsZero() {
		return decimal.Decimal{}, errors.New("total tokens signalled is zero, cannot divide by zero")
	}
	signalledRatio := subgraphSignalledTokens.Div(totalTokensSignalled)
	return signalledRatio, nil
}

func (s *SubgraphDeployment) CurrentRatio(graphNetwork *GraphNetwork) (currentRatio decimal.Decimal, err error) {
	signalledRatio, err := s.SignalledRatio(graphNetwork)
	if err != nil {
		return
	}
	stakedRatio, err := s.StakedRatio(graphNetwork)
	if err != nil {
		return
	}
	if stakedRatio.IsZero() {
		currentRatio, err = decimal.NewFromString("0")
		if err != nil {
			return currentRatio, errors.New("failed to create decimal from string '0'")
		}
	} else {
		currentRatio = signalledRatio.Div(stakedRatio)
	}
	return
}

func (s *SubgraphDeployment) Capacity(graphNetwork *GraphNetwork) (capacity decimal.Decimal, err error) {
	signalledRatio, err := s.SignalledRatio(graphNetwork)
	if err != nil {
		return decimal.Decimal{}, err
	}
	totalTokensAllocated, err := utils.ToDecimal(graphNetwork.TotalTokensAllocated, 18)
	if err != nil {
		return decimal.Decimal{}, errors.New("failed to convert signal amount to big.Float")
	}
	capacity = totalTokensAllocated.Mul(signalledRatio)

	return
}

func (s *SubgraphDeployment) AvailableCapacity(graphNetwork *GraphNetwork) (availaleCapacity decimal.Decimal, err error) {
	capacity, err := s.Capacity(graphNetwork)
	if err != nil {
		return
	}
	subgraphStakedTokens, err := utils.ToDecimal(s.StakedTokens, 18)
	if err != nil {
		return availaleCapacity, errors.New("failed to convert string to big.Int")
	}
	availaleCapacity = capacity.Sub(subgraphStakedTokens)
	return
}
