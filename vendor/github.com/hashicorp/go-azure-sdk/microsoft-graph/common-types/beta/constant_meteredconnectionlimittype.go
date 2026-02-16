package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeteredConnectionLimitType string

const (
	MeteredConnectionLimitType_Fixed        MeteredConnectionLimitType = "fixed"
	MeteredConnectionLimitType_Unrestricted MeteredConnectionLimitType = "unrestricted"
	MeteredConnectionLimitType_Variable     MeteredConnectionLimitType = "variable"
)

func PossibleValuesForMeteredConnectionLimitType() []string {
	return []string{
		string(MeteredConnectionLimitType_Fixed),
		string(MeteredConnectionLimitType_Unrestricted),
		string(MeteredConnectionLimitType_Variable),
	}
}

func (s *MeteredConnectionLimitType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMeteredConnectionLimitType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMeteredConnectionLimitType(input string) (*MeteredConnectionLimitType, error) {
	vals := map[string]MeteredConnectionLimitType{
		"fixed":        MeteredConnectionLimitType_Fixed,
		"unrestricted": MeteredConnectionLimitType_Unrestricted,
		"variable":     MeteredConnectionLimitType_Variable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MeteredConnectionLimitType(input)
	return &out, nil
}
