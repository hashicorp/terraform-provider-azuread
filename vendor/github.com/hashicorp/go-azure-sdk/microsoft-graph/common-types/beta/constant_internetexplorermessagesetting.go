package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InternetExplorerMessageSetting string

const (
	InternetExplorerMessageSetting_Disabled      InternetExplorerMessageSetting = "disabled"
	InternetExplorerMessageSetting_Enabled       InternetExplorerMessageSetting = "enabled"
	InternetExplorerMessageSetting_KeepGoing     InternetExplorerMessageSetting = "keepGoing"
	InternetExplorerMessageSetting_NotConfigured InternetExplorerMessageSetting = "notConfigured"
)

func PossibleValuesForInternetExplorerMessageSetting() []string {
	return []string{
		string(InternetExplorerMessageSetting_Disabled),
		string(InternetExplorerMessageSetting_Enabled),
		string(InternetExplorerMessageSetting_KeepGoing),
		string(InternetExplorerMessageSetting_NotConfigured),
	}
}

func (s *InternetExplorerMessageSetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInternetExplorerMessageSetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInternetExplorerMessageSetting(input string) (*InternetExplorerMessageSetting, error) {
	vals := map[string]InternetExplorerMessageSetting{
		"disabled":      InternetExplorerMessageSetting_Disabled,
		"enabled":       InternetExplorerMessageSetting_Enabled,
		"keepgoing":     InternetExplorerMessageSetting_KeepGoing,
		"notconfigured": InternetExplorerMessageSetting_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InternetExplorerMessageSetting(input)
	return &out, nil
}
