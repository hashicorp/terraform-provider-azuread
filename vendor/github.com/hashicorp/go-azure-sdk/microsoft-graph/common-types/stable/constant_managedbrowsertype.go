package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedBrowserType string

const (
	ManagedBrowserType_MicrosoftEdge ManagedBrowserType = "microsoftEdge"
	ManagedBrowserType_NotConfigured ManagedBrowserType = "notConfigured"
)

func PossibleValuesForManagedBrowserType() []string {
	return []string{
		string(ManagedBrowserType_MicrosoftEdge),
		string(ManagedBrowserType_NotConfigured),
	}
}

func (s *ManagedBrowserType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedBrowserType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedBrowserType(input string) (*ManagedBrowserType, error) {
	vals := map[string]ManagedBrowserType{
		"microsoftedge": ManagedBrowserType_MicrosoftEdge,
		"notconfigured": ManagedBrowserType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedBrowserType(input)
	return &out, nil
}
