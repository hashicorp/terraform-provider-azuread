package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExchangeAccessStateReason string

const (
	DeviceManagementExchangeAccessStateReason_AzureADBlockDueToAccessPolicy DeviceManagementExchangeAccessStateReason = "azureADBlockDueToAccessPolicy"
	DeviceManagementExchangeAccessStateReason_Compliant                     DeviceManagementExchangeAccessStateReason = "compliant"
	DeviceManagementExchangeAccessStateReason_CompromisedPassword           DeviceManagementExchangeAccessStateReason = "compromisedPassword"
	DeviceManagementExchangeAccessStateReason_DeviceNotKnownWithManagedApp  DeviceManagementExchangeAccessStateReason = "deviceNotKnownWithManagedApp"
	DeviceManagementExchangeAccessStateReason_ExchangeDeviceRule            DeviceManagementExchangeAccessStateReason = "exchangeDeviceRule"
	DeviceManagementExchangeAccessStateReason_ExchangeGlobalRule            DeviceManagementExchangeAccessStateReason = "exchangeGlobalRule"
	DeviceManagementExchangeAccessStateReason_ExchangeIndividualRule        DeviceManagementExchangeAccessStateReason = "exchangeIndividualRule"
	DeviceManagementExchangeAccessStateReason_ExchangeMailboxPolicy         DeviceManagementExchangeAccessStateReason = "exchangeMailboxPolicy"
	DeviceManagementExchangeAccessStateReason_ExchangeUpgrade               DeviceManagementExchangeAccessStateReason = "exchangeUpgrade"
	DeviceManagementExchangeAccessStateReason_MfaRequired                   DeviceManagementExchangeAccessStateReason = "mfaRequired"
	DeviceManagementExchangeAccessStateReason_None                          DeviceManagementExchangeAccessStateReason = "none"
	DeviceManagementExchangeAccessStateReason_NotCompliant                  DeviceManagementExchangeAccessStateReason = "notCompliant"
	DeviceManagementExchangeAccessStateReason_NotEnrolled                   DeviceManagementExchangeAccessStateReason = "notEnrolled"
	DeviceManagementExchangeAccessStateReason_Other                         DeviceManagementExchangeAccessStateReason = "other"
	DeviceManagementExchangeAccessStateReason_Unknown                       DeviceManagementExchangeAccessStateReason = "unknown"
	DeviceManagementExchangeAccessStateReason_UnknownLocation               DeviceManagementExchangeAccessStateReason = "unknownLocation"
)

func PossibleValuesForDeviceManagementExchangeAccessStateReason() []string {
	return []string{
		string(DeviceManagementExchangeAccessStateReason_AzureADBlockDueToAccessPolicy),
		string(DeviceManagementExchangeAccessStateReason_Compliant),
		string(DeviceManagementExchangeAccessStateReason_CompromisedPassword),
		string(DeviceManagementExchangeAccessStateReason_DeviceNotKnownWithManagedApp),
		string(DeviceManagementExchangeAccessStateReason_ExchangeDeviceRule),
		string(DeviceManagementExchangeAccessStateReason_ExchangeGlobalRule),
		string(DeviceManagementExchangeAccessStateReason_ExchangeIndividualRule),
		string(DeviceManagementExchangeAccessStateReason_ExchangeMailboxPolicy),
		string(DeviceManagementExchangeAccessStateReason_ExchangeUpgrade),
		string(DeviceManagementExchangeAccessStateReason_MfaRequired),
		string(DeviceManagementExchangeAccessStateReason_None),
		string(DeviceManagementExchangeAccessStateReason_NotCompliant),
		string(DeviceManagementExchangeAccessStateReason_NotEnrolled),
		string(DeviceManagementExchangeAccessStateReason_Other),
		string(DeviceManagementExchangeAccessStateReason_Unknown),
		string(DeviceManagementExchangeAccessStateReason_UnknownLocation),
	}
}

func (s *DeviceManagementExchangeAccessStateReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementExchangeAccessStateReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementExchangeAccessStateReason(input string) (*DeviceManagementExchangeAccessStateReason, error) {
	vals := map[string]DeviceManagementExchangeAccessStateReason{
		"azureadblockduetoaccesspolicy": DeviceManagementExchangeAccessStateReason_AzureADBlockDueToAccessPolicy,
		"compliant":                     DeviceManagementExchangeAccessStateReason_Compliant,
		"compromisedpassword":           DeviceManagementExchangeAccessStateReason_CompromisedPassword,
		"devicenotknownwithmanagedapp":  DeviceManagementExchangeAccessStateReason_DeviceNotKnownWithManagedApp,
		"exchangedevicerule":            DeviceManagementExchangeAccessStateReason_ExchangeDeviceRule,
		"exchangeglobalrule":            DeviceManagementExchangeAccessStateReason_ExchangeGlobalRule,
		"exchangeindividualrule":        DeviceManagementExchangeAccessStateReason_ExchangeIndividualRule,
		"exchangemailboxpolicy":         DeviceManagementExchangeAccessStateReason_ExchangeMailboxPolicy,
		"exchangeupgrade":               DeviceManagementExchangeAccessStateReason_ExchangeUpgrade,
		"mfarequired":                   DeviceManagementExchangeAccessStateReason_MfaRequired,
		"none":                          DeviceManagementExchangeAccessStateReason_None,
		"notcompliant":                  DeviceManagementExchangeAccessStateReason_NotCompliant,
		"notenrolled":                   DeviceManagementExchangeAccessStateReason_NotEnrolled,
		"other":                         DeviceManagementExchangeAccessStateReason_Other,
		"unknown":                       DeviceManagementExchangeAccessStateReason_Unknown,
		"unknownlocation":               DeviceManagementExchangeAccessStateReason_UnknownLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExchangeAccessStateReason(input)
	return &out, nil
}
