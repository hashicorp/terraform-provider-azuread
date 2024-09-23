package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataAdditionalUserAttributes string

const (
	IndustryDataAdditionalUserAttributes_UserGradeLevel IndustryDataAdditionalUserAttributes = "userGradeLevel"
	IndustryDataAdditionalUserAttributes_UserNumber     IndustryDataAdditionalUserAttributes = "userNumber"
)

func PossibleValuesForIndustryDataAdditionalUserAttributes() []string {
	return []string{
		string(IndustryDataAdditionalUserAttributes_UserGradeLevel),
		string(IndustryDataAdditionalUserAttributes_UserNumber),
	}
}

func (s *IndustryDataAdditionalUserAttributes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataAdditionalUserAttributes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataAdditionalUserAttributes(input string) (*IndustryDataAdditionalUserAttributes, error) {
	vals := map[string]IndustryDataAdditionalUserAttributes{
		"usergradelevel": IndustryDataAdditionalUserAttributes_UserGradeLevel,
		"usernumber":     IndustryDataAdditionalUserAttributes_UserNumber,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataAdditionalUserAttributes(input)
	return &out, nil
}
