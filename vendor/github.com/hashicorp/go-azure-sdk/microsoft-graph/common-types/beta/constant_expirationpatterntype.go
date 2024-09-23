package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExpirationPatternType string

const (
	ExpirationPatternType_AfterDateTime ExpirationPatternType = "afterDateTime"
	ExpirationPatternType_AfterDuration ExpirationPatternType = "afterDuration"
	ExpirationPatternType_NoExpiration  ExpirationPatternType = "noExpiration"
	ExpirationPatternType_NotSpecified  ExpirationPatternType = "notSpecified"
)

func PossibleValuesForExpirationPatternType() []string {
	return []string{
		string(ExpirationPatternType_AfterDateTime),
		string(ExpirationPatternType_AfterDuration),
		string(ExpirationPatternType_NoExpiration),
		string(ExpirationPatternType_NotSpecified),
	}
}

func (s *ExpirationPatternType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExpirationPatternType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExpirationPatternType(input string) (*ExpirationPatternType, error) {
	vals := map[string]ExpirationPatternType{
		"afterdatetime": ExpirationPatternType_AfterDateTime,
		"afterduration": ExpirationPatternType_AfterDuration,
		"noexpiration":  ExpirationPatternType_NoExpiration,
		"notspecified":  ExpirationPatternType_NotSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExpirationPatternType(input)
	return &out, nil
}
