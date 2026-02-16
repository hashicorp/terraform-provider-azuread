package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventStatus string

const (
	VirtualEventStatus_Canceled  VirtualEventStatus = "canceled"
	VirtualEventStatus_Draft     VirtualEventStatus = "draft"
	VirtualEventStatus_Published VirtualEventStatus = "published"
)

func PossibleValuesForVirtualEventStatus() []string {
	return []string{
		string(VirtualEventStatus_Canceled),
		string(VirtualEventStatus_Draft),
		string(VirtualEventStatus_Published),
	}
}

func (s *VirtualEventStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVirtualEventStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVirtualEventStatus(input string) (*VirtualEventStatus, error) {
	vals := map[string]VirtualEventStatus{
		"canceled":  VirtualEventStatus_Canceled,
		"draft":     VirtualEventStatus_Draft,
		"published": VirtualEventStatus_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VirtualEventStatus(input)
	return &out, nil
}
