package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppDependencyType string

const (
	MobileAppDependencyType_AutoInstall MobileAppDependencyType = "autoInstall"
	MobileAppDependencyType_Detect      MobileAppDependencyType = "detect"
)

func PossibleValuesForMobileAppDependencyType() []string {
	return []string{
		string(MobileAppDependencyType_AutoInstall),
		string(MobileAppDependencyType_Detect),
	}
}

func (s *MobileAppDependencyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppDependencyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppDependencyType(input string) (*MobileAppDependencyType, error) {
	vals := map[string]MobileAppDependencyType{
		"autoinstall": MobileAppDependencyType_AutoInstall,
		"detect":      MobileAppDependencyType_Detect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppDependencyType(input)
	return &out, nil
}
