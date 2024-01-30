package crossbell

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rss3-network/node/provider/ethereum/contract"
)

//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./profile/Profile.abi --pkg profile --type Profile --out ./profile/profile.go
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./character/Character.abi --pkg character --type Character --out ./character/character.go
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./event/Event.abi --pkg event --type Event --out ./event/event.go
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./linklist/Linklist.abi --pkg linklist --type LinkList --out ./linklist/linklist.go
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./periphery/Periphery.abi --pkg periphery --type Periphery --out ./periphery/periphery.go
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./tips/Tips.abi --pkg tips --type Tips --out ./tips/tips.go

var (
	AddressWeb3Entry     = common.HexToAddress("0xa6f969045641Cf486a747A2688F3a5A6d43cd0D8")
	AddressLinkList      = common.HexToAddress("0xFc8C75bD5c26F50798758f387B698f207a016b6A")
	AddressPeriphery     = common.HexToAddress("0x96e96b7AF62D628cE7eb2016D2c1D2786614eA73")
	AddressXSyncOperator = common.HexToAddress("0x0f588318a494e4508a121a32b6670b5494ca3357")
	AddressTips          = common.HexToAddress("0x0058be0845952D887D1668B5545de995E12e8783")

	// profile events
	EventHashProfileCreated = contract.EventHash("ProfileCreated(uint256,address,address,string,uint256)")
	EventHashSetProfileURI  = contract.EventHash("SetProfileUri(uint256,string)")

	// events
	EventHashCharacterCreated         = contract.EventHash("CharacterCreated(uint256,address,address,string,uint256)")
	EventHashSetHandle                = contract.EventHash("SetHandle(address,uint256,string)")
	EventHashSetCharacterURI          = contract.EventHash("SetCharacterUri(uint256,string)")
	EventHashPostNote                 = contract.EventHash("PostNote(uint256,uint256,bytes32,bytes32,bytes)")
	EventHashSetNoteURI               = contract.EventHash("SetNoteUri(uint256,uint256,string)")
	EventHashDeleteNote               = contract.EventHash("DeleteNote(uint256,uint256)")
	EventHashMintNote                 = contract.EventHash("MintNote(address,uint256,uint256,address,uint256)")
	EventHashSetOperator              = contract.EventHash("SetOperator(uint256,address,uint256)")
	EventHashAddOperator              = contract.EventHash("AddOperator(uint256,address,uint256)")
	EventHashRemoveOperator           = contract.EventHash("RemoveOperator(uint256,address,uint256)")
	EventHashGrantOperatorPermissions = contract.EventHash("GrantOperatorPermissions(uint256,address,uint256)")
	EventHashTipCharacterForNote      = contract.EventHash("TipCharacterForNote(uint256,uint256,uint256,address,uint256)")
)
