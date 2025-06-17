package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCloudAppInfoState string

const (
	SecurityCloudAppInfoState_False   SecurityCloudAppInfoState = "false"
	SecurityCloudAppInfoState_True    SecurityCloudAppInfoState = "true"
	SecurityCloudAppInfoState_Unknown SecurityCloudAppInfoState = "unknown"
)

func PossibleValuesForSecurityCloudAppInfoState() []string {
	return []string{
		string(SecurityCloudAppInfoState_False),
		string(SecurityCloudAppInfoState_True),
		string(SecurityCloudAppInfoState_Unknown),
	}
}

func (s *SecurityCloudAppInfoState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityCloudAppInfoState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityCloudAppInfoState(input string) (*SecurityCloudAppInfoState, error) {
	vals := map[string]SecurityCloudAppInfoState{
		"false":   SecurityCloudAppInfoState_False,
		"true":    SecurityCloudAppInfoState_True,
		"unknown": SecurityCloudAppInfoState_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCloudAppInfoState(input)
	return &out, nil
}
