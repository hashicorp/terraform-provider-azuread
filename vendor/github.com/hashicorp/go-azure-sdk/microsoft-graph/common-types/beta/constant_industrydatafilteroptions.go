package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataFilterOptions string

const (
	IndustryDataFilterOptions_OrgExternalId IndustryDataFilterOptions = "orgExternalId"
)

func PossibleValuesForIndustryDataFilterOptions() []string {
	return []string{
		string(IndustryDataFilterOptions_OrgExternalId),
	}
}

func (s *IndustryDataFilterOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIndustryDataFilterOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIndustryDataFilterOptions(input string) (*IndustryDataFilterOptions, error) {
	vals := map[string]IndustryDataFilterOptions{
		"orgexternalid": IndustryDataFilterOptions_OrgExternalId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IndustryDataFilterOptions(input)
	return &out, nil
}
