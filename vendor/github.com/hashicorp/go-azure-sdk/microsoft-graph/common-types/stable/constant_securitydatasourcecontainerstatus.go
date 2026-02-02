package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDataSourceContainerStatus string

const (
	SecurityDataSourceContainerStatus_Active   SecurityDataSourceContainerStatus = "active"
	SecurityDataSourceContainerStatus_Released SecurityDataSourceContainerStatus = "released"
)

func PossibleValuesForSecurityDataSourceContainerStatus() []string {
	return []string{
		string(SecurityDataSourceContainerStatus_Active),
		string(SecurityDataSourceContainerStatus_Released),
	}
}

func (s *SecurityDataSourceContainerStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDataSourceContainerStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDataSourceContainerStatus(input string) (*SecurityDataSourceContainerStatus, error) {
	vals := map[string]SecurityDataSourceContainerStatus{
		"active":   SecurityDataSourceContainerStatus_Active,
		"released": SecurityDataSourceContainerStatus_Released,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDataSourceContainerStatus(input)
	return &out, nil
}
