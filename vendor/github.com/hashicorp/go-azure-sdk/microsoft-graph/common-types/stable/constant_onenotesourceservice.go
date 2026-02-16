package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSourceService string

const (
	OnenoteSourceService_OnPremOneDriveForBusiness OnenoteSourceService = "OnPremOneDriveForBusiness"
	OnenoteSourceService_OneDrive                  OnenoteSourceService = "OneDrive"
	OnenoteSourceService_OneDriveForBusiness       OnenoteSourceService = "OneDriveForBusiness"
	OnenoteSourceService_Unknown                   OnenoteSourceService = "Unknown"
)

func PossibleValuesForOnenoteSourceService() []string {
	return []string{
		string(OnenoteSourceService_OnPremOneDriveForBusiness),
		string(OnenoteSourceService_OneDrive),
		string(OnenoteSourceService_OneDriveForBusiness),
		string(OnenoteSourceService_Unknown),
	}
}

func (s *OnenoteSourceService) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnenoteSourceService(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnenoteSourceService(input string) (*OnenoteSourceService, error) {
	vals := map[string]OnenoteSourceService{
		"onpremonedriveforbusiness": OnenoteSourceService_OnPremOneDriveForBusiness,
		"onedrive":                  OnenoteSourceService_OneDrive,
		"onedriveforbusiness":       OnenoteSourceService_OneDriveForBusiness,
		"unknown":                   OnenoteSourceService_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnenoteSourceService(input)
	return &out, nil
}
