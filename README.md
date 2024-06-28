# reorg-fix_comparison
This repository used to store testcase and experiment result for our reorg-fix protocol.

## dictionary
```
.
├── LICENSE
├── README.md
├── case			// store docker-compose.yml for every testcase.
│   ├── attack-tpstest
│   ├── blockcost
│   └── tpstest
├── config			// common config files for all service in system.
├── experiment          	// our own experiment results.
│   ├── attack-tpstest
│   ├── blockcost
│   └── tpstest
├── source			// source code for all service.
│   ├── attacker-service
│   ├── go-ethereum
│   ├── prysm
│   ├── prysm-reorg-fix
│   ├── strategy-gen
│   └── txpress
└── txpress.sh			// script to run a transaction press test.
```


## cases
There are three testcases in the repository.

- GenerateBlockTimeCost Comparison: Compare time cost for generate a block between origin and reorg-fixed version.
- Normal TPS Comparison: Compare tps between origin and reorg-fixed version.
- Attack TPS Comparison: Compare tps between origin and reorg-fixed version while implementing malicious attacks. 


## how to run testcase
