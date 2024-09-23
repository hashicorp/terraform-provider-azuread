package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceTypes string

const (
	DeviceTypes_Android           DeviceTypes = "android"
	DeviceTypes_AndroidEnterprise DeviceTypes = "androidEnterprise"
	DeviceTypes_AndroidForWork    DeviceTypes = "androidForWork"
	DeviceTypes_Blackberry        DeviceTypes = "blackberry"
	DeviceTypes_Desktop           DeviceTypes = "desktop"
	DeviceTypes_HoloLens          DeviceTypes = "holoLens"
	DeviceTypes_IPad              DeviceTypes = "iPad"
	DeviceTypes_IPhone            DeviceTypes = "iPhone"
	DeviceTypes_IPod              DeviceTypes = "iPod"
	DeviceTypes_ISocConsumer      DeviceTypes = "iSocConsumer"
	DeviceTypes_Mac               DeviceTypes = "mac"
	DeviceTypes_MacMDM            DeviceTypes = "macMDM"
	DeviceTypes_Nokia             DeviceTypes = "nokia"
	DeviceTypes_Palm              DeviceTypes = "palm"
	DeviceTypes_SurfaceHub        DeviceTypes = "surfaceHub"
	DeviceTypes_Unix              DeviceTypes = "unix"
	DeviceTypes_Unknown           DeviceTypes = "unknown"
	DeviceTypes_WinCE             DeviceTypes = "winCE"
	DeviceTypes_WinEmbedded       DeviceTypes = "winEmbedded"
	DeviceTypes_WinMO6            DeviceTypes = "winMO6"
	DeviceTypes_WindowsPhone      DeviceTypes = "windowsPhone"
	DeviceTypes_WindowsRT         DeviceTypes = "windowsRT"
)

func PossibleValuesForDeviceTypes() []string {
	return []string{
		string(DeviceTypes_Android),
		string(DeviceTypes_AndroidEnterprise),
		string(DeviceTypes_AndroidForWork),
		string(DeviceTypes_Blackberry),
		string(DeviceTypes_Desktop),
		string(DeviceTypes_HoloLens),
		string(DeviceTypes_IPad),
		string(DeviceTypes_IPhone),
		string(DeviceTypes_IPod),
		string(DeviceTypes_ISocConsumer),
		string(DeviceTypes_Mac),
		string(DeviceTypes_MacMDM),
		string(DeviceTypes_Nokia),
		string(DeviceTypes_Palm),
		string(DeviceTypes_SurfaceHub),
		string(DeviceTypes_Unix),
		string(DeviceTypes_Unknown),
		string(DeviceTypes_WinCE),
		string(DeviceTypes_WinEmbedded),
		string(DeviceTypes_WinMO6),
		string(DeviceTypes_WindowsPhone),
		string(DeviceTypes_WindowsRT),
	}
}

func (s *DeviceTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceTypes(input string) (*DeviceTypes, error) {
	vals := map[string]DeviceTypes{
		"android":           DeviceTypes_Android,
		"androidenterprise": DeviceTypes_AndroidEnterprise,
		"androidforwork":    DeviceTypes_AndroidForWork,
		"blackberry":        DeviceTypes_Blackberry,
		"desktop":           DeviceTypes_Desktop,
		"hololens":          DeviceTypes_HoloLens,
		"ipad":              DeviceTypes_IPad,
		"iphone":            DeviceTypes_IPhone,
		"ipod":              DeviceTypes_IPod,
		"isocconsumer":      DeviceTypes_ISocConsumer,
		"mac":               DeviceTypes_Mac,
		"macmdm":            DeviceTypes_MacMDM,
		"nokia":             DeviceTypes_Nokia,
		"palm":              DeviceTypes_Palm,
		"surfacehub":        DeviceTypes_SurfaceHub,
		"unix":              DeviceTypes_Unix,
		"unknown":           DeviceTypes_Unknown,
		"wince":             DeviceTypes_WinCE,
		"winembedded":       DeviceTypes_WinEmbedded,
		"winmo6":            DeviceTypes_WinMO6,
		"windowsphone":      DeviceTypes_WindowsPhone,
		"windowsrt":         DeviceTypes_WindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceTypes(input)
	return &out, nil
}
