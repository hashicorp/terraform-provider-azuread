package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadDeliveryPlatform string

const (
	PayloadDeliveryPlatform_Email   PayloadDeliveryPlatform = "email"
	PayloadDeliveryPlatform_Sms     PayloadDeliveryPlatform = "sms"
	PayloadDeliveryPlatform_Teams   PayloadDeliveryPlatform = "teams"
	PayloadDeliveryPlatform_Unknown PayloadDeliveryPlatform = "unknown"
)

func PossibleValuesForPayloadDeliveryPlatform() []string {
	return []string{
		string(PayloadDeliveryPlatform_Email),
		string(PayloadDeliveryPlatform_Sms),
		string(PayloadDeliveryPlatform_Teams),
		string(PayloadDeliveryPlatform_Unknown),
	}
}

func (s *PayloadDeliveryPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadDeliveryPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadDeliveryPlatform(input string) (*PayloadDeliveryPlatform, error) {
	vals := map[string]PayloadDeliveryPlatform{
		"email":   PayloadDeliveryPlatform_Email,
		"sms":     PayloadDeliveryPlatform_Sms,
		"teams":   PayloadDeliveryPlatform_Teams,
		"unknown": PayloadDeliveryPlatform_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadDeliveryPlatform(input)
	return &out, nil
}
