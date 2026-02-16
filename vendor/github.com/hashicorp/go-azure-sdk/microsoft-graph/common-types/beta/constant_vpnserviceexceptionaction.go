package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnServiceExceptionAction string

const (
	VpnServiceExceptionAction_AllowTrafficOutside VpnServiceExceptionAction = "allowTrafficOutside"
	VpnServiceExceptionAction_DropTraffic         VpnServiceExceptionAction = "dropTraffic"
	VpnServiceExceptionAction_ForceTrafficViaVPN  VpnServiceExceptionAction = "forceTrafficViaVPN"
)

func PossibleValuesForVpnServiceExceptionAction() []string {
	return []string{
		string(VpnServiceExceptionAction_AllowTrafficOutside),
		string(VpnServiceExceptionAction_DropTraffic),
		string(VpnServiceExceptionAction_ForceTrafficViaVPN),
	}
}

func (s *VpnServiceExceptionAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnServiceExceptionAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnServiceExceptionAction(input string) (*VpnServiceExceptionAction, error) {
	vals := map[string]VpnServiceExceptionAction{
		"allowtrafficoutside": VpnServiceExceptionAction_AllowTrafficOutside,
		"droptraffic":         VpnServiceExceptionAction_DropTraffic,
		"forcetrafficviavpn":  VpnServiceExceptionAction_ForceTrafficViaVPN,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnServiceExceptionAction(input)
	return &out, nil
}
