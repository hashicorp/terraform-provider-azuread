package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidKeyguardFeature string

const (
	AndroidKeyguardFeature_AllFeatures             AndroidKeyguardFeature = "allFeatures"
	AndroidKeyguardFeature_Biometrics              AndroidKeyguardFeature = "biometrics"
	AndroidKeyguardFeature_Camera                  AndroidKeyguardFeature = "camera"
	AndroidKeyguardFeature_Face                    AndroidKeyguardFeature = "face"
	AndroidKeyguardFeature_Fingerprint             AndroidKeyguardFeature = "fingerprint"
	AndroidKeyguardFeature_Iris                    AndroidKeyguardFeature = "iris"
	AndroidKeyguardFeature_NotConfigured           AndroidKeyguardFeature = "notConfigured"
	AndroidKeyguardFeature_Notifications           AndroidKeyguardFeature = "notifications"
	AndroidKeyguardFeature_RemoteInput             AndroidKeyguardFeature = "remoteInput"
	AndroidKeyguardFeature_TrustAgents             AndroidKeyguardFeature = "trustAgents"
	AndroidKeyguardFeature_UnredactedNotifications AndroidKeyguardFeature = "unredactedNotifications"
)

func PossibleValuesForAndroidKeyguardFeature() []string {
	return []string{
		string(AndroidKeyguardFeature_AllFeatures),
		string(AndroidKeyguardFeature_Biometrics),
		string(AndroidKeyguardFeature_Camera),
		string(AndroidKeyguardFeature_Face),
		string(AndroidKeyguardFeature_Fingerprint),
		string(AndroidKeyguardFeature_Iris),
		string(AndroidKeyguardFeature_NotConfigured),
		string(AndroidKeyguardFeature_Notifications),
		string(AndroidKeyguardFeature_RemoteInput),
		string(AndroidKeyguardFeature_TrustAgents),
		string(AndroidKeyguardFeature_UnredactedNotifications),
	}
}

func (s *AndroidKeyguardFeature) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidKeyguardFeature(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidKeyguardFeature(input string) (*AndroidKeyguardFeature, error) {
	vals := map[string]AndroidKeyguardFeature{
		"allfeatures":             AndroidKeyguardFeature_AllFeatures,
		"biometrics":              AndroidKeyguardFeature_Biometrics,
		"camera":                  AndroidKeyguardFeature_Camera,
		"face":                    AndroidKeyguardFeature_Face,
		"fingerprint":             AndroidKeyguardFeature_Fingerprint,
		"iris":                    AndroidKeyguardFeature_Iris,
		"notconfigured":           AndroidKeyguardFeature_NotConfigured,
		"notifications":           AndroidKeyguardFeature_Notifications,
		"remoteinput":             AndroidKeyguardFeature_RemoteInput,
		"trustagents":             AndroidKeyguardFeature_TrustAgents,
		"unredactednotifications": AndroidKeyguardFeature_UnredactedNotifications,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidKeyguardFeature(input)
	return &out, nil
}
