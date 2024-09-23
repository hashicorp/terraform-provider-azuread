package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAdditionalDataOptions string

const (
	SecurityAdditionalDataOptions_AllVersions SecurityAdditionalDataOptions = "allVersions"
	SecurityAdditionalDataOptions_LinkedFiles SecurityAdditionalDataOptions = "linkedFiles"
)

func PossibleValuesForSecurityAdditionalDataOptions() []string {
	return []string{
		string(SecurityAdditionalDataOptions_AllVersions),
		string(SecurityAdditionalDataOptions_LinkedFiles),
	}
}

func (s *SecurityAdditionalDataOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAdditionalDataOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAdditionalDataOptions(input string) (*SecurityAdditionalDataOptions, error) {
	vals := map[string]SecurityAdditionalDataOptions{
		"allversions": SecurityAdditionalDataOptions_AllVersions,
		"linkedfiles": SecurityAdditionalDataOptions_LinkedFiles,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAdditionalDataOptions(input)
	return &out, nil
}
