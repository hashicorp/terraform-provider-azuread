package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesAzureADDeviceRegistrationErrorReason string

const (
	WindowsUpdatesAzureADDeviceRegistrationErrorReason_InvalidAzureADDeviceId WindowsUpdatesAzureADDeviceRegistrationErrorReason = "invalidAzureADDeviceId"
	WindowsUpdatesAzureADDeviceRegistrationErrorReason_InvalidAzureADJoin     WindowsUpdatesAzureADDeviceRegistrationErrorReason = "invalidAzureADJoin"
	WindowsUpdatesAzureADDeviceRegistrationErrorReason_InvalidGlobalDeviceId  WindowsUpdatesAzureADDeviceRegistrationErrorReason = "invalidGlobalDeviceId"
	WindowsUpdatesAzureADDeviceRegistrationErrorReason_MissingTrustType       WindowsUpdatesAzureADDeviceRegistrationErrorReason = "missingTrustType"
)

func PossibleValuesForWindowsUpdatesAzureADDeviceRegistrationErrorReason() []string {
	return []string{
		string(WindowsUpdatesAzureADDeviceRegistrationErrorReason_InvalidAzureADDeviceId),
		string(WindowsUpdatesAzureADDeviceRegistrationErrorReason_InvalidAzureADJoin),
		string(WindowsUpdatesAzureADDeviceRegistrationErrorReason_InvalidGlobalDeviceId),
		string(WindowsUpdatesAzureADDeviceRegistrationErrorReason_MissingTrustType),
	}
}

func (s *WindowsUpdatesAzureADDeviceRegistrationErrorReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdatesAzureADDeviceRegistrationErrorReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdatesAzureADDeviceRegistrationErrorReason(input string) (*WindowsUpdatesAzureADDeviceRegistrationErrorReason, error) {
	vals := map[string]WindowsUpdatesAzureADDeviceRegistrationErrorReason{
		"invalidazureaddeviceid": WindowsUpdatesAzureADDeviceRegistrationErrorReason_InvalidAzureADDeviceId,
		"invalidazureadjoin":     WindowsUpdatesAzureADDeviceRegistrationErrorReason_InvalidAzureADJoin,
		"invalidglobaldeviceid":  WindowsUpdatesAzureADDeviceRegistrationErrorReason_InvalidGlobalDeviceId,
		"missingtrusttype":       WindowsUpdatesAzureADDeviceRegistrationErrorReason_MissingTrustType,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesAzureADDeviceRegistrationErrorReason(input)
	return &out, nil
}
