package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {
	var (
		rootFlagSet     = flag.NewFlagSet("graph-indexer", flag.ExitOnError)
		rulesFlagSet    = flag.NewFlagSet("graph-indexer rules", flag.ExitOnError)
		verbose         = rootFlagSet.Bool("v", false, "increase log verbosity")
		agentHost       = rootFlagSet.String("indexer-agent", "http://localhost:8000", "Indexer Agent Mgmt API Host")
		networkSubgraph = rootFlagSet.String("network-subgraph", "https://gateway.thegraph.com/network", "Network Subgraph Endpoint")
		indexNode       = rootFlagSet.String("index-node", "http://localhost:8030/graphql", "Index or query node graphql endpoint")
		ethNode         = rootFlagSet.String("eth-node", "https://cloudflare-eth.com/", "Ethereum Node address")
		networkID       = rootFlagSet.String("eth-network-id", "mainnet", "Ethereum Network ID")
		httpTimeout     = rootFlagSet.Int("timeout", 15, "HTTP timeout")
		hdWalletPath    = rootFlagSet.String("hd-wallet-path", "m/44'/60'/0'/0/0", "HD Wallet Path")
		mnemonic        = rootFlagSet.String("mnemonic", "", "Operator mnemonic phrase")
	)

	httpClient := http.Client{
		Timeout: time.Duration(*httpTimeout) * time.Second,
	}
	statusIndexing := &ffcli.Command{
		Name:       "indexing",
		ShortUsage: "graph-indexer status indexing",
		ShortHelp:  "Get subgraphs indexing statuses",
		Exec: func(ctx context.Context, args []string) error {
			return getIndexingStatuses(ctx, *indexNode, httpClient)
		},
	}

	status := &ffcli.Command{
		Name:        "status",
		ShortUsage:  "graph-indexer status",
		ShortHelp:   "Check the status of an indexer",
		Subcommands: []*ffcli.Command{statusIndexing},

		Exec: func(ctx context.Context, args []string) error {
			return status(ctx, *agentHost, *networkSubgraph, httpClient)
		},
	}

	rulesGet := &ffcli.Command{
		Name:       "get",
		ShortUsage: "graph-indexer rules get <deploymentID>",
		ShortHelp:  "Get one or more indexing rules",
		//	FlagSet:    stakeFlagSet,
		Exec: func(ctx context.Context, args []string) error {
			return getIndexingRule(ctx, *agentHost, args, httpClient)
		},
	}

	rulesSet := &ffcli.Command{
		Name:       "set",
		ShortUsage: "graph-indexer rules set",
		ShortHelp:  "Set one or more indexing rules",
		// Subcommands: []*ffcli.Command{rulesSetDecisionBasis, rulesSetAllocationAmmount},

		Exec: func(ctx context.Context, args []string) error {
			return setIndexingRule(ctx, *agentHost, args[0], args[1:], httpClient)
		},
	}
	rulesStart := &ffcli.Command{
		Name:       "start",
		ShortUsage: "graph-indexer rules start",
		ShortHelp:  "Always index a deployment (and start indexing it if necessary)",
		//	FlagSet:    stakeFlagSet,
		Exec: func(_ context.Context, args []string) error {
			return nil
		},
	}

	rulesStop := &ffcli.Command{
		Name:       "stop",
		ShortUsage: "graph-indexer rules stop",
		ShortHelp:  "Never index a deployment (and stop indexing it if necessary)",
		//	FlagSet:    stakeFlagSet,
		Exec: func(_ context.Context, args []string) error {
			return nil
		},
	}

	rulesClear := &ffcli.Command{
		Name:       "clear",
		ShortUsage: "graph-indexer rules clear",
		ShortHelp:  "Clear one or more indexing rules",
		//	FlagSet:    stakeFlagSet,
		Exec: func(_ context.Context, args []string) error {
			return nil
		},
	}

	rulesDelete := &ffcli.Command{
		Name:       "delete",
		ShortUsage: "graph-indexer rules delete",
		ShortHelp:  "Remove one or many indexing rules",
		FlagSet:    rulesFlagSet,
		//	FlagSet:    stakeFlagSet,
		Exec: func(ctx context.Context, args []string) error {
			return deleteIndexingRule(ctx, *agentHost, args, httpClient)
		},
	}

	rules := &ffcli.Command{
		Name:        "rules",
		ShortUsage:  "graph-indexer rules [<arg> ...]",
		ShortHelp:   "Configure indexing rules",
		FlagSet:     rulesFlagSet,
		Subcommands: []*ffcli.Command{rulesGet, rulesSet, rulesStop, rulesStart, rulesClear, rulesDelete},
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}
	costSetModel := &ffcli.Command{
		Name:       "model",
		ShortUsage: "graph-indexer cost set model",
		ShortHelp:  "Update a cost model",
		//	FlagSet:    stakeFlagSet,
		Exec: func(_ context.Context, args []string) error {
			return setCostModel(*agentHost, args[0], args[1:], httpClient)
		},
	}
	costSetVariables := &ffcli.Command{
		Name:       "variables",
		ShortUsage: "graph-indexer cost set variables",
		ShortHelp:  "Update cost model variables",
		//	FlagSet:    stakeFlagSet,
		Exec: func(_ context.Context, args []string) error {
			return nil
		},
	}
	costGet := &ffcli.Command{
		Name:       "get",
		ShortUsage: "graph-indexer cost get <deploymentID>",
		ShortHelp:  "Get cost models and/or variables for one or all subgraphs",
		//	FlagSet:    stakeFlagSet,
		Exec: func(_ context.Context, args []string) error {
			return getAllModelsWithVariables(*agentHost, httpClient)
		},
	}
	costSet := &ffcli.Command{
		Name:        "set",
		ShortUsage:  "graph-indexer cost set",
		ShortHelp:   "Never index a deployment (and stop indexing it if necessary)",
		Subcommands: []*ffcli.Command{costSetModel, costSetVariables},

		//	FlagSet:    stakeFlagSet,
		Exec: func(_ context.Context, args []string) error {
			return nil
		},
	}

	cost := &ffcli.Command{
		Name:        "cost",
		ShortUsage:  "graph-indexer cost [<arg> ...]",
		ShortHelp:   "Manage costing for subgraphs",
		Subcommands: []*ffcli.Command{costSet, costGet},
		Exec: func(context.Context, []string) error {
			fmt.Printf(*agentHost)
			return flag.ErrHelp
		},
	}

	signals := &ffcli.Command{
		Name:       "signals",
		ShortUsage: "graph-indexer signals",
		ShortHelp:  "Get list of subgraph deployments with signals",
		Exec: func(ctx context.Context, args []string) error {
			return signals(ctx, *networkSubgraph, httpClient)
		},
	}

	poiCompare := &ffcli.Command{
		Name:       "compare",
		ShortUsage: "graph-indexer-cli poi compare",
		Exec: func(ctx context.Context, args []string) error {
			count := 1
			var err error
			if len(args) > 0 && args[0] != "" {
				count, err = strconv.Atoi(args[0])
				if err != nil {
					return err
				}
			}
			return comparePoi(ctx, *agentHost, *indexNode, *ethNode, *networkSubgraph, *verbose, count, httpClient)
		},
	}

	poi := &ffcli.Command{
		Name:        "poi",
		ShortUsage:  "graph-indexer poi <indexer> <blockNumber>",
		ShortHelp:   "Get ProofOfIndexing ",
		Subcommands: []*ffcli.Command{poiCompare},
		Exec: func(ctx context.Context, args []string) error {
			return getPoi(ctx, *ethNode, *indexNode, *networkSubgraph, args[0], args[1], args[2], httpClient)
		},
	}

	allocationsClose := &ffcli.Command{
		Name:       "close",
		ShortUsage: "graph-indexer allocations close",
		ShortHelp:  "Close an allocation",
		Exec: func(ctx context.Context, args []string) error {
			var stakingContractAddress string
			switch *networkID {
			case "mainnet":
				stakingContractAddress = "0xF55041E37E12cD407ad00CE2910B8269B01263b9"
			case "rinkeby":
				stakingContractAddress = "0x2d44C0e097F6cD0f514edAC633d82E01280B4A5c"
			default:
				fmt.Printf("Unsupported Network ID: %s.\n", *networkID)
				os.Exit(1)
			}
			return closeAllocation(ctx, *ethNode, stakingContractAddress, *mnemonic, *hdWalletPath, args[0], args[1])
		},
	}

	allocationsGet := &ffcli.Command{
		Name:       "get",
		ShortUsage: "graph-indexer allocations close <AllocationID> <POI>",
		ShortHelp:  "Get allocation info",
		Exec: func(ctx context.Context, args []string) error {
			var stakingContractAddress string
			switch *networkID {
			case "mainnet":
				stakingContractAddress = "0xF55041E37E12cD407ad00CE2910B8269B01263b9"
			case "rinkeby":
				stakingContractAddress = "0x2d44C0e097F6cD0f514edAC633d82E01280B4A5c"
			default:
				fmt.Printf("Unsupported Network ID: %s.\n", *networkID)
				os.Exit(1)
			}
			return getAllocation(ctx, *ethNode, stakingContractAddress, args[0])
		},
	}

	allocations := &ffcli.Command{
		Name:        "allocations",
		ShortUsage:  "graph-indxer allocations [<arg> ...]",
		ShortHelp:   "Get info and manage allocations ",
		Subcommands: []*ffcli.Command{allocationsGet, allocationsClose},
		Exec: func(context.Context, []string) error {
			fmt.Printf(*agentHost)
			return flag.ErrHelp
		},
	}

	root := &ffcli.Command{
		ShortUsage:  "graph-indexer [flags] <subcommand>",
		FlagSet:     rootFlagSet,
		Subcommands: []*ffcli.Command{status, rules, cost, signals, poi, allocations},
		Options:     []ff.Option{ff.WithEnvVarPrefix("GRAPH")},
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
