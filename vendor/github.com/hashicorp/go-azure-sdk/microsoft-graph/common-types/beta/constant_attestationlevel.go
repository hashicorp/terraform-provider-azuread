package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttestationLevel string

const (
	AttestationLevel_Attested    AttestationLevel = "attested"
	AttestationLevel_NotAttested AttestationLevel = "notAttested"
)

func PossibleValuesForAttestationLevel() []string {
	return []string{
		string(AttestationLevel_Attested),
		string(AttestationLevel_NotAttested),
	}
}

func (s *AttestationLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAttestationLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAttestationLevel(input string) (*AttestationLevel, error) {
	vals := map[string]AttestationLevel{
		"attested":    AttestationLevel_Attested,
		"notattested": AttestationLevel_NotAttested,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttestationLevel(input)
	return &out, nil
}
