package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageRequestType string

const (
	AccessPackageRequestType_AdminAdd     AccessPackageRequestType = "adminAdd"
	AccessPackageRequestType_AdminRemove  AccessPackageRequestType = "adminRemove"
	AccessPackageRequestType_AdminUpdate  AccessPackageRequestType = "adminUpdate"
	AccessPackageRequestType_NotSpecified AccessPackageRequestType = "notSpecified"
	AccessPackageRequestType_OnBehalfAdd  AccessPackageRequestType = "onBehalfAdd"
	AccessPackageRequestType_SystemAdd    AccessPackageRequestType = "systemAdd"
	AccessPackageRequestType_SystemRemove AccessPackageRequestType = "systemRemove"
	AccessPackageRequestType_SystemUpdate AccessPackageRequestType = "systemUpdate"
	AccessPackageRequestType_UserAdd      AccessPackageRequestType = "userAdd"
	AccessPackageRequestType_UserRemove   AccessPackageRequestType = "userRemove"
	AccessPackageRequestType_UserUpdate   AccessPackageRequestType = "userUpdate"
)

func PossibleValuesForAccessPackageRequestType() []string {
	return []string{
		string(AccessPackageRequestType_AdminAdd),
		string(AccessPackageRequestType_AdminRemove),
		string(AccessPackageRequestType_AdminUpdate),
		string(AccessPackageRequestType_NotSpecified),
		string(AccessPackageRequestType_OnBehalfAdd),
		string(AccessPackageRequestType_SystemAdd),
		string(AccessPackageRequestType_SystemRemove),
		string(AccessPackageRequestType_SystemUpdate),
		string(AccessPackageRequestType_UserAdd),
		string(AccessPackageRequestType_UserRemove),
		string(AccessPackageRequestType_UserUpdate),
	}
}

func (s *AccessPackageRequestType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageRequestType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageRequestType(input string) (*AccessPackageRequestType, error) {
	vals := map[string]AccessPackageRequestType{
		"adminadd":     AccessPackageRequestType_AdminAdd,
		"adminremove":  AccessPackageRequestType_AdminRemove,
		"adminupdate":  AccessPackageRequestType_AdminUpdate,
		"notspecified": AccessPackageRequestType_NotSpecified,
		"onbehalfadd":  AccessPackageRequestType_OnBehalfAdd,
		"systemadd":    AccessPackageRequestType_SystemAdd,
		"systemremove": AccessPackageRequestType_SystemRemove,
		"systemupdate": AccessPackageRequestType_SystemUpdate,
		"useradd":      AccessPackageRequestType_UserAdd,
		"userremove":   AccessPackageRequestType_UserRemove,
		"userupdate":   AccessPackageRequestType_UserUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageRequestType(input)
	return &out, nil
}
