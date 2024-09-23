package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerCertificateAccessType string

const (
	AndroidDeviceOwnerCertificateAccessType_SpecificApps AndroidDeviceOwnerCertificateAccessType = "specificApps"
	AndroidDeviceOwnerCertificateAccessType_UserApproval AndroidDeviceOwnerCertificateAccessType = "userApproval"
)

func PossibleValuesForAndroidDeviceOwnerCertificateAccessType() []string {
	return []string{
		string(AndroidDeviceOwnerCertificateAccessType_SpecificApps),
		string(AndroidDeviceOwnerCertificateAccessType_UserApproval),
	}
}

func (s *AndroidDeviceOwnerCertificateAccessType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerCertificateAccessType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerCertificateAccessType(input string) (*AndroidDeviceOwnerCertificateAccessType, error) {
	vals := map[string]AndroidDeviceOwnerCertificateAccessType{
		"specificapps": AndroidDeviceOwnerCertificateAccessType_SpecificApps,
		"userapproval": AndroidDeviceOwnerCertificateAccessType_UserApproval,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerCertificateAccessType(input)
	return &out, nil
}
