package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConditionCategory string

const (
	DeviceManagementConditionCategory_AzureNetworkConnectionCheckFailures DeviceManagementConditionCategory = "azureNetworkConnectionCheckFailures"
	DeviceManagementConditionCategory_CloudPCConnectionErrors             DeviceManagementConditionCategory = "cloudPcConnectionErrors"
	DeviceManagementConditionCategory_CloudPCHostHealthCheckFailures      DeviceManagementConditionCategory = "cloudPcHostHealthCheckFailures"
	DeviceManagementConditionCategory_CloudPCInGracePeriod                DeviceManagementConditionCategory = "cloudPcInGracePeriod"
	DeviceManagementConditionCategory_CloudPCZoneOutage                   DeviceManagementConditionCategory = "cloudPcZoneOutage"
	DeviceManagementConditionCategory_FrontlineInsufficientLicenses       DeviceManagementConditionCategory = "frontlineInsufficientLicenses"
	DeviceManagementConditionCategory_ImageUploadFailures                 DeviceManagementConditionCategory = "imageUploadFailures"
	DeviceManagementConditionCategory_ProvisionFailures                   DeviceManagementConditionCategory = "provisionFailures"
)

func PossibleValuesForDeviceManagementConditionCategory() []string {
	return []string{
		string(DeviceManagementConditionCategory_AzureNetworkConnectionCheckFailures),
		string(DeviceManagementConditionCategory_CloudPCConnectionErrors),
		string(DeviceManagementConditionCategory_CloudPCHostHealthCheckFailures),
		string(DeviceManagementConditionCategory_CloudPCInGracePeriod),
		string(DeviceManagementConditionCategory_CloudPCZoneOutage),
		string(DeviceManagementConditionCategory_FrontlineInsufficientLicenses),
		string(DeviceManagementConditionCategory_ImageUploadFailures),
		string(DeviceManagementConditionCategory_ProvisionFailures),
	}
}

func (s *DeviceManagementConditionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConditionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConditionCategory(input string) (*DeviceManagementConditionCategory, error) {
	vals := map[string]DeviceManagementConditionCategory{
		"azurenetworkconnectioncheckfailures": DeviceManagementConditionCategory_AzureNetworkConnectionCheckFailures,
		"cloudpcconnectionerrors":             DeviceManagementConditionCategory_CloudPCConnectionErrors,
		"cloudpchosthealthcheckfailures":      DeviceManagementConditionCategory_CloudPCHostHealthCheckFailures,
		"cloudpcingraceperiod":                DeviceManagementConditionCategory_CloudPCInGracePeriod,
		"cloudpczoneoutage":                   DeviceManagementConditionCategory_CloudPCZoneOutage,
		"frontlineinsufficientlicenses":       DeviceManagementConditionCategory_FrontlineInsufficientLicenses,
		"imageuploadfailures":                 DeviceManagementConditionCategory_ImageUploadFailures,
		"provisionfailures":                   DeviceManagementConditionCategory_ProvisionFailures,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConditionCategory(input)
	return &out, nil
}
