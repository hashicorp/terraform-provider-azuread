package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryOptimizationRestrictPeerSelectionByOptions string

const (
	DeliveryOptimizationRestrictPeerSelectionByOptions_NotConfigured DeliveryOptimizationRestrictPeerSelectionByOptions = "notConfigured"
	DeliveryOptimizationRestrictPeerSelectionByOptions_SubnetMask    DeliveryOptimizationRestrictPeerSelectionByOptions = "subnetMask"
)

func PossibleValuesForDeliveryOptimizationRestrictPeerSelectionByOptions() []string {
	return []string{
		string(DeliveryOptimizationRestrictPeerSelectionByOptions_NotConfigured),
		string(DeliveryOptimizationRestrictPeerSelectionByOptions_SubnetMask),
	}
}

func (s *DeliveryOptimizationRestrictPeerSelectionByOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeliveryOptimizationRestrictPeerSelectionByOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeliveryOptimizationRestrictPeerSelectionByOptions(input string) (*DeliveryOptimizationRestrictPeerSelectionByOptions, error) {
	vals := map[string]DeliveryOptimizationRestrictPeerSelectionByOptions{
		"notconfigured": DeliveryOptimizationRestrictPeerSelectionByOptions_NotConfigured,
		"subnetmask":    DeliveryOptimizationRestrictPeerSelectionByOptions_SubnetMask,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeliveryOptimizationRestrictPeerSelectionByOptions(input)
	return &out, nil
}
