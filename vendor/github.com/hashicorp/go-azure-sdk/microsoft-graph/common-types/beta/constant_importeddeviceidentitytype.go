package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityType string

const (
	ImportedDeviceIdentityType_Imei                    ImportedDeviceIdentityType = "imei"
	ImportedDeviceIdentityType_ManufacturerModelSerial ImportedDeviceIdentityType = "manufacturerModelSerial"
	ImportedDeviceIdentityType_SerialNumber            ImportedDeviceIdentityType = "serialNumber"
	ImportedDeviceIdentityType_Unknown                 ImportedDeviceIdentityType = "unknown"
)

func PossibleValuesForImportedDeviceIdentityType() []string {
	return []string{
		string(ImportedDeviceIdentityType_Imei),
		string(ImportedDeviceIdentityType_ManufacturerModelSerial),
		string(ImportedDeviceIdentityType_SerialNumber),
		string(ImportedDeviceIdentityType_Unknown),
	}
}

func (s *ImportedDeviceIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedDeviceIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedDeviceIdentityType(input string) (*ImportedDeviceIdentityType, error) {
	vals := map[string]ImportedDeviceIdentityType{
		"imei":                    ImportedDeviceIdentityType_Imei,
		"manufacturermodelserial": ImportedDeviceIdentityType_ManufacturerModelSerial,
		"serialnumber":            ImportedDeviceIdentityType_SerialNumber,
		"unknown":                 ImportedDeviceIdentityType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityType(input)
	return &out, nil
}
