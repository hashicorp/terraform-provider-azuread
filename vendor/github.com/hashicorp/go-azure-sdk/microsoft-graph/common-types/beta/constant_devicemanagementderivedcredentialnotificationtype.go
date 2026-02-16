package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementDerivedCredentialNotificationType string

const (
	DeviceManagementDerivedCredentialNotificationType_CompanyPortal DeviceManagementDerivedCredentialNotificationType = "companyPortal"
	DeviceManagementDerivedCredentialNotificationType_Email         DeviceManagementDerivedCredentialNotificationType = "email"
	DeviceManagementDerivedCredentialNotificationType_None          DeviceManagementDerivedCredentialNotificationType = "none"
)

func PossibleValuesForDeviceManagementDerivedCredentialNotificationType() []string {
	return []string{
		string(DeviceManagementDerivedCredentialNotificationType_CompanyPortal),
		string(DeviceManagementDerivedCredentialNotificationType_Email),
		string(DeviceManagementDerivedCredentialNotificationType_None),
	}
}

func (s *DeviceManagementDerivedCredentialNotificationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementDerivedCredentialNotificationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementDerivedCredentialNotificationType(input string) (*DeviceManagementDerivedCredentialNotificationType, error) {
	vals := map[string]DeviceManagementDerivedCredentialNotificationType{
		"companyportal": DeviceManagementDerivedCredentialNotificationType_CompanyPortal,
		"email":         DeviceManagementDerivedCredentialNotificationType_Email,
		"none":          DeviceManagementDerivedCredentialNotificationType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementDerivedCredentialNotificationType(input)
	return &out, nil
}
