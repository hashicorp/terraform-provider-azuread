package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadComplexity string

const (
	PayloadComplexity_High    PayloadComplexity = "high"
	PayloadComplexity_Low     PayloadComplexity = "low"
	PayloadComplexity_Medium  PayloadComplexity = "medium"
	PayloadComplexity_Unknown PayloadComplexity = "unknown"
)

func PossibleValuesForPayloadComplexity() []string {
	return []string{
		string(PayloadComplexity_High),
		string(PayloadComplexity_Low),
		string(PayloadComplexity_Medium),
		string(PayloadComplexity_Unknown),
	}
}

func (s *PayloadComplexity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadComplexity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadComplexity(input string) (*PayloadComplexity, error) {
	vals := map[string]PayloadComplexity{
		"high":    PayloadComplexity_High,
		"low":     PayloadComplexity_Low,
		"medium":  PayloadComplexity_Medium,
		"unknown": PayloadComplexity_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadComplexity(input)
	return &out, nil
}
