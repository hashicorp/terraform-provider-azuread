package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessTlsAction string

const (
	NetworkaccessTlsAction_Bypassed    NetworkaccessTlsAction = "bypassed"
	NetworkaccessTlsAction_Intercepted NetworkaccessTlsAction = "intercepted"
)

func PossibleValuesForNetworkaccessTlsAction() []string {
	return []string{
		string(NetworkaccessTlsAction_Bypassed),
		string(NetworkaccessTlsAction_Intercepted),
	}
}

func (s *NetworkaccessTlsAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessTlsAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessTlsAction(input string) (*NetworkaccessTlsAction, error) {
	vals := map[string]NetworkaccessTlsAction{
		"bypassed":    NetworkaccessTlsAction_Bypassed,
		"intercepted": NetworkaccessTlsAction_Intercepted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessTlsAction(input)
	return &out, nil
}
