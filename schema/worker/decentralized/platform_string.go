// Code generated by "enumer --values --type=Platform --linecomment --output platform_string.go --json --yaml --sql"; DO NOT EDIT.

package decentralized

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _PlatformName = "Unknown1inchAAVEAavegotchiCrossbellCurveENSFarcasterHighlightIQWikiKiwiStandLensLidoLooksRareMattersMirrorOpenSeaOptimismParagraphRSS3SAVMStargateUniswapVSL"

var _PlatformIndex = [...]uint8{0, 7, 12, 16, 26, 35, 40, 43, 52, 61, 67, 76, 80, 84, 93, 100, 106, 113, 121, 130, 134, 138, 146, 153, 156}

const _PlatformLowerName = "unknown1inchaaveaavegotchicrossbellcurveensfarcasterhighlightiqwikikiwistandlenslidolooksraremattersmirroropenseaoptimismparagraphrss3savmstargateuniswapvsl"

func (i Platform) String() string {
	if i >= Platform(len(_PlatformIndex)-1) {
		return fmt.Sprintf("Platform(%d)", i)
	}
	return _PlatformName[_PlatformIndex[i]:_PlatformIndex[i+1]]
}

func (Platform) Values() []string {
	return PlatformStrings()
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PlatformNoOp() {
	var x [1]struct{}
	_ = x[PlatformUnknown-(0)]
	_ = x[Platform1Inch-(1)]
	_ = x[PlatformAAVE-(2)]
	_ = x[PlatformAavegotchi-(3)]
	_ = x[PlatformCrossbell-(4)]
	_ = x[PlatformCurve-(5)]
	_ = x[PlatformENS-(6)]
	_ = x[PlatformFarcaster-(7)]
	_ = x[PlatformHighlight-(8)]
	_ = x[PlatformIQWiki-(9)]
	_ = x[PlatformKiwiStand-(10)]
	_ = x[PlatformLens-(11)]
	_ = x[PlatformLido-(12)]
	_ = x[PlatformLooksRare-(13)]
	_ = x[PlatformMatters-(14)]
	_ = x[PlatformMirror-(15)]
	_ = x[PlatformOpenSea-(16)]
	_ = x[PlatformOptimism-(17)]
	_ = x[PlatformParagraph-(18)]
	_ = x[PlatformRSS3-(19)]
	_ = x[PlatformSAVM-(20)]
	_ = x[PlatformStargate-(21)]
	_ = x[PlatformUniswap-(22)]
	_ = x[PlatformVSL-(23)]
}

var _PlatformValues = []Platform{PlatformUnknown, Platform1Inch, PlatformAAVE, PlatformAavegotchi, PlatformCrossbell, PlatformCurve, PlatformENS, PlatformFarcaster, PlatformHighlight, PlatformIQWiki, PlatformKiwiStand, PlatformLens, PlatformLido, PlatformLooksRare, PlatformMatters, PlatformMirror, PlatformOpenSea, PlatformOptimism, PlatformParagraph, PlatformRSS3, PlatformSAVM, PlatformStargate, PlatformUniswap, PlatformVSL}

var _PlatformNameToValueMap = map[string]Platform{
	_PlatformName[0:7]:          PlatformUnknown,
	_PlatformLowerName[0:7]:     PlatformUnknown,
	_PlatformName[7:12]:         Platform1Inch,
	_PlatformLowerName[7:12]:    Platform1Inch,
	_PlatformName[12:16]:        PlatformAAVE,
	_PlatformLowerName[12:16]:   PlatformAAVE,
	_PlatformName[16:26]:        PlatformAavegotchi,
	_PlatformLowerName[16:26]:   PlatformAavegotchi,
	_PlatformName[26:35]:        PlatformCrossbell,
	_PlatformLowerName[26:35]:   PlatformCrossbell,
	_PlatformName[35:40]:        PlatformCurve,
	_PlatformLowerName[35:40]:   PlatformCurve,
	_PlatformName[40:43]:        PlatformENS,
	_PlatformLowerName[40:43]:   PlatformENS,
	_PlatformName[43:52]:        PlatformFarcaster,
	_PlatformLowerName[43:52]:   PlatformFarcaster,
	_PlatformName[52:61]:        PlatformHighlight,
	_PlatformLowerName[52:61]:   PlatformHighlight,
	_PlatformName[61:67]:        PlatformIQWiki,
	_PlatformLowerName[61:67]:   PlatformIQWiki,
	_PlatformName[67:76]:        PlatformKiwiStand,
	_PlatformLowerName[67:76]:   PlatformKiwiStand,
	_PlatformName[76:80]:        PlatformLens,
	_PlatformLowerName[76:80]:   PlatformLens,
	_PlatformName[80:84]:        PlatformLido,
	_PlatformLowerName[80:84]:   PlatformLido,
	_PlatformName[84:93]:        PlatformLooksRare,
	_PlatformLowerName[84:93]:   PlatformLooksRare,
	_PlatformName[93:100]:       PlatformMatters,
	_PlatformLowerName[93:100]:  PlatformMatters,
	_PlatformName[100:106]:      PlatformMirror,
	_PlatformLowerName[100:106]: PlatformMirror,
	_PlatformName[106:113]:      PlatformOpenSea,
	_PlatformLowerName[106:113]: PlatformOpenSea,
	_PlatformName[113:121]:      PlatformOptimism,
	_PlatformLowerName[113:121]: PlatformOptimism,
	_PlatformName[121:130]:      PlatformParagraph,
	_PlatformLowerName[121:130]: PlatformParagraph,
	_PlatformName[130:134]:      PlatformRSS3,
	_PlatformLowerName[130:134]: PlatformRSS3,
	_PlatformName[134:138]:      PlatformSAVM,
	_PlatformLowerName[134:138]: PlatformSAVM,
	_PlatformName[138:146]:      PlatformStargate,
	_PlatformLowerName[138:146]: PlatformStargate,
	_PlatformName[146:153]:      PlatformUniswap,
	_PlatformLowerName[146:153]: PlatformUniswap,
	_PlatformName[153:156]:      PlatformVSL,
	_PlatformLowerName[153:156]: PlatformVSL,
}

var _PlatformNames = []string{
	_PlatformName[0:7],
	_PlatformName[7:12],
	_PlatformName[12:16],
	_PlatformName[16:26],
	_PlatformName[26:35],
	_PlatformName[35:40],
	_PlatformName[40:43],
	_PlatformName[43:52],
	_PlatformName[52:61],
	_PlatformName[61:67],
	_PlatformName[67:76],
	_PlatformName[76:80],
	_PlatformName[80:84],
	_PlatformName[84:93],
	_PlatformName[93:100],
	_PlatformName[100:106],
	_PlatformName[106:113],
	_PlatformName[113:121],
	_PlatformName[121:130],
	_PlatformName[130:134],
	_PlatformName[134:138],
	_PlatformName[138:146],
	_PlatformName[146:153],
	_PlatformName[153:156],
}

// PlatformString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PlatformString(s string) (Platform, error) {
	if val, ok := _PlatformNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _PlatformNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Platform values", s)
}

// PlatformValues returns all values of the enum
func PlatformValues() []Platform {
	return _PlatformValues
}

// PlatformStrings returns a slice of all String values of the enum
func PlatformStrings() []string {
	strs := make([]string, len(_PlatformNames))
	copy(strs, _PlatformNames)
	return strs
}

// IsAPlatform returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Platform) IsAPlatform() bool {
	for _, v := range _PlatformValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for Platform
func (i Platform) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Platform
func (i *Platform) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Platform should be a string, got %s", data)
	}

	var err error
	*i, err = PlatformString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for Platform
func (i Platform) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Platform
func (i *Platform) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = PlatformString(s)
	return err
}

func (i Platform) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Platform) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of Platform: %[1]T(%[1]v)", value)
	}

	val, err := PlatformString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
