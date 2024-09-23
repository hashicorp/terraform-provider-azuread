package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceUpdateCategory string

const (
	ServiceUpdateCategory_PlanForChange     ServiceUpdateCategory = "planForChange"
	ServiceUpdateCategory_PreventOrFixIssue ServiceUpdateCategory = "preventOrFixIssue"
	ServiceUpdateCategory_StayInformed      ServiceUpdateCategory = "stayInformed"
)

func PossibleValuesForServiceUpdateCategory() []string {
	return []string{
		string(ServiceUpdateCategory_PlanForChange),
		string(ServiceUpdateCategory_PreventOrFixIssue),
		string(ServiceUpdateCategory_StayInformed),
	}
}

func (s *ServiceUpdateCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceUpdateCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceUpdateCategory(input string) (*ServiceUpdateCategory, error) {
	vals := map[string]ServiceUpdateCategory{
		"planforchange":     ServiceUpdateCategory_PlanForChange,
		"preventorfixissue": ServiceUpdateCategory_PreventOrFixIssue,
		"stayinformed":      ServiceUpdateCategory_StayInformed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceUpdateCategory(input)
	return &out, nil
}
