package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceStartType string

const (
	ServiceStartType_Automatic ServiceStartType = "automatic"
	ServiceStartType_Disabled  ServiceStartType = "disabled"
	ServiceStartType_Manual    ServiceStartType = "manual"
)

func PossibleValuesForServiceStartType() []string {
	return []string{
		string(ServiceStartType_Automatic),
		string(ServiceStartType_Disabled),
		string(ServiceStartType_Manual),
	}
}

func (s *ServiceStartType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceStartType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceStartType(input string) (*ServiceStartType, error) {
	vals := map[string]ServiceStartType{
		"automatic": ServiceStartType_Automatic,
		"disabled":  ServiceStartType_Disabled,
		"manual":    ServiceStartType_Manual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceStartType(input)
	return &out, nil
}
