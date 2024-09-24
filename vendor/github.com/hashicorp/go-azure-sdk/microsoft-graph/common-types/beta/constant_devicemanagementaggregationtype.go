package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAggregationType string

const (
	DeviceManagementAggregationType_AffectedCloudPCCount      DeviceManagementAggregationType = "affectedCloudPcCount"
	DeviceManagementAggregationType_AffectedCloudPCPercentage DeviceManagementAggregationType = "affectedCloudPcPercentage"
	DeviceManagementAggregationType_Count                     DeviceManagementAggregationType = "count"
	DeviceManagementAggregationType_Percentage                DeviceManagementAggregationType = "percentage"
)

func PossibleValuesForDeviceManagementAggregationType() []string {
	return []string{
		string(DeviceManagementAggregationType_AffectedCloudPCCount),
		string(DeviceManagementAggregationType_AffectedCloudPCPercentage),
		string(DeviceManagementAggregationType_Count),
		string(DeviceManagementAggregationType_Percentage),
	}
}

func (s *DeviceManagementAggregationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAggregationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAggregationType(input string) (*DeviceManagementAggregationType, error) {
	vals := map[string]DeviceManagementAggregationType{
		"affectedcloudpccount":      DeviceManagementAggregationType_AffectedCloudPCCount,
		"affectedcloudpcpercentage": DeviceManagementAggregationType_AffectedCloudPCPercentage,
		"count":                     DeviceManagementAggregationType_Count,
		"percentage":                DeviceManagementAggregationType_Percentage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAggregationType(input)
	return &out, nil
}
