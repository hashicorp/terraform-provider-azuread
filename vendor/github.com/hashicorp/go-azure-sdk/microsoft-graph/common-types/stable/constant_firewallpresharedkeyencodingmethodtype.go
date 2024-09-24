package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FirewallPreSharedKeyEncodingMethodType string

const (
	FirewallPreSharedKeyEncodingMethodType_DeviceDefault FirewallPreSharedKeyEncodingMethodType = "deviceDefault"
	FirewallPreSharedKeyEncodingMethodType_None          FirewallPreSharedKeyEncodingMethodType = "none"
	FirewallPreSharedKeyEncodingMethodType_UtF8          FirewallPreSharedKeyEncodingMethodType = "utF8"
)

func PossibleValuesForFirewallPreSharedKeyEncodingMethodType() []string {
	return []string{
		string(FirewallPreSharedKeyEncodingMethodType_DeviceDefault),
		string(FirewallPreSharedKeyEncodingMethodType_None),
		string(FirewallPreSharedKeyEncodingMethodType_UtF8),
	}
}

func (s *FirewallPreSharedKeyEncodingMethodType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFirewallPreSharedKeyEncodingMethodType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFirewallPreSharedKeyEncodingMethodType(input string) (*FirewallPreSharedKeyEncodingMethodType, error) {
	vals := map[string]FirewallPreSharedKeyEncodingMethodType{
		"devicedefault": FirewallPreSharedKeyEncodingMethodType_DeviceDefault,
		"none":          FirewallPreSharedKeyEncodingMethodType_None,
		"utf8":          FirewallPreSharedKeyEncodingMethodType_UtF8,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FirewallPreSharedKeyEncodingMethodType(input)
	return &out, nil
}
