package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrganizationAllowedAudiences string

const (
	OrganizationAllowedAudiences_Everyone               OrganizationAllowedAudiences = "everyone"
	OrganizationAllowedAudiences_FederatedOrganizations OrganizationAllowedAudiences = "federatedOrganizations"
	OrganizationAllowedAudiences_Me                     OrganizationAllowedAudiences = "me"
	OrganizationAllowedAudiences_Organization           OrganizationAllowedAudiences = "organization"
)

func PossibleValuesForOrganizationAllowedAudiences() []string {
	return []string{
		string(OrganizationAllowedAudiences_Everyone),
		string(OrganizationAllowedAudiences_FederatedOrganizations),
		string(OrganizationAllowedAudiences_Me),
		string(OrganizationAllowedAudiences_Organization),
	}
}

func (s *OrganizationAllowedAudiences) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOrganizationAllowedAudiences(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOrganizationAllowedAudiences(input string) (*OrganizationAllowedAudiences, error) {
	vals := map[string]OrganizationAllowedAudiences{
		"everyone":               OrganizationAllowedAudiences_Everyone,
		"federatedorganizations": OrganizationAllowedAudiences_FederatedOrganizations,
		"me":                     OrganizationAllowedAudiences_Me,
		"organization":           OrganizationAllowedAudiences_Organization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OrganizationAllowedAudiences(input)
	return &out, nil
}
