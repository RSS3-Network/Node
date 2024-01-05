// Code generated by "enumer --values --type=EthereumChainID --linecomment --output ethereum_chain_string.go --json --yaml --sql"; DO NOT EDIT.

package filter

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const (
	_EthereumChainIDName_0      = "ethereum"
	_EthereumChainIDLowerName_0 = "ethereum"
	_EthereumChainIDName_1      = "optimism"
	_EthereumChainIDLowerName_1 = "optimism"
	_EthereumChainIDName_2      = "polygon"
	_EthereumChainIDLowerName_2 = "polygon"
)

var (
	_EthereumChainIDIndex_0 = [...]uint8{0, 8}
	_EthereumChainIDIndex_1 = [...]uint8{0, 8}
	_EthereumChainIDIndex_2 = [...]uint8{0, 7}
)

func (i EthereumChainID) String() string {
	switch {
	case i == 1:
		return _EthereumChainIDName_0
	case i == 10:
		return _EthereumChainIDName_1
	case i == 137:
		return _EthereumChainIDName_2
	default:
		return fmt.Sprintf("EthereumChainID(%d)", i)
	}
}

func (EthereumChainID) Values() []string {
	return EthereumChainIDStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _EthereumChainIDNoOp() {
	var x [1]struct{}
	_ = x[EthereumChainIDMainnet-(1)]
	_ = x[EthereumChainIDOptimism-(10)]
	_ = x[EthereumChainIDPolygon-(137)]
}

var _EthereumChainIDValues = []EthereumChainID{EthereumChainIDMainnet, EthereumChainIDOptimism, EthereumChainIDPolygon}

var _EthereumChainIDNameToValueMap = map[string]EthereumChainID{
	_EthereumChainIDName_0[0:8]:      EthereumChainIDMainnet,
	_EthereumChainIDLowerName_0[0:8]: EthereumChainIDMainnet,
	_EthereumChainIDName_1[0:8]:      EthereumChainIDOptimism,
	_EthereumChainIDLowerName_1[0:8]: EthereumChainIDOptimism,
	_EthereumChainIDName_2[0:7]:      EthereumChainIDPolygon,
	_EthereumChainIDLowerName_2[0:7]: EthereumChainIDPolygon,
}

var _EthereumChainIDNames = []string{
	_EthereumChainIDName_0[0:8],
	_EthereumChainIDName_1[0:8],
	_EthereumChainIDName_2[0:7],
}

// EthereumChainIDString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func EthereumChainIDString(s string) (EthereumChainID, error) {
	if val, ok := _EthereumChainIDNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _EthereumChainIDNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to EthereumChainID values", s)
}

// EthereumChainIDValues returns all values of the enum
func EthereumChainIDValues() []EthereumChainID {
	return _EthereumChainIDValues
}

// EthereumChainIDStrings returns a slice of all String values of the enum
func EthereumChainIDStrings() []string {
	strs := make([]string, len(_EthereumChainIDNames))
	copy(strs, _EthereumChainIDNames)
	return strs
}

// IsAEthereumChainID returns "true" if the value is listed in the enum definition. "false" otherwise
func (i EthereumChainID) IsAEthereumChainID() bool {
	for _, v := range _EthereumChainIDValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for EthereumChainID
func (i EthereumChainID) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for EthereumChainID
func (i *EthereumChainID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("EthereumChainID should be a string, got %s", data)
	}

	var err error
	*i, err = EthereumChainIDString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for EthereumChainID
func (i EthereumChainID) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for EthereumChainID
func (i *EthereumChainID) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = EthereumChainIDString(s)
	return err
}

func (i EthereumChainID) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *EthereumChainID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of EthereumChainID: %[1]T(%[1]v)", value)
	}

	val, err := EthereumChainIDString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
