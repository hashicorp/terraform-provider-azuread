package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsConnectionState string

const (
	ExternalConnectorsConnectionState_Draft         ExternalConnectorsConnectionState = "draft"
	ExternalConnectorsConnectionState_LimitExceeded ExternalConnectorsConnectionState = "limitExceeded"
	ExternalConnectorsConnectionState_Obsolete      ExternalConnectorsConnectionState = "obsolete"
	ExternalConnectorsConnectionState_Ready         ExternalConnectorsConnectionState = "ready"
)

func PossibleValuesForExternalConnectorsConnectionState() []string {
	return []string{
		string(ExternalConnectorsConnectionState_Draft),
		string(ExternalConnectorsConnectionState_LimitExceeded),
		string(ExternalConnectorsConnectionState_Obsolete),
		string(ExternalConnectorsConnectionState_Ready),
	}
}

func (s *ExternalConnectorsConnectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsConnectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsConnectionState(input string) (*ExternalConnectorsConnectionState, error) {
	vals := map[string]ExternalConnectorsConnectionState{
		"draft":         ExternalConnectorsConnectionState_Draft,
		"limitexceeded": ExternalConnectorsConnectionState_LimitExceeded,
		"obsolete":      ExternalConnectorsConnectionState_Obsolete,
		"ready":         ExternalConnectorsConnectionState_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsConnectionState(input)
	return &out, nil
}
