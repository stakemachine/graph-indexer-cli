package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {
	var (
		rootFlagSet  = flag.NewFlagSet("graph-indexer", flag.ExitOnError)
		rulesFlagSet = flag.NewFlagSet("graph-indexer rules", flag.ExitOnError)
		/*costFlagSet   = flag.NewFlagSet("graph-indexer cost", flag.ExitOnError)
		statusFlagSet = flag.NewFlagSet("graph-indexer status", flag.ExitOnError) */
		agentHost       = rootFlagSet.String("indexer-agent", "http://localhost:8000", "Indexer Agent Mgmt API Host")
		networkSubgraph = rootFlagSet.String("network-subgraph", "https://api.thegraph.com/subgraphs/name/graphprotocol/graph-network-testnet", "Network Subgraph Endpoint")
		indexNode       = rootFlagSet.String("index-node", "http://localhost:8030/graphql", "Index or query node graphql endpoint")
	)
	statusIndexing := &ffcli.Command{
		Name:       "indexing",
		ShortUsage: "graph-indexer status indexing",
		ShortHelp:  "Get subgraphs indexing statuses",
		Exec: func(ctx context.Context, args []string) error {
			return getIndexingStatuses(ctx, *indexNode)
		},
	}

	status := &ffcli.Command{
		Name:        "status",
		ShortUsage:  "graph-indexer status",
		ShortHelp:   "Check the status of an indexer",
		Subcommands: []*ffcli.Command{statusIndexing},

		Exec: func(ctx context.Context, args []string) error {
			return status(ctx, *agentHost, *networkSubgraph)
		},
	}

	rulesGet := &ffcli.Command{
		Name:       "get",
		ShortUsage: "graph-indexer rules get <deploymentID>",
		ShortHelp:  "Get one or more indexing rules",
		//	FlagSet:    stakeFlagSet,
		Exec: func(ctx context.Context, args []string) error {
			return getRule(ctx, *agentHost, args)
		},
	}

	rulesSet := &ffcli.Command{
		Name:       "set",
		ShortUsage: "graph-indexer rules set",
		ShortHelp:  "Set one or more indexing rules",
		// Subcommands: []*ffcli.Command{rulesSetDecisionBasis, rulesSetAllocationAmmount},

		Exec: func(ctx context.Context, args []string) error {
			return setRule(ctx, *agentHost, args[0], args[1:])
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
			return deleteRule(ctx, *agentHost, args)
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
			return setCostModel(*agentHost, args[0], args[1:])
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
			return getAllModelsWithVariables(*agentHost)
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
			return signals(ctx, *networkSubgraph)
		},
	}

	root := &ffcli.Command{
		ShortUsage:  "graph-indexer [flags] <subcommand>",
		FlagSet:     rootFlagSet,
		Subcommands: []*ffcli.Command{status, rules, cost, signals},
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
