package contract

//go:generate go run --mod=mod github.com/dmarkham/enumer@v1.5.9 --values --type=Standard --output standard_string.go --linecomment --json --sql
type Standard int

const (
	StandardUnknown Standard = 0    // Unknown
	StandardERC20   Standard = 20   // ERC-20
	StandardERC165  Standard = 165  // ERC-165
	StandardERC721  Standard = 721  // ERC-721
	StandardERC1155 Standard = 1155 // ERC-1155
	StandardERC1967 Standard = 1967 // ERC-1967
)
