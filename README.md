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

- I. GenerateBlockTimeCost Comparison: Compare time cost for generate a block between origin and reorg-fixed version.
- II. Normal TPS Comparison: Compare tps between origin and reorg-fixed version.
- III. Attack TPS Comparison: Compare tps between origin and reorg-fixed version while implementing malicious attacks. 


## how to run testcase

#### 0. environment dependent
- linux os (ubuntu 22.04 is best)
- docker

#### 1. clone the repository.
```
$ git clone https://github.com/tsinghua-cel/reorg-fix_comparison
```

#### 2. run testcase.
Run testcase with `./runtest.sh casenum` to run the testcase, valid `casenum` is `1,2,3`.
```shell
./runtest.sh 1
./runtest.sh 2
./runtest.sh 3
```
