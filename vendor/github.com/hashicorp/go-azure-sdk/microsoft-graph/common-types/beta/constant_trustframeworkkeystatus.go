package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrustFrameworkKeyStatus string

const (
	TrustFrameworkKeyStatus_Disabled TrustFrameworkKeyStatus = "disabled"
	TrustFrameworkKeyStatus_Enabled  TrustFrameworkKeyStatus = "enabled"
)

func PossibleValuesForTrustFrameworkKeyStatus() []string {
	return []string{
		string(TrustFrameworkKeyStatus_Disabled),
		string(TrustFrameworkKeyStatus_Enabled),
	}
}

func (s *TrustFrameworkKeyStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrustFrameworkKeyStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrustFrameworkKeyStatus(input string) (*TrustFrameworkKeyStatus, error) {
	vals := map[string]TrustFrameworkKeyStatus{
		"disabled": TrustFrameworkKeyStatus_Disabled,
		"enabled":  TrustFrameworkKeyStatus_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrustFrameworkKeyStatus(input)
	return &out, nil
}
