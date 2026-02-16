package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StagedFeatureName string

const (
	StagedFeatureName_CertificateBasedAuthentication StagedFeatureName = "certificateBasedAuthentication"
	StagedFeatureName_EmailAsAlternateId             StagedFeatureName = "emailAsAlternateId"
	StagedFeatureName_MultiFactorAuthentication      StagedFeatureName = "multiFactorAuthentication"
	StagedFeatureName_PassthroughAuthentication      StagedFeatureName = "passthroughAuthentication"
	StagedFeatureName_PasswordHashSync               StagedFeatureName = "passwordHashSync"
	StagedFeatureName_SeamlessSso                    StagedFeatureName = "seamlessSso"
)

func PossibleValuesForStagedFeatureName() []string {
	return []string{
		string(StagedFeatureName_CertificateBasedAuthentication),
		string(StagedFeatureName_EmailAsAlternateId),
		string(StagedFeatureName_MultiFactorAuthentication),
		string(StagedFeatureName_PassthroughAuthentication),
		string(StagedFeatureName_PasswordHashSync),
		string(StagedFeatureName_SeamlessSso),
	}
}

func (s *StagedFeatureName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStagedFeatureName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStagedFeatureName(input string) (*StagedFeatureName, error) {
	vals := map[string]StagedFeatureName{
		"certificatebasedauthentication": StagedFeatureName_CertificateBasedAuthentication,
		"emailasalternateid":             StagedFeatureName_EmailAsAlternateId,
		"multifactorauthentication":      StagedFeatureName_MultiFactorAuthentication,
		"passthroughauthentication":      StagedFeatureName_PassthroughAuthentication,
		"passwordhashsync":               StagedFeatureName_PasswordHashSync,
		"seamlesssso":                    StagedFeatureName_SeamlessSso,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StagedFeatureName(input)
	return &out, nil
}
