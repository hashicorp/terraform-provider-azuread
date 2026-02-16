package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsStatementEffect string

const (
	AwsStatementEffect_Allow AwsStatementEffect = "allow"
	AwsStatementEffect_Deny  AwsStatementEffect = "deny"
)

func PossibleValuesForAwsStatementEffect() []string {
	return []string{
		string(AwsStatementEffect_Allow),
		string(AwsStatementEffect_Deny),
	}
}

func (s *AwsStatementEffect) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAwsStatementEffect(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAwsStatementEffect(input string) (*AwsStatementEffect, error) {
	vals := map[string]AwsStatementEffect{
		"allow": AwsStatementEffect_Allow,
		"deny":  AwsStatementEffect_Deny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AwsStatementEffect(input)
	return &out, nil
}
