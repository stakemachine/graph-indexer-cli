package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/shurcooL/graphql"
	"github.com/stakemachine/graph-indexer-cli/eth"
	mgmt "github.com/stakemachine/graph-indexer-cli/graphql"
	"github.com/stakemachine/graph-indexer-cli/utils"
)

func askForConfirmation() bool {
	var response string

	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}

	switch strings.ToLower(response) {
	case "y", "yes":
		return true
	case "n", "no":
		return false
	default:
		fmt.Println("Please type only (y)es or (n)o and then press enter:")
		return askForConfirmation()
	}
}

func printIndexerDeployments(indexerDeployments []mgmt.IndexerDeployment) error {
	if len(indexerDeployments) >= 1 {
		fmt.Println("Indexer Deployments")
	} else {
		fmt.Println("No active deployments")
	}
	td := table.NewWriter()
	td.SetOutputMirror(os.Stdout)
	td.AppendHeader(table.Row{"subgraphDeployment", "synced", "health", "fatalError", "node", "network", "latestBlock", "chainHeadBlock", "earliestBlock"})
	for _, d := range indexerDeployments {
		td.AppendRow(table.Row{d.SubgraphDeployment, d.Synced, d.Health, d.FatalError, d.Node, d.Chains[0].Network, d.Chains[0].LatestBlock.Number, d.Chains[0].ChainHeadBlock.Number, d.Chains[0].EarliestBlock.Number})
		td.AppendSeparator()
	}
	td.SetStyle(table.StyleLight)
	td.Style().Format.Header = text.FormatDefault
	td.Render()
	return nil
}

func printIndexerAllocations(indexerAllocations []mgmt.IndexerAllocation) error {
	if len(indexerAllocations) >= 1 {
		fmt.Println("Indexer Allocations")
	} else {
		fmt.Println("No active allocations")
	}
	ta := table.NewWriter()
	ta.SetOutputMirror(os.Stdout)
	ta.AppendHeader(table.Row{"id", "subgraphDeployment", "allocatedTokens", "createdAtEpoch", "signalledTokens", "stakedTokens"})
	for _, a := range indexerAllocations {
		allocatedTokens, err := utils.ToDecimal(a.AllocatedTokens, 18)
		if err != nil {
			log.Fatalln(err)
		}
		signalledTokens, err := utils.ToDecimal(a.SignalledTokens, 18)
		if err != nil {
			log.Fatalln(err)
		}
		stakedTokens, err := utils.ToDecimal(a.StakedTokens, 18)
		if err != nil {
			log.Fatalln(err)
		}
		ta.AppendRow(table.Row{a.ID, a.SubgraphDeployment, allocatedTokens, a.CreatedAtEpoch, signalledTokens, stakedTokens})
		ta.AppendSeparator()
	}
	ta.SetStyle(table.StyleLight)
	ta.Style().Format.Header = text.FormatDefault
	ta.Render()
	return nil
}

func printRules(rules []mgmt.IndexingRule) error {
	if len(rules) >= 1 {
		fmt.Println("Indexing Rules")
	}
	ti := table.NewWriter()
	ti.SetOutputMirror(os.Stdout)
	ti.AppendHeader(table.Row{"identifier", "identifierType", "allocationAmount", "allocationLifetime", "autoRenewal", "parallelAllocations", "maxAllocationPercentage", "minSignal", "maxSignal", "minStake", "minAverageQueryFees", "custom", "decisionBasis", "requireSupported"})
	for _, r := range rules {

		encoded := ""
		var err error
		if r.Identifier != "global" && strings.HasPrefix(r.Identifier, "0x") {

			encoded, err = utils.SubgraphHexToHash(r.Identifier)

			if err != nil {
				return err
			}

		} else {
			encoded = r.Identifier
		}
		allocationAmount, err := utils.ToDecimal(r.AllocationAmount, 18)
		if err != nil {
			fmt.Println("Error:", err, encoded)
		}
		ti.AppendRow(table.Row{encoded, r.IdentifierType, allocationAmount, r.AllocationLifetime, r.AutoRenewal, r.ParallelAllocations, r.MaxAllocationPercentage, r.MinSignal, r.MaxSignal, r.MinStake, r.MinAverageQueryFees, r.Custom, r.DecisionBasis, r.RequireSupported})
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

func status(ctx context.Context, agentHost string, networkSubgraph string, httpClient http.Client) error {
	mgmtAPI := graphql.NewClient(agentHost, &httpClient)
	gqlClient := mgmt.GraphService{Client: mgmtAPI}

	status, err := gqlClient.GetStatus()
	if err != nil {
		log.Fatalln(err)
	}
	subgraphAPI := graphql.NewClient(networkSubgraph, &httpClient)
	subgraphAPIClient := mgmt.GraphService{Client: subgraphAPI}
	allos, err := subgraphAPIClient.GetActiveAllocations(status.IndexerRegistration.Address)
	if err != nil {
		log.Fatalln(err)
	}
	indexerInfo, err := subgraphAPIClient.GetIndexerInfo(status.IndexerRegistration.Address)
	if err != nil {
		log.Fatalln(err)
	}
	currentEpoch, err := subgraphAPIClient.GetCurrentEpoch()
	if err != nil {
		log.Fatalln(err)
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
	stakedTokens, err := utils.ToDecimal(indexerInfo.StakedTokens, 18)
	if err != nil {
		log.Fatalln(err)
	}
	allocatedTokens, err := utils.ToDecimal(indexerInfo.AllocatedTokens, 18)
	if err != nil {
		log.Fatalln(err)
	}
	availableStake, err := utils.ToDecimal(indexerInfo.AvailableStake, 18)
	if err != nil {
		log.Fatalln(err)
	}
	ti.AppendRows([]table.Row{
		{stakedTokens, allocatedTokens, availableStake},
	})
	ti.SetStyle(table.StyleLight)
	ti.Render()
	fmt.Println("Current epoch:", currentEpoch.CurrentEpoch)

	err = printIndexerAllocations(status.IndexerAllocations)
	if err != nil {
		return err
	}

	if len(allos) > 0 {
		fmt.Printf("Active allocations (%d)\n", len(allos))
		ta := table.NewWriter()
		ta.SetOutputMirror(os.Stdout)
		ta.AppendHeader(table.Row{"Allocation ID", "Subgraph Deployment ID", "Subgraph Name", "Created at Epoch", "Allocated tokens", "Signalled", "Staked Tokens"})
		for _, a := range allos {
			subgraphHash, e := utils.SubgraphHexToHash(a.SubgraphDeployment.ID)
			if e != nil {
				log.Fatalln(e)
			}
			allocatedTokens, e := utils.ToDecimal(a.AllocatedTokens, 18)
			if e != nil {
				log.Fatalln(e)
			}
			ta.AppendRows([]table.Row{
				{a.ID, subgraphHash, a.SubgraphDeployment.OriginalName, a.CreatedAtEpoch, allocatedTokens},
			})
		}
		ta.SetStyle(table.StyleLight)
		ta.Render()
	} else {
		fmt.Println("No active allocations")
	}

	err = printIndexerDeployments(status.IndexerDeployments)
	if err != nil {
		return err
	}

	err = printRules(status.IndexingRules)
	if err != nil {
		return err
	}

	return nil
}

func getRule(ctx context.Context, agentHost string, args []string, httpClient http.Client) error {
	if len(args) != 1 {
		return errors.New("rules get requires one hash argument")
	}
	mgmtAPI := graphql.NewClient(agentHost, &httpClient)
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

func setRule(ctx context.Context, agentHost string, deploymentID string, args []string, httpClient http.Client) error {
	mgmtAPI := graphql.NewClient(agentHost, &httpClient)
	gqlClient := mgmt.GraphService{Client: mgmtAPI}
	err := gqlClient.SetIndexingRule(deploymentID, args)
	if err != nil {
		return err
	}
	var rules []mgmt.IndexingRule
	ir, err := gqlClient.GetIndexingRule(deploymentID)
	if err != nil {
		return err
	}
	rules = append(rules, ir)
	err = printRules(rules)
	if err != nil {
		return err
	}
	return nil
}

func deleteRule(ctx context.Context, agentHost string, args []string, httpClient http.Client) error {
	if len(args) != 1 {
		return errors.New("rules get requires one hash argument")
	}
	mgmtAPI := graphql.NewClient(agentHost, &httpClient)
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

func getAllModelsWithVariables(agentHost string, httpClient http.Client) error {
	mgmtAPI := graphql.NewClient(agentHost, &httpClient)
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

func setCostModel(agentHost string, deploymentID string, args []string, httpClient http.Client) error {
	if len(args)%2 != 0 {
		return errors.New("an uneven number of key/value pairs was passed")
	}
	modelMap := make(map[string]string)
	for i := 0; i < len(args); i += 2 {
		modelMap[args[i]] = args[i+1]
	}
	deployment, err := utils.SubgraphHashToHex(deploymentID)
	if err != nil {
		return fmt.Errorf("error SubgraphHashToHex: %w", err)
	}
	costModel := mgmt.CostModel{
		Deployment: deployment,
		Model:      modelMap["costModel"],
		Variables:  modelMap["variables"],
	}

	mgmtAPI := graphql.NewClient(agentHost, &httpClient)
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

func signals(ctx context.Context, networkSubgraph string, httpClient http.Client) error {
	subgraphAPI := graphql.NewClient(networkSubgraph, &httpClient)
	subgraphAPIClient := mgmt.GraphService{Client: subgraphAPI}
	subgraphDeployments, err := subgraphAPIClient.GetSubgraphDeploymentsSignalled()
	if err != nil {
		return err
	}
	totalSignalAmount := big.NewFloat(0)
	totalSignalledTokens := big.NewFloat(0)
	for _, s := range subgraphDeployments {
		subgraphSignalAmount, ok := new(big.Float).SetString(s.SignalAmount)
		if !ok {
			return errors.New("failed to convert signal amount to big.Float")
		}
		totalSignalAmount = totalSignalAmount.Add(totalSignalAmount, subgraphSignalAmount)
		subgraphSignalledTokens, ok := new(big.Float).SetString(s.SignalledTokens)
		if !ok {
			return errors.New("failed to convert signalled tokens to big.Float")
		}
		totalSignalledTokens = totalSignalledTokens.Add(totalSignalledTokens, subgraphSignalledTokens)
	}

	sort.SliceStable(subgraphDeployments, func(i, j int) bool {
		numA, _ := new(big.Int).SetString(subgraphDeployments[i].SignalAmount, 10)
		numB, _ := new(big.Int).SetString(subgraphDeployments[j].SignalAmount, 10)
		return numA.Cmp(numB) > 0
	})

	ts := table.NewWriter()
	ts.SetOutputMirror(os.Stdout)
	ts.AppendHeader(table.Row{"#", "Subgraph Deployment ID", "Subgraph Original Name", "Signal Amount", "%", "Signalled Tokens", "%"})
	for i, s := range subgraphDeployments {
		subgraphDeploymentHash, err := utils.SubgraphHexToHash(s.ID)
		if err != nil {
			return err
		}
		subgraphSignalAmount, ok := new(big.Float).SetString(s.SignalAmount)
		if !ok {
			return errors.New("failed to convert signal amount to big.Int")
		}
		subgraphSignalledTokens, ok := new(big.Float).SetString(s.SignalledTokens)
		if !ok {
			return errors.New("failed to convert signal amount to big.Int")
		}
		percentageSignalAmount := fmt.Sprintf("%.2f", new(big.Float).Quo(subgraphSignalAmount, totalSignalAmount))
		percentageSignalledTokens := fmt.Sprintf("%.2f", new(big.Float).Quo(subgraphSignalledTokens, totalSignalledTokens))

		signalledTokens, err := utils.ToDecimal(s.SignalledTokens, 18)
		if err != nil {
			return err
		}
		ts.AppendRows([]table.Row{
			{i, subgraphDeploymentHash, s.OriginalName, s.SignalAmount, percentageSignalAmount, signalledTokens.Round(2).String(), percentageSignalledTokens},
		})
	}

	ts.SetStyle(table.StyleLight)
	ts.Render()

	return nil
}

func getIndexingStatuses(ctx context.Context, indexNode string, httpClient http.Client) error {
	indexNodeAPI := graphql.NewClient(indexNode, &httpClient)
	indexNodeAPIClient := mgmt.GraphService{Client: indexNodeAPI}
	indexingStatuses, err := indexNodeAPIClient.GetIndexingStatuses()
	if err != nil {
		return err
	}
	tis := table.NewWriter()
	tis.SetOutputMirror(os.Stdout)
	tis.AppendHeader(table.Row{"#", "Subgraph Deployment ID", "Chain", "Node ID", "Latest Block", "Chain Head", "Health"})
	for i, is := range indexingStatuses {
		tis.AppendRows([]table.Row{
			{i, is.Subgraph, is.Chains[0].Network, is.Node, is.Chains[0].LatestBlock.Number, is.Chains[0].ChainHeadBlock.Number, is.Health},
		})
	}
	tis.SetStyle(table.StyleLight)
	tis.Render()
	return nil
}

func getPoi(ctx context.Context, ethNode, indexNode, networkSubgraph, indexerAddress, subgraph, epoch string, httpClient http.Client) error {
	epochInt, err := strconv.Atoi(epoch)
	if err != nil {
		return err
	}
	subgraphAPI := graphql.NewClient(networkSubgraph, &httpClient)
	subgraphAPIClient := mgmt.GraphService{Client: subgraphAPI}
	epochInfo, err := subgraphAPIClient.GetEpochInfo(epochInt)
	if err != nil {
		return err
	}
	ethClient, err := ethclient.Dial(ethNode)
	if err != nil {
		return err
	}
	blockInfo, err := ethClient.BlockByNumber(ctx, big.NewInt(int64(epochInfo.StartBlock)))
	if err != nil {
		return err
	}

	indexNodeAPI := graphql.NewClient(indexNode, &httpClient)
	indexNodeAPIClient := mgmt.GraphService{Client: indexNodeAPI}
	poi, err := indexNodeAPIClient.GetProofOfIndexing(epochInfo.StartBlock, blockInfo.Hash().String(), indexerAddress, subgraph)
	if err != nil {
		return err
	}

	tp := table.NewWriter()
	tp.SetOutputMirror(os.Stdout)
	tp.AppendHeader(table.Row{"Start Block", "Block Hash", "Proof Of Indexing"})

	tp.AppendRows([]table.Row{
		{epochInfo.StartBlock, blockInfo.Hash().Hex(), poi},
	})

	tp.SetStyle(table.StyleLight)
	tp.Render()

	return nil
}

func comparePoi(ctx context.Context, agentHost, indexNode, ethNode, networkSubgraph string, verbose bool, count int, httpClient http.Client) error {
	mgmtAPI := graphql.NewClient(agentHost, &httpClient)
	gqlClient := mgmt.GraphService{Client: mgmtAPI}

	indexNodeAPI := graphql.NewClient(indexNode, &httpClient)
	indexNodeAPIClient := mgmt.GraphService{Client: indexNodeAPI}

	status, err := gqlClient.GetStatus()
	if err != nil {
		return err
	}
	fmt.Println("Indexer: ", status.IndexerRegistration.Address)
	subgraphAPI := graphql.NewClient(networkSubgraph, &httpClient)
	subgraphAPIClient := mgmt.GraphService{Client: subgraphAPI}
	currentEpoch, err := subgraphAPIClient.GetCurrentEpoch()
	if err != nil {
		return err
	}

	ethClient, err := ethclient.Dial(ethNode)
	if err != nil {
		return err
	}

	fmt.Println("Current epoch: ", currentEpoch.CurrentEpoch)
	allos, err := subgraphAPIClient.GetActiveAllocations(status.IndexerRegistration.Address)
	if err != nil {
		return err
	}
	epoch := currentEpoch.CurrentEpoch
	for i := 0; i < count; i++ {
		fmt.Println("Checking epoch", epoch)
		epochInfo, err := subgraphAPIClient.GetEpochInfo(epoch)
		if err != nil {
			return err
		}
		blockInfo, err := ethClient.BlockByNumber(ctx, big.NewInt(int64(epochInfo.StartBlock)))
		if err != nil {
			return err
		}
		for _, a := range allos {
			subgraphHash, e := utils.SubgraphHexToHash(a.SubgraphDeployment.ID)
			if e != nil {
				return e
			}

			closedAllos, e := subgraphAPIClient.GetClosedAllocations(a.SubgraphDeployment.ID, epoch)
			if e != nil {
				return e
			}
			totalNumberOfPoi := len(closedAllos)
			matches := 0
			for _, ca := range closedAllos {
				poi, err := indexNodeAPIClient.GetProofOfIndexing(epochInfo.StartBlock, blockInfo.Hash().String(), ca.Indexer.ID, subgraphHash)
				if err != nil {
					return err
				}

				if ca.Poi == string(poi) {
					matches++
				}
				if ca.Poi == "0x0000000000000000000000000000000000000000000000000000000000000000" {
					if !verbose {
						totalNumberOfPoi--
					}
				}

				if verbose {
					if ca.Poi != string(poi) {
						allocatedTokens, e := utils.ToDecimal(ca.AllocatedTokens, 18)
						if e != nil {
							return e
						}
						fmt.Printf("%s Indexer %s submitted: %s we got: %s subgraph: %s allocated tokens: %s\n", text.FgRed.Sprint("Mismatch found!"), ca.Indexer.ID, ca.Poi, poi, subgraphHash, allocatedTokens)

						previousEpochInfo, err := subgraphAPIClient.GetEpochInfo(epoch - 1)
						if err != nil {
							return err
						}
						previousBlockInfo, err := ethClient.BlockByNumber(ctx, big.NewInt(int64(previousEpochInfo.StartBlock)))
						if err != nil {
							return err
						}

						previousPoi, err := indexNodeAPIClient.GetProofOfIndexing(previousEpochInfo.StartBlock, previousBlockInfo.Hash().String(), status.IndexerRegistration.Address, subgraphHash)
						if err != nil {
							return err
						}

						if ca.Poi != string(previousPoi) {
							fmt.Printf("previous poi not ok %s Indexer %s submitted: %s we got: %s subgraph: %s allocated tokens: %s epoch: %s\n", text.FgRed.Sprint("Mismatch found!"), ca.Indexer.ID, ca.Poi, previousPoi, subgraphHash, allocatedTokens, previousEpochInfo.ID)
						} else {
							fmt.Printf("previous poi ok %s Indexer %s submitted: %s we got: %s subgraph: %s allocated tokens: %s epoch: %s\n", text.FgRed.Sprint("Mismatch found!"), ca.Indexer.ID, ca.Poi, previousPoi, subgraphHash, allocatedTokens, previousEpochInfo.ID)
						}
					}
				}

			}

			var match string

			switch {
			case totalNumberOfPoi == 0:
				continue
			case totalNumberOfPoi == matches:
				match = text.FgGreen.Sprintf("%d/%d", matches, totalNumberOfPoi)
			case matches == 0:
				match = text.FgRed.Sprintf("%d/%d", matches, totalNumberOfPoi)
			default:
				match = text.FgYellow.Sprintf("%d/%d", matches, totalNumberOfPoi)
			}

			fmt.Printf("Subgraph: %s\tPOI Matches: %s\n", subgraphHash, match)
		}
		epoch--
	}
	return nil
}

func getAllocation(ctx context.Context, ethNode, contractAddress, allocationID string) error {
	ethClient, err := ethclient.Dial(ethNode)
	if err != nil {
		return err
	}
	e := eth.Service{
		Client: ethClient,
	}
	allocationData, err := e.GetAllocation(contractAddress, allocationID)
	if err != nil {
		return err
	}

	allocationState, err := e.GetAllocationState(contractAddress, allocationID)
	if err != nil {
		return err
	}

	subgraphHash, err := utils.SubgraphHexToHash(common.BytesToHash(allocationData.SubgraphDeploymentID[:]).String())
	if err != nil {
		return err
	}
	allocatedTokens, err := utils.ToDecimal(allocationData.Tokens, 18)
	if err != nil {
		return err
	}

	accRewardsPerAllocatedToken, err := utils.ToDecimal(allocationData.AccRewardsPerAllocatedToken, 18)
	if err != nil {
		return err
	}

	effectiveAllocation, err := utils.ToDecimal(allocationData.EffectiveAllocation, 18)
	if err != nil {
		return err
	}

	ta := table.NewWriter()
	ta.SetOutputMirror(os.Stdout)
	ta.AppendRows([]table.Row{
		{"Allocation ID", allocationID},
		{"Indexer Address", allocationData.Indexer},
		{"Subgraph ID", subgraphHash},
		{"Allocated Tokens", allocatedTokens.String() + " GRT"},
		{"Created at Epoch", allocationData.CreatedAtEpoch.String()},
		{"Closed at Epoch", allocationData.ClosedAtEpoch.String()},
		{"AccRewards Per Allocated Tokens", accRewardsPerAllocatedToken.String() + " GRT"},
		{"Effective Allocation", effectiveAllocation.String() + " GRT"},
		{"Collected Fees", allocationData.CollectedFees},
		{"State", allocationState},
	})
	ta.SetCaption("State: 1 - Active, 2 - ? ,3 - Closed, 4 - Claimed")
	ta.SetStyle(table.StyleLight)
	ta.Style().Options.DrawBorder = true
	ta.Style().Options.SeparateRows = true
	ta.Render()
	return nil
}

func closeAllocation(ctx context.Context, ethNode, contractAddress, mnemonic, hdWalletPath, allocationID, poi string) error {
	if (len(poi) != 3) && (len(poi) != 66) {
		return errors.New("invalid POI provided")
	}

	account, privateKey, err := eth.GetWalletAccount(mnemonic, hdWalletPath)
	if err != nil {
		return err
	}

	rpcClient, err := rpc.Dial(ethNode)
	if err != nil {
		return err
	}

	ethClient := ethclient.NewClient(rpcClient)

	e := eth.Service{
		Client: ethClient,
	}

	blockNumber, err := e.Client.BlockNumber(context.Background())
	if err != nil {
		return err
	}

	balance, err := e.Client.BalanceAt(context.Background(), account.Address, big.NewInt(int64(blockNumber)))
	if err != nil {
		return err
	}

	normalizedBalance, err := utils.ToDecimal(balance, 18)
	if err != nil {
		return err
	}

	tx, err := e.CloseAllocation(privateKey, account.Address.String(), contractAddress, allocationID, poi)
	if err != nil {
		return err
	}
	estimatedFee, err := utils.ToDecimal(tx.Cost(), 18)
	if err != nil {
		return err
	}

	normalizedGasFeeCap, err := utils.ToDecimal(tx.GasFeeCap(), 18)
	if err != nil {
		return err
	}

	normalizedGasTipCap, err := utils.ToDecimal(tx.GasTipCap(), 18)
	if err != nil {
		return err
	}

	tc := table.NewWriter()
	tc.SetOutputMirror(os.Stdout)
	tc.AppendHeader(table.Row{"Transaction details"})
	tc.AppendRows([]table.Row{
		{"From Address", account.Address},
		{"ETH Balance", normalizedBalance.String() + " ETH"},
		{"To Contract", contractAddress},
		{"Nonce", tx.Nonce()},
		{"ChainID", tx.ChainId()},
		{"GasLimit", tx.Gas()},
		{"GasFeeCap", normalizedGasFeeCap.String() + " ETH"},
		{"GasTipCap", normalizedGasTipCap.String() + " ETH"},
		{"Estimated Fee", estimatedFee.String() + " ETH"},
		{"Allocation ID", allocationID},
		{"POI", poi},
		{"ETH Node", ethNode},
	})
	tc.SetStyle(table.StyleLight)
	tc.Style().Options.DrawBorder = true
	tc.Style().Options.SeparateRows = true
	tc.Render()

	fmt.Println("You are using not reference software, beware of a possible violation of the arbitration charter. ")
	fmt.Printf("Confirm transaction to send? (y)es or (n)o: ")
	if askForConfirmation() {
		err := e.Client.SendTransaction(context.Background(), tx)
		if err != nil {
			return err
		}
		fmt.Println("\nTransaction sent, txid: ", tx.Hash().String())
	}
	return nil
}
