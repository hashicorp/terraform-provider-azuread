package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileAccountUse string

const (
	AndroidWorkProfileAccountUse_AllowAll                     AndroidWorkProfileAccountUse = "allowAll"
	AndroidWorkProfileAccountUse_AllowAllExceptGoogleAccounts AndroidWorkProfileAccountUse = "allowAllExceptGoogleAccounts"
	AndroidWorkProfileAccountUse_BlockAll                     AndroidWorkProfileAccountUse = "blockAll"
)

func PossibleValuesForAndroidWorkProfileAccountUse() []string {
	return []string{
		string(AndroidWorkProfileAccountUse_AllowAll),
		string(AndroidWorkProfileAccountUse_AllowAllExceptGoogleAccounts),
		string(AndroidWorkProfileAccountUse_BlockAll),
	}
}

func (s *AndroidWorkProfileAccountUse) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileAccountUse(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileAccountUse(input string) (*AndroidWorkProfileAccountUse, error) {
	vals := map[string]AndroidWorkProfileAccountUse{
		"allowall":                     AndroidWorkProfileAccountUse_AllowAll,
		"allowallexceptgoogleaccounts": AndroidWorkProfileAccountUse_AllowAllExceptGoogleAccounts,
		"blockall":                     AndroidWorkProfileAccountUse_BlockAll,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileAccountUse(input)
	return &out, nil
}
