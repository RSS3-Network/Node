package erc165

// https://eips.ethereum.org/EIPS/eip-165
//go:generate go run --mod=mod github.com/ethereum/go-ethereum/cmd/abigen@v1.13.5 --abi ./abi/ERC165.abi --pkg erc165 --type ERC165 --out contract_erc165.go
