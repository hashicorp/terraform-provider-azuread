package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTeamsMessageDeliveryAction string

const (
	SecurityTeamsMessageDeliveryAction_Blocked         SecurityTeamsMessageDeliveryAction = "blocked"
	SecurityTeamsMessageDeliveryAction_Delivered       SecurityTeamsMessageDeliveryAction = "delivered"
	SecurityTeamsMessageDeliveryAction_DeliveredAsSpam SecurityTeamsMessageDeliveryAction = "deliveredAsSpam"
	SecurityTeamsMessageDeliveryAction_Replaced        SecurityTeamsMessageDeliveryAction = "replaced"
	SecurityTeamsMessageDeliveryAction_Unknown         SecurityTeamsMessageDeliveryAction = "unknown"
)

func PossibleValuesForSecurityTeamsMessageDeliveryAction() []string {
	return []string{
		string(SecurityTeamsMessageDeliveryAction_Blocked),
		string(SecurityTeamsMessageDeliveryAction_Delivered),
		string(SecurityTeamsMessageDeliveryAction_DeliveredAsSpam),
		string(SecurityTeamsMessageDeliveryAction_Replaced),
		string(SecurityTeamsMessageDeliveryAction_Unknown),
	}
}

func (s *SecurityTeamsMessageDeliveryAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityTeamsMessageDeliveryAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityTeamsMessageDeliveryAction(input string) (*SecurityTeamsMessageDeliveryAction, error) {
	vals := map[string]SecurityTeamsMessageDeliveryAction{
		"blocked":         SecurityTeamsMessageDeliveryAction_Blocked,
		"delivered":       SecurityTeamsMessageDeliveryAction_Delivered,
		"deliveredasspam": SecurityTeamsMessageDeliveryAction_DeliveredAsSpam,
		"replaced":        SecurityTeamsMessageDeliveryAction_Replaced,
		"unknown":         SecurityTeamsMessageDeliveryAction_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityTeamsMessageDeliveryAction(input)
	return &out, nil
}
