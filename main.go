package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/peterbourgon/ff/v3"
	"github.com/peterbourgon/ff/v3/ffcli"
	"github.com/shopspring/decimal"
)

var errNotImplemented = errors.New("not implemented")

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
		httpTimeout     = rootFlagSet.Int("timeout", 15, "HTTP timeout in seconds")
		hdWalletPath    = rootFlagSet.String("hd-wallet-path", "m/44'/60'/0'/0/0", "HD Wallet Path")
		mnemonic        = rootFlagSet.String("mnemonic", "", "Operator mnemonic phrase")
	)

	// newHTTPClient constructs a client using the parsed timeout flag value.
	// Must be called inside Exec closures (after flag parsing), not at setup time.
	newHTTPClient := func() http.Client {
		return http.Client{
			Timeout: time.Duration(*httpTimeout) * time.Second,
		}
	}

	stakingContractAddress := func() (string, error) {
		switch *networkID {
		case "mainnet":
			return "0xF55041E37E12cD407ad00CE2910B8269B01263b9", nil
		default:
			return "", fmt.Errorf("unsupported network ID: %s", *networkID)
		}
	}

	statusIndexing := &ffcli.Command{
		Name:       "indexing",
		ShortUsage: "graph-indexer status indexing",
		ShortHelp:  "Get subgraphs indexing statuses",
		Exec: func(ctx context.Context, _ []string) error {
			return getIndexingStatuses(ctx, *indexNode, newHTTPClient())
		},
	}

	status := &ffcli.Command{
		Name:        "status",
		ShortUsage:  "graph-indexer status",
		ShortHelp:   "Check the status of an indexer",
		Subcommands: []*ffcli.Command{statusIndexing},

		Exec: func(ctx context.Context, _ []string) error {
			return status(ctx, *agentHost, *networkSubgraph, *networkID, newHTTPClient())
		},
	}

	rulesGet := &ffcli.Command{
		Name:       "get",
		ShortUsage: "graph-indexer rules get <deploymentID>",
		ShortHelp:  "Get one or more indexing rules",
		Exec: func(ctx context.Context, args []string) error {
			return getIndexingRule(ctx, *agentHost, args, newHTTPClient())
		},
	}

	rulesSet := &ffcli.Command{
		Name:       "set",
		ShortUsage: "graph-indexer rules set",
		ShortHelp:  "Set one or more indexing rules",
		Exec: func(ctx context.Context, args []string) error {
			return setIndexingRule(ctx, *agentHost, args[0], args[1:], newHTTPClient())
		},
	}
	rulesStart := &ffcli.Command{
		Name:       "start",
		ShortUsage: "graph-indexer rules start",
		ShortHelp:  "Always index a deployment (and start indexing it if necessary)",
		Exec: func(_ context.Context, _ []string) error {
			return errNotImplemented
		},
	}

	rulesStop := &ffcli.Command{
		Name:       "stop",
		ShortUsage: "graph-indexer rules stop",
		ShortHelp:  "Never index a deployment (and stop indexing it if necessary)",
		Exec: func(_ context.Context, _ []string) error {
			return errNotImplemented
		},
	}

	rulesClear := &ffcli.Command{
		Name:       "clear",
		ShortUsage: "graph-indexer rules clear",
		ShortHelp:  "Clear one or more indexing rules",
		Exec: func(_ context.Context, _ []string) error {
			return errNotImplemented
		},
	}

	rulesDelete := &ffcli.Command{
		Name:       "delete",
		ShortUsage: "graph-indexer rules delete",
		ShortHelp:  "Remove one or many indexing rules",
		FlagSet:    rulesFlagSet,
		Exec: func(ctx context.Context, args []string) error {
			return deleteIndexingRule(ctx, *agentHost, args, newHTTPClient())
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
		Exec: func(_ context.Context, args []string) error {
			return setCostModel(*agentHost, args[0], args[1:], newHTTPClient())
		},
	}
	costSetVariables := &ffcli.Command{
		Name:       "variables",
		ShortUsage: "graph-indexer cost set variables",
		ShortHelp:  "Update cost model variables",
		Exec: func(_ context.Context, _ []string) error {
			return errNotImplemented
		},
	}
	costGet := &ffcli.Command{
		Name:       "get",
		ShortUsage: "graph-indexer cost get <deploymentID>",
		ShortHelp:  "Get cost models and/or variables for one or all subgraphs",
		Exec: func(_ context.Context, _ []string) error {
			return getAllModelsWithVariables(*agentHost, newHTTPClient())
		},
	}
	costSet := &ffcli.Command{
		Name:        "set",
		ShortUsage:  "graph-indexer cost set",
		ShortHelp:   "Never index a deployment (and stop indexing it if necessary)",
		Subcommands: []*ffcli.Command{costSetModel, costSetVariables},
		Exec: func(_ context.Context, _ []string) error {
			return flag.ErrHelp
		},
	}

	cost := &ffcli.Command{
		Name:        "cost",
		ShortUsage:  "graph-indexer cost [<arg> ...]",
		ShortHelp:   "Manage costing for subgraphs",
		Subcommands: []*ffcli.Command{costSet, costGet},
		Exec: func(context.Context, []string) error {
			return flag.ErrHelp
		},
	}

	signals := &ffcli.Command{
		Name:       "signals",
		ShortUsage: "graph-indexer signals",
		ShortHelp:  "Get list of subgraph deployments with signals",
		Exec: func(ctx context.Context, _ []string) error {
			return signals(ctx, *networkSubgraph, *indexNode, newHTTPClient())
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
			return comparePoi(ctx, *agentHost, *indexNode, *ethNode, *networkSubgraph, *verbose, count, newHTTPClient())
		},
	}

	poi := &ffcli.Command{
		Name:        "poi",
		ShortUsage:  "graph-indexer poi <indexer> <blockNumber>",
		ShortHelp:   "Get ProofOfIndexing ",
		Subcommands: []*ffcli.Command{poiCompare},
		Exec: func(ctx context.Context, args []string) error {
			return getPoi(ctx, *ethNode, *indexNode, *networkSubgraph, args[0], args[1], args[2], newHTTPClient())
		},
	}
	allocationsAdvice := &ffcli.Command{
		Name:       "advice",
		ShortUsage: "graph-indexer allocations advice <stakeAmount>",
		ShortHelp:  "Allocations advice with specified stake amount",
		Exec: func(ctx context.Context, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("stake amount is required")
			}
			stakeAmount, err := decimal.NewFromString(args[0])
			if err != nil {
				return fmt.Errorf("invalid stake amount: %w", err)
			}
			return allocationsAdvice(ctx, *networkSubgraph, *indexNode, stakeAmount, newHTTPClient())
		},
	}
	allocationsClose := &ffcli.Command{
		Name:       "close",
		ShortUsage: "graph-indexer allocations close",
		ShortHelp:  "Close an allocation",
		Exec: func(ctx context.Context, args []string) error {
			addr, err := stakingContractAddress()
			if err != nil {
				return err
			}
			return closeAllocation(ctx, *ethNode, addr, *mnemonic, *hdWalletPath, args[0], args[1])
		},
	}

	allocationsGet := &ffcli.Command{
		Name:       "get",
		ShortUsage: "graph-indexer allocations get <AllocationID>",
		ShortHelp:  "Get allocation info",
		Exec: func(ctx context.Context, args []string) error {
			addr, err := stakingContractAddress()
			if err != nil {
				return err
			}
			return getAllocation(ctx, *ethNode, addr, args[0])
		},
	}

	allocations := &ffcli.Command{
		Name:        "allocations",
		ShortUsage:  "graph-indexer allocations [<arg> ...]",
		ShortHelp:   "Get info and manage allocations",
		Subcommands: []*ffcli.Command{allocationsGet, allocationsClose, allocationsAdvice},
		Exec: func(context.Context, []string) error {
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
