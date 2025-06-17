package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCBlobAccessTier string

const (
	CloudPCBlobAccessTier_Archive CloudPCBlobAccessTier = "archive"
	CloudPCBlobAccessTier_Cold    CloudPCBlobAccessTier = "cold"
	CloudPCBlobAccessTier_Cool    CloudPCBlobAccessTier = "cool"
	CloudPCBlobAccessTier_Hot     CloudPCBlobAccessTier = "hot"
)

func PossibleValuesForCloudPCBlobAccessTier() []string {
	return []string{
		string(CloudPCBlobAccessTier_Archive),
		string(CloudPCBlobAccessTier_Cold),
		string(CloudPCBlobAccessTier_Cool),
		string(CloudPCBlobAccessTier_Hot),
	}
}

func (s *CloudPCBlobAccessTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCBlobAccessTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCBlobAccessTier(input string) (*CloudPCBlobAccessTier, error) {
	vals := map[string]CloudPCBlobAccessTier{
		"archive": CloudPCBlobAccessTier_Archive,
		"cold":    CloudPCBlobAccessTier_Cold,
		"cool":    CloudPCBlobAccessTier_Cool,
		"hot":     CloudPCBlobAccessTier_Hot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCBlobAccessTier(input)
	return &out, nil
}
