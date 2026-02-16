package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEntityType string

const (
	SecurityEntityType_IPAddress   SecurityEntityType = "ipAddress"
	SecurityEntityType_MachineName SecurityEntityType = "machineName"
	SecurityEntityType_Other       SecurityEntityType = "other"
	SecurityEntityType_Unknown     SecurityEntityType = "unknown"
	SecurityEntityType_UserName    SecurityEntityType = "userName"
)

func PossibleValuesForSecurityEntityType() []string {
	return []string{
		string(SecurityEntityType_IPAddress),
		string(SecurityEntityType_MachineName),
		string(SecurityEntityType_Other),
		string(SecurityEntityType_Unknown),
		string(SecurityEntityType_UserName),
	}
}

func (s *SecurityEntityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEntityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEntityType(input string) (*SecurityEntityType, error) {
	vals := map[string]SecurityEntityType{
		"ipaddress":   SecurityEntityType_IPAddress,
		"machinename": SecurityEntityType_MachineName,
		"other":       SecurityEntityType_Other,
		"unknown":     SecurityEntityType_Unknown,
		"username":    SecurityEntityType_UserName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEntityType(input)
	return &out, nil
}
