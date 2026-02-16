package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectAlternativeNameType string

const (
	SubjectAlternativeNameType_CustomAzureADAttribute      SubjectAlternativeNameType = "customAzureADAttribute"
	SubjectAlternativeNameType_DomainNameService           SubjectAlternativeNameType = "domainNameService"
	SubjectAlternativeNameType_EmailAddress                SubjectAlternativeNameType = "emailAddress"
	SubjectAlternativeNameType_None                        SubjectAlternativeNameType = "none"
	SubjectAlternativeNameType_UniversalResourceIdentifier SubjectAlternativeNameType = "universalResourceIdentifier"
	SubjectAlternativeNameType_UserPrincipalName           SubjectAlternativeNameType = "userPrincipalName"
)

func PossibleValuesForSubjectAlternativeNameType() []string {
	return []string{
		string(SubjectAlternativeNameType_CustomAzureADAttribute),
		string(SubjectAlternativeNameType_DomainNameService),
		string(SubjectAlternativeNameType_EmailAddress),
		string(SubjectAlternativeNameType_None),
		string(SubjectAlternativeNameType_UniversalResourceIdentifier),
		string(SubjectAlternativeNameType_UserPrincipalName),
	}
}

func (s *SubjectAlternativeNameType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubjectAlternativeNameType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubjectAlternativeNameType(input string) (*SubjectAlternativeNameType, error) {
	vals := map[string]SubjectAlternativeNameType{
		"customazureadattribute":      SubjectAlternativeNameType_CustomAzureADAttribute,
		"domainnameservice":           SubjectAlternativeNameType_DomainNameService,
		"emailaddress":                SubjectAlternativeNameType_EmailAddress,
		"none":                        SubjectAlternativeNameType_None,
		"universalresourceidentifier": SubjectAlternativeNameType_UniversalResourceIdentifier,
		"userprincipalname":           SubjectAlternativeNameType_UserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubjectAlternativeNameType(input)
	return &out, nil
}
