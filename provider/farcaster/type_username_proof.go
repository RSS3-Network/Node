// Code generated by "enumer --values --type=UsernameProofType --output type_username_proof.go --transform=snake-upper"; DO NOT EDIT.

package farcaster

import (
	"fmt"
	"strings"
)

const _UsernameProofTypeName = "USERNAME_TYPE_NONEUSERNAME_TYPE_FNAMEUSERNAME_TYPE_ENS_L1"

var _UsernameProofTypeIndex = [...]uint8{0, 18, 37, 57}

const _UsernameProofTypeLowerName = "username_type_noneusername_type_fnameusername_type_ens_l1"

func (i UsernameProofType) String() string {
	if i < 0 || i >= UsernameProofType(len(_UsernameProofTypeIndex)-1) {
		return fmt.Sprintf("UsernameProofType(%d)", i)
	}
	return _UsernameProofTypeName[_UsernameProofTypeIndex[i]:_UsernameProofTypeIndex[i+1]]
}

func (UsernameProofType) Values() []string {
	return UsernameProofTypeStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _UsernameProofTypeNoOp() {
	var x [1]struct{}
	_ = x[UsernameTypeNone-(0)]
	_ = x[UsernameTypeFname-(1)]
	_ = x[UsernameTypeEnsL1-(2)]
}

var _UsernameProofTypeValues = []UsernameProofType{UsernameTypeNone, UsernameTypeFname, UsernameTypeEnsL1}

var _UsernameProofTypeNameToValueMap = map[string]UsernameProofType{
	_UsernameProofTypeName[0:18]:       UsernameTypeNone,
	_UsernameProofTypeLowerName[0:18]:  UsernameTypeNone,
	_UsernameProofTypeName[18:37]:      UsernameTypeFname,
	_UsernameProofTypeLowerName[18:37]: UsernameTypeFname,
	_UsernameProofTypeName[37:57]:      UsernameTypeEnsL1,
	_UsernameProofTypeLowerName[37:57]: UsernameTypeEnsL1,
}

var _UsernameProofTypeNames = []string{
	_UsernameProofTypeName[0:18],
	_UsernameProofTypeName[18:37],
	_UsernameProofTypeName[37:57],
}

// UsernameProofTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func UsernameProofTypeString(s string) (UsernameProofType, error) {
	if val, ok := _UsernameProofTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _UsernameProofTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to UsernameProofType values", s)
}

// UsernameProofTypeValues returns all values of the enum
func UsernameProofTypeValues() []UsernameProofType {
	return _UsernameProofTypeValues
}

// UsernameProofTypeStrings returns a slice of all String values of the enum
func UsernameProofTypeStrings() []string {
	strs := make([]string, len(_UsernameProofTypeNames))
	copy(strs, _UsernameProofTypeNames)
	return strs
}

// IsAUsernameProofType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i UsernameProofType) IsAUsernameProofType() bool {
	for _, v := range _UsernameProofTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
