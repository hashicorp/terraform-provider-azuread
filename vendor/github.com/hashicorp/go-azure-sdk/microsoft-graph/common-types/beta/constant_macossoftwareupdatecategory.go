package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateCategory string

const (
	MacOSSoftwareUpdateCategory_ConfigurationDataFile MacOSSoftwareUpdateCategory = "configurationDataFile"
	MacOSSoftwareUpdateCategory_Critical              MacOSSoftwareUpdateCategory = "critical"
	MacOSSoftwareUpdateCategory_Firmware              MacOSSoftwareUpdateCategory = "firmware"
	MacOSSoftwareUpdateCategory_Other                 MacOSSoftwareUpdateCategory = "other"
)

func PossibleValuesForMacOSSoftwareUpdateCategory() []string {
	return []string{
		string(MacOSSoftwareUpdateCategory_ConfigurationDataFile),
		string(MacOSSoftwareUpdateCategory_Critical),
		string(MacOSSoftwareUpdateCategory_Firmware),
		string(MacOSSoftwareUpdateCategory_Other),
	}
}

func (s *MacOSSoftwareUpdateCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSoftwareUpdateCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSoftwareUpdateCategory(input string) (*MacOSSoftwareUpdateCategory, error) {
	vals := map[string]MacOSSoftwareUpdateCategory{
		"configurationdatafile": MacOSSoftwareUpdateCategory_ConfigurationDataFile,
		"critical":              MacOSSoftwareUpdateCategory_Critical,
		"firmware":              MacOSSoftwareUpdateCategory_Firmware,
		"other":                 MacOSSoftwareUpdateCategory_Other,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateCategory(input)
	return &out, nil
}
