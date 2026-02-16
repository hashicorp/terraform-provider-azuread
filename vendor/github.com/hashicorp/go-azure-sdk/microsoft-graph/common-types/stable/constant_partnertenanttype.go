package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerTenantType string

const (
	PartnerTenantType_BreadthPartner                          PartnerTenantType = "breadthPartner"
	PartnerTenantType_BreadthPartnerDelegatedAdmin            PartnerTenantType = "breadthPartnerDelegatedAdmin"
	PartnerTenantType_MicrosoftSupport                        PartnerTenantType = "microsoftSupport"
	PartnerTenantType_ResellerPartnerDelegatedAdmin           PartnerTenantType = "resellerPartnerDelegatedAdmin"
	PartnerTenantType_SyndicatePartner                        PartnerTenantType = "syndicatePartner"
	PartnerTenantType_ValueAddedResellerPartnerDelegatedAdmin PartnerTenantType = "valueAddedResellerPartnerDelegatedAdmin"
)

func PossibleValuesForPartnerTenantType() []string {
	return []string{
		string(PartnerTenantType_BreadthPartner),
		string(PartnerTenantType_BreadthPartnerDelegatedAdmin),
		string(PartnerTenantType_MicrosoftSupport),
		string(PartnerTenantType_ResellerPartnerDelegatedAdmin),
		string(PartnerTenantType_SyndicatePartner),
		string(PartnerTenantType_ValueAddedResellerPartnerDelegatedAdmin),
	}
}

func (s *PartnerTenantType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePartnerTenantType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePartnerTenantType(input string) (*PartnerTenantType, error) {
	vals := map[string]PartnerTenantType{
		"breadthpartner":                          PartnerTenantType_BreadthPartner,
		"breadthpartnerdelegatedadmin":            PartnerTenantType_BreadthPartnerDelegatedAdmin,
		"microsoftsupport":                        PartnerTenantType_MicrosoftSupport,
		"resellerpartnerdelegatedadmin":           PartnerTenantType_ResellerPartnerDelegatedAdmin,
		"syndicatepartner":                        PartnerTenantType_SyndicatePartner,
		"valueaddedresellerpartnerdelegatedadmin": PartnerTenantType_ValueAddedResellerPartnerDelegatedAdmin,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PartnerTenantType(input)
	return &out, nil
}
