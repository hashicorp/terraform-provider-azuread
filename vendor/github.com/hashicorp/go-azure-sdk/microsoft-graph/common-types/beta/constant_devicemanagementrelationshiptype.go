package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementRelationshipType string

const (
	DeviceManagementRelationshipType_And DeviceManagementRelationshipType = "and"
	DeviceManagementRelationshipType_Or  DeviceManagementRelationshipType = "or"
)

func PossibleValuesForDeviceManagementRelationshipType() []string {
	return []string{
		string(DeviceManagementRelationshipType_And),
		string(DeviceManagementRelationshipType_Or),
	}
}

func (s *DeviceManagementRelationshipType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementRelationshipType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementRelationshipType(input string) (*DeviceManagementRelationshipType, error) {
	vals := map[string]DeviceManagementRelationshipType{
		"and": DeviceManagementRelationshipType_And,
		"or":  DeviceManagementRelationshipType_Or,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementRelationshipType(input)
	return &out, nil
}
