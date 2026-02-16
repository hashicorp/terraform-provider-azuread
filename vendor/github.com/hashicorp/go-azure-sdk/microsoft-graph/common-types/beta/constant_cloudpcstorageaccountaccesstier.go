package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCStorageAccountAccessTier string

const (
	CloudPCStorageAccountAccessTier_Cold    CloudPCStorageAccountAccessTier = "cold"
	CloudPCStorageAccountAccessTier_Cool    CloudPCStorageAccountAccessTier = "cool"
	CloudPCStorageAccountAccessTier_Hot     CloudPCStorageAccountAccessTier = "hot"
	CloudPCStorageAccountAccessTier_Premium CloudPCStorageAccountAccessTier = "premium"
)

func PossibleValuesForCloudPCStorageAccountAccessTier() []string {
	return []string{
		string(CloudPCStorageAccountAccessTier_Cold),
		string(CloudPCStorageAccountAccessTier_Cool),
		string(CloudPCStorageAccountAccessTier_Hot),
		string(CloudPCStorageAccountAccessTier_Premium),
	}
}

func (s *CloudPCStorageAccountAccessTier) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCStorageAccountAccessTier(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCStorageAccountAccessTier(input string) (*CloudPCStorageAccountAccessTier, error) {
	vals := map[string]CloudPCStorageAccountAccessTier{
		"cold":    CloudPCStorageAccountAccessTier_Cold,
		"cool":    CloudPCStorageAccountAccessTier_Cool,
		"hot":     CloudPCStorageAccountAccessTier_Hot,
		"premium": CloudPCStorageAccountAccessTier_Premium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCStorageAccountAccessTier(input)
	return &out, nil
}
