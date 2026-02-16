package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElevationRequestState string

const (
	ElevationRequestState_Approved  ElevationRequestState = "approved"
	ElevationRequestState_Completed ElevationRequestState = "completed"
	ElevationRequestState_Denied    ElevationRequestState = "denied"
	ElevationRequestState_Expired   ElevationRequestState = "expired"
	ElevationRequestState_None      ElevationRequestState = "none"
	ElevationRequestState_Pending   ElevationRequestState = "pending"
	ElevationRequestState_Revoked   ElevationRequestState = "revoked"
)

func PossibleValuesForElevationRequestState() []string {
	return []string{
		string(ElevationRequestState_Approved),
		string(ElevationRequestState_Completed),
		string(ElevationRequestState_Denied),
		string(ElevationRequestState_Expired),
		string(ElevationRequestState_None),
		string(ElevationRequestState_Pending),
		string(ElevationRequestState_Revoked),
	}
}

func (s *ElevationRequestState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseElevationRequestState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseElevationRequestState(input string) (*ElevationRequestState, error) {
	vals := map[string]ElevationRequestState{
		"approved":  ElevationRequestState_Approved,
		"completed": ElevationRequestState_Completed,
		"denied":    ElevationRequestState_Denied,
		"expired":   ElevationRequestState_Expired,
		"none":      ElevationRequestState_None,
		"pending":   ElevationRequestState_Pending,
		"revoked":   ElevationRequestState_Revoked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ElevationRequestState(input)
	return &out, nil
}
