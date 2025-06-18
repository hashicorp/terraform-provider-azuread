package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAppInfoDataAtRestEncryptionMethod string

const (
	SecurityAppInfoDataAtRestEncryptionMethod_Aes          SecurityAppInfoDataAtRestEncryptionMethod = "aes"
	SecurityAppInfoDataAtRestEncryptionMethod_BitLocker    SecurityAppInfoDataAtRestEncryptionMethod = "bitLocker"
	SecurityAppInfoDataAtRestEncryptionMethod_Blowfish     SecurityAppInfoDataAtRestEncryptionMethod = "blowfish"
	SecurityAppInfoDataAtRestEncryptionMethod_Des          SecurityAppInfoDataAtRestEncryptionMethod = "des"
	SecurityAppInfoDataAtRestEncryptionMethod_Des3         SecurityAppInfoDataAtRestEncryptionMethod = "des3"
	SecurityAppInfoDataAtRestEncryptionMethod_NotSupported SecurityAppInfoDataAtRestEncryptionMethod = "notSupported"
	SecurityAppInfoDataAtRestEncryptionMethod_Rc4          SecurityAppInfoDataAtRestEncryptionMethod = "rc4"
	SecurityAppInfoDataAtRestEncryptionMethod_RsA          SecurityAppInfoDataAtRestEncryptionMethod = "rsA"
	SecurityAppInfoDataAtRestEncryptionMethod_Unknown      SecurityAppInfoDataAtRestEncryptionMethod = "unknown"
)

func PossibleValuesForSecurityAppInfoDataAtRestEncryptionMethod() []string {
	return []string{
		string(SecurityAppInfoDataAtRestEncryptionMethod_Aes),
		string(SecurityAppInfoDataAtRestEncryptionMethod_BitLocker),
		string(SecurityAppInfoDataAtRestEncryptionMethod_Blowfish),
		string(SecurityAppInfoDataAtRestEncryptionMethod_Des),
		string(SecurityAppInfoDataAtRestEncryptionMethod_Des3),
		string(SecurityAppInfoDataAtRestEncryptionMethod_NotSupported),
		string(SecurityAppInfoDataAtRestEncryptionMethod_Rc4),
		string(SecurityAppInfoDataAtRestEncryptionMethod_RsA),
		string(SecurityAppInfoDataAtRestEncryptionMethod_Unknown),
	}
}

func (s *SecurityAppInfoDataAtRestEncryptionMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAppInfoDataAtRestEncryptionMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAppInfoDataAtRestEncryptionMethod(input string) (*SecurityAppInfoDataAtRestEncryptionMethod, error) {
	vals := map[string]SecurityAppInfoDataAtRestEncryptionMethod{
		"aes":          SecurityAppInfoDataAtRestEncryptionMethod_Aes,
		"bitlocker":    SecurityAppInfoDataAtRestEncryptionMethod_BitLocker,
		"blowfish":     SecurityAppInfoDataAtRestEncryptionMethod_Blowfish,
		"des":          SecurityAppInfoDataAtRestEncryptionMethod_Des,
		"des3":         SecurityAppInfoDataAtRestEncryptionMethod_Des3,
		"notsupported": SecurityAppInfoDataAtRestEncryptionMethod_NotSupported,
		"rc4":          SecurityAppInfoDataAtRestEncryptionMethod_Rc4,
		"rsa":          SecurityAppInfoDataAtRestEncryptionMethod_RsA,
		"unknown":      SecurityAppInfoDataAtRestEncryptionMethod_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAppInfoDataAtRestEncryptionMethod(input)
	return &out, nil
}
