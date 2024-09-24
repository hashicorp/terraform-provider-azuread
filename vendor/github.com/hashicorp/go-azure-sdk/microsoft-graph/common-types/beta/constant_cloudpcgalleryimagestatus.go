package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCGalleryImageStatus string

const (
	CloudPCGalleryImageStatus_NotSupported         CloudPCGalleryImageStatus = "notSupported"
	CloudPCGalleryImageStatus_Supported            CloudPCGalleryImageStatus = "supported"
	CloudPCGalleryImageStatus_SupportedWithWarning CloudPCGalleryImageStatus = "supportedWithWarning"
)

func PossibleValuesForCloudPCGalleryImageStatus() []string {
	return []string{
		string(CloudPCGalleryImageStatus_NotSupported),
		string(CloudPCGalleryImageStatus_Supported),
		string(CloudPCGalleryImageStatus_SupportedWithWarning),
	}
}

func (s *CloudPCGalleryImageStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCGalleryImageStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCGalleryImageStatus(input string) (*CloudPCGalleryImageStatus, error) {
	vals := map[string]CloudPCGalleryImageStatus{
		"notsupported":         CloudPCGalleryImageStatus_NotSupported,
		"supported":            CloudPCGalleryImageStatus_Supported,
		"supportedwithwarning": CloudPCGalleryImageStatus_SupportedWithWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCGalleryImageStatus(input)
	return &out, nil
}
