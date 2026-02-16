package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageCatalogType string

const (
	AccessPackageCatalogType_ServiceDefault AccessPackageCatalogType = "serviceDefault"
	AccessPackageCatalogType_ServiceManaged AccessPackageCatalogType = "serviceManaged"
	AccessPackageCatalogType_UserManaged    AccessPackageCatalogType = "userManaged"
)

func PossibleValuesForAccessPackageCatalogType() []string {
	return []string{
		string(AccessPackageCatalogType_ServiceDefault),
		string(AccessPackageCatalogType_ServiceManaged),
		string(AccessPackageCatalogType_UserManaged),
	}
}

func (s *AccessPackageCatalogType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAccessPackageCatalogType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAccessPackageCatalogType(input string) (*AccessPackageCatalogType, error) {
	vals := map[string]AccessPackageCatalogType{
		"servicedefault": AccessPackageCatalogType_ServiceDefault,
		"servicemanaged": AccessPackageCatalogType_ServiceManaged,
		"usermanaged":    AccessPackageCatalogType_UserManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageCatalogType(input)
	return &out, nil
}
