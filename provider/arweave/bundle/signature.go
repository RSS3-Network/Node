package bundle

const (
	SignatureTypeArweave       = 1
	SignatureED25519           = 2
	SignatureEthereum          = 3
	SignatureSolana            = 4
	SignatureInjectedAptos     = 5
	SignatureMultiAptos        = 6
	SignatureTypeTypedEthereum = 7
)

type SignatureType struct {
	SignatureLength uint16
	PublicKeyLength uint16
}

var signatureTypeMap = map[uint16]SignatureType{
	SignatureTypeArweave: {
		SignatureLength: 512,
		PublicKeyLength: 512,
	},
	SignatureED25519: {
		SignatureLength: 64,
		PublicKeyLength: 32,
	},
	SignatureEthereum: {
		SignatureLength: 64 + 1,
		PublicKeyLength: 65,
	},
	SignatureSolana: {
		SignatureLength: 64,
		PublicKeyLength: 32,
	},
	SignatureInjectedAptos: {
		SignatureLength: 64,
		PublicKeyLength: 32,
	},
	SignatureMultiAptos: {
		SignatureLength: 64*32 + 4,
		PublicKeyLength: 32*32 + 1,
	},
	SignatureTypeTypedEthereum: {
		SignatureLength: 64 + 1,
		PublicKeyLength: 42,
	},
}
