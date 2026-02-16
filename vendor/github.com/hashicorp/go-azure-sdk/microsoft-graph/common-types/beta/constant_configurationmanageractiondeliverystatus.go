package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerActionDeliveryStatus string

const (
	ConfigurationManagerActionDeliveryStatus_DeliveredToConnectorService       ConfigurationManagerActionDeliveryStatus = "deliveredToConnectorService"
	ConfigurationManagerActionDeliveryStatus_DeliveredToOnPremisesServer       ConfigurationManagerActionDeliveryStatus = "deliveredToOnPremisesServer"
	ConfigurationManagerActionDeliveryStatus_FailedToDeliverToConnectorService ConfigurationManagerActionDeliveryStatus = "failedToDeliverToConnectorService"
	ConfigurationManagerActionDeliveryStatus_PendingDelivery                   ConfigurationManagerActionDeliveryStatus = "pendingDelivery"
	ConfigurationManagerActionDeliveryStatus_Unknown                           ConfigurationManagerActionDeliveryStatus = "unknown"
)

func PossibleValuesForConfigurationManagerActionDeliveryStatus() []string {
	return []string{
		string(ConfigurationManagerActionDeliveryStatus_DeliveredToConnectorService),
		string(ConfigurationManagerActionDeliveryStatus_DeliveredToOnPremisesServer),
		string(ConfigurationManagerActionDeliveryStatus_FailedToDeliverToConnectorService),
		string(ConfigurationManagerActionDeliveryStatus_PendingDelivery),
		string(ConfigurationManagerActionDeliveryStatus_Unknown),
	}
}

func (s *ConfigurationManagerActionDeliveryStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConfigurationManagerActionDeliveryStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConfigurationManagerActionDeliveryStatus(input string) (*ConfigurationManagerActionDeliveryStatus, error) {
	vals := map[string]ConfigurationManagerActionDeliveryStatus{
		"deliveredtoconnectorservice":       ConfigurationManagerActionDeliveryStatus_DeliveredToConnectorService,
		"deliveredtoonpremisesserver":       ConfigurationManagerActionDeliveryStatus_DeliveredToOnPremisesServer,
		"failedtodelivertoconnectorservice": ConfigurationManagerActionDeliveryStatus_FailedToDeliverToConnectorService,
		"pendingdelivery":                   ConfigurationManagerActionDeliveryStatus_PendingDelivery,
		"unknown":                           ConfigurationManagerActionDeliveryStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationManagerActionDeliveryStatus(input)
	return &out, nil
}
