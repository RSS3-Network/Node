package ens

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

// getEnsName get ens name by name hash (node)
func (w *worker) getEnsName(ctx context.Context, _ *big.Int, namehash common.Hash) (string, error) {
	// Load name from dataset
	namehashRecord, err := w.databaseClient.LoadDatasetENSNamehash(ctx, namehash)
	if err != nil {
		zap.L().Error("Fail to find ens namehash from dataset", zap.String("namehash", namehash.Hex()), zap.Error(err))

		return "", fmt.Errorf("fail to find ens namehash %v, all methods failed", namehash.String())
	}

	return namehashRecord.Name, nil
}

// decodeName decode ens name
func (w *worker) decodeName(buf []byte) ([]string, error) {
	var (
		offset     = 0
		list       []string
		firstLabel string
	)

	// check first label length
	length := int(buf[offset])
	if length == 0 {
		return []string{firstLabel, "."}, nil
	}

	// decode every label
	for length > 0 {
		// extract label from buf
		label := string(buf[offset+1 : offset+1+length])

		// check label
		if !w.checkValidLabel(label) {
			return nil, fmt.Errorf("invalid label")
		}

		// Put dot between the gap if not the first label
		if offset > 1 {
			list = append(list, ".")
		} else {
			firstLabel = label
		}

		// label in list
		list = append(list, label)

		// update offset and len
		offset += length + 1
		length = int(buf[offset])
	}

	// return first label and complete domain string
	return []string{firstLabel, strings.Join(list, "")}, nil
}

// checkValidLabel: check label
func (w *worker) checkValidLabel(name string) bool {
	for i := 0; i < len(name); i++ {
		c := name[i]
		if c == 0 {
			zap.L().Error("Invalid label contained null byte. Skipping.", zap.String("name", name))
			return false
		} else if c == 46 {
			zap.L().Error("Invalid label contained separator char '.'. Skipping.", zap.String("name", name))
			return false
		}
	}

	return true
}

// CompressPubkey encodes a public key to 33-byte compressed format.
func CompressPubkey(x, y *big.Int) []byte {
	fx := new(secp256k1.FieldVal)
	fy := new(secp256k1.FieldVal)

	fx.SetBytes((*[32]byte)(x.Bytes()))
	fy.SetBytes((*[32]byte)(y.Bytes()))

	return secp256k1.NewPublicKey(fx, fy).SerializeCompressed()
}
