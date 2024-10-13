// Code generated by "enumer -type=NodeType -linecomment -json=true -text=true -sql=true"; DO NOT EDIT.

package blockchain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _NodeTypeName = "infura"

var _NodeTypeIndex = [...]uint8{0, 6}

const _NodeTypeLowerName = "infura"

func (i NodeType) String() string {
	i -= 1
	if i >= NodeType(len(_NodeTypeIndex)-1) {
		return fmt.Sprintf("NodeType(%d)", i+1)
	}
	return _NodeTypeName[_NodeTypeIndex[i]:_NodeTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _NodeTypeNoOp() {
	var x [1]struct{}
	_ = x[NodeInfura-(1)]
}

var _NodeTypeValues = []NodeType{NodeInfura}

var _NodeTypeNameToValueMap = map[string]NodeType{
	_NodeTypeName[0:6]:      NodeInfura,
	_NodeTypeLowerName[0:6]: NodeInfura,
}

var _NodeTypeNames = []string{
	_NodeTypeName[0:6],
}

// NodeTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NodeTypeString(s string) (NodeType, error) {
	if val, ok := _NodeTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _NodeTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to NodeType values", s)
}

// NodeTypeValues returns all values of the enum
func NodeTypeValues() []NodeType {
	return _NodeTypeValues
}

// NodeTypeStrings returns a slice of all String values of the enum
func NodeTypeStrings() []string {
	strs := make([]string, len(_NodeTypeNames))
	copy(strs, _NodeTypeNames)
	return strs
}

// IsANodeType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i NodeType) IsANodeType() bool {
	for _, v := range _NodeTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for NodeType
func (i NodeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for NodeType
func (i *NodeType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("NodeType should be a string, got %s", data)
	}

	var err error
	*i, err = NodeTypeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for NodeType
func (i NodeType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for NodeType
func (i *NodeType) UnmarshalText(text []byte) error {
	var err error
	*i, err = NodeTypeString(string(text))
	return err
}

func (i NodeType) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *NodeType) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of NodeType: %[1]T(%[1]v)", value)
	}

	val, err := NodeTypeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
