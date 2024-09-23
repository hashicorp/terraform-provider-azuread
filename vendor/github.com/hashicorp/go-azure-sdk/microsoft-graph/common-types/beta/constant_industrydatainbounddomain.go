package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataInboundDomain string

const (
	IndustryDataInboundDomain_EducationRostering IndustryDataInboundDomain = "educationRostering"
)

func PossibleValuesForIndustryDataInboundDomain() []string {
	return []string{
		string(IndustryDataInboundDomain_EducationRostering),
	}
}

func (s *IndustryDataInboundDomain) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataInboundDomain(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataInboundDomain(input string) (*IndustryDataInboundDomain, error) {
	vals := map[string]IndustryDataInboundDomain{
		"educationrostering": IndustryDataInboundDomain_EducationRostering,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataInboundDomain(input)
	return &out, nil
}
