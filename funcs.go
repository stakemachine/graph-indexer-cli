package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/shurcooL/graphql"
	mgmt "github.com/stakemachine/graph-indexer-cli/graphql"
	"github.com/stakemachine/graph-indexer-cli/utils"
)

func printRules(rules []mgmt.IndexingRule) error {
	if len(rules) > 1 {
		fmt.Println("Indexing Rules")
	}
	ti := table.NewWriter()
	ti.SetOutputMirror(os.Stdout)
	ti.AppendHeader(table.Row{"deployment", "allocationAmount", "parallelAllocations", "maxAllocationPercentage", "minSignal", "maxSignal", "minStake", "minAverageQueryFees", "custom", "decisionBasis"})
	for _, r := range rules {

		encoded := ""
		var err error
		if r.Deployment != "global" {

			encoded, err = utils.SubgraphHexToHash(r.Deployment)

			if err != nil {
				return err
			}

		} else {
			encoded = r.Deployment
		}
		ti.AppendRow(table.Row{encoded, utils.ToDecimal(r.AllocationAmount, 18), r.ParallelAllocations, r.MaxAllocationPercentage, r.MinSignal, r.MaxSignal, r.MinStake, r.MinAverageQueryFees, r.Custom, r.DecisionBasis})
		ti.AppendSeparator()
	}
	ti.SetStyle(table.StyleLight)
	ti.Style().Format.Header = text.FormatDefault

	ti.Render()
	return nil
}

func printCostModels(costModels []mgmt.CostModel) error {
	if len(costModels) > 1 {
		fmt.Println("Cost Models")
	}

	tcm := table.NewWriter()
	tcm.SetOutputMirror(os.Stdout)
	tcm.AppendHeader(table.Row{"deployment", "costModel", "variables"})

	for _, cm := range costModels {
		deployment, err := utils.SubgraphHexToHash(cm.Deployment)
		if err != nil {
			return err
		}
		tcm.AppendRow(table.Row{deployment, cm.Model, cm.Variables})
		tcm.AppendSeparator()
	}

	tcm.SetStyle(table.StyleLight)
	tcm.Style().Format.Header = text.FormatDefault

	tcm.Render()
	return nil
}

func status(ctx context.Context, agentHost string, networkSubgraph string) error {
	mgmtAPI := graphql.NewClient(agentHost, nil)
	gqlClient := mgmt.GraphService{Client: mgmtAPI}

	status, err := gqlClient.GetStatus()
	if err != nil {
		panic(err)
	}
	subgraphAPI := graphql.NewClient(networkSubgraph, nil)
	subgraphAPIClient := mgmt.GraphService{Client: subgraphAPI}
	allos, err := subgraphAPIClient.GetActiveAllocations(status.IndexerRegistration.Address)
	if err != nil {
		panic(err)
	}
	indexerInfo, err := subgraphAPIClient.GetIndexerInfo(status.IndexerRegistration.Address)
	if err != nil {
		panic(err)
	}
	fmt.Println("Registration")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"URL", "Address", "Registered", "Location"})
	t.AppendRows([]table.Row{
		{status.IndexerRegistration.URL, status.IndexerRegistration.Address, status.IndexerRegistration.Registered, status.IndexerRegistration.Location},
	})
	t.SetStyle(table.StyleLight)
	t.Render()
	fmt.Println("Endpoints")
	te := table.NewWriter()
	te.SetOutputMirror(os.Stdout)
	te.AppendHeader(table.Row{"Name", "URL", "Status"})
	te.AppendRows([]table.Row{
		{"service", status.IndexerEndpoints.Service.URL, status.IndexerEndpoints.Service.Healthy},
		{"status", status.IndexerEndpoints.Status.URL, status.IndexerEndpoints.Status.Healthy},
	})
	te.SetStyle(table.StyleLight)
	te.Render()
	fmt.Println("Indexer info")
	ti := table.NewWriter()
	ti.SetOutputMirror(os.Stdout)
	ti.AppendHeader(table.Row{"Staked tokens", "Allocated tokens", "Available stake"})
	ti.AppendRows([]table.Row{
		{utils.ToDecimal(indexerInfo.StakedTokens, 18), utils.ToDecimal(indexerInfo.AllocatedTokens, 18), utils.ToDecimal(indexerInfo.AvailableStake, 18)},
	})
	ti.SetStyle(table.StyleLight)
	ti.Render()
	if len(allos) > 0 {
		fmt.Println("Active allocations")
		ta := table.NewWriter()
		ta.SetOutputMirror(os.Stdout)
		ta.AppendHeader(table.Row{"Allocation ID", "Subgraph Deployment ID", "Subgraph Name", "Created at Epoch", "Allocated tokens"})
		for _, a := range allos {
			subgraphHash, err := utils.SubgraphHexToHash(a.SubgraphDeployment.ID)
			if err != nil {
				panic(err)
			}
			ta.AppendRows([]table.Row{
				{a.ID, subgraphHash, a.SubgraphDeployment.OriginalName, a.CreatedAtEpoch, utils.ToDecimal(a.AllocatedTokens, 18)},
			})
		}
		ta.SetStyle(table.StyleLight)
		ta.Render()
	} else {
		fmt.Println("No active allocations")
	}

	err = printRules(status.IndexingRules)
	if err != nil {
		return err
	}

	return nil
}

func getRule(ctx context.Context, agentHost string, args []string) error {
	if len(args) != 1 {
		return errors.New("rules get requires one hash argument")
	}
	mgmtAPI := graphql.NewClient(agentHost, nil)
	gqlClient := mgmt.GraphService{Client: mgmtAPI}
	deployment := args[0]
	var id string
	var err error
	if strings.HasPrefix(deployment, "Qm") {
		id, err = utils.SubgraphHashToHex(deployment)
		if err != nil {
			return err
		}
	} else {
		id = deployment
	}
	r, err := gqlClient.GetIndexingRule(id)
	if err != nil {
		return err
	}
	var rules []mgmt.IndexingRule
	rules = append(rules, r)
	err = printRules(rules)
	if err != nil {
		return err
	}
	return nil
}

func setRule(ctx context.Context, agentHost string, deploymentID string, args []string) error {
	if len(args)%2 != 0 {
		return errors.New("An uneven number of key/value pairs was passed")
	}
	mgmtAPI := graphql.NewClient(agentHost, nil)
	gqlClient := mgmt.GraphService{Client: mgmtAPI}
	rulesMap := make(map[string]string)
	for i := 0; i < len(args); i += 2 {
		rulesMap[args[i]] = args[i+1]
	}

	var m struct {
		IndexingRule mgmt.IndexingRule `graphql:"setIndexingRule(rule:{deployment:$deployment, allocationAmount:$allocationAmount,parallelAllocations:$parallelAllocations,maxAllocationPercentage:$maxAllocationPercentage,minSignal:$minSignal,maxSignal:$maxSignal,minStake:$minStake,minAverageQueryFees:$minAverageQueryFees,custom:$custom,decisionBasis:$decisionBasis})"`
	}
	var deployment string
	var err error
	if deploymentID != "global" {
		deployment, err = utils.SubgraphHashToHex(deploymentID)
		if err != nil {
			return err
		}
	} else {
		deployment = deploymentID
	}

	variables := make(map[string]interface{})
	variables["deployment"] = graphql.String(deployment)
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
			variables[i] = graphql.String(utils.ToWei(p, 18).String())
		default:
			variables[i] = graphql.String(p)
		}
	}

	err = gqlClient.Client.Mutate(context.Background(), &m, variables)
	if err != nil {
		return err
	}
	var rules []mgmt.IndexingRule
	rules = append(rules, m.IndexingRule)
	err = printRules(rules)
	if err != nil {
		return err
	}
	return nil
}

func deleteRule(ctx context.Context, agentHost string, args []string) error {
	if len(args) != 1 {
		return errors.New("rules get requires one hash argument")
	}
	mgmtAPI := graphql.NewClient(agentHost, nil)
	gqlClient := mgmt.GraphService{Client: mgmtAPI}
	deploymentID := args[0]
	ok, err := gqlClient.DeleteIndexingRule(deploymentID)
	if err != nil {
		return err
	}
	if ok {
		fmt.Println("Rule for ", deploymentID, "has been deleted")
	}
	return nil
}

func getAllModelsWithVariables(agentHost string) error {
	mgmtAPI := graphql.NewClient(agentHost, nil)
	gqlClient := mgmt.GraphService{Client: mgmtAPI}
	modelsWithVariables, err := gqlClient.GetModelsWithVariables()
	if err != nil {
		return err
	}
	err = printCostModels(modelsWithVariables)
	if err != nil {
		return err
	}
	return nil
}

func setCostModel(agentHost string, deploymentID string, args []string) error {
	if len(args)%2 != 0 {
		return errors.New("An uneven number of key/value pairs was passed")
	}
	modelMap := make(map[string]string)
	for i := 0; i < len(args); i += 2 {
		modelMap[args[i]] = args[i+1]
	}
	deployment, err := utils.SubgraphHashToHex(deploymentID)
	if err != nil {
		return fmt.Errorf("Error SubgraphHashToHex: %w", err)
	}
	costModel := mgmt.CostModel{
		Deployment: deployment,
		Model:      modelMap["costModel"],
		Variables:  modelMap["variables"],
	}

	mgmtAPI := graphql.NewClient(agentHost, nil)
	gqlClient := mgmt.GraphService{Client: mgmtAPI}
	modelResponse, err := gqlClient.SetModel(costModel)
	if err != nil {
		return err
	}
	var costModels []mgmt.CostModel
	costModels = append(costModels, modelResponse)
	err = printCostModels(costModels)
	if err != nil {
		return err
	}
	return nil
}
