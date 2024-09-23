package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsPolicyType string

const (
	AwsPolicyType_Custom AwsPolicyType = "custom"
	AwsPolicyType_System AwsPolicyType = "system"
)

func PossibleValuesForAwsPolicyType() []string {
	return []string{
		string(AwsPolicyType_Custom),
		string(AwsPolicyType_System),
	}
}

func (s *AwsPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAwsPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAwsPolicyType(input string) (*AwsPolicyType, error) {
	vals := map[string]AwsPolicyType{
		"custom": AwsPolicyType_Custom,
		"system": AwsPolicyType_System,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AwsPolicyType(input)
	return &out, nil
}
