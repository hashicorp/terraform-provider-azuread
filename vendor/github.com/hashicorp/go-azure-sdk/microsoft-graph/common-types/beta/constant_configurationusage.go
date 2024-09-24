package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationUsage string

const (
	ConfigurationUsage_Allowed       ConfigurationUsage = "allowed"
	ConfigurationUsage_Blocked       ConfigurationUsage = "blocked"
	ConfigurationUsage_NotConfigured ConfigurationUsage = "notConfigured"
	ConfigurationUsage_Required      ConfigurationUsage = "required"
)

func PossibleValuesForConfigurationUsage() []string {
	return []string{
		string(ConfigurationUsage_Allowed),
		string(ConfigurationUsage_Blocked),
		string(ConfigurationUsage_NotConfigured),
		string(ConfigurationUsage_Required),
	}
}

func (s *ConfigurationUsage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConfigurationUsage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConfigurationUsage(input string) (*ConfigurationUsage, error) {
	vals := map[string]ConfigurationUsage{
		"allowed":       ConfigurationUsage_Allowed,
		"blocked":       ConfigurationUsage_Blocked,
		"notconfigured": ConfigurationUsage_NotConfigured,
		"required":      ConfigurationUsage_Required,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConfigurationUsage(input)
	return &out, nil
}
