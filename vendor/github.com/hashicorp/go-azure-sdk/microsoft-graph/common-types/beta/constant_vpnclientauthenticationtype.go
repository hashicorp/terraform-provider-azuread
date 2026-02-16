package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnClientAuthenticationType string

const (
	VpnClientAuthenticationType_DeviceAuthentication VpnClientAuthenticationType = "deviceAuthentication"
	VpnClientAuthenticationType_UserAuthentication   VpnClientAuthenticationType = "userAuthentication"
)

func PossibleValuesForVpnClientAuthenticationType() []string {
	return []string{
		string(VpnClientAuthenticationType_DeviceAuthentication),
		string(VpnClientAuthenticationType_UserAuthentication),
	}
}

func (s *VpnClientAuthenticationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnClientAuthenticationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnClientAuthenticationType(input string) (*VpnClientAuthenticationType, error) {
	vals := map[string]VpnClientAuthenticationType{
		"deviceauthentication": VpnClientAuthenticationType_DeviceAuthentication,
		"userauthentication":   VpnClientAuthenticationType_UserAuthentication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnClientAuthenticationType(input)
	return &out, nil
}
