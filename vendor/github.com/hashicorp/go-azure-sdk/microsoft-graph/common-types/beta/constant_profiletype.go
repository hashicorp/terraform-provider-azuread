package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileType string

const (
	ProfileType_AdministrativeTemplates ProfileType = "administrativeTemplates"
	ProfileType_DcV1DeviceRestrictions  ProfileType = "dcV1DeviceRestrictions"
	ProfileType_DcV1EndpointProtection  ProfileType = "dcV1EndpointProtection"
	ProfileType_HardwareConfig          ProfileType = "hardwareConfig"
	ProfileType_ImportedADMXTemplates   ProfileType = "importedADMXTemplates"
	ProfileType_OemAppConfig            ProfileType = "oemAppConfig"
	ProfileType_SettingsCatalog         ProfileType = "settingsCatalog"
)

func PossibleValuesForProfileType() []string {
	return []string{
		string(ProfileType_AdministrativeTemplates),
		string(ProfileType_DcV1DeviceRestrictions),
		string(ProfileType_DcV1EndpointProtection),
		string(ProfileType_HardwareConfig),
		string(ProfileType_ImportedADMXTemplates),
		string(ProfileType_OemAppConfig),
		string(ProfileType_SettingsCatalog),
	}
}

func (s *ProfileType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProfileType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProfileType(input string) (*ProfileType, error) {
	vals := map[string]ProfileType{
		"administrativetemplates": ProfileType_AdministrativeTemplates,
		"dcv1devicerestrictions":  ProfileType_DcV1DeviceRestrictions,
		"dcv1endpointprotection":  ProfileType_DcV1EndpointProtection,
		"hardwareconfig":          ProfileType_HardwareConfig,
		"importedadmxtemplates":   ProfileType_ImportedADMXTemplates,
		"oemappconfig":            ProfileType_OemAppConfig,
		"settingscatalog":         ProfileType_SettingsCatalog,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProfileType(input)
	return &out, nil
}
