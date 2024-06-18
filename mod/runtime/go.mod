module github.com/berachain/beacon-kit/mod/runtime

go 1.22.4

replace (
	// The following are required to build with the latest version of the cosmos-sdk main branch:
	cosmossdk.io/api => cosmossdk.io/api v0.7.3-0.20240530104414-90cbb022d5f6
	cosmossdk.io/x/consensus => cosmossdk.io/x/consensus v0.0.0-20240530104414-90cbb022d5f6
	github.com/cosmos/cosmos-sdk => github.com/berachain/cosmos-sdk v0.46.0-beta2.0.20240529213909-58c32d695e1a
)

require (
	cosmossdk.io/core v0.12.1-0.20240530104414-90cbb022d5f6
	github.com/berachain/beacon-kit/mod/errors v0.0.0-20240617161612-ab1257fcf5a1
	github.com/berachain/beacon-kit/mod/log v0.0.0-20240610210054-bfdc14c4013c
	github.com/berachain/beacon-kit/mod/primitives v0.0.0-20240617161612-ab1257fcf5a1
	github.com/cometbft/cometbft v1.0.0-alpha.2.0.20240613135100-716d8f8c592d
	github.com/cometbft/cometbft/api v1.0.0-rc.1
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.3.3 // indirect
	github.com/btcsuite/btcd/chaincfg/chainhash v1.1.0 // indirect
	github.com/cockroachdb/errors v1.11.3 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/cosmos/crypto v0.0.0-20240312084433-de8f9c76030d // indirect
	github.com/cosmos/gogoproto v1.5.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	github.com/ethereum/go-ethereum v1.14.5 // indirect
	github.com/ferranbt/fastssz v0.1.4-0.20240422063434-a4db75388da1 // indirect
	github.com/getsentry/sentry-go v0.28.1 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/holiman/uint256 v1.2.5-0.20240612125212-75a520988c94 // indirect
	github.com/klauspost/cpuid/v2 v2.2.8 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/minio/sha256-simd v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/oasisprotocol/curve25519-voi v0.0.0-20230904125328-1f23a7beb09a // indirect
	github.com/petermattis/goid v0.0.0-20240503122002-4b96552b8156 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prysmaticlabs/gohashtree v0.0.4-beta // indirect
	github.com/rogpeppe/go-internal v1.12.0 // indirect
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/stretchr/objx v0.5.2 // indirect
	github.com/supranational/blst v0.3.12 // indirect
	golang.org/x/crypto v0.24.0 // indirect
	golang.org/x/exp v0.0.0-20240613232115-7f521ea00fb8 // indirect
	golang.org/x/net v0.26.0 // indirect
	golang.org/x/sync v0.7.0 // indirect
	golang.org/x/sys v0.21.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/grpc v1.64.0 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
