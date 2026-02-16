package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceNowConnectionStatus string

const (
	ServiceNowConnectionStatus_Disabled ServiceNowConnectionStatus = "disabled"
	ServiceNowConnectionStatus_Enabled  ServiceNowConnectionStatus = "enabled"
)

func PossibleValuesForServiceNowConnectionStatus() []string {
	return []string{
		string(ServiceNowConnectionStatus_Disabled),
		string(ServiceNowConnectionStatus_Enabled),
	}
}

func (s *ServiceNowConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceNowConnectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceNowConnectionStatus(input string) (*ServiceNowConnectionStatus, error) {
	vals := map[string]ServiceNowConnectionStatus{
		"disabled": ServiceNowConnectionStatus_Disabled,
		"enabled":  ServiceNowConnectionStatus_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceNowConnectionStatus(input)
	return &out, nil
}
