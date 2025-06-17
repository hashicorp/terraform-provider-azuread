package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareOathTokenHashFunction string

const (
	HardwareOathTokenHashFunction_Hmacsha1   HardwareOathTokenHashFunction = "hmacsha1"
	HardwareOathTokenHashFunction_Hmacsha256 HardwareOathTokenHashFunction = "hmacsha256"
)

func PossibleValuesForHardwareOathTokenHashFunction() []string {
	return []string{
		string(HardwareOathTokenHashFunction_Hmacsha1),
		string(HardwareOathTokenHashFunction_Hmacsha256),
	}
}

func (s *HardwareOathTokenHashFunction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHardwareOathTokenHashFunction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHardwareOathTokenHashFunction(input string) (*HardwareOathTokenHashFunction, error) {
	vals := map[string]HardwareOathTokenHashFunction{
		"hmacsha1":   HardwareOathTokenHashFunction_Hmacsha1,
		"hmacsha256": HardwareOathTokenHashFunction_Hmacsha256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareOathTokenHashFunction(input)
	return &out, nil
}
