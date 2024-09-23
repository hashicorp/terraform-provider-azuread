package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCrossProfileDataSharing string

const (
	AndroidDeviceOwnerCrossProfileDataSharing_CrossProfileDataSharingAllowed       AndroidDeviceOwnerCrossProfileDataSharing = "crossProfileDataSharingAllowed"
	AndroidDeviceOwnerCrossProfileDataSharing_CrossProfileDataSharingBlocked       AndroidDeviceOwnerCrossProfileDataSharing = "crossProfileDataSharingBlocked"
	AndroidDeviceOwnerCrossProfileDataSharing_DataSharingFromWorkToPersonalBlocked AndroidDeviceOwnerCrossProfileDataSharing = "dataSharingFromWorkToPersonalBlocked"
	AndroidDeviceOwnerCrossProfileDataSharing_NotConfigured                        AndroidDeviceOwnerCrossProfileDataSharing = "notConfigured"
	AndroidDeviceOwnerCrossProfileDataSharing_UnkownFutureValue                    AndroidDeviceOwnerCrossProfileDataSharing = "unkownFutureValue"
)

func PossibleValuesForAndroidDeviceOwnerCrossProfileDataSharing() []string {
	return []string{
		string(AndroidDeviceOwnerCrossProfileDataSharing_CrossProfileDataSharingAllowed),
		string(AndroidDeviceOwnerCrossProfileDataSharing_CrossProfileDataSharingBlocked),
		string(AndroidDeviceOwnerCrossProfileDataSharing_DataSharingFromWorkToPersonalBlocked),
		string(AndroidDeviceOwnerCrossProfileDataSharing_NotConfigured),
		string(AndroidDeviceOwnerCrossProfileDataSharing_UnkownFutureValue),
	}
}

func (s *AndroidDeviceOwnerCrossProfileDataSharing) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerCrossProfileDataSharing(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerCrossProfileDataSharing(input string) (*AndroidDeviceOwnerCrossProfileDataSharing, error) {
	vals := map[string]AndroidDeviceOwnerCrossProfileDataSharing{
		"crossprofiledatasharingallowed":       AndroidDeviceOwnerCrossProfileDataSharing_CrossProfileDataSharingAllowed,
		"crossprofiledatasharingblocked":       AndroidDeviceOwnerCrossProfileDataSharing_CrossProfileDataSharingBlocked,
		"datasharingfromworktopersonalblocked": AndroidDeviceOwnerCrossProfileDataSharing_DataSharingFromWorkToPersonalBlocked,
		"notconfigured":                        AndroidDeviceOwnerCrossProfileDataSharing_NotConfigured,
		"unkownfuturevalue":                    AndroidDeviceOwnerCrossProfileDataSharing_UnkownFutureValue,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerCrossProfileDataSharing(input)
	return &out, nil
}
