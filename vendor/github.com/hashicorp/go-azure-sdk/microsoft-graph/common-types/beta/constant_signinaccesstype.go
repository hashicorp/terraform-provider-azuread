package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInAccessType string

const (
	SignInAccessType_B2bCollaboration SignInAccessType = "b2bCollaboration"
	SignInAccessType_B2bDirectConnect SignInAccessType = "b2bDirectConnect"
	SignInAccessType_MicrosoftSupport SignInAccessType = "microsoftSupport"
	SignInAccessType_None             SignInAccessType = "none"
	SignInAccessType_Passthrough      SignInAccessType = "passthrough"
	SignInAccessType_ServiceProvider  SignInAccessType = "serviceProvider"
)

func PossibleValuesForSignInAccessType() []string {
	return []string{
		string(SignInAccessType_B2bCollaboration),
		string(SignInAccessType_B2bDirectConnect),
		string(SignInAccessType_MicrosoftSupport),
		string(SignInAccessType_None),
		string(SignInAccessType_Passthrough),
		string(SignInAccessType_ServiceProvider),
	}
}

func (s *SignInAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInAccessType(input string) (*SignInAccessType, error) {
	vals := map[string]SignInAccessType{
		"b2bcollaboration": SignInAccessType_B2bCollaboration,
		"b2bdirectconnect": SignInAccessType_B2bDirectConnect,
		"microsoftsupport": SignInAccessType_MicrosoftSupport,
		"none":             SignInAccessType_None,
		"passthrough":      SignInAccessType_Passthrough,
		"serviceprovider":  SignInAccessType_ServiceProvider,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInAccessType(input)
	return &out, nil
}
