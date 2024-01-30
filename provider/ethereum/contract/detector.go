package contract

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/rss3-network/node/provider/ethereum"
	"github.com/rss3-network/node/provider/ethereum/contract/erc165"
	"github.com/rss3-network/node/provider/ethereum/proxy"
	"github.com/rss3-network/protocol-go/schema/metadata"
)

var (
	// InterfaceIDERC721 method ID from https://eips.ethereum.org/EIPS/eip-721#:~:text=0x80ac58cd.
	InterfaceIDERC721 = [4]byte(hexutil.MustDecode("0x80ac58cd"))
	// InterfaceIDERC1155 method ID from https://eips.ethereum.org/EIPS/eip-1155#:~:text=0xd9b67a26.
	InterfaceIDERC1155 = [4]byte(hexutil.MustDecode("0xd9b67a26"))
)

var exceptionalContracts = map[uint64]map[common.Address]metadata.Standard{}

// DetectTokenStandard detects the token standard of the contract by bytecode.
// It supports the ERC-1967 proxy contract.
func DetectTokenStandard(ctx context.Context, chainID uint64, address common.Address, blockNumber *big.Int, ethereumClient ethereum.Client) (metadata.Standard, error) {
	if addressMap, exists := exceptionalContracts[chainID]; exists {
		if standard, exists := addressMap[address]; exists {
			return standard, nil
		}
	}

	code, err := ethereumClient.CodeAt(ctx, address, blockNumber)
	if err != nil {
		return metadata.StandardUnknown, fmt.Errorf("get code: %w", err)
	}

	// If the contract is ERC-1967, get the implementation contract first
	if DetectERC1967WithCode(chainID, address, code) {
		implementation, err := proxy.GetImplementation(ctx, address, blockNumber, ethereumClient)
		if err != nil {
			return metadata.StandardUnknown, fmt.Errorf("get %s implementation: %w", address, err)
		}

		if code, err = ethereumClient.CodeAt(ctx, *implementation, blockNumber); err != nil {
			return metadata.StandardUnknown, fmt.Errorf("get %s code: %w", implementation, err)
		}
	}

	// It can cover most ERC-20 tokens
	if DetectERC20WithCode(chainID, address, code) {
		return metadata.StandardERC20, nil
	}

	// The result may be influenced by compiler optimization
	if DetectERC721WithCode(chainID, address, code) {
		return metadata.StandardERC721, nil
	}

	// The result may be influenced by compiler optimization
	if DetectERC1155WithCode(chainID, address, code) {
		return metadata.StandardERC1155, nil
	}

	// Precise determination of contract criteria by Interface ID, but will require RPC requests
	if DetectERC165WithCode(chainID, address, code) {
		return DetectNFTStandard(ctx, chainID, address, blockNumber, ethereumClient)
	}

	// If ERC-165 is not implemented, it is not a standard NFT contract
	return metadata.StandardUnknown, nil
}

// DetectNFTStandard detects the NFT standard of the contract by ERC-165.
func DetectNFTStandard(ctx context.Context, _ uint64, address common.Address, blockNumber *big.Int, ethereumClient ethereum.Client) (metadata.Standard, error) {
	caller, err := erc165.NewERC165Caller(address, ethereumClient)
	if err != nil {
		return metadata.StandardUnknown, fmt.Errorf("initialize ERC-165 caller: %w", err)
	}

	callOptions := bind.CallOpts{
		BlockNumber: blockNumber,
		Context:     ctx,
	}

	isERC721, err := caller.SupportsInterface(&callOptions, InterfaceIDERC721)
	if err != nil {
		return metadata.StandardUnknown, fmt.Errorf("detect ERC-721 interface: %w", err)
	}

	if isERC721 {
		return metadata.StandardERC721, nil
	}

	isERC1155, err := caller.SupportsInterface(&callOptions, InterfaceIDERC1155)
	if err != nil {
		return metadata.StandardUnknown, fmt.Errorf("detect ERC-1155 interface: %w", err)
	}

	if isERC1155 {
		return metadata.StandardERC1155, nil
	}

	return metadata.StandardUnknown, nil
}

// DetectERC20WithCode detects if bytecode of the contract is ERC-20.
func DetectERC20WithCode(_ uint64, address common.Address, code []byte) bool {
	exceptionalAddresses := map[common.Address]bool{}

	return exceptionalAddresses[address] || ContainsMethodIDs(code,
		MethodID("name()"),
		MethodID("symbol()"),
		MethodID("decimals()"),
		MethodID("totalSupply()"),
		MethodID("decimals()"),
		MethodID("balanceOf(address)"),
		MethodID("transfer(address,uint256)"),
		MethodID("transferFrom(address,address,uint256)"),
		MethodID("approve(address,uint256)"),
		MethodID("allowance(address,address)"),
		MethodID("Transfer(address,address,uint256)"),
		MethodID("Approval(address,address,uint256)"),
	)
}

// DetectERC165WithCode detects if bytecode of the contract is ERC-165.
func DetectERC165WithCode(_ uint64, _ common.Address, code []byte) bool {
	return ContainsMethodIDs(code, MethodID("supportsInterface(bytes4)"))
}

// DetectERC721WithCode detects if bytecode of the contract is ERC-721.
func DetectERC721WithCode(_ uint64, _ common.Address, code []byte) bool {
	return ContainsMethodIDs(code, // 0x80ac58cd
		MethodID("balanceOf(address)"),
		MethodID("ownerOf(uint256)"),
		MethodID("approve(address,uint256)"),
		MethodID("getApproved(uint256)"),
		MethodID("setApprovalForAll(address,bool)"),
		MethodID("isApprovedForAll(address,address)"),
		MethodID("transferFrom(address,address,uint256)"),
		MethodID("safeTransferFrom(address,address,uint256)"),
		MethodID("safeTransferFrom(address,address,uint256,bytes)"),
	)
}

// DetectERC1155WithCode detects if bytecode of the contract is ERC-1155.
func DetectERC1155WithCode(_ uint64, _ common.Address, code []byte) bool {
	return ContainsMethodIDs(code, // 0xd9b67a26
		MethodID("balanceOf(address,uint256)"),
		MethodID("balanceOfBatch(address[],uint256[])"),
		MethodID("setApprovalForAll(address,bool)"),
		MethodID("isApprovedForAll(address,address)"),
		MethodID("safeTransferFrom(address,address,uint256,uint256,bytes)"),
		MethodID("safeBatchTransferFrom(address,address,uint256[],uint256[],bytes)"),
	)
}

// DetectERC1967WithCode detects if bytecode of the contract is ERC-1967.
func DetectERC1967WithCode(_ uint64, _ common.Address, code []byte) bool {
	return ContainsMethodIDs(code, MethodID("implementation()"))
}
