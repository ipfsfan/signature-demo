module demo

go 1.15

require (
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-state-types v0.1.1-0.20210506134452-99b279731c48
	github.com/filecoin-project/lotus v1.10.1
)

replace (
	github.com/filecoin-project/go-address v0.0.5 => github.com/wd-002/go-address v0.0.6-0.20210712031549-df4e0bb30482
	github.com/filecoin-project/lotus => github.com/wd-002/lotus v1.10.2-0.20210716032019-8eb29d8eb8a6
)
