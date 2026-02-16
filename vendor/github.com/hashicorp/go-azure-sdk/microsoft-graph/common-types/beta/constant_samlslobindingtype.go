package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SamlSLOBindingType string

const (
	SamlSLOBindingType_HttpPost     SamlSLOBindingType = "httpPost"
	SamlSLOBindingType_HttpRedirect SamlSLOBindingType = "httpRedirect"
)

func PossibleValuesForSamlSLOBindingType() []string {
	return []string{
		string(SamlSLOBindingType_HttpPost),
		string(SamlSLOBindingType_HttpRedirect),
	}
}

func (s *SamlSLOBindingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSamlSLOBindingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSamlSLOBindingType(input string) (*SamlSLOBindingType, error) {
	vals := map[string]SamlSLOBindingType{
		"httppost":     SamlSLOBindingType_HttpPost,
		"httpredirect": SamlSLOBindingType_HttpRedirect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SamlSLOBindingType(input)
	return &out, nil
}
