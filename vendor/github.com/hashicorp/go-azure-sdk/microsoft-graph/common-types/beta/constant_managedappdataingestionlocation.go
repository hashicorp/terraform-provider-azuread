package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppDataIngestionLocation string

const (
	ManagedAppDataIngestionLocation_Camera              ManagedAppDataIngestionLocation = "camera"
	ManagedAppDataIngestionLocation_OneDriveForBusiness ManagedAppDataIngestionLocation = "oneDriveForBusiness"
	ManagedAppDataIngestionLocation_PhotoLibrary        ManagedAppDataIngestionLocation = "photoLibrary"
	ManagedAppDataIngestionLocation_SharePoint          ManagedAppDataIngestionLocation = "sharePoint"
)

func PossibleValuesForManagedAppDataIngestionLocation() []string {
	return []string{
		string(ManagedAppDataIngestionLocation_Camera),
		string(ManagedAppDataIngestionLocation_OneDriveForBusiness),
		string(ManagedAppDataIngestionLocation_PhotoLibrary),
		string(ManagedAppDataIngestionLocation_SharePoint),
	}
}

func (s *ManagedAppDataIngestionLocation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppDataIngestionLocation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppDataIngestionLocation(input string) (*ManagedAppDataIngestionLocation, error) {
	vals := map[string]ManagedAppDataIngestionLocation{
		"camera":              ManagedAppDataIngestionLocation_Camera,
		"onedriveforbusiness": ManagedAppDataIngestionLocation_OneDriveForBusiness,
		"photolibrary":        ManagedAppDataIngestionLocation_PhotoLibrary,
		"sharepoint":          ManagedAppDataIngestionLocation_SharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppDataIngestionLocation(input)
	return &out, nil
}
