package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessIntentCategory string

const (
	NetworkaccessIntentCategory_Collection              NetworkaccessIntentCategory = "collection"
	NetworkaccessIntentCategory_CommandAndControl       NetworkaccessIntentCategory = "commandAndControl"
	NetworkaccessIntentCategory_CredentialAccess        NetworkaccessIntentCategory = "credentialAccess"
	NetworkaccessIntentCategory_DefenseEvasion          NetworkaccessIntentCategory = "defenseEvasion"
	NetworkaccessIntentCategory_Discovery               NetworkaccessIntentCategory = "discovery"
	NetworkaccessIntentCategory_Evasion                 NetworkaccessIntentCategory = "evasion"
	NetworkaccessIntentCategory_Execution               NetworkaccessIntentCategory = "execution"
	NetworkaccessIntentCategory_Exfiltration            NetworkaccessIntentCategory = "exfiltration"
	NetworkaccessIntentCategory_Impact                  NetworkaccessIntentCategory = "impact"
	NetworkaccessIntentCategory_ImpairProcessControl    NetworkaccessIntentCategory = "impairProcessControl"
	NetworkaccessIntentCategory_InhibitResponseFunction NetworkaccessIntentCategory = "inhibitResponseFunction"
	NetworkaccessIntentCategory_InitialAccess           NetworkaccessIntentCategory = "initialAccess"
	NetworkaccessIntentCategory_LateralMovement         NetworkaccessIntentCategory = "lateralMovement"
	NetworkaccessIntentCategory_Persistence             NetworkaccessIntentCategory = "persistence"
	NetworkaccessIntentCategory_PrivilegeEscalation     NetworkaccessIntentCategory = "privilegeEscalation"
	NetworkaccessIntentCategory_Reconnaissance          NetworkaccessIntentCategory = "reconnaissance"
	NetworkaccessIntentCategory_ResourceDevelopment     NetworkaccessIntentCategory = "resourceDevelopment"
)

func PossibleValuesForNetworkaccessIntentCategory() []string {
	return []string{
		string(NetworkaccessIntentCategory_Collection),
		string(NetworkaccessIntentCategory_CommandAndControl),
		string(NetworkaccessIntentCategory_CredentialAccess),
		string(NetworkaccessIntentCategory_DefenseEvasion),
		string(NetworkaccessIntentCategory_Discovery),
		string(NetworkaccessIntentCategory_Evasion),
		string(NetworkaccessIntentCategory_Execution),
		string(NetworkaccessIntentCategory_Exfiltration),
		string(NetworkaccessIntentCategory_Impact),
		string(NetworkaccessIntentCategory_ImpairProcessControl),
		string(NetworkaccessIntentCategory_InhibitResponseFunction),
		string(NetworkaccessIntentCategory_InitialAccess),
		string(NetworkaccessIntentCategory_LateralMovement),
		string(NetworkaccessIntentCategory_Persistence),
		string(NetworkaccessIntentCategory_PrivilegeEscalation),
		string(NetworkaccessIntentCategory_Reconnaissance),
		string(NetworkaccessIntentCategory_ResourceDevelopment),
	}
}

func (s *NetworkaccessIntentCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessIntentCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessIntentCategory(input string) (*NetworkaccessIntentCategory, error) {
	vals := map[string]NetworkaccessIntentCategory{
		"collection":              NetworkaccessIntentCategory_Collection,
		"commandandcontrol":       NetworkaccessIntentCategory_CommandAndControl,
		"credentialaccess":        NetworkaccessIntentCategory_CredentialAccess,
		"defenseevasion":          NetworkaccessIntentCategory_DefenseEvasion,
		"discovery":               NetworkaccessIntentCategory_Discovery,
		"evasion":                 NetworkaccessIntentCategory_Evasion,
		"execution":               NetworkaccessIntentCategory_Execution,
		"exfiltration":            NetworkaccessIntentCategory_Exfiltration,
		"impact":                  NetworkaccessIntentCategory_Impact,
		"impairprocesscontrol":    NetworkaccessIntentCategory_ImpairProcessControl,
		"inhibitresponsefunction": NetworkaccessIntentCategory_InhibitResponseFunction,
		"initialaccess":           NetworkaccessIntentCategory_InitialAccess,
		"lateralmovement":         NetworkaccessIntentCategory_LateralMovement,
		"persistence":             NetworkaccessIntentCategory_Persistence,
		"privilegeescalation":     NetworkaccessIntentCategory_PrivilegeEscalation,
		"reconnaissance":          NetworkaccessIntentCategory_Reconnaissance,
		"resourcedevelopment":     NetworkaccessIntentCategory_ResourceDevelopment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessIntentCategory(input)
	return &out, nil
}
