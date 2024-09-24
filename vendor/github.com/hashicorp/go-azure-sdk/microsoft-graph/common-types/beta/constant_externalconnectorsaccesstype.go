package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsAccessType string

const (
	ExternalConnectorsAccessType_Deny  ExternalConnectorsAccessType = "deny"
	ExternalConnectorsAccessType_Grant ExternalConnectorsAccessType = "grant"
)

func PossibleValuesForExternalConnectorsAccessType() []string {
	return []string{
		string(ExternalConnectorsAccessType_Deny),
		string(ExternalConnectorsAccessType_Grant),
	}
}

func (s *ExternalConnectorsAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsAccessType(input string) (*ExternalConnectorsAccessType, error) {
	vals := map[string]ExternalConnectorsAccessType{
		"deny":  ExternalConnectorsAccessType_Deny,
		"grant": ExternalConnectorsAccessType_Grant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsAccessType(input)
	return &out, nil
}
