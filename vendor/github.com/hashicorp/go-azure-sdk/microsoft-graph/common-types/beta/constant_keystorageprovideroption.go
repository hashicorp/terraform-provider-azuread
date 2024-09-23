package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyStorageProviderOption string

const (
	KeyStorageProviderOption_UsePassportForWorkKspOtherwiseFail KeyStorageProviderOption = "usePassportForWorkKspOtherwiseFail"
	KeyStorageProviderOption_UseSoftwareKsp                     KeyStorageProviderOption = "useSoftwareKsp"
	KeyStorageProviderOption_UseTpmKspOtherwiseFail             KeyStorageProviderOption = "useTpmKspOtherwiseFail"
	KeyStorageProviderOption_UseTpmKspOtherwiseUseSoftwareKsp   KeyStorageProviderOption = "useTpmKspOtherwiseUseSoftwareKsp"
)

func PossibleValuesForKeyStorageProviderOption() []string {
	return []string{
		string(KeyStorageProviderOption_UsePassportForWorkKspOtherwiseFail),
		string(KeyStorageProviderOption_UseSoftwareKsp),
		string(KeyStorageProviderOption_UseTpmKspOtherwiseFail),
		string(KeyStorageProviderOption_UseTpmKspOtherwiseUseSoftwareKsp),
	}
}

func (s *KeyStorageProviderOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseKeyStorageProviderOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseKeyStorageProviderOption(input string) (*KeyStorageProviderOption, error) {
	vals := map[string]KeyStorageProviderOption{
		"usepassportforworkkspotherwisefail": KeyStorageProviderOption_UsePassportForWorkKspOtherwiseFail,
		"usesoftwareksp":                     KeyStorageProviderOption_UseSoftwareKsp,
		"usetpmkspotherwisefail":             KeyStorageProviderOption_UseTpmKspOtherwiseFail,
		"usetpmkspotherwiseusesoftwareksp":   KeyStorageProviderOption_UseTpmKspOtherwiseUseSoftwareKsp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := KeyStorageProviderOption(input)
	return &out, nil
}
