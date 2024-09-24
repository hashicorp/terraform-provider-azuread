package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessAudienceReason string

const (
	ConditionalAccessAudienceReason_ConfidentialClientIdToken    ConditionalAccessAudienceReason = "confidentialClientIdToken"
	ConditionalAccessAudienceReason_ConfidentialClientNonIdToken ConditionalAccessAudienceReason = "confidentialClientNonIdToken"
	ConditionalAccessAudienceReason_DelegatedScope               ConditionalAccessAudienceReason = "delegatedScope"
	ConditionalAccessAudienceReason_FirstPartyResourceDefault    ConditionalAccessAudienceReason = "firstPartyResourceDefault"
	ConditionalAccessAudienceReason_None                         ConditionalAccessAudienceReason = "none"
	ConditionalAccessAudienceReason_ResourceMapping              ConditionalAccessAudienceReason = "resourceMapping"
	ConditionalAccessAudienceReason_ResourceMappingDefault       ConditionalAccessAudienceReason = "resourceMappingDefault"
	ConditionalAccessAudienceReason_ResourcelessRequest          ConditionalAccessAudienceReason = "resourcelessRequest"
	ConditionalAccessAudienceReason_ScopeMapping                 ConditionalAccessAudienceReason = "scopeMapping"
	ConditionalAccessAudienceReason_ScopeMappingDefault          ConditionalAccessAudienceReason = "scopeMappingDefault"
	ConditionalAccessAudienceReason_ThirdPartyResourceDefault    ConditionalAccessAudienceReason = "thirdPartyResourceDefault"
)

func PossibleValuesForConditionalAccessAudienceReason() []string {
	return []string{
		string(ConditionalAccessAudienceReason_ConfidentialClientIdToken),
		string(ConditionalAccessAudienceReason_ConfidentialClientNonIdToken),
		string(ConditionalAccessAudienceReason_DelegatedScope),
		string(ConditionalAccessAudienceReason_FirstPartyResourceDefault),
		string(ConditionalAccessAudienceReason_None),
		string(ConditionalAccessAudienceReason_ResourceMapping),
		string(ConditionalAccessAudienceReason_ResourceMappingDefault),
		string(ConditionalAccessAudienceReason_ResourcelessRequest),
		string(ConditionalAccessAudienceReason_ScopeMapping),
		string(ConditionalAccessAudienceReason_ScopeMappingDefault),
		string(ConditionalAccessAudienceReason_ThirdPartyResourceDefault),
	}
}

func (s *ConditionalAccessAudienceReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConditionalAccessAudienceReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConditionalAccessAudienceReason(input string) (*ConditionalAccessAudienceReason, error) {
	vals := map[string]ConditionalAccessAudienceReason{
		"confidentialclientidtoken":    ConditionalAccessAudienceReason_ConfidentialClientIdToken,
		"confidentialclientnonidtoken": ConditionalAccessAudienceReason_ConfidentialClientNonIdToken,
		"delegatedscope":               ConditionalAccessAudienceReason_DelegatedScope,
		"firstpartyresourcedefault":    ConditionalAccessAudienceReason_FirstPartyResourceDefault,
		"none":                         ConditionalAccessAudienceReason_None,
		"resourcemapping":              ConditionalAccessAudienceReason_ResourceMapping,
		"resourcemappingdefault":       ConditionalAccessAudienceReason_ResourceMappingDefault,
		"resourcelessrequest":          ConditionalAccessAudienceReason_ResourcelessRequest,
		"scopemapping":                 ConditionalAccessAudienceReason_ScopeMapping,
		"scopemappingdefault":          ConditionalAccessAudienceReason_ScopeMappingDefault,
		"thirdpartyresourcedefault":    ConditionalAccessAudienceReason_ThirdPartyResourceDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessAudienceReason(input)
	return &out, nil
}
