package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedTargetScope string

const (
	AllowedTargetScope_AllConfiguredConnectedOrganizationUsers AllowedTargetScope = "allConfiguredConnectedOrganizationUsers"
	AllowedTargetScope_AllDirectoryServicePrincipals           AllowedTargetScope = "allDirectoryServicePrincipals"
	AllowedTargetScope_AllDirectoryUsers                       AllowedTargetScope = "allDirectoryUsers"
	AllowedTargetScope_AllExternalUsers                        AllowedTargetScope = "allExternalUsers"
	AllowedTargetScope_AllMemberUsers                          AllowedTargetScope = "allMemberUsers"
	AllowedTargetScope_NotSpecified                            AllowedTargetScope = "notSpecified"
	AllowedTargetScope_SpecificConnectedOrganizationUsers      AllowedTargetScope = "specificConnectedOrganizationUsers"
	AllowedTargetScope_SpecificDirectoryServicePrincipals      AllowedTargetScope = "specificDirectoryServicePrincipals"
	AllowedTargetScope_SpecificDirectoryUsers                  AllowedTargetScope = "specificDirectoryUsers"
)

func PossibleValuesForAllowedTargetScope() []string {
	return []string{
		string(AllowedTargetScope_AllConfiguredConnectedOrganizationUsers),
		string(AllowedTargetScope_AllDirectoryServicePrincipals),
		string(AllowedTargetScope_AllDirectoryUsers),
		string(AllowedTargetScope_AllExternalUsers),
		string(AllowedTargetScope_AllMemberUsers),
		string(AllowedTargetScope_NotSpecified),
		string(AllowedTargetScope_SpecificConnectedOrganizationUsers),
		string(AllowedTargetScope_SpecificDirectoryServicePrincipals),
		string(AllowedTargetScope_SpecificDirectoryUsers),
	}
}

func (s *AllowedTargetScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAllowedTargetScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAllowedTargetScope(input string) (*AllowedTargetScope, error) {
	vals := map[string]AllowedTargetScope{
		"allconfiguredconnectedorganizationusers": AllowedTargetScope_AllConfiguredConnectedOrganizationUsers,
		"alldirectoryserviceprincipals":           AllowedTargetScope_AllDirectoryServicePrincipals,
		"alldirectoryusers":                       AllowedTargetScope_AllDirectoryUsers,
		"allexternalusers":                        AllowedTargetScope_AllExternalUsers,
		"allmemberusers":                          AllowedTargetScope_AllMemberUsers,
		"notspecified":                            AllowedTargetScope_NotSpecified,
		"specificconnectedorganizationusers":      AllowedTargetScope_SpecificConnectedOrganizationUsers,
		"specificdirectoryserviceprincipals":      AllowedTargetScope_SpecificDirectoryServicePrincipals,
		"specificdirectoryusers":                  AllowedTargetScope_SpecificDirectoryUsers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllowedTargetScope(input)
	return &out, nil
}
