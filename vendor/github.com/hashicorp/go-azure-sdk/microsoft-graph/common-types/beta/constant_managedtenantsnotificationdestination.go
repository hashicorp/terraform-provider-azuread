package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsNotificationDestination string

const (
	ManagedTenantsNotificationDestination_Api   ManagedTenantsNotificationDestination = "api"
	ManagedTenantsNotificationDestination_Email ManagedTenantsNotificationDestination = "email"
	ManagedTenantsNotificationDestination_None  ManagedTenantsNotificationDestination = "none"
	ManagedTenantsNotificationDestination_Sms   ManagedTenantsNotificationDestination = "sms"
)

func PossibleValuesForManagedTenantsNotificationDestination() []string {
	return []string{
		string(ManagedTenantsNotificationDestination_Api),
		string(ManagedTenantsNotificationDestination_Email),
		string(ManagedTenantsNotificationDestination_None),
		string(ManagedTenantsNotificationDestination_Sms),
	}
}

func (s *ManagedTenantsNotificationDestination) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsNotificationDestination(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsNotificationDestination(input string) (*ManagedTenantsNotificationDestination, error) {
	vals := map[string]ManagedTenantsNotificationDestination{
		"api":   ManagedTenantsNotificationDestination_Api,
		"email": ManagedTenantsNotificationDestination_Email,
		"none":  ManagedTenantsNotificationDestination_None,
		"sms":   ManagedTenantsNotificationDestination_Sms,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsNotificationDestination(input)
	return &out, nil
}
