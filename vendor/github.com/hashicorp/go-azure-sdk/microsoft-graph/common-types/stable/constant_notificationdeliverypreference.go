package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationDeliveryPreference string

const (
	NotificationDeliveryPreference_DeliverAfterCampaignEnd NotificationDeliveryPreference = "deliverAfterCampaignEnd"
	NotificationDeliveryPreference_DeliverImmedietly       NotificationDeliveryPreference = "deliverImmedietly"
	NotificationDeliveryPreference_Unknown                 NotificationDeliveryPreference = "unknown"
)

func PossibleValuesForNotificationDeliveryPreference() []string {
	return []string{
		string(NotificationDeliveryPreference_DeliverAfterCampaignEnd),
		string(NotificationDeliveryPreference_DeliverImmedietly),
		string(NotificationDeliveryPreference_Unknown),
	}
}

func (s *NotificationDeliveryPreference) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNotificationDeliveryPreference(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNotificationDeliveryPreference(input string) (*NotificationDeliveryPreference, error) {
	vals := map[string]NotificationDeliveryPreference{
		"deliveraftercampaignend": NotificationDeliveryPreference_DeliverAfterCampaignEnd,
		"deliverimmedietly":       NotificationDeliveryPreference_DeliverImmedietly,
		"unknown":                 NotificationDeliveryPreference_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NotificationDeliveryPreference(input)
	return &out, nil
}
