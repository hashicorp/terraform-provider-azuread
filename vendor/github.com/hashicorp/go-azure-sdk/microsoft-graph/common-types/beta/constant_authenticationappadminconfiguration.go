package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAppAdminConfiguration string

const (
	AuthenticationAppAdminConfiguration_Disabled      AuthenticationAppAdminConfiguration = "disabled"
	AuthenticationAppAdminConfiguration_Enabled       AuthenticationAppAdminConfiguration = "enabled"
	AuthenticationAppAdminConfiguration_NotApplicable AuthenticationAppAdminConfiguration = "notApplicable"
)

func PossibleValuesForAuthenticationAppAdminConfiguration() []string {
	return []string{
		string(AuthenticationAppAdminConfiguration_Disabled),
		string(AuthenticationAppAdminConfiguration_Enabled),
		string(AuthenticationAppAdminConfiguration_NotApplicable),
	}
}

func (s *AuthenticationAppAdminConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationAppAdminConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationAppAdminConfiguration(input string) (*AuthenticationAppAdminConfiguration, error) {
	vals := map[string]AuthenticationAppAdminConfiguration{
		"disabled":      AuthenticationAppAdminConfiguration_Disabled,
		"enabled":       AuthenticationAppAdminConfiguration_Enabled,
		"notapplicable": AuthenticationAppAdminConfiguration_NotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationAppAdminConfiguration(input)
	return &out, nil
}
