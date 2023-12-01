// Code generated by "enumer --values --type=ReactionType --output type_reaction.go --transform=snake-upper"; DO NOT EDIT.

package farcaster

import (
	"fmt"
	"strings"
)

const _ReactionTypeName = "REACTION_TYPE_NONEREACTION_TYPE_LIKEREACTION_TYPE_RECAST"

var _ReactionTypeIndex = [...]uint8{0, 18, 36, 56}

const _ReactionTypeLowerName = "reaction_type_nonereaction_type_likereaction_type_recast"

func (i ReactionType) String() string {
	if i < 0 || i >= ReactionType(len(_ReactionTypeIndex)-1) {
		return fmt.Sprintf("ReactionType(%d)", i)
	}
	return _ReactionTypeName[_ReactionTypeIndex[i]:_ReactionTypeIndex[i+1]]
}

func (ReactionType) Values() []string {
	return ReactionTypeStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ReactionTypeNoOp() {
	var x [1]struct{}
	_ = x[ReactionTypeNone-(0)]
	_ = x[ReactionTypeLike-(1)]
	_ = x[ReactionTypeRecast-(2)]
}

var _ReactionTypeValues = []ReactionType{ReactionTypeNone, ReactionTypeLike, ReactionTypeRecast}

var _ReactionTypeNameToValueMap = map[string]ReactionType{
	_ReactionTypeName[0:18]:       ReactionTypeNone,
	_ReactionTypeLowerName[0:18]:  ReactionTypeNone,
	_ReactionTypeName[18:36]:      ReactionTypeLike,
	_ReactionTypeLowerName[18:36]: ReactionTypeLike,
	_ReactionTypeName[36:56]:      ReactionTypeRecast,
	_ReactionTypeLowerName[36:56]: ReactionTypeRecast,
}

var _ReactionTypeNames = []string{
	_ReactionTypeName[0:18],
	_ReactionTypeName[18:36],
	_ReactionTypeName[36:56],
}

// ReactionTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ReactionTypeString(s string) (ReactionType, error) {
	if val, ok := _ReactionTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ReactionTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ReactionType values", s)
}

// ReactionTypeValues returns all values of the enum
func ReactionTypeValues() []ReactionType {
	return _ReactionTypeValues
}

// ReactionTypeStrings returns a slice of all String values of the enum
func ReactionTypeStrings() []string {
	strs := make([]string, len(_ReactionTypeNames))
	copy(strs, _ReactionTypeNames)
	return strs
}

// IsAReactionType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ReactionType) IsAReactionType() bool {
	for _, v := range _ReactionTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
