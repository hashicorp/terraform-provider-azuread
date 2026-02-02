package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealthClassificationType string

const (
	ServiceHealthClassificationType_Advisory ServiceHealthClassificationType = "advisory"
	ServiceHealthClassificationType_Incident ServiceHealthClassificationType = "incident"
)

func PossibleValuesForServiceHealthClassificationType() []string {
	return []string{
		string(ServiceHealthClassificationType_Advisory),
		string(ServiceHealthClassificationType_Incident),
	}
}

func (s *ServiceHealthClassificationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceHealthClassificationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceHealthClassificationType(input string) (*ServiceHealthClassificationType, error) {
	vals := map[string]ServiceHealthClassificationType{
		"advisory": ServiceHealthClassificationType_Advisory,
		"incident": ServiceHealthClassificationType_Incident,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceHealthClassificationType(input)
	return &out, nil
}
