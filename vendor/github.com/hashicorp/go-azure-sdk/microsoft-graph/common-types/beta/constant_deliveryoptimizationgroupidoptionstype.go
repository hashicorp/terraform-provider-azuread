package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeliveryOptimizationGroupIdOptionsType string

const (
	DeliveryOptimizationGroupIdOptionsType_AdSite                 DeliveryOptimizationGroupIdOptionsType = "adSite"
	DeliveryOptimizationGroupIdOptionsType_AuthenticatedDomainSid DeliveryOptimizationGroupIdOptionsType = "authenticatedDomainSid"
	DeliveryOptimizationGroupIdOptionsType_DhcpUserOption         DeliveryOptimizationGroupIdOptionsType = "dhcpUserOption"
	DeliveryOptimizationGroupIdOptionsType_DnsSuffix              DeliveryOptimizationGroupIdOptionsType = "dnsSuffix"
	DeliveryOptimizationGroupIdOptionsType_NotConfigured          DeliveryOptimizationGroupIdOptionsType = "notConfigured"
)

func PossibleValuesForDeliveryOptimizationGroupIdOptionsType() []string {
	return []string{
		string(DeliveryOptimizationGroupIdOptionsType_AdSite),
		string(DeliveryOptimizationGroupIdOptionsType_AuthenticatedDomainSid),
		string(DeliveryOptimizationGroupIdOptionsType_DhcpUserOption),
		string(DeliveryOptimizationGroupIdOptionsType_DnsSuffix),
		string(DeliveryOptimizationGroupIdOptionsType_NotConfigured),
	}
}

func (s *DeliveryOptimizationGroupIdOptionsType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeliveryOptimizationGroupIdOptionsType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeliveryOptimizationGroupIdOptionsType(input string) (*DeliveryOptimizationGroupIdOptionsType, error) {
	vals := map[string]DeliveryOptimizationGroupIdOptionsType{
		"adsite":                 DeliveryOptimizationGroupIdOptionsType_AdSite,
		"authenticateddomainsid": DeliveryOptimizationGroupIdOptionsType_AuthenticatedDomainSid,
		"dhcpuseroption":         DeliveryOptimizationGroupIdOptionsType_DhcpUserOption,
		"dnssuffix":              DeliveryOptimizationGroupIdOptionsType_DnsSuffix,
		"notconfigured":          DeliveryOptimizationGroupIdOptionsType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeliveryOptimizationGroupIdOptionsType(input)
	return &out, nil
}
