// Code generated by "enumer --values --type=Worker --linecomment --output worker_string.go --json --yaml --sql"; DO NOT EDIT.

package decentralized

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _WorkerName = "aaveaavegotchibasecorecrossbellcurveenshighlightiqwikikiwistandlenslidolooksraremattersmirrormomoka1inchopenseaoptimismparagraphrss3savmstargateuniswapvsl"

var _WorkerIndex = [...]uint8{0, 4, 14, 18, 22, 31, 36, 39, 48, 54, 63, 67, 71, 80, 87, 93, 99, 104, 111, 119, 128, 132, 136, 144, 151, 154}

const _WorkerLowerName = "aaveaavegotchibasecorecrossbellcurveenshighlightiqwikikiwistandlenslidolooksraremattersmirrormomoka1inchopenseaoptimismparagraphrss3savmstargateuniswapvsl"

func (i Worker) String() string {
	i -= 1
	if i < 0 || i >= Worker(len(_WorkerIndex)-1) {
		return fmt.Sprintf("Worker(%d)", i+1)
	}
	return _WorkerName[_WorkerIndex[i]:_WorkerIndex[i+1]]
}

func (Worker) Values() []string {
	return WorkerStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _WorkerNoOp() {
	var x [1]struct{}
	_ = x[Aave-(1)]
	_ = x[Aavegotchi-(2)]
	_ = x[Base-(3)]
	_ = x[Core-(4)]
	_ = x[Crossbell-(5)]
	_ = x[Curve-(6)]
	_ = x[ENS-(7)]
	_ = x[Highlight-(8)]
	_ = x[IQWiki-(9)]
	_ = x[KiwiStand-(10)]
	_ = x[Lens-(11)]
	_ = x[Lido-(12)]
	_ = x[Looksrare-(13)]
	_ = x[Matters-(14)]
	_ = x[Mirror-(15)]
	_ = x[Momoka-(16)]
	_ = x[Oneinch-(17)]
	_ = x[OpenSea-(18)]
	_ = x[Optimism-(19)]
	_ = x[Paragraph-(20)]
	_ = x[RSS3-(21)]
	_ = x[SAVM-(22)]
	_ = x[Stargate-(23)]
	_ = x[Uniswap-(24)]
	_ = x[VSL-(25)]
}

var _WorkerValues = []Worker{Aave, Aavegotchi, Base, Core, Crossbell, Curve, ENS, Highlight, IQWiki, KiwiStand, Lens, Lido, Looksrare, Matters, Mirror, Momoka, Oneinch, OpenSea, Optimism, Paragraph, RSS3, SAVM, Stargate, Uniswap, VSL}

var _WorkerNameToValueMap = map[string]Worker{
	_WorkerName[0:4]:          Aave,
	_WorkerLowerName[0:4]:     Aave,
	_WorkerName[4:14]:         Aavegotchi,
	_WorkerLowerName[4:14]:    Aavegotchi,
	_WorkerName[14:18]:        Base,
	_WorkerLowerName[14:18]:   Base,
	_WorkerName[18:22]:        Core,
	_WorkerLowerName[18:22]:   Core,
	_WorkerName[22:31]:        Crossbell,
	_WorkerLowerName[22:31]:   Crossbell,
	_WorkerName[31:36]:        Curve,
	_WorkerLowerName[31:36]:   Curve,
	_WorkerName[36:39]:        ENS,
	_WorkerLowerName[36:39]:   ENS,
	_WorkerName[39:48]:        Highlight,
	_WorkerLowerName[39:48]:   Highlight,
	_WorkerName[48:54]:        IQWiki,
	_WorkerLowerName[48:54]:   IQWiki,
	_WorkerName[54:63]:        KiwiStand,
	_WorkerLowerName[54:63]:   KiwiStand,
	_WorkerName[63:67]:        Lens,
	_WorkerLowerName[63:67]:   Lens,
	_WorkerName[67:71]:        Lido,
	_WorkerLowerName[67:71]:   Lido,
	_WorkerName[71:80]:        Looksrare,
	_WorkerLowerName[71:80]:   Looksrare,
	_WorkerName[80:87]:        Matters,
	_WorkerLowerName[80:87]:   Matters,
	_WorkerName[87:93]:        Mirror,
	_WorkerLowerName[87:93]:   Mirror,
	_WorkerName[93:99]:        Momoka,
	_WorkerLowerName[93:99]:   Momoka,
	_WorkerName[99:104]:       Oneinch,
	_WorkerLowerName[99:104]:  Oneinch,
	_WorkerName[104:111]:      OpenSea,
	_WorkerLowerName[104:111]: OpenSea,
	_WorkerName[111:119]:      Optimism,
	_WorkerLowerName[111:119]: Optimism,
	_WorkerName[119:128]:      Paragraph,
	_WorkerLowerName[119:128]: Paragraph,
	_WorkerName[128:132]:      RSS3,
	_WorkerLowerName[128:132]: RSS3,
	_WorkerName[132:136]:      SAVM,
	_WorkerLowerName[132:136]: SAVM,
	_WorkerName[136:144]:      Stargate,
	_WorkerLowerName[136:144]: Stargate,
	_WorkerName[144:151]:      Uniswap,
	_WorkerLowerName[144:151]: Uniswap,
	_WorkerName[151:154]:      VSL,
	_WorkerLowerName[151:154]: VSL,
}

var _WorkerNames = []string{
	_WorkerName[0:4],
	_WorkerName[4:14],
	_WorkerName[14:18],
	_WorkerName[18:22],
	_WorkerName[22:31],
	_WorkerName[31:36],
	_WorkerName[36:39],
	_WorkerName[39:48],
	_WorkerName[48:54],
	_WorkerName[54:63],
	_WorkerName[63:67],
	_WorkerName[67:71],
	_WorkerName[71:80],
	_WorkerName[80:87],
	_WorkerName[87:93],
	_WorkerName[93:99],
	_WorkerName[99:104],
	_WorkerName[104:111],
	_WorkerName[111:119],
	_WorkerName[119:128],
	_WorkerName[128:132],
	_WorkerName[132:136],
	_WorkerName[136:144],
	_WorkerName[144:151],
	_WorkerName[151:154],
}

// WorkerString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func WorkerString(s string) (Worker, error) {
	if val, ok := _WorkerNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _WorkerNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Worker values", s)
}

// WorkerValues returns all values of the enum
func WorkerValues() []Worker {
	return _WorkerValues
}

// WorkerStrings returns a slice of all String values of the enum
func WorkerStrings() []string {
	strs := make([]string, len(_WorkerNames))
	copy(strs, _WorkerNames)
	return strs
}

// IsAWorker returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Worker) IsAWorker() bool {
	for _, v := range _WorkerValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Worker
func (i Worker) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Worker
func (i *Worker) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Worker should be a string, got %s", data)
	}

	var err error
	*i, err = WorkerString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for Worker
func (i Worker) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Worker
func (i *Worker) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = WorkerString(s)
	return err
}

func (i Worker) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Worker) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of Worker: %[1]T(%[1]v)", value)
	}

	val, err := WorkerString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
