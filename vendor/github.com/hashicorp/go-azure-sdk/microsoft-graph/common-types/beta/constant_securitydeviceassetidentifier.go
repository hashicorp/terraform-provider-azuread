package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceAssetIdentifier string

const (
	SecurityDeviceAssetIdentifier_DestinationDeviceName SecurityDeviceAssetIdentifier = "destinationDeviceName"
	SecurityDeviceAssetIdentifier_DeviceId              SecurityDeviceAssetIdentifier = "deviceId"
	SecurityDeviceAssetIdentifier_DeviceName            SecurityDeviceAssetIdentifier = "deviceName"
	SecurityDeviceAssetIdentifier_RemoteDeviceName      SecurityDeviceAssetIdentifier = "remoteDeviceName"
	SecurityDeviceAssetIdentifier_TargetDeviceName      SecurityDeviceAssetIdentifier = "targetDeviceName"
)

func PossibleValuesForSecurityDeviceAssetIdentifier() []string {
	return []string{
		string(SecurityDeviceAssetIdentifier_DestinationDeviceName),
		string(SecurityDeviceAssetIdentifier_DeviceId),
		string(SecurityDeviceAssetIdentifier_DeviceName),
		string(SecurityDeviceAssetIdentifier_RemoteDeviceName),
		string(SecurityDeviceAssetIdentifier_TargetDeviceName),
	}
}

func (s *SecurityDeviceAssetIdentifier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDeviceAssetIdentifier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDeviceAssetIdentifier(input string) (*SecurityDeviceAssetIdentifier, error) {
	vals := map[string]SecurityDeviceAssetIdentifier{
		"destinationdevicename": SecurityDeviceAssetIdentifier_DestinationDeviceName,
		"deviceid":              SecurityDeviceAssetIdentifier_DeviceId,
		"devicename":            SecurityDeviceAssetIdentifier_DeviceName,
		"remotedevicename":      SecurityDeviceAssetIdentifier_RemoteDeviceName,
		"targetdevicename":      SecurityDeviceAssetIdentifier_TargetDeviceName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceAssetIdentifier(input)
	return &out, nil
}
