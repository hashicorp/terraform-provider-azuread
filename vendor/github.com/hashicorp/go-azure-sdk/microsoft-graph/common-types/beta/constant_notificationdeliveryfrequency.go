package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationDeliveryFrequency string

const (
	NotificationDeliveryFrequency_BiWeekly NotificationDeliveryFrequency = "biWeekly"
	NotificationDeliveryFrequency_Unknown  NotificationDeliveryFrequency = "unknown"
	NotificationDeliveryFrequency_Weekly   NotificationDeliveryFrequency = "weekly"
)

func PossibleValuesForNotificationDeliveryFrequency() []string {
	return []string{
		string(NotificationDeliveryFrequency_BiWeekly),
		string(NotificationDeliveryFrequency_Unknown),
		string(NotificationDeliveryFrequency_Weekly),
	}
}

func (s *NotificationDeliveryFrequency) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNotificationDeliveryFrequency(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNotificationDeliveryFrequency(input string) (*NotificationDeliveryFrequency, error) {
	vals := map[string]NotificationDeliveryFrequency{
		"biweekly": NotificationDeliveryFrequency_BiWeekly,
		"unknown":  NotificationDeliveryFrequency_Unknown,
		"weekly":   NotificationDeliveryFrequency_Weekly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NotificationDeliveryFrequency(input)
	return &out, nil
}
