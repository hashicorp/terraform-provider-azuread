package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChassisType string

const (
	ChassisType_Desktop          ChassisType = "desktop"
	ChassisType_EnterpriseServer ChassisType = "enterpriseServer"
	ChassisType_Laptop           ChassisType = "laptop"
	ChassisType_MobileOther      ChassisType = "mobileOther"
	ChassisType_MobileUnknown    ChassisType = "mobileUnknown"
	ChassisType_Phone            ChassisType = "phone"
	ChassisType_Tablet           ChassisType = "tablet"
	ChassisType_Unknown          ChassisType = "unknown"
	ChassisType_WorksWorkstation ChassisType = "worksWorkstation"
)

func PossibleValuesForChassisType() []string {
	return []string{
		string(ChassisType_Desktop),
		string(ChassisType_EnterpriseServer),
		string(ChassisType_Laptop),
		string(ChassisType_MobileOther),
		string(ChassisType_MobileUnknown),
		string(ChassisType_Phone),
		string(ChassisType_Tablet),
		string(ChassisType_Unknown),
		string(ChassisType_WorksWorkstation),
	}
}

func (s *ChassisType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseChassisType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseChassisType(input string) (*ChassisType, error) {
	vals := map[string]ChassisType{
		"desktop":          ChassisType_Desktop,
		"enterpriseserver": ChassisType_EnterpriseServer,
		"laptop":           ChassisType_Laptop,
		"mobileother":      ChassisType_MobileOther,
		"mobileunknown":    ChassisType_MobileUnknown,
		"phone":            ChassisType_Phone,
		"tablet":           ChassisType_Tablet,
		"unknown":          ChassisType_Unknown,
		"worksworkstation": ChassisType_WorksWorkstation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ChassisType(input)
	return &out, nil
}
