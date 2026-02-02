package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotEnrollmentType string

const (
	WindowsAutopilotEnrollmentType_AzureADJoinedUsingDeviceAuthWithAutopilotProfile    WindowsAutopilotEnrollmentType = "azureADJoinedUsingDeviceAuthWithAutopilotProfile"
	WindowsAutopilotEnrollmentType_AzureADJoinedUsingDeviceAuthWithoutAutopilotProfile WindowsAutopilotEnrollmentType = "azureADJoinedUsingDeviceAuthWithoutAutopilotProfile"
	WindowsAutopilotEnrollmentType_AzureADJoinedWithAutopilotProfile                   WindowsAutopilotEnrollmentType = "azureADJoinedWithAutopilotProfile"
	WindowsAutopilotEnrollmentType_AzureADJoinedWithOfflineAutopilotProfile            WindowsAutopilotEnrollmentType = "azureADJoinedWithOfflineAutopilotProfile"
	WindowsAutopilotEnrollmentType_AzureADJoinedWithWhiteGlove                         WindowsAutopilotEnrollmentType = "azureADJoinedWithWhiteGlove"
	WindowsAutopilotEnrollmentType_OfflineDomainJoined                                 WindowsAutopilotEnrollmentType = "offlineDomainJoined"
	WindowsAutopilotEnrollmentType_OfflineDomainJoinedWithOfflineAutopilotProfile      WindowsAutopilotEnrollmentType = "offlineDomainJoinedWithOfflineAutopilotProfile"
	WindowsAutopilotEnrollmentType_OfflineDomainJoinedWithWhiteGlove                   WindowsAutopilotEnrollmentType = "offlineDomainJoinedWithWhiteGlove"
	WindowsAutopilotEnrollmentType_Unknown                                             WindowsAutopilotEnrollmentType = "unknown"
)

func PossibleValuesForWindowsAutopilotEnrollmentType() []string {
	return []string{
		string(WindowsAutopilotEnrollmentType_AzureADJoinedUsingDeviceAuthWithAutopilotProfile),
		string(WindowsAutopilotEnrollmentType_AzureADJoinedUsingDeviceAuthWithoutAutopilotProfile),
		string(WindowsAutopilotEnrollmentType_AzureADJoinedWithAutopilotProfile),
		string(WindowsAutopilotEnrollmentType_AzureADJoinedWithOfflineAutopilotProfile),
		string(WindowsAutopilotEnrollmentType_AzureADJoinedWithWhiteGlove),
		string(WindowsAutopilotEnrollmentType_OfflineDomainJoined),
		string(WindowsAutopilotEnrollmentType_OfflineDomainJoinedWithOfflineAutopilotProfile),
		string(WindowsAutopilotEnrollmentType_OfflineDomainJoinedWithWhiteGlove),
		string(WindowsAutopilotEnrollmentType_Unknown),
	}
}

func (s *WindowsAutopilotEnrollmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsAutopilotEnrollmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsAutopilotEnrollmentType(input string) (*WindowsAutopilotEnrollmentType, error) {
	vals := map[string]WindowsAutopilotEnrollmentType{
		"azureadjoinedusingdeviceauthwithautopilotprofile":    WindowsAutopilotEnrollmentType_AzureADJoinedUsingDeviceAuthWithAutopilotProfile,
		"azureadjoinedusingdeviceauthwithoutautopilotprofile": WindowsAutopilotEnrollmentType_AzureADJoinedUsingDeviceAuthWithoutAutopilotProfile,
		"azureadjoinedwithautopilotprofile":                   WindowsAutopilotEnrollmentType_AzureADJoinedWithAutopilotProfile,
		"azureadjoinedwithofflineautopilotprofile":            WindowsAutopilotEnrollmentType_AzureADJoinedWithOfflineAutopilotProfile,
		"azureadjoinedwithwhiteglove":                         WindowsAutopilotEnrollmentType_AzureADJoinedWithWhiteGlove,
		"offlinedomainjoined":                                 WindowsAutopilotEnrollmentType_OfflineDomainJoined,
		"offlinedomainjoinedwithofflineautopilotprofile":      WindowsAutopilotEnrollmentType_OfflineDomainJoinedWithOfflineAutopilotProfile,
		"offlinedomainjoinedwithwhiteglove":                   WindowsAutopilotEnrollmentType_OfflineDomainJoinedWithWhiteGlove,
		"unknown":                                             WindowsAutopilotEnrollmentType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotEnrollmentType(input)
	return &out, nil
}
