package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorGroupType string

const (
	ConnectorGroupType_ApplicationProxy ConnectorGroupType = "applicationProxy"
)

func PossibleValuesForConnectorGroupType() []string {
	return []string{
		string(ConnectorGroupType_ApplicationProxy),
	}
}

func (s *ConnectorGroupType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectorGroupType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectorGroupType(input string) (*ConnectorGroupType, error) {
	vals := map[string]ConnectorGroupType{
		"applicationproxy": ConnectorGroupType_ApplicationProxy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectorGroupType(input)
	return &out, nil
}
