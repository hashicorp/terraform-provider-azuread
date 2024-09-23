package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAttributeCollectionInputType string

const (
	AuthenticationAttributeCollectionInputType_Boolean             AuthenticationAttributeCollectionInputType = "boolean"
	AuthenticationAttributeCollectionInputType_CheckboxMultiSelect AuthenticationAttributeCollectionInputType = "checkboxMultiSelect"
	AuthenticationAttributeCollectionInputType_RadioSingleSelect   AuthenticationAttributeCollectionInputType = "radioSingleSelect"
	AuthenticationAttributeCollectionInputType_Text                AuthenticationAttributeCollectionInputType = "text"
)

func PossibleValuesForAuthenticationAttributeCollectionInputType() []string {
	return []string{
		string(AuthenticationAttributeCollectionInputType_Boolean),
		string(AuthenticationAttributeCollectionInputType_CheckboxMultiSelect),
		string(AuthenticationAttributeCollectionInputType_RadioSingleSelect),
		string(AuthenticationAttributeCollectionInputType_Text),
	}
}

func (s *AuthenticationAttributeCollectionInputType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationAttributeCollectionInputType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationAttributeCollectionInputType(input string) (*AuthenticationAttributeCollectionInputType, error) {
	vals := map[string]AuthenticationAttributeCollectionInputType{
		"boolean":             AuthenticationAttributeCollectionInputType_Boolean,
		"checkboxmultiselect": AuthenticationAttributeCollectionInputType_CheckboxMultiSelect,
		"radiosingleselect":   AuthenticationAttributeCollectionInputType_RadioSingleSelect,
		"text":                AuthenticationAttributeCollectionInputType_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationAttributeCollectionInputType(input)
	return &out, nil
}
