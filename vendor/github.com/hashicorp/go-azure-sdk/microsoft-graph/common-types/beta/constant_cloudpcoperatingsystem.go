package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOperatingSystem string

const (
	CloudPCOperatingSystem_Windows10 CloudPCOperatingSystem = "windows10"
	CloudPCOperatingSystem_Windows11 CloudPCOperatingSystem = "windows11"
)

func PossibleValuesForCloudPCOperatingSystem() []string {
	return []string{
		string(CloudPCOperatingSystem_Windows10),
		string(CloudPCOperatingSystem_Windows11),
	}
}

func (s *CloudPCOperatingSystem) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCOperatingSystem(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCOperatingSystem(input string) (*CloudPCOperatingSystem, error) {
	vals := map[string]CloudPCOperatingSystem{
		"windows10": CloudPCOperatingSystem_Windows10,
		"windows11": CloudPCOperatingSystem_Windows11,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOperatingSystem(input)
	return &out, nil
}
