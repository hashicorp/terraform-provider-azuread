package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EapFastConfiguration string

const (
	EapFastConfiguration_NoProtectedAccessCredential                         EapFastConfiguration = "noProtectedAccessCredential"
	EapFastConfiguration_UseProtectedAccessCredential                        EapFastConfiguration = "useProtectedAccessCredential"
	EapFastConfiguration_UseProtectedAccessCredentialAndProvision            EapFastConfiguration = "useProtectedAccessCredentialAndProvision"
	EapFastConfiguration_UseProtectedAccessCredentialAndProvisionAnonymously EapFastConfiguration = "useProtectedAccessCredentialAndProvisionAnonymously"
)

func PossibleValuesForEapFastConfiguration() []string {
	return []string{
		string(EapFastConfiguration_NoProtectedAccessCredential),
		string(EapFastConfiguration_UseProtectedAccessCredential),
		string(EapFastConfiguration_UseProtectedAccessCredentialAndProvision),
		string(EapFastConfiguration_UseProtectedAccessCredentialAndProvisionAnonymously),
	}
}

func (s *EapFastConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEapFastConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEapFastConfiguration(input string) (*EapFastConfiguration, error) {
	vals := map[string]EapFastConfiguration{
		"noprotectedaccesscredential":                         EapFastConfiguration_NoProtectedAccessCredential,
		"useprotectedaccesscredential":                        EapFastConfiguration_UseProtectedAccessCredential,
		"useprotectedaccesscredentialandprovision":            EapFastConfiguration_UseProtectedAccessCredentialAndProvision,
		"useprotectedaccesscredentialandprovisionanonymously": EapFastConfiguration_UseProtectedAccessCredentialAndProvisionAnonymously,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EapFastConfiguration(input)
	return &out, nil
}
