package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEapType string

const (
	AndroidEapType_EapTls  AndroidEapType = "eapTls"
	AndroidEapType_EapTtls AndroidEapType = "eapTtls"
	AndroidEapType_Peap    AndroidEapType = "peap"
)

func PossibleValuesForAndroidEapType() []string {
	return []string{
		string(AndroidEapType_EapTls),
		string(AndroidEapType_EapTtls),
		string(AndroidEapType_Peap),
	}
}

func (s *AndroidEapType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidEapType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidEapType(input string) (*AndroidEapType, error) {
	vals := map[string]AndroidEapType{
		"eaptls":  AndroidEapType_EapTls,
		"eapttls": AndroidEapType_EapTtls,
		"peap":    AndroidEapType_Peap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEapType(input)
	return &out, nil
}
