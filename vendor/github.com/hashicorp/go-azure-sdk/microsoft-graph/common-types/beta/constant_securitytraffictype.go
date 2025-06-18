package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTrafficType string

const (
	SecurityTrafficType_DownloadedBytes SecurityTrafficType = "downloadedBytes"
	SecurityTrafficType_Unknown         SecurityTrafficType = "unknown"
	SecurityTrafficType_UploadedBytes   SecurityTrafficType = "uploadedBytes"
)

func PossibleValuesForSecurityTrafficType() []string {
	return []string{
		string(SecurityTrafficType_DownloadedBytes),
		string(SecurityTrafficType_Unknown),
		string(SecurityTrafficType_UploadedBytes),
	}
}

func (s *SecurityTrafficType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityTrafficType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityTrafficType(input string) (*SecurityTrafficType, error) {
	vals := map[string]SecurityTrafficType{
		"downloadedbytes": SecurityTrafficType_DownloadedBytes,
		"unknown":         SecurityTrafficType_Unknown,
		"uploadedbytes":   SecurityTrafficType_UploadedBytes,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityTrafficType(input)
	return &out, nil
}
