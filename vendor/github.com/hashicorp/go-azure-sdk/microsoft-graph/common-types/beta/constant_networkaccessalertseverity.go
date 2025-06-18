package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessAlertSeverity string

const (
	NetworkaccessAlertSeverity_High          NetworkaccessAlertSeverity = "high"
	NetworkaccessAlertSeverity_Informational NetworkaccessAlertSeverity = "informational"
	NetworkaccessAlertSeverity_Low           NetworkaccessAlertSeverity = "low"
	NetworkaccessAlertSeverity_Medium        NetworkaccessAlertSeverity = "medium"
)

func PossibleValuesForNetworkaccessAlertSeverity() []string {
	return []string{
		string(NetworkaccessAlertSeverity_High),
		string(NetworkaccessAlertSeverity_Informational),
		string(NetworkaccessAlertSeverity_Low),
		string(NetworkaccessAlertSeverity_Medium),
	}
}

func (s *NetworkaccessAlertSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessAlertSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessAlertSeverity(input string) (*NetworkaccessAlertSeverity, error) {
	vals := map[string]NetworkaccessAlertSeverity{
		"high":          NetworkaccessAlertSeverity_High,
		"informational": NetworkaccessAlertSeverity_Informational,
		"low":           NetworkaccessAlertSeverity_Low,
		"medium":        NetworkaccessAlertSeverity_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessAlertSeverity(input)
	return &out, nil
}
