package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceHealthOrigin string

const (
	ServiceHealthOrigin_Customer   ServiceHealthOrigin = "customer"
	ServiceHealthOrigin_Microsoft  ServiceHealthOrigin = "microsoft"
	ServiceHealthOrigin_ThirdParty ServiceHealthOrigin = "thirdParty"
)

func PossibleValuesForServiceHealthOrigin() []string {
	return []string{
		string(ServiceHealthOrigin_Customer),
		string(ServiceHealthOrigin_Microsoft),
		string(ServiceHealthOrigin_ThirdParty),
	}
}

func (s *ServiceHealthOrigin) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceHealthOrigin(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceHealthOrigin(input string) (*ServiceHealthOrigin, error) {
	vals := map[string]ServiceHealthOrigin{
		"customer":   ServiceHealthOrigin_Customer,
		"microsoft":  ServiceHealthOrigin_Microsoft,
		"thirdparty": ServiceHealthOrigin_ThirdParty,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceHealthOrigin(input)
	return &out, nil
}
