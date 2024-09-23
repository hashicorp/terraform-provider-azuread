package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDefenderApplicationControlSupplementalPolicyStatuses string

const (
	WindowsDefenderApplicationControlSupplementalPolicyStatuses_NotAuthorizedByToken WindowsDefenderApplicationControlSupplementalPolicyStatuses = "notAuthorizedByToken"
	WindowsDefenderApplicationControlSupplementalPolicyStatuses_PolicyNotFound       WindowsDefenderApplicationControlSupplementalPolicyStatuses = "policyNotFound"
	WindowsDefenderApplicationControlSupplementalPolicyStatuses_Success              WindowsDefenderApplicationControlSupplementalPolicyStatuses = "success"
	WindowsDefenderApplicationControlSupplementalPolicyStatuses_TokenError           WindowsDefenderApplicationControlSupplementalPolicyStatuses = "tokenError"
	WindowsDefenderApplicationControlSupplementalPolicyStatuses_Unknown              WindowsDefenderApplicationControlSupplementalPolicyStatuses = "unknown"
)

func PossibleValuesForWindowsDefenderApplicationControlSupplementalPolicyStatuses() []string {
	return []string{
		string(WindowsDefenderApplicationControlSupplementalPolicyStatuses_NotAuthorizedByToken),
		string(WindowsDefenderApplicationControlSupplementalPolicyStatuses_PolicyNotFound),
		string(WindowsDefenderApplicationControlSupplementalPolicyStatuses_Success),
		string(WindowsDefenderApplicationControlSupplementalPolicyStatuses_TokenError),
		string(WindowsDefenderApplicationControlSupplementalPolicyStatuses_Unknown),
	}
}

func (s *WindowsDefenderApplicationControlSupplementalPolicyStatuses) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDefenderApplicationControlSupplementalPolicyStatuses(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDefenderApplicationControlSupplementalPolicyStatuses(input string) (*WindowsDefenderApplicationControlSupplementalPolicyStatuses, error) {
	vals := map[string]WindowsDefenderApplicationControlSupplementalPolicyStatuses{
		"notauthorizedbytoken": WindowsDefenderApplicationControlSupplementalPolicyStatuses_NotAuthorizedByToken,
		"policynotfound":       WindowsDefenderApplicationControlSupplementalPolicyStatuses_PolicyNotFound,
		"success":              WindowsDefenderApplicationControlSupplementalPolicyStatuses_Success,
		"tokenerror":           WindowsDefenderApplicationControlSupplementalPolicyStatuses_TokenError,
		"unknown":              WindowsDefenderApplicationControlSupplementalPolicyStatuses_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDefenderApplicationControlSupplementalPolicyStatuses(input)
	return &out, nil
}
