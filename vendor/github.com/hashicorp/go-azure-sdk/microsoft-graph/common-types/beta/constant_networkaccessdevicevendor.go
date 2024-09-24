package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDeviceVendor string

const (
	NetworkaccessDeviceVendor_BarracudaNetworks  NetworkaccessDeviceVendor = "barracudaNetworks"
	NetworkaccessDeviceVendor_CheckPoint         NetworkaccessDeviceVendor = "checkPoint"
	NetworkaccessDeviceVendor_CiscoCatalyst      NetworkaccessDeviceVendor = "ciscoCatalyst"
	NetworkaccessDeviceVendor_CiscoMeraki        NetworkaccessDeviceVendor = "ciscoMeraki"
	NetworkaccessDeviceVendor_Citrix             NetworkaccessDeviceVendor = "citrix"
	NetworkaccessDeviceVendor_Fortinet           NetworkaccessDeviceVendor = "fortinet"
	NetworkaccessDeviceVendor_HpeAruba           NetworkaccessDeviceVendor = "hpeAruba"
	NetworkaccessDeviceVendor_NetFoundry         NetworkaccessDeviceVendor = "netFoundry"
	NetworkaccessDeviceVendor_Nuage              NetworkaccessDeviceVendor = "nuage"
	NetworkaccessDeviceVendor_OpenSystems        NetworkaccessDeviceVendor = "openSystems"
	NetworkaccessDeviceVendor_Other              NetworkaccessDeviceVendor = "other"
	NetworkaccessDeviceVendor_PaloAltoNetworks   NetworkaccessDeviceVendor = "paloAltoNetworks"
	NetworkaccessDeviceVendor_RiverbedTechnology NetworkaccessDeviceVendor = "riverbedTechnology"
	NetworkaccessDeviceVendor_SilverPeak         NetworkaccessDeviceVendor = "silverPeak"
	NetworkaccessDeviceVendor_Versa              NetworkaccessDeviceVendor = "versa"
	NetworkaccessDeviceVendor_VmWareSdWan        NetworkaccessDeviceVendor = "vmWareSdWan"
)

func PossibleValuesForNetworkaccessDeviceVendor() []string {
	return []string{
		string(NetworkaccessDeviceVendor_BarracudaNetworks),
		string(NetworkaccessDeviceVendor_CheckPoint),
		string(NetworkaccessDeviceVendor_CiscoCatalyst),
		string(NetworkaccessDeviceVendor_CiscoMeraki),
		string(NetworkaccessDeviceVendor_Citrix),
		string(NetworkaccessDeviceVendor_Fortinet),
		string(NetworkaccessDeviceVendor_HpeAruba),
		string(NetworkaccessDeviceVendor_NetFoundry),
		string(NetworkaccessDeviceVendor_Nuage),
		string(NetworkaccessDeviceVendor_OpenSystems),
		string(NetworkaccessDeviceVendor_Other),
		string(NetworkaccessDeviceVendor_PaloAltoNetworks),
		string(NetworkaccessDeviceVendor_RiverbedTechnology),
		string(NetworkaccessDeviceVendor_SilverPeak),
		string(NetworkaccessDeviceVendor_Versa),
		string(NetworkaccessDeviceVendor_VmWareSdWan),
	}
}

func (s *NetworkaccessDeviceVendor) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessDeviceVendor(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessDeviceVendor(input string) (*NetworkaccessDeviceVendor, error) {
	vals := map[string]NetworkaccessDeviceVendor{
		"barracudanetworks":  NetworkaccessDeviceVendor_BarracudaNetworks,
		"checkpoint":         NetworkaccessDeviceVendor_CheckPoint,
		"ciscocatalyst":      NetworkaccessDeviceVendor_CiscoCatalyst,
		"ciscomeraki":        NetworkaccessDeviceVendor_CiscoMeraki,
		"citrix":             NetworkaccessDeviceVendor_Citrix,
		"fortinet":           NetworkaccessDeviceVendor_Fortinet,
		"hpearuba":           NetworkaccessDeviceVendor_HpeAruba,
		"netfoundry":         NetworkaccessDeviceVendor_NetFoundry,
		"nuage":              NetworkaccessDeviceVendor_Nuage,
		"opensystems":        NetworkaccessDeviceVendor_OpenSystems,
		"other":              NetworkaccessDeviceVendor_Other,
		"paloaltonetworks":   NetworkaccessDeviceVendor_PaloAltoNetworks,
		"riverbedtechnology": NetworkaccessDeviceVendor_RiverbedTechnology,
		"silverpeak":         NetworkaccessDeviceVendor_SilverPeak,
		"versa":              NetworkaccessDeviceVendor_Versa,
		"vmwaresdwan":        NetworkaccessDeviceVendor_VmWareSdWan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessDeviceVendor(input)
	return &out, nil
}
