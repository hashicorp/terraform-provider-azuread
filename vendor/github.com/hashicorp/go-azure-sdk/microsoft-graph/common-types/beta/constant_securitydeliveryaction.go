package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeliveryAction string

const (
	SecurityDeliveryAction_Blocked         SecurityDeliveryAction = "blocked"
	SecurityDeliveryAction_Delivered       SecurityDeliveryAction = "delivered"
	SecurityDeliveryAction_DeliveredToJunk SecurityDeliveryAction = "deliveredToJunk"
	SecurityDeliveryAction_Replaced        SecurityDeliveryAction = "replaced"
	SecurityDeliveryAction_Unknown         SecurityDeliveryAction = "unknown"
)

func PossibleValuesForSecurityDeliveryAction() []string {
	return []string{
		string(SecurityDeliveryAction_Blocked),
		string(SecurityDeliveryAction_Delivered),
		string(SecurityDeliveryAction_DeliveredToJunk),
		string(SecurityDeliveryAction_Replaced),
		string(SecurityDeliveryAction_Unknown),
	}
}

func (s *SecurityDeliveryAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDeliveryAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDeliveryAction(input string) (*SecurityDeliveryAction, error) {
	vals := map[string]SecurityDeliveryAction{
		"blocked":         SecurityDeliveryAction_Blocked,
		"delivered":       SecurityDeliveryAction_Delivered,
		"deliveredtojunk": SecurityDeliveryAction_DeliveredToJunk,
		"replaced":        SecurityDeliveryAction_Replaced,
		"unknown":         SecurityDeliveryAction_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeliveryAction(input)
	return &out, nil
}
