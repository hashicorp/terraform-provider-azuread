package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceLogCollectionTemplateType string

const (
	DeviceLogCollectionTemplateType_Predefined DeviceLogCollectionTemplateType = "predefined"
)

func PossibleValuesForDeviceLogCollectionTemplateType() []string {
	return []string{
		string(DeviceLogCollectionTemplateType_Predefined),
	}
}

func (s *DeviceLogCollectionTemplateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceLogCollectionTemplateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceLogCollectionTemplateType(input string) (*DeviceLogCollectionTemplateType, error) {
	vals := map[string]DeviceLogCollectionTemplateType{
		"predefined": DeviceLogCollectionTemplateType_Predefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceLogCollectionTemplateType(input)
	return &out, nil
}
