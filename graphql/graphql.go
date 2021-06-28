package graphql

import (
	"context"
	"errors"
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
		IndexingRules []IndexingRule `graphql:"indexingRules(merged: true)"`
	}
	err := gs.Client.Query(context.Background(), &q, nil)
	if err != nil {
		return Status{}, err
	}
	status := Status{
		q.IndexerRegistration,
		q.IndexerEndpoints,
		q.IndexingRules,
	}
	return status, nil
}

func (gs *GraphService) GetIndexingRule(id string) (IndexingRule, error) {
	variables := make(map[string]interface{})
	if id == "all" {
		return IndexingRule{}, errors.New("Use status command to get all rules")
	}
	variables["deployment"] = graphql.String(id)
	var q struct {
		IndexingRule IndexingRule `graphql:"indexingRule(deployment:$deployment, merged: true)"`
	}
	err := gs.Client.Query(context.Background(), &q, variables)
	if err != nil {
		return IndexingRule{}, err
	}
	return q.IndexingRule, nil
}

func (gs *GraphService) DeleteIndexingRule(deploymentID string) (bool, error) {
	var m struct {
		DeleteIndexingRule graphql.Boolean `graphql:"deleteIndexingRule(deployment:$deployment)"`
	}
	var deployment string
	var err error
	if deploymentID != "global" {
		deployment, err = utils.SubgraphHashToHex(deploymentID)
		if err != nil {
			return false, err
		}
	} else {
		return false, errors.New("Cannot delete global rule")
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

func (gs *GraphService) GetCurrentEpoch() (Epoch, error) {
	var q struct {
		CurrentEpoch Epoch `graphql:"graphNetwork(id: 1)"`
	}
	err := gs.Client.Query(context.Background(), &q, nil)
	if err != nil {
		return Epoch{}, err
	}
	return q.CurrentEpoch, nil
}
