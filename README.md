# The Graph Indexer CLI
[![golangci-lint](https://github.com/stakemachine/graph-indexer-cli/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/stakemachine/graph-indexer-cli/actions/workflows/golangci-lint.yml)
[![CodeQL](https://github.com/stakemachine/graph-indexer-cli/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/stakemachine/graph-indexer-cli/actions/workflows/codeql-analysis.yml)  
This repository is WIP rework of [indexer-cli](https://github.com/graphprotocol/indexer/tree/main/packages/indexer-cli).  
Not all commands are implemented. Alpha software, use at your own risk.

## Install
```
go get -u github.com/stakemachine/graph-indexer-cli
```
## How to use
First of all you need to set ENV variable for your indexer-agent.
```
export GRAPH_INDEXER_AGENT="http://<indexer-agent>:port"
```

### Status
```
graph-indexer-cli status
```
By default we use testnet subgraph, to see active allocations in mainnet:
```
export GRAPH_NETWORK_SUBGRAPH="https://api.thegraph.com/subgraphs/name/graphprotocol/graph-network-mainnet"
```

### Rules
#### Set
```
graph-indexer-cli rules set QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF allocationAmount 100000 decisionBasis always
```
#### Get
```
graph-indexer-cli rules get QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF
```
#### Delete
```
graph-indexer-cli rules delete QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF
```

### Cost Models
#### Set
```
graph-indexer-cli cost set model QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF costModel 'default => 0.1;'
```
also you can set variables for model
```
graph-indexer-cli cost set model QmRhYzT8HEZ9LziQhP6JfNfd4co9A7muUYQhPMJsMUojSF variables '{"DAI":"0.5"}'
```
#### Get
```
graph-indexer-cli cost get all # or Deployment ID
```