package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppNotificationRestriction string

const (
	ManagedAppNotificationRestriction_Allow                   ManagedAppNotificationRestriction = "allow"
	ManagedAppNotificationRestriction_Block                   ManagedAppNotificationRestriction = "block"
	ManagedAppNotificationRestriction_BlockOrganizationalData ManagedAppNotificationRestriction = "blockOrganizationalData"
)

func PossibleValuesForManagedAppNotificationRestriction() []string {
	return []string{
		string(ManagedAppNotificationRestriction_Allow),
		string(ManagedAppNotificationRestriction_Block),
		string(ManagedAppNotificationRestriction_BlockOrganizationalData),
	}
}

func (s *ManagedAppNotificationRestriction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppNotificationRestriction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppNotificationRestriction(input string) (*ManagedAppNotificationRestriction, error) {
	vals := map[string]ManagedAppNotificationRestriction{
		"allow":                   ManagedAppNotificationRestriction_Allow,
		"block":                   ManagedAppNotificationRestriction_Block,
		"blockorganizationaldata": ManagedAppNotificationRestriction_BlockOrganizationalData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppNotificationRestriction(input)
	return &out, nil
}
