package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaErrorCode string

const (
	ZebraFotaErrorCode_NoDevicesFoundInSelectedAadGroups                  ZebraFotaErrorCode = "noDevicesFoundInSelectedAadGroups"
	ZebraFotaErrorCode_NoIntuneDevicesFoundInSelectedAadGroups            ZebraFotaErrorCode = "noIntuneDevicesFoundInSelectedAadGroups"
	ZebraFotaErrorCode_NoZebraFotaDevicesFoundForSelectedDeviceModel      ZebraFotaErrorCode = "noZebraFotaDevicesFoundForSelectedDeviceModel"
	ZebraFotaErrorCode_NoZebraFotaEnrolledDevicesFoundForCurrentTenant    ZebraFotaErrorCode = "noZebraFotaEnrolledDevicesFoundForCurrentTenant"
	ZebraFotaErrorCode_NoZebraFotaEnrolledDevicesFoundInSelectedAadGroups ZebraFotaErrorCode = "noZebraFotaEnrolledDevicesFoundInSelectedAadGroups"
	ZebraFotaErrorCode_Success                                            ZebraFotaErrorCode = "success"
	ZebraFotaErrorCode_ZebraFotaCreateDeploymentRequestFailure            ZebraFotaErrorCode = "zebraFotaCreateDeploymentRequestFailure"
)

func PossibleValuesForZebraFotaErrorCode() []string {
	return []string{
		string(ZebraFotaErrorCode_NoDevicesFoundInSelectedAadGroups),
		string(ZebraFotaErrorCode_NoIntuneDevicesFoundInSelectedAadGroups),
		string(ZebraFotaErrorCode_NoZebraFotaDevicesFoundForSelectedDeviceModel),
		string(ZebraFotaErrorCode_NoZebraFotaEnrolledDevicesFoundForCurrentTenant),
		string(ZebraFotaErrorCode_NoZebraFotaEnrolledDevicesFoundInSelectedAadGroups),
		string(ZebraFotaErrorCode_Success),
		string(ZebraFotaErrorCode_ZebraFotaCreateDeploymentRequestFailure),
	}
}

func (s *ZebraFotaErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseZebraFotaErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseZebraFotaErrorCode(input string) (*ZebraFotaErrorCode, error) {
	vals := map[string]ZebraFotaErrorCode{
		"nodevicesfoundinselectedaadgroups":                  ZebraFotaErrorCode_NoDevicesFoundInSelectedAadGroups,
		"nointunedevicesfoundinselectedaadgroups":            ZebraFotaErrorCode_NoIntuneDevicesFoundInSelectedAadGroups,
		"nozebrafotadevicesfoundforselecteddevicemodel":      ZebraFotaErrorCode_NoZebraFotaDevicesFoundForSelectedDeviceModel,
		"nozebrafotaenrolleddevicesfoundforcurrenttenant":    ZebraFotaErrorCode_NoZebraFotaEnrolledDevicesFoundForCurrentTenant,
		"nozebrafotaenrolleddevicesfoundinselectedaadgroups": ZebraFotaErrorCode_NoZebraFotaEnrolledDevicesFoundInSelectedAadGroups,
		"success": ZebraFotaErrorCode_Success,
		"zebrafotacreatedeploymentrequestfailure": ZebraFotaErrorCode_ZebraFotaCreateDeploymentRequestFailure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaErrorCode(input)
	return &out, nil
}
