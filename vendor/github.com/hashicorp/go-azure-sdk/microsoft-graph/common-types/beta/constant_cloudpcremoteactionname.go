package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRemoteActionName string

const (
	CloudPCRemoteActionName_ChangeUserAccountType CloudPCRemoteActionName = "changeUserAccountType"
	CloudPCRemoteActionName_CreateSnapshot        CloudPCRemoteActionName = "createSnapshot"
	CloudPCRemoteActionName_MoveRegion            CloudPCRemoteActionName = "moveRegion"
	CloudPCRemoteActionName_PlaceUnderReview      CloudPCRemoteActionName = "placeUnderReview"
	CloudPCRemoteActionName_PowerOff              CloudPCRemoteActionName = "powerOff"
	CloudPCRemoteActionName_PowerOn               CloudPCRemoteActionName = "powerOn"
	CloudPCRemoteActionName_Rename                CloudPCRemoteActionName = "rename"
	CloudPCRemoteActionName_Reprovision           CloudPCRemoteActionName = "reprovision"
	CloudPCRemoteActionName_Resize                CloudPCRemoteActionName = "resize"
	CloudPCRemoteActionName_Restart               CloudPCRemoteActionName = "restart"
	CloudPCRemoteActionName_Restore               CloudPCRemoteActionName = "restore"
	CloudPCRemoteActionName_Troubleshoot          CloudPCRemoteActionName = "troubleshoot"
	CloudPCRemoteActionName_Unknown               CloudPCRemoteActionName = "unknown"
)

func PossibleValuesForCloudPCRemoteActionName() []string {
	return []string{
		string(CloudPCRemoteActionName_ChangeUserAccountType),
		string(CloudPCRemoteActionName_CreateSnapshot),
		string(CloudPCRemoteActionName_MoveRegion),
		string(CloudPCRemoteActionName_PlaceUnderReview),
		string(CloudPCRemoteActionName_PowerOff),
		string(CloudPCRemoteActionName_PowerOn),
		string(CloudPCRemoteActionName_Rename),
		string(CloudPCRemoteActionName_Reprovision),
		string(CloudPCRemoteActionName_Resize),
		string(CloudPCRemoteActionName_Restart),
		string(CloudPCRemoteActionName_Restore),
		string(CloudPCRemoteActionName_Troubleshoot),
		string(CloudPCRemoteActionName_Unknown),
	}
}

func (s *CloudPCRemoteActionName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCRemoteActionName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCRemoteActionName(input string) (*CloudPCRemoteActionName, error) {
	vals := map[string]CloudPCRemoteActionName{
		"changeuseraccounttype": CloudPCRemoteActionName_ChangeUserAccountType,
		"createsnapshot":        CloudPCRemoteActionName_CreateSnapshot,
		"moveregion":            CloudPCRemoteActionName_MoveRegion,
		"placeunderreview":      CloudPCRemoteActionName_PlaceUnderReview,
		"poweroff":              CloudPCRemoteActionName_PowerOff,
		"poweron":               CloudPCRemoteActionName_PowerOn,
		"rename":                CloudPCRemoteActionName_Rename,
		"reprovision":           CloudPCRemoteActionName_Reprovision,
		"resize":                CloudPCRemoteActionName_Resize,
		"restart":               CloudPCRemoteActionName_Restart,
		"restore":               CloudPCRemoteActionName_Restore,
		"troubleshoot":          CloudPCRemoteActionName_Troubleshoot,
		"unknown":               CloudPCRemoteActionName_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCRemoteActionName(input)
	return &out, nil
}
