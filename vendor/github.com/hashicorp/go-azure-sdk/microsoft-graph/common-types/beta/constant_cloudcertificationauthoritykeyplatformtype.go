package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudCertificationAuthorityKeyPlatformType string

const (
	CloudCertificationAuthorityKeyPlatformType_HardwareSecurityModule CloudCertificationAuthorityKeyPlatformType = "hardwareSecurityModule"
	CloudCertificationAuthorityKeyPlatformType_Software               CloudCertificationAuthorityKeyPlatformType = "software"
	CloudCertificationAuthorityKeyPlatformType_Unknown                CloudCertificationAuthorityKeyPlatformType = "unknown"
)

func PossibleValuesForCloudCertificationAuthorityKeyPlatformType() []string {
	return []string{
		string(CloudCertificationAuthorityKeyPlatformType_HardwareSecurityModule),
		string(CloudCertificationAuthorityKeyPlatformType_Software),
		string(CloudCertificationAuthorityKeyPlatformType_Unknown),
	}
}

func (s *CloudCertificationAuthorityKeyPlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudCertificationAuthorityKeyPlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudCertificationAuthorityKeyPlatformType(input string) (*CloudCertificationAuthorityKeyPlatformType, error) {
	vals := map[string]CloudCertificationAuthorityKeyPlatformType{
		"hardwaresecuritymodule": CloudCertificationAuthorityKeyPlatformType_HardwareSecurityModule,
		"software":               CloudCertificationAuthorityKeyPlatformType_Software,
		"unknown":                CloudCertificationAuthorityKeyPlatformType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudCertificationAuthorityKeyPlatformType(input)
	return &out, nil
}
