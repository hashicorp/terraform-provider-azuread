package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertRuleTemplate string

const (
	DeviceManagementAlertRuleTemplate_CloudPCFrontlineConcurrencyScenario            DeviceManagementAlertRuleTemplate = "cloudPcFrontlineConcurrencyScenario"
	DeviceManagementAlertRuleTemplate_CloudPCFrontlineInsufficientLicensesScenario   DeviceManagementAlertRuleTemplate = "cloudPcFrontlineInsufficientLicensesScenario"
	DeviceManagementAlertRuleTemplate_CloudPCImageUploadScenario                     DeviceManagementAlertRuleTemplate = "cloudPcImageUploadScenario"
	DeviceManagementAlertRuleTemplate_CloudPCInGracePeriodScenario                   DeviceManagementAlertRuleTemplate = "cloudPcInGracePeriodScenario"
	DeviceManagementAlertRuleTemplate_CloudPCInaccessibleScenario                    DeviceManagementAlertRuleTemplate = "cloudPcInaccessibleScenario"
	DeviceManagementAlertRuleTemplate_CloudPCOnPremiseNetworkConnectionCheckScenario DeviceManagementAlertRuleTemplate = "cloudPcOnPremiseNetworkConnectionCheckScenario"
	DeviceManagementAlertRuleTemplate_CloudPCProvisionScenario                       DeviceManagementAlertRuleTemplate = "cloudPcProvisionScenario"
)

func PossibleValuesForDeviceManagementAlertRuleTemplate() []string {
	return []string{
		string(DeviceManagementAlertRuleTemplate_CloudPCFrontlineConcurrencyScenario),
		string(DeviceManagementAlertRuleTemplate_CloudPCFrontlineInsufficientLicensesScenario),
		string(DeviceManagementAlertRuleTemplate_CloudPCImageUploadScenario),
		string(DeviceManagementAlertRuleTemplate_CloudPCInGracePeriodScenario),
		string(DeviceManagementAlertRuleTemplate_CloudPCInaccessibleScenario),
		string(DeviceManagementAlertRuleTemplate_CloudPCOnPremiseNetworkConnectionCheckScenario),
		string(DeviceManagementAlertRuleTemplate_CloudPCProvisionScenario),
	}
}

func (s *DeviceManagementAlertRuleTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAlertRuleTemplate(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAlertRuleTemplate(input string) (*DeviceManagementAlertRuleTemplate, error) {
	vals := map[string]DeviceManagementAlertRuleTemplate{
		"cloudpcfrontlineconcurrencyscenario":            DeviceManagementAlertRuleTemplate_CloudPCFrontlineConcurrencyScenario,
		"cloudpcfrontlineinsufficientlicensesscenario":   DeviceManagementAlertRuleTemplate_CloudPCFrontlineInsufficientLicensesScenario,
		"cloudpcimageuploadscenario":                     DeviceManagementAlertRuleTemplate_CloudPCImageUploadScenario,
		"cloudpcingraceperiodscenario":                   DeviceManagementAlertRuleTemplate_CloudPCInGracePeriodScenario,
		"cloudpcinaccessiblescenario":                    DeviceManagementAlertRuleTemplate_CloudPCInaccessibleScenario,
		"cloudpconpremisenetworkconnectioncheckscenario": DeviceManagementAlertRuleTemplate_CloudPCOnPremiseNetworkConnectionCheckScenario,
		"cloudpcprovisionscenario":                       DeviceManagementAlertRuleTemplate_CloudPCProvisionScenario,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAlertRuleTemplate(input)
	return &out, nil
}
