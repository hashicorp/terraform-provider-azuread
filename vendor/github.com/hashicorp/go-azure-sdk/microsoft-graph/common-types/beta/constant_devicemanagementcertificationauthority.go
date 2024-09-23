package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementCertificationAuthority string

const (
	DeviceManagementCertificationAuthority_DigiCert      DeviceManagementCertificationAuthority = "digiCert"
	DeviceManagementCertificationAuthority_Microsoft     DeviceManagementCertificationAuthority = "microsoft"
	DeviceManagementCertificationAuthority_NotConfigured DeviceManagementCertificationAuthority = "notConfigured"
)

func PossibleValuesForDeviceManagementCertificationAuthority() []string {
	return []string{
		string(DeviceManagementCertificationAuthority_DigiCert),
		string(DeviceManagementCertificationAuthority_Microsoft),
		string(DeviceManagementCertificationAuthority_NotConfigured),
	}
}

func (s *DeviceManagementCertificationAuthority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementCertificationAuthority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementCertificationAuthority(input string) (*DeviceManagementCertificationAuthority, error) {
	vals := map[string]DeviceManagementCertificationAuthority{
		"digicert":      DeviceManagementCertificationAuthority_DigiCert,
		"microsoft":     DeviceManagementCertificationAuthority_Microsoft,
		"notconfigured": DeviceManagementCertificationAuthority_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementCertificationAuthority(input)
	return &out, nil
}
