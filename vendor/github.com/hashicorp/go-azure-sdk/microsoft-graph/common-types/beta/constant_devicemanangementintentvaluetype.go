package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManangementIntentValueType string

const (
	DeviceManangementIntentValueType_AbstractComplex DeviceManangementIntentValueType = "abstractComplex"
	DeviceManangementIntentValueType_Boolean         DeviceManangementIntentValueType = "boolean"
	DeviceManangementIntentValueType_Collection      DeviceManangementIntentValueType = "collection"
	DeviceManangementIntentValueType_Complex         DeviceManangementIntentValueType = "complex"
	DeviceManangementIntentValueType_Integer         DeviceManangementIntentValueType = "integer"
	DeviceManangementIntentValueType_String          DeviceManangementIntentValueType = "string"
)

func PossibleValuesForDeviceManangementIntentValueType() []string {
	return []string{
		string(DeviceManangementIntentValueType_AbstractComplex),
		string(DeviceManangementIntentValueType_Boolean),
		string(DeviceManangementIntentValueType_Collection),
		string(DeviceManangementIntentValueType_Complex),
		string(DeviceManangementIntentValueType_Integer),
		string(DeviceManangementIntentValueType_String),
	}
}

func (s *DeviceManangementIntentValueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManangementIntentValueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManangementIntentValueType(input string) (*DeviceManangementIntentValueType, error) {
	vals := map[string]DeviceManangementIntentValueType{
		"abstractcomplex": DeviceManangementIntentValueType_AbstractComplex,
		"boolean":         DeviceManangementIntentValueType_Boolean,
		"collection":      DeviceManangementIntentValueType_Collection,
		"complex":         DeviceManangementIntentValueType_Complex,
		"integer":         DeviceManangementIntentValueType_Integer,
		"string":          DeviceManangementIntentValueType_String,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManangementIntentValueType(input)
	return &out, nil
}
