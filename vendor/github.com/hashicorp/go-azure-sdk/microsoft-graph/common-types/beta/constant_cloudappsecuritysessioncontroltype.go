package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudAppSecuritySessionControlType string

const (
	CloudAppSecuritySessionControlType_BlockDownloads CloudAppSecuritySessionControlType = "blockDownloads"
	CloudAppSecuritySessionControlType_McasConfigured CloudAppSecuritySessionControlType = "mcasConfigured"
	CloudAppSecuritySessionControlType_MonitorOnly    CloudAppSecuritySessionControlType = "monitorOnly"
)

func PossibleValuesForCloudAppSecuritySessionControlType() []string {
	return []string{
		string(CloudAppSecuritySessionControlType_BlockDownloads),
		string(CloudAppSecuritySessionControlType_McasConfigured),
		string(CloudAppSecuritySessionControlType_MonitorOnly),
	}
}

func (s *CloudAppSecuritySessionControlType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudAppSecuritySessionControlType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudAppSecuritySessionControlType(input string) (*CloudAppSecuritySessionControlType, error) {
	vals := map[string]CloudAppSecuritySessionControlType{
		"blockdownloads": CloudAppSecuritySessionControlType_BlockDownloads,
		"mcasconfigured": CloudAppSecuritySessionControlType_McasConfigured,
		"monitoronly":    CloudAppSecuritySessionControlType_MonitorOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudAppSecuritySessionControlType(input)
	return &out, nil
}
