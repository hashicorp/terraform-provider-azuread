package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPrivacyDataCategory string

const (
	WindowsPrivacyDataCategory_AccountInfo         WindowsPrivacyDataCategory = "accountInfo"
	WindowsPrivacyDataCategory_AppsRunInBackground WindowsPrivacyDataCategory = "appsRunInBackground"
	WindowsPrivacyDataCategory_Calendar            WindowsPrivacyDataCategory = "calendar"
	WindowsPrivacyDataCategory_CallHistory         WindowsPrivacyDataCategory = "callHistory"
	WindowsPrivacyDataCategory_Camera              WindowsPrivacyDataCategory = "camera"
	WindowsPrivacyDataCategory_Contacts            WindowsPrivacyDataCategory = "contacts"
	WindowsPrivacyDataCategory_DiagnosticsInfo     WindowsPrivacyDataCategory = "diagnosticsInfo"
	WindowsPrivacyDataCategory_Email               WindowsPrivacyDataCategory = "email"
	WindowsPrivacyDataCategory_Location            WindowsPrivacyDataCategory = "location"
	WindowsPrivacyDataCategory_Messaging           WindowsPrivacyDataCategory = "messaging"
	WindowsPrivacyDataCategory_Microphone          WindowsPrivacyDataCategory = "microphone"
	WindowsPrivacyDataCategory_Motion              WindowsPrivacyDataCategory = "motion"
	WindowsPrivacyDataCategory_NotConfigured       WindowsPrivacyDataCategory = "notConfigured"
	WindowsPrivacyDataCategory_Notifications       WindowsPrivacyDataCategory = "notifications"
	WindowsPrivacyDataCategory_Phone               WindowsPrivacyDataCategory = "phone"
	WindowsPrivacyDataCategory_Radios              WindowsPrivacyDataCategory = "radios"
	WindowsPrivacyDataCategory_SyncWithDevices     WindowsPrivacyDataCategory = "syncWithDevices"
	WindowsPrivacyDataCategory_Tasks               WindowsPrivacyDataCategory = "tasks"
	WindowsPrivacyDataCategory_TrustedDevices      WindowsPrivacyDataCategory = "trustedDevices"
)

func PossibleValuesForWindowsPrivacyDataCategory() []string {
	return []string{
		string(WindowsPrivacyDataCategory_AccountInfo),
		string(WindowsPrivacyDataCategory_AppsRunInBackground),
		string(WindowsPrivacyDataCategory_Calendar),
		string(WindowsPrivacyDataCategory_CallHistory),
		string(WindowsPrivacyDataCategory_Camera),
		string(WindowsPrivacyDataCategory_Contacts),
		string(WindowsPrivacyDataCategory_DiagnosticsInfo),
		string(WindowsPrivacyDataCategory_Email),
		string(WindowsPrivacyDataCategory_Location),
		string(WindowsPrivacyDataCategory_Messaging),
		string(WindowsPrivacyDataCategory_Microphone),
		string(WindowsPrivacyDataCategory_Motion),
		string(WindowsPrivacyDataCategory_NotConfigured),
		string(WindowsPrivacyDataCategory_Notifications),
		string(WindowsPrivacyDataCategory_Phone),
		string(WindowsPrivacyDataCategory_Radios),
		string(WindowsPrivacyDataCategory_SyncWithDevices),
		string(WindowsPrivacyDataCategory_Tasks),
		string(WindowsPrivacyDataCategory_TrustedDevices),
	}
}

func (s *WindowsPrivacyDataCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPrivacyDataCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPrivacyDataCategory(input string) (*WindowsPrivacyDataCategory, error) {
	vals := map[string]WindowsPrivacyDataCategory{
		"accountinfo":         WindowsPrivacyDataCategory_AccountInfo,
		"appsruninbackground": WindowsPrivacyDataCategory_AppsRunInBackground,
		"calendar":            WindowsPrivacyDataCategory_Calendar,
		"callhistory":         WindowsPrivacyDataCategory_CallHistory,
		"camera":              WindowsPrivacyDataCategory_Camera,
		"contacts":            WindowsPrivacyDataCategory_Contacts,
		"diagnosticsinfo":     WindowsPrivacyDataCategory_DiagnosticsInfo,
		"email":               WindowsPrivacyDataCategory_Email,
		"location":            WindowsPrivacyDataCategory_Location,
		"messaging":           WindowsPrivacyDataCategory_Messaging,
		"microphone":          WindowsPrivacyDataCategory_Microphone,
		"motion":              WindowsPrivacyDataCategory_Motion,
		"notconfigured":       WindowsPrivacyDataCategory_NotConfigured,
		"notifications":       WindowsPrivacyDataCategory_Notifications,
		"phone":               WindowsPrivacyDataCategory_Phone,
		"radios":              WindowsPrivacyDataCategory_Radios,
		"syncwithdevices":     WindowsPrivacyDataCategory_SyncWithDevices,
		"tasks":               WindowsPrivacyDataCategory_Tasks,
		"trusteddevices":      WindowsPrivacyDataCategory_TrustedDevices,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPrivacyDataCategory(input)
	return &out, nil
}
