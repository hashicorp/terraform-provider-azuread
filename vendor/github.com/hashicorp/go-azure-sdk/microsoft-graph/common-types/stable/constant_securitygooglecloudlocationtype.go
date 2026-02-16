package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGoogleCloudLocationType string

const (
	SecurityGoogleCloudLocationType_Global   SecurityGoogleCloudLocationType = "global"
	SecurityGoogleCloudLocationType_Regional SecurityGoogleCloudLocationType = "regional"
	SecurityGoogleCloudLocationType_Unknown  SecurityGoogleCloudLocationType = "unknown"
	SecurityGoogleCloudLocationType_Zonal    SecurityGoogleCloudLocationType = "zonal"
)

func PossibleValuesForSecurityGoogleCloudLocationType() []string {
	return []string{
		string(SecurityGoogleCloudLocationType_Global),
		string(SecurityGoogleCloudLocationType_Regional),
		string(SecurityGoogleCloudLocationType_Unknown),
		string(SecurityGoogleCloudLocationType_Zonal),
	}
}

func (s *SecurityGoogleCloudLocationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityGoogleCloudLocationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityGoogleCloudLocationType(input string) (*SecurityGoogleCloudLocationType, error) {
	vals := map[string]SecurityGoogleCloudLocationType{
		"global":   SecurityGoogleCloudLocationType_Global,
		"regional": SecurityGoogleCloudLocationType_Regional,
		"unknown":  SecurityGoogleCloudLocationType_Unknown,
		"zonal":    SecurityGoogleCloudLocationType_Zonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGoogleCloudLocationType(input)
	return &out, nil
}
