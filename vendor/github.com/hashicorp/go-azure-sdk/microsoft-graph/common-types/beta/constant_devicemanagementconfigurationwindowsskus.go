package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationWindowsSkus string

const (
	DeviceManagementConfigurationWindowsSkus_HoloLens                DeviceManagementConfigurationWindowsSkus = "holoLens"
	DeviceManagementConfigurationWindowsSkus_HoloLensEnterprise      DeviceManagementConfigurationWindowsSkus = "holoLensEnterprise"
	DeviceManagementConfigurationWindowsSkus_HolographicForBusiness  DeviceManagementConfigurationWindowsSkus = "holographicForBusiness"
	DeviceManagementConfigurationWindowsSkus_Iot                     DeviceManagementConfigurationWindowsSkus = "iot"
	DeviceManagementConfigurationWindowsSkus_IotEnterprise           DeviceManagementConfigurationWindowsSkus = "iotEnterprise"
	DeviceManagementConfigurationWindowsSkus_SurfaceHub              DeviceManagementConfigurationWindowsSkus = "surfaceHub"
	DeviceManagementConfigurationWindowsSkus_Unknown                 DeviceManagementConfigurationWindowsSkus = "unknown"
	DeviceManagementConfigurationWindowsSkus_WindowsEducation        DeviceManagementConfigurationWindowsSkus = "windowsEducation"
	DeviceManagementConfigurationWindowsSkus_WindowsEnterprise       DeviceManagementConfigurationWindowsSkus = "windowsEnterprise"
	DeviceManagementConfigurationWindowsSkus_WindowsHome             DeviceManagementConfigurationWindowsSkus = "windowsHome"
	DeviceManagementConfigurationWindowsSkus_WindowsMobile           DeviceManagementConfigurationWindowsSkus = "windowsMobile"
	DeviceManagementConfigurationWindowsSkus_WindowsMobileEnterprise DeviceManagementConfigurationWindowsSkus = "windowsMobileEnterprise"
	DeviceManagementConfigurationWindowsSkus_WindowsMultiSession     DeviceManagementConfigurationWindowsSkus = "windowsMultiSession"
	DeviceManagementConfigurationWindowsSkus_WindowsProfessional     DeviceManagementConfigurationWindowsSkus = "windowsProfessional"
	DeviceManagementConfigurationWindowsSkus_WindowsTeamSurface      DeviceManagementConfigurationWindowsSkus = "windowsTeamSurface"
)

func PossibleValuesForDeviceManagementConfigurationWindowsSkus() []string {
	return []string{
		string(DeviceManagementConfigurationWindowsSkus_HoloLens),
		string(DeviceManagementConfigurationWindowsSkus_HoloLensEnterprise),
		string(DeviceManagementConfigurationWindowsSkus_HolographicForBusiness),
		string(DeviceManagementConfigurationWindowsSkus_Iot),
		string(DeviceManagementConfigurationWindowsSkus_IotEnterprise),
		string(DeviceManagementConfigurationWindowsSkus_SurfaceHub),
		string(DeviceManagementConfigurationWindowsSkus_Unknown),
		string(DeviceManagementConfigurationWindowsSkus_WindowsEducation),
		string(DeviceManagementConfigurationWindowsSkus_WindowsEnterprise),
		string(DeviceManagementConfigurationWindowsSkus_WindowsHome),
		string(DeviceManagementConfigurationWindowsSkus_WindowsMobile),
		string(DeviceManagementConfigurationWindowsSkus_WindowsMobileEnterprise),
		string(DeviceManagementConfigurationWindowsSkus_WindowsMultiSession),
		string(DeviceManagementConfigurationWindowsSkus_WindowsProfessional),
		string(DeviceManagementConfigurationWindowsSkus_WindowsTeamSurface),
	}
}

func (s *DeviceManagementConfigurationWindowsSkus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationWindowsSkus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationWindowsSkus(input string) (*DeviceManagementConfigurationWindowsSkus, error) {
	vals := map[string]DeviceManagementConfigurationWindowsSkus{
		"hololens":                DeviceManagementConfigurationWindowsSkus_HoloLens,
		"hololensenterprise":      DeviceManagementConfigurationWindowsSkus_HoloLensEnterprise,
		"holographicforbusiness":  DeviceManagementConfigurationWindowsSkus_HolographicForBusiness,
		"iot":                     DeviceManagementConfigurationWindowsSkus_Iot,
		"iotenterprise":           DeviceManagementConfigurationWindowsSkus_IotEnterprise,
		"surfacehub":              DeviceManagementConfigurationWindowsSkus_SurfaceHub,
		"unknown":                 DeviceManagementConfigurationWindowsSkus_Unknown,
		"windowseducation":        DeviceManagementConfigurationWindowsSkus_WindowsEducation,
		"windowsenterprise":       DeviceManagementConfigurationWindowsSkus_WindowsEnterprise,
		"windowshome":             DeviceManagementConfigurationWindowsSkus_WindowsHome,
		"windowsmobile":           DeviceManagementConfigurationWindowsSkus_WindowsMobile,
		"windowsmobileenterprise": DeviceManagementConfigurationWindowsSkus_WindowsMobileEnterprise,
		"windowsmultisession":     DeviceManagementConfigurationWindowsSkus_WindowsMultiSession,
		"windowsprofessional":     DeviceManagementConfigurationWindowsSkus_WindowsProfessional,
		"windowsteamsurface":      DeviceManagementConfigurationWindowsSkus_WindowsTeamSurface,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationWindowsSkus(input)
	return &out, nil
}
