package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCustomAttributeValueType string

const (
	DeviceCustomAttributeValueType_DateTime DeviceCustomAttributeValueType = "dateTime"
	DeviceCustomAttributeValueType_Integer  DeviceCustomAttributeValueType = "integer"
	DeviceCustomAttributeValueType_String   DeviceCustomAttributeValueType = "string"
)

func PossibleValuesForDeviceCustomAttributeValueType() []string {
	return []string{
		string(DeviceCustomAttributeValueType_DateTime),
		string(DeviceCustomAttributeValueType_Integer),
		string(DeviceCustomAttributeValueType_String),
	}
}

func (s *DeviceCustomAttributeValueType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceCustomAttributeValueType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceCustomAttributeValueType(input string) (*DeviceCustomAttributeValueType, error) {
	vals := map[string]DeviceCustomAttributeValueType{
		"datetime": DeviceCustomAttributeValueType_DateTime,
		"integer":  DeviceCustomAttributeValueType_Integer,
		"string":   DeviceCustomAttributeValueType_String,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCustomAttributeValueType(input)
	return &out, nil
}
