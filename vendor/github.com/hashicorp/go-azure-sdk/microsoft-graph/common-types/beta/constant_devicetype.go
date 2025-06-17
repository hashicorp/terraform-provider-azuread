package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceType string

const (
	DeviceType_Android           DeviceType = "android"
	DeviceType_AndroidEnterprise DeviceType = "androidEnterprise"
	DeviceType_AndroidForWork    DeviceType = "androidForWork"
	DeviceType_AndroidnGMS       DeviceType = "androidnGMS"
	DeviceType_Blackberry        DeviceType = "blackberry"
	DeviceType_ChromeOS          DeviceType = "chromeOS"
	DeviceType_CloudPC           DeviceType = "cloudPC"
	DeviceType_Desktop           DeviceType = "desktop"
	DeviceType_HoloLens          DeviceType = "holoLens"
	DeviceType_IPad              DeviceType = "iPad"
	DeviceType_IPhone            DeviceType = "iPhone"
	DeviceType_IPod              DeviceType = "iPod"
	DeviceType_ISocConsumer      DeviceType = "iSocConsumer"
	DeviceType_Linux             DeviceType = "linux"
	DeviceType_Mac               DeviceType = "mac"
	DeviceType_MacMDM            DeviceType = "macMDM"
	DeviceType_Nokia             DeviceType = "nokia"
	DeviceType_Palm              DeviceType = "palm"
	DeviceType_SurfaceHub        DeviceType = "surfaceHub"
	DeviceType_TvOS              DeviceType = "tvOS"
	DeviceType_Unix              DeviceType = "unix"
	DeviceType_Unknown           DeviceType = "unknown"
	DeviceType_VisionOS          DeviceType = "visionOS"
	DeviceType_WinCE             DeviceType = "winCE"
	DeviceType_WinEmbedded       DeviceType = "winEmbedded"
	DeviceType_WinMO6            DeviceType = "winMO6"
	DeviceType_Windows10x        DeviceType = "windows10x"
	DeviceType_WindowsPhone      DeviceType = "windowsPhone"
	DeviceType_WindowsRT         DeviceType = "windowsRT"
)

func PossibleValuesForDeviceType() []string {
	return []string{
		string(DeviceType_Android),
		string(DeviceType_AndroidEnterprise),
		string(DeviceType_AndroidForWork),
		string(DeviceType_AndroidnGMS),
		string(DeviceType_Blackberry),
		string(DeviceType_ChromeOS),
		string(DeviceType_CloudPC),
		string(DeviceType_Desktop),
		string(DeviceType_HoloLens),
		string(DeviceType_IPad),
		string(DeviceType_IPhone),
		string(DeviceType_IPod),
		string(DeviceType_ISocConsumer),
		string(DeviceType_Linux),
		string(DeviceType_Mac),
		string(DeviceType_MacMDM),
		string(DeviceType_Nokia),
		string(DeviceType_Palm),
		string(DeviceType_SurfaceHub),
		string(DeviceType_TvOS),
		string(DeviceType_Unix),
		string(DeviceType_Unknown),
		string(DeviceType_VisionOS),
		string(DeviceType_WinCE),
		string(DeviceType_WinEmbedded),
		string(DeviceType_WinMO6),
		string(DeviceType_Windows10x),
		string(DeviceType_WindowsPhone),
		string(DeviceType_WindowsRT),
	}
}

func (s *DeviceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceType(input string) (*DeviceType, error) {
	vals := map[string]DeviceType{
		"android":           DeviceType_Android,
		"androidenterprise": DeviceType_AndroidEnterprise,
		"androidforwork":    DeviceType_AndroidForWork,
		"androidngms":       DeviceType_AndroidnGMS,
		"blackberry":        DeviceType_Blackberry,
		"chromeos":          DeviceType_ChromeOS,
		"cloudpc":           DeviceType_CloudPC,
		"desktop":           DeviceType_Desktop,
		"hololens":          DeviceType_HoloLens,
		"ipad":              DeviceType_IPad,
		"iphone":            DeviceType_IPhone,
		"ipod":              DeviceType_IPod,
		"isocconsumer":      DeviceType_ISocConsumer,
		"linux":             DeviceType_Linux,
		"mac":               DeviceType_Mac,
		"macmdm":            DeviceType_MacMDM,
		"nokia":             DeviceType_Nokia,
		"palm":              DeviceType_Palm,
		"surfacehub":        DeviceType_SurfaceHub,
		"tvos":              DeviceType_TvOS,
		"unix":              DeviceType_Unix,
		"unknown":           DeviceType_Unknown,
		"visionos":          DeviceType_VisionOS,
		"wince":             DeviceType_WinCE,
		"winembedded":       DeviceType_WinEmbedded,
		"winmo6":            DeviceType_WinMO6,
		"windows10x":        DeviceType_Windows10x,
		"windowsphone":      DeviceType_WindowsPhone,
		"windowsrt":         DeviceType_WindowsRT,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceType(input)
	return &out, nil
}
