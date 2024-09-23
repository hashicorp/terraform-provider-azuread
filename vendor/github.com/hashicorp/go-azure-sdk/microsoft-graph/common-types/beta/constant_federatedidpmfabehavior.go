package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FederatedIdpMfaBehavior string

const (
	FederatedIdpMfaBehavior_AcceptIfMfaDoneByFederatedIdp FederatedIdpMfaBehavior = "acceptIfMfaDoneByFederatedIdp"
	FederatedIdpMfaBehavior_EnforceMfaByFederatedIdp      FederatedIdpMfaBehavior = "enforceMfaByFederatedIdp"
	FederatedIdpMfaBehavior_RejectMfaByFederatedIdp       FederatedIdpMfaBehavior = "rejectMfaByFederatedIdp"
)

func PossibleValuesForFederatedIdpMfaBehavior() []string {
	return []string{
		string(FederatedIdpMfaBehavior_AcceptIfMfaDoneByFederatedIdp),
		string(FederatedIdpMfaBehavior_EnforceMfaByFederatedIdp),
		string(FederatedIdpMfaBehavior_RejectMfaByFederatedIdp),
	}
}

func (s *FederatedIdpMfaBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFederatedIdpMfaBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFederatedIdpMfaBehavior(input string) (*FederatedIdpMfaBehavior, error) {
	vals := map[string]FederatedIdpMfaBehavior{
		"acceptifmfadonebyfederatedidp": FederatedIdpMfaBehavior_AcceptIfMfaDoneByFederatedIdp,
		"enforcemfabyfederatedidp":      FederatedIdpMfaBehavior_EnforceMfaByFederatedIdp,
		"rejectmfabyfederatedidp":       FederatedIdpMfaBehavior_RejectMfaByFederatedIdp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FederatedIdpMfaBehavior(input)
	return &out, nil
}
