package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerDelegatedAppScopeType string

const (
	AndroidDeviceOwnerDelegatedAppScopeType_CaptureNetworkActivityLog AndroidDeviceOwnerDelegatedAppScopeType = "captureNetworkActivityLog"
	AndroidDeviceOwnerDelegatedAppScopeType_CaptureSecurityLog        AndroidDeviceOwnerDelegatedAppScopeType = "captureSecurityLog"
	AndroidDeviceOwnerDelegatedAppScopeType_CertificateInstall        AndroidDeviceOwnerDelegatedAppScopeType = "certificateInstall"
	AndroidDeviceOwnerDelegatedAppScopeType_Unspecified               AndroidDeviceOwnerDelegatedAppScopeType = "unspecified"
)

func PossibleValuesForAndroidDeviceOwnerDelegatedAppScopeType() []string {
	return []string{
		string(AndroidDeviceOwnerDelegatedAppScopeType_CaptureNetworkActivityLog),
		string(AndroidDeviceOwnerDelegatedAppScopeType_CaptureSecurityLog),
		string(AndroidDeviceOwnerDelegatedAppScopeType_CertificateInstall),
		string(AndroidDeviceOwnerDelegatedAppScopeType_Unspecified),
	}
}

func (s *AndroidDeviceOwnerDelegatedAppScopeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerDelegatedAppScopeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerDelegatedAppScopeType(input string) (*AndroidDeviceOwnerDelegatedAppScopeType, error) {
	vals := map[string]AndroidDeviceOwnerDelegatedAppScopeType{
		"capturenetworkactivitylog": AndroidDeviceOwnerDelegatedAppScopeType_CaptureNetworkActivityLog,
		"capturesecuritylog":        AndroidDeviceOwnerDelegatedAppScopeType_CaptureSecurityLog,
		"certificateinstall":        AndroidDeviceOwnerDelegatedAppScopeType_CertificateInstall,
		"unspecified":               AndroidDeviceOwnerDelegatedAppScopeType_Unspecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerDelegatedAppScopeType(input)
	return &out, nil
}
