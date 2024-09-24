package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionState string

const (
	ConnectionState_Draft         ConnectionState = "draft"
	ConnectionState_LimitExceeded ConnectionState = "limitExceeded"
	ConnectionState_Obsolete      ConnectionState = "obsolete"
	ConnectionState_Ready         ConnectionState = "ready"
)

func PossibleValuesForConnectionState() []string {
	return []string{
		string(ConnectionState_Draft),
		string(ConnectionState_LimitExceeded),
		string(ConnectionState_Obsolete),
		string(ConnectionState_Ready),
	}
}

func (s *ConnectionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectionState(input string) (*ConnectionState, error) {
	vals := map[string]ConnectionState{
		"draft":         ConnectionState_Draft,
		"limitexceeded": ConnectionState_LimitExceeded,
		"obsolete":      ConnectionState_Obsolete,
		"ready":         ConnectionState_Ready,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectionState(input)
	return &out, nil
}
