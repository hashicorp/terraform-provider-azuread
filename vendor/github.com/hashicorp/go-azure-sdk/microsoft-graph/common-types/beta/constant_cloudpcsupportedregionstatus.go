package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSupportedRegionStatus string

const (
	CloudPCSupportedRegionStatus_Available   CloudPCSupportedRegionStatus = "available"
	CloudPCSupportedRegionStatus_Restricted  CloudPCSupportedRegionStatus = "restricted"
	CloudPCSupportedRegionStatus_Unavailable CloudPCSupportedRegionStatus = "unavailable"
)

func PossibleValuesForCloudPCSupportedRegionStatus() []string {
	return []string{
		string(CloudPCSupportedRegionStatus_Available),
		string(CloudPCSupportedRegionStatus_Restricted),
		string(CloudPCSupportedRegionStatus_Unavailable),
	}
}

func (s *CloudPCSupportedRegionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCSupportedRegionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCSupportedRegionStatus(input string) (*CloudPCSupportedRegionStatus, error) {
	vals := map[string]CloudPCSupportedRegionStatus{
		"available":   CloudPCSupportedRegionStatus_Available,
		"restricted":  CloudPCSupportedRegionStatus_Restricted,
		"unavailable": CloudPCSupportedRegionStatus_Unavailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCSupportedRegionStatus(input)
	return &out, nil
}
