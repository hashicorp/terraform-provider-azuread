package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppSupersedenceType string

const (
	MobileAppSupersedenceType_Replace MobileAppSupersedenceType = "replace"
	MobileAppSupersedenceType_Update  MobileAppSupersedenceType = "update"
)

func PossibleValuesForMobileAppSupersedenceType() []string {
	return []string{
		string(MobileAppSupersedenceType_Replace),
		string(MobileAppSupersedenceType_Update),
	}
}

func (s *MobileAppSupersedenceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMobileAppSupersedenceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMobileAppSupersedenceType(input string) (*MobileAppSupersedenceType, error) {
	vals := map[string]MobileAppSupersedenceType{
		"replace": MobileAppSupersedenceType_Replace,
		"update":  MobileAppSupersedenceType_Update,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MobileAppSupersedenceType(input)
	return &out, nil
}
