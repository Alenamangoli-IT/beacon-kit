package genesis

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

var _ Genesis = (*Geth)(nil)

// Geth is a struct that holds the genesis configuration for Geth.
type Geth struct {
	*core.Genesis
}

func NewGeth() *Geth {
	return &Geth{
		&core.Genesis{
			Config:    defaultGethChainConfig(),
			Nonce:     0x0000000000000000,
			Timestamp: 0x0,
			ExtraData: common.FromHex(
				// nolint:lll // default genesis extra data
				"0x0000000000000000000000000000000000000000000000000000000000000000658bdf435d810c91414ec09147daa6db624063790000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			),
			GasLimit:   0x1c9c380,
			Difficulty: big.NewInt(0x0),
			Alloc:      make(types.GenesisAlloc),
			Coinbase:   common.Address{},
		},
	}
}

func (g *Geth) AddAccount(address string, balance *big.Int) error {
	addr := common.HexToAddress(address)
	if _, ok := g.Alloc[addr]; ok {
		return errAccountAlreadyExists
	}
	g.Alloc[addr] = types.Account{
		Balance: balance,
	}
	return nil
}

func (g *Geth) AddPredeploy(address string, code []byte, balance *big.Int, nonce uint64) error {
	addr := common.HexToAddress(address)
	if _, ok := g.Alloc[addr]; ok {
		return errPredeployAlreadyExists
	}
	g.Alloc[addr] = types.Account{
		Code:    code,
		Balance: balance,
		Nonce:   nonce,
	}
	return nil
}

func (g *Geth) WriteJSON(filename string) error {
	return writeGenesisToJSON(g, filename)
}

func defaultGethChainConfig() *params.ChainConfig {
	return &params.ChainConfig{
		ChainID:                       big.NewInt(80086), // 80086 is the chain ID for Berachain
		HomesteadBlock:                big.NewInt(0),
		DAOForkBlock:                  big.NewInt(0),
		DAOForkSupport:                true,
		EIP150Block:                   big.NewInt(0),
		EIP155Block:                   big.NewInt(0),
		EIP158Block:                   big.NewInt(0),
		ByzantiumBlock:                big.NewInt(0),
		ConstantinopleBlock:           big.NewInt(0),
		PetersburgBlock:               big.NewInt(0),
		IstanbulBlock:                 big.NewInt(0),
		MuirGlacierBlock:              big.NewInt(0),
		BerlinBlock:                   big.NewInt(0),
		LondonBlock:                   big.NewInt(0),
		ArrowGlacierBlock:             big.NewInt(0),
		GrayGlacierBlock:              big.NewInt(0),
		MergeNetsplitBlock:            big.NewInt(0),
		ShanghaiTime:                  new(uint64),
		CancunTime:                    new(uint64),
		TerminalTotalDifficulty:       big.NewInt(0),
		TerminalTotalDifficultyPassed: true,
	}
}