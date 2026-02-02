package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LifecycleEventType string

const (
	LifecycleEventType_Missed                  LifecycleEventType = "missed"
	LifecycleEventType_ReauthorizationRequired LifecycleEventType = "reauthorizationRequired"
	LifecycleEventType_SubscriptionRemoved     LifecycleEventType = "subscriptionRemoved"
)

func PossibleValuesForLifecycleEventType() []string {
	return []string{
		string(LifecycleEventType_Missed),
		string(LifecycleEventType_ReauthorizationRequired),
		string(LifecycleEventType_SubscriptionRemoved),
	}
}

func (s *LifecycleEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLifecycleEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLifecycleEventType(input string) (*LifecycleEventType, error) {
	vals := map[string]LifecycleEventType{
		"missed":                  LifecycleEventType_Missed,
		"reauthorizationrequired": LifecycleEventType_ReauthorizationRequired,
		"subscriptionremoved":     LifecycleEventType_SubscriptionRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LifecycleEventType(input)
	return &out, nil
}
