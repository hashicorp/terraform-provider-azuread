package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryAdditionalDataOptions string

const (
	EdiscoveryAdditionalDataOptions_AllVersions EdiscoveryAdditionalDataOptions = "allVersions"
	EdiscoveryAdditionalDataOptions_LinkedFiles EdiscoveryAdditionalDataOptions = "linkedFiles"
)

func PossibleValuesForEdiscoveryAdditionalDataOptions() []string {
	return []string{
		string(EdiscoveryAdditionalDataOptions_AllVersions),
		string(EdiscoveryAdditionalDataOptions_LinkedFiles),
	}
}

func (s *EdiscoveryAdditionalDataOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryAdditionalDataOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryAdditionalDataOptions(input string) (*EdiscoveryAdditionalDataOptions, error) {
	vals := map[string]EdiscoveryAdditionalDataOptions{
		"allversions": EdiscoveryAdditionalDataOptions_AllVersions,
		"linkedfiles": EdiscoveryAdditionalDataOptions_LinkedFiles,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryAdditionalDataOptions(input)
	return &out, nil
}
