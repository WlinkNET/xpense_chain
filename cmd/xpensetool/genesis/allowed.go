package genesis

import (
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/WlinkNET/xpense_chain/opera/genesis"
	"github.com/WlinkNET/xpense_chain/opera/genesisstore"
	"github.com/ethereum/go-ethereum/common"
)

type GenesisTemplate struct {
	Name   string
	Header genesis.Header
	Hashes genesis.Hashes
}

var (
	allowedGenesisSigners = []common.Address{
		common.HexToAddress("0xAe529c0F574a987a0c2C702FF684385572d0F46d"), //Allowed genesis signer...
	}

	mainnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash("0xa82e9aa1bd9e064299cfe56b20b5d3c4b87f7a5b95ce9305c75fe40d6d37913b"),
		NetworkID:   41245,
		NetworkName: "xpense",
	}

	/*testnetHeader = genesis.Header{
		GenesisID:   hash.HexToHash(""), //GenesisID first Epoch....
		NetworkID:   41246,							//retrieve data from testnet-genesis.json
		NetworkName: "xpense-testnet",
	}*/

	allowedGenesis = []GenesisTemplate{

		{
			Name:   "Xpense-Genesis-Export",
			Header: mainnetHeader,
			Hashes: genesis.Hashes{
				genesisstore.EpochsSection(0):     hash.HexToHash("0xa82e9aa1bd9e064299cfe56b20b5d3c4b87f7a5b95ce9305c75fe40d6d37913b"),
				genesisstore.BlocksSection(0):     hash.HexToHash("0xbcd876993a08c113addb708277497a355f4ea03561bd6d302a8479d4db19622e"),
				genesisstore.FwsLiveSection(0):    hash.HexToHash("0x06a05e3f0e6d68dcd51296dbe3f0c65f1fa4f0f4af10b30de6881178a83bbe92"),
				genesisstore.FwsArchiveSection(0): hash.HexToHash("0xdb1681d7d3e8c933a6b4d132fefaa9f7db49b7f089177778e62a59a5cae6c43f"),
			},
		},
		//TODO Add for testnet data
	}
)
