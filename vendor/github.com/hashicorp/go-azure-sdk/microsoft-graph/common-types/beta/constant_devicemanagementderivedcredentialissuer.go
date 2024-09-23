package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementDerivedCredentialIssuer string

const (
	DeviceManagementDerivedCredentialIssuer_EntrustDatacard DeviceManagementDerivedCredentialIssuer = "entrustDatacard"
	DeviceManagementDerivedCredentialIssuer_Intercede       DeviceManagementDerivedCredentialIssuer = "intercede"
	DeviceManagementDerivedCredentialIssuer_Purebred        DeviceManagementDerivedCredentialIssuer = "purebred"
	DeviceManagementDerivedCredentialIssuer_XTec            DeviceManagementDerivedCredentialIssuer = "xTec"
)

func PossibleValuesForDeviceManagementDerivedCredentialIssuer() []string {
	return []string{
		string(DeviceManagementDerivedCredentialIssuer_EntrustDatacard),
		string(DeviceManagementDerivedCredentialIssuer_Intercede),
		string(DeviceManagementDerivedCredentialIssuer_Purebred),
		string(DeviceManagementDerivedCredentialIssuer_XTec),
	}
}

func (s *DeviceManagementDerivedCredentialIssuer) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementDerivedCredentialIssuer(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementDerivedCredentialIssuer(input string) (*DeviceManagementDerivedCredentialIssuer, error) {
	vals := map[string]DeviceManagementDerivedCredentialIssuer{
		"entrustdatacard": DeviceManagementDerivedCredentialIssuer_EntrustDatacard,
		"intercede":       DeviceManagementDerivedCredentialIssuer_Intercede,
		"purebred":        DeviceManagementDerivedCredentialIssuer_Purebred,
		"xtec":            DeviceManagementDerivedCredentialIssuer_XTec,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementDerivedCredentialIssuer(input)
	return &out, nil
}
