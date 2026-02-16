package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KerberosSignOnMappingAttributeType string

const (
	KerberosSignOnMappingAttributeType_OnPremisesSAMAccountName        KerberosSignOnMappingAttributeType = "onPremisesSAMAccountName"
	KerberosSignOnMappingAttributeType_OnPremisesUserPrincipalName     KerberosSignOnMappingAttributeType = "onPremisesUserPrincipalName"
	KerberosSignOnMappingAttributeType_OnPremisesUserPrincipalUsername KerberosSignOnMappingAttributeType = "onPremisesUserPrincipalUsername"
	KerberosSignOnMappingAttributeType_UserPrincipalName               KerberosSignOnMappingAttributeType = "userPrincipalName"
	KerberosSignOnMappingAttributeType_UserPrincipalUsername           KerberosSignOnMappingAttributeType = "userPrincipalUsername"
)

func PossibleValuesForKerberosSignOnMappingAttributeType() []string {
	return []string{
		string(KerberosSignOnMappingAttributeType_OnPremisesSAMAccountName),
		string(KerberosSignOnMappingAttributeType_OnPremisesUserPrincipalName),
		string(KerberosSignOnMappingAttributeType_OnPremisesUserPrincipalUsername),
		string(KerberosSignOnMappingAttributeType_UserPrincipalName),
		string(KerberosSignOnMappingAttributeType_UserPrincipalUsername),
	}
}

func (s *KerberosSignOnMappingAttributeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKerberosSignOnMappingAttributeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKerberosSignOnMappingAttributeType(input string) (*KerberosSignOnMappingAttributeType, error) {
	vals := map[string]KerberosSignOnMappingAttributeType{
		"onpremisessamaccountname":        KerberosSignOnMappingAttributeType_OnPremisesSAMAccountName,
		"onpremisesuserprincipalname":     KerberosSignOnMappingAttributeType_OnPremisesUserPrincipalName,
		"onpremisesuserprincipalusername": KerberosSignOnMappingAttributeType_OnPremisesUserPrincipalUsername,
		"userprincipalname":               KerberosSignOnMappingAttributeType_UserPrincipalName,
		"userprincipalusername":           KerberosSignOnMappingAttributeType_UserPrincipalUsername,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KerberosSignOnMappingAttributeType(input)
	return &out, nil
}
