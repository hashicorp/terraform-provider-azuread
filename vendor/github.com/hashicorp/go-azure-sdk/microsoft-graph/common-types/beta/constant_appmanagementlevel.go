package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppManagementLevel string

const (
	AppManagementLevel_AndroidEnterprise                                      AppManagementLevel = "androidEnterprise"
	AppManagementLevel_AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode AppManagementLevel = "androidEnterpriseDedicatedDevicesWithAzureAdSharedMode"
	AppManagementLevel_AndroidOpenSourceProjectUserAssociated                 AppManagementLevel = "androidOpenSourceProjectUserAssociated"
	AppManagementLevel_AndroidOpenSourceProjectUserless                       AppManagementLevel = "androidOpenSourceProjectUserless"
	AppManagementLevel_Mdm                                                    AppManagementLevel = "mdm"
	AppManagementLevel_Unmanaged                                              AppManagementLevel = "unmanaged"
	AppManagementLevel_Unspecified                                            AppManagementLevel = "unspecified"
)

func PossibleValuesForAppManagementLevel() []string {
	return []string{
		string(AppManagementLevel_AndroidEnterprise),
		string(AppManagementLevel_AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode),
		string(AppManagementLevel_AndroidOpenSourceProjectUserAssociated),
		string(AppManagementLevel_AndroidOpenSourceProjectUserless),
		string(AppManagementLevel_Mdm),
		string(AppManagementLevel_Unmanaged),
		string(AppManagementLevel_Unspecified),
	}
}

func (s *AppManagementLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppManagementLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppManagementLevel(input string) (*AppManagementLevel, error) {
	vals := map[string]AppManagementLevel{
		"androidenterprise": AppManagementLevel_AndroidEnterprise,
		"androidenterprisededicateddeviceswithazureadsharedmode": AppManagementLevel_AndroidEnterpriseDedicatedDevicesWithAzureAdSharedMode,
		"androidopensourceprojectuserassociated":                 AppManagementLevel_AndroidOpenSourceProjectUserAssociated,
		"androidopensourceprojectuserless":                       AppManagementLevel_AndroidOpenSourceProjectUserless,
		"mdm":                                                    AppManagementLevel_Mdm,
		"unmanaged":                                              AppManagementLevel_Unmanaged,
		"unspecified":                                            AppManagementLevel_Unspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppManagementLevel(input)
	return &out, nil
}
