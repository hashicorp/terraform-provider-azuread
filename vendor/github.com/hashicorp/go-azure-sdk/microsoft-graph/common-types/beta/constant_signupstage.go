package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignUpStage string

const (
	SignUpStage_AttributeCollectionAndValidation SignUpStage = "attributeCollectionAndValidation"
	SignUpStage_Consent                          SignUpStage = "consent"
	SignUpStage_CredentialCollection             SignUpStage = "credentialCollection"
	SignUpStage_CredentialFederation             SignUpStage = "credentialFederation"
	SignUpStage_CredentialValidation             SignUpStage = "credentialValidation"
	SignUpStage_TenantConsent                    SignUpStage = "tenantConsent"
	SignUpStage_UserCreation                     SignUpStage = "userCreation"
)

func PossibleValuesForSignUpStage() []string {
	return []string{
		string(SignUpStage_AttributeCollectionAndValidation),
		string(SignUpStage_Consent),
		string(SignUpStage_CredentialCollection),
		string(SignUpStage_CredentialFederation),
		string(SignUpStage_CredentialValidation),
		string(SignUpStage_TenantConsent),
		string(SignUpStage_UserCreation),
	}
}

func (s *SignUpStage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignUpStage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignUpStage(input string) (*SignUpStage, error) {
	vals := map[string]SignUpStage{
		"attributecollectionandvalidation": SignUpStage_AttributeCollectionAndValidation,
		"consent":                          SignUpStage_Consent,
		"credentialcollection":             SignUpStage_CredentialCollection,
		"credentialfederation":             SignUpStage_CredentialFederation,
		"credentialvalidation":             SignUpStage_CredentialValidation,
		"tenantconsent":                    SignUpStage_TenantConsent,
		"usercreation":                     SignUpStage_UserCreation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignUpStage(input)
	return &out, nil
}
