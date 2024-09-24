package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceAppStatus string

const (
	ServiceAppStatus_Active          ServiceAppStatus = "active"
	ServiceAppStatus_Inactive        ServiceAppStatus = "inactive"
	ServiceAppStatus_PendingActive   ServiceAppStatus = "pendingActive"
	ServiceAppStatus_PendingInactive ServiceAppStatus = "pendingInactive"
)

func PossibleValuesForServiceAppStatus() []string {
	return []string{
		string(ServiceAppStatus_Active),
		string(ServiceAppStatus_Inactive),
		string(ServiceAppStatus_PendingActive),
		string(ServiceAppStatus_PendingInactive),
	}
}

func (s *ServiceAppStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceAppStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceAppStatus(input string) (*ServiceAppStatus, error) {
	vals := map[string]ServiceAppStatus{
		"active":          ServiceAppStatus_Active,
		"inactive":        ServiceAppStatus_Inactive,
		"pendingactive":   ServiceAppStatus_PendingActive,
		"pendinginactive": ServiceAppStatus_PendingInactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceAppStatus(input)
	return &out, nil
}
