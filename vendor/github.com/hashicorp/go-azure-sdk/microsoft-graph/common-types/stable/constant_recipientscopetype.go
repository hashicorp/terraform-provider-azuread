package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecipientScopeType string

const (
	RecipientScopeType_External           RecipientScopeType = "external"
	RecipientScopeType_ExternalNonPartner RecipientScopeType = "externalNonPartner"
	RecipientScopeType_ExternalPartner    RecipientScopeType = "externalPartner"
	RecipientScopeType_Internal           RecipientScopeType = "internal"
	RecipientScopeType_None               RecipientScopeType = "none"
)

func PossibleValuesForRecipientScopeType() []string {
	return []string{
		string(RecipientScopeType_External),
		string(RecipientScopeType_ExternalNonPartner),
		string(RecipientScopeType_ExternalPartner),
		string(RecipientScopeType_Internal),
		string(RecipientScopeType_None),
	}
}

func (s *RecipientScopeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecipientScopeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecipientScopeType(input string) (*RecipientScopeType, error) {
	vals := map[string]RecipientScopeType{
		"external":           RecipientScopeType_External,
		"externalnonpartner": RecipientScopeType_ExternalNonPartner,
		"externalpartner":    RecipientScopeType_ExternalPartner,
		"internal":           RecipientScopeType_Internal,
		"none":               RecipientScopeType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecipientScopeType(input)
	return &out, nil
}
