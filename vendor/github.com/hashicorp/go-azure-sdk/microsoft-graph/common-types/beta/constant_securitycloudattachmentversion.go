package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityCloudAttachmentVersion string

const (
	SecurityCloudAttachmentVersion_All       SecurityCloudAttachmentVersion = "all"
	SecurityCloudAttachmentVersion_Latest    SecurityCloudAttachmentVersion = "latest"
	SecurityCloudAttachmentVersion_Recent10  SecurityCloudAttachmentVersion = "recent10"
	SecurityCloudAttachmentVersion_Recent100 SecurityCloudAttachmentVersion = "recent100"
)

func PossibleValuesForSecurityCloudAttachmentVersion() []string {
	return []string{
		string(SecurityCloudAttachmentVersion_All),
		string(SecurityCloudAttachmentVersion_Latest),
		string(SecurityCloudAttachmentVersion_Recent10),
		string(SecurityCloudAttachmentVersion_Recent100),
	}
}

func (s *SecurityCloudAttachmentVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityCloudAttachmentVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityCloudAttachmentVersion(input string) (*SecurityCloudAttachmentVersion, error) {
	vals := map[string]SecurityCloudAttachmentVersion{
		"all":       SecurityCloudAttachmentVersion_All,
		"latest":    SecurityCloudAttachmentVersion_Latest,
		"recent10":  SecurityCloudAttachmentVersion_Recent10,
		"recent100": SecurityCloudAttachmentVersion_Recent100,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityCloudAttachmentVersion(input)
	return &out, nil
}
