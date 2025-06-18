package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppDataStorageLocation string

const (
	ManagedAppDataStorageLocation_Box                 ManagedAppDataStorageLocation = "box"
	ManagedAppDataStorageLocation_Egnyte              ManagedAppDataStorageLocation = "egnyte"
	ManagedAppDataStorageLocation_IManage             ManagedAppDataStorageLocation = "iManage"
	ManagedAppDataStorageLocation_LocalStorage        ManagedAppDataStorageLocation = "localStorage"
	ManagedAppDataStorageLocation_OneDriveForBusiness ManagedAppDataStorageLocation = "oneDriveForBusiness"
	ManagedAppDataStorageLocation_PhotoLibrary        ManagedAppDataStorageLocation = "photoLibrary"
	ManagedAppDataStorageLocation_SharePoint          ManagedAppDataStorageLocation = "sharePoint"
)

func PossibleValuesForManagedAppDataStorageLocation() []string {
	return []string{
		string(ManagedAppDataStorageLocation_Box),
		string(ManagedAppDataStorageLocation_Egnyte),
		string(ManagedAppDataStorageLocation_IManage),
		string(ManagedAppDataStorageLocation_LocalStorage),
		string(ManagedAppDataStorageLocation_OneDriveForBusiness),
		string(ManagedAppDataStorageLocation_PhotoLibrary),
		string(ManagedAppDataStorageLocation_SharePoint),
	}
}

func (s *ManagedAppDataStorageLocation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppDataStorageLocation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppDataStorageLocation(input string) (*ManagedAppDataStorageLocation, error) {
	vals := map[string]ManagedAppDataStorageLocation{
		"box":                 ManagedAppDataStorageLocation_Box,
		"egnyte":              ManagedAppDataStorageLocation_Egnyte,
		"imanage":             ManagedAppDataStorageLocation_IManage,
		"localstorage":        ManagedAppDataStorageLocation_LocalStorage,
		"onedriveforbusiness": ManagedAppDataStorageLocation_OneDriveForBusiness,
		"photolibrary":        ManagedAppDataStorageLocation_PhotoLibrary,
		"sharepoint":          ManagedAppDataStorageLocation_SharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppDataStorageLocation(input)
	return &out, nil
}
