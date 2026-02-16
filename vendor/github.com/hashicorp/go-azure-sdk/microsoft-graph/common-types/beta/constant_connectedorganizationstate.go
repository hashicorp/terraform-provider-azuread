package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectedOrganizationState string

const (
	ConnectedOrganizationState_Configured ConnectedOrganizationState = "configured"
	ConnectedOrganizationState_Proposed   ConnectedOrganizationState = "proposed"
)

func PossibleValuesForConnectedOrganizationState() []string {
	return []string{
		string(ConnectedOrganizationState_Configured),
		string(ConnectedOrganizationState_Proposed),
	}
}

func (s *ConnectedOrganizationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectedOrganizationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectedOrganizationState(input string) (*ConnectedOrganizationState, error) {
	vals := map[string]ConnectedOrganizationState{
		"configured": ConnectedOrganizationState_Configured,
		"proposed":   ConnectedOrganizationState_Proposed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectedOrganizationState(input)
	return &out, nil
}
