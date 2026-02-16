package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoverySourceType string

const (
	EdiscoverySourceType_Mailbox EdiscoverySourceType = "mailbox"
	EdiscoverySourceType_Site    EdiscoverySourceType = "site"
)

func PossibleValuesForEdiscoverySourceType() []string {
	return []string{
		string(EdiscoverySourceType_Mailbox),
		string(EdiscoverySourceType_Site),
	}
}

func (s *EdiscoverySourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoverySourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoverySourceType(input string) (*EdiscoverySourceType, error) {
	vals := map[string]EdiscoverySourceType{
		"mailbox": EdiscoverySourceType_Mailbox,
		"site":    EdiscoverySourceType_Site,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoverySourceType(input)
	return &out, nil
}
