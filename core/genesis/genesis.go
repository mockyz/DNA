package genesis

import (
	"errors"
	"time"
	"github.com/Ontology/common"
	"github.com/Ontology/common/config"
	"github.com/Ontology/core/types"
	"github.com/Ontology/core/utils"
	"github.com/Ontology/crypto"
	vmtypes "github.com/Ontology/vm/types"
)

const (
	BlockVersion      uint32 = 0
	GenesisNonce      uint64 = 2083236893
	DecrementInterval uint32 = 2000000

	OntRegisterAmount = 1000000000
	OngRegisterAmount = 1000000000
)

var (
	GenerationAmount = [17]uint32{80, 70, 60, 50, 40, 30, 20, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}

	OntContractCode = &vmtypes.VmCode{VmType: vmtypes.NativeVM, Code: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}}
	OngContractCode = &vmtypes.VmCode{VmType: vmtypes.NativeVM, Code: []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 2}}
	OntContractAddress = OntContractCode.AddressFromVmCode()
	OngContractAddress = OngContractCode.AddressFromVmCode()

	ONTToken   = NewGoverningToken()
	ONGToken   = NewUtilityToken()
	ONTTokenID = ONTToken.Hash()
	ONGTokenID = ONGToken.Hash()
)

var GenBlockTime = (config.DEFAULTGENBLOCKTIME * time.Second)

var GenesisBookKeepers []*crypto.PubKey

func GenesisBlockInit(defaultBookKeeper []*crypto.PubKey) (*types.Block, error) {
	//getBookKeeper
	GenesisBookKeepers = defaultBookKeeper
	nextBookKeeper, err := types.AddressFromBookKeepers(defaultBookKeeper)
	if err != nil {
		return nil, errors.New("[Block],GenesisBlockInit err with GetBookKeeperAddress")
	}
	//blockdata
	genesisHeader := &types.Header{
		Version:          BlockVersion,
		PrevBlockHash:    common.Uint256{},
		TransactionsRoot: common.Uint256{},
		Timestamp:        uint32(uint32(time.Date(2017, time.February, 23, 0, 0, 0, 0, time.UTC).Unix())),
		Height:           uint32(0),
		ConsensusData:    GenesisNonce,
		NextBookKeeper:   nextBookKeeper,

		BookKeepers: nil,
		SigData:     nil,
	}

	//block
	ont := NewGoverningToken()
	ong := NewUtilityToken()

	genesisBlock := &types.Block{
		Header: genesisHeader,
		Transactions: []*types.Transaction{
			ont,
			ong,
			NewGoverningInit(),
		},
	}
	return genesisBlock, nil
}

func NewGoverningToken() *types.Transaction {
	tx := utils.NewDeployTransaction([]byte("ONT Token"), "ONT", "1.0",
		"Ontology Team", "contact@ont.io", "Ontology Network ONT Token", vmtypes.NativeVM, true)
	return tx
}

func NewUtilityToken() *types.Transaction {
	tx := utils.NewDeployTransaction([]byte("ONT Token"), "ONG", "1.0",
		"Ontology Team", "contact@ont.io", "Ontology Network ONG Token", vmtypes.NativeVM, true)
	return tx
}

func NewGoverningInit() *types.Transaction {
	vmCode := vmtypes.VmCode{
		VmType: vmtypes.NativeVM,
		Code: []byte{21, 67, 111, 109, 109, 111, 110, 46, 84, 111, 107, 101, 110, 46, 84, 114, 97, 110, 115, 102, 101, 114},
	}
	tx := utils.NewInvokeTransaction(vmCode)
	return tx
}
