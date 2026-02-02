package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RootDomains string

const (
	RootDomains_All                              RootDomains = "all"
	RootDomains_AllFederated                     RootDomains = "allFederated"
	RootDomains_AllManaged                       RootDomains = "allManaged"
	RootDomains_AllManagedAndEnumeratedFederated RootDomains = "allManagedAndEnumeratedFederated"
	RootDomains_Enumerated                       RootDomains = "enumerated"
	RootDomains_None                             RootDomains = "none"
)

func PossibleValuesForRootDomains() []string {
	return []string{
		string(RootDomains_All),
		string(RootDomains_AllFederated),
		string(RootDomains_AllManaged),
		string(RootDomains_AllManagedAndEnumeratedFederated),
		string(RootDomains_Enumerated),
		string(RootDomains_None),
	}
}

func (s *RootDomains) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRootDomains(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRootDomains(input string) (*RootDomains, error) {
	vals := map[string]RootDomains{
		"all":                              RootDomains_All,
		"allfederated":                     RootDomains_AllFederated,
		"allmanaged":                       RootDomains_AllManaged,
		"allmanagedandenumeratedfederated": RootDomains_AllManagedAndEnumeratedFederated,
		"enumerated":                       RootDomains_Enumerated,
		"none":                             RootDomains_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RootDomains(input)
	return &out, nil
}
