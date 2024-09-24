package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EditionUpgradeLicenseType string

const (
	EditionUpgradeLicenseType_LicenseFile   EditionUpgradeLicenseType = "licenseFile"
	EditionUpgradeLicenseType_NotConfigured EditionUpgradeLicenseType = "notConfigured"
	EditionUpgradeLicenseType_ProductKey    EditionUpgradeLicenseType = "productKey"
)

func PossibleValuesForEditionUpgradeLicenseType() []string {
	return []string{
		string(EditionUpgradeLicenseType_LicenseFile),
		string(EditionUpgradeLicenseType_NotConfigured),
		string(EditionUpgradeLicenseType_ProductKey),
	}
}

func (s *EditionUpgradeLicenseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEditionUpgradeLicenseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEditionUpgradeLicenseType(input string) (*EditionUpgradeLicenseType, error) {
	vals := map[string]EditionUpgradeLicenseType{
		"licensefile":   EditionUpgradeLicenseType_LicenseFile,
		"notconfigured": EditionUpgradeLicenseType_NotConfigured,
		"productkey":    EditionUpgradeLicenseType_ProductKey,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EditionUpgradeLicenseType(input)
	return &out, nil
}
