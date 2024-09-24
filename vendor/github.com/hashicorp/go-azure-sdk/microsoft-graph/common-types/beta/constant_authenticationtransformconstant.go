package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationTransformConstant string

const (
	AuthenticationTransformConstant_Aes128Gcm AuthenticationTransformConstant = "aes128Gcm"
	AuthenticationTransformConstant_Aes192Gcm AuthenticationTransformConstant = "aes192Gcm"
	AuthenticationTransformConstant_Aes256Gcm AuthenticationTransformConstant = "aes256Gcm"
	AuthenticationTransformConstant_Md596     AuthenticationTransformConstant = "md5_96"
	AuthenticationTransformConstant_Sha196    AuthenticationTransformConstant = "sha1_96"
	AuthenticationTransformConstant_Sha256128 AuthenticationTransformConstant = "sha_256_128"
)

func PossibleValuesForAuthenticationTransformConstant() []string {
	return []string{
		string(AuthenticationTransformConstant_Aes128Gcm),
		string(AuthenticationTransformConstant_Aes192Gcm),
		string(AuthenticationTransformConstant_Aes256Gcm),
		string(AuthenticationTransformConstant_Md596),
		string(AuthenticationTransformConstant_Sha196),
		string(AuthenticationTransformConstant_Sha256128),
	}
}

func (s *AuthenticationTransformConstant) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationTransformConstant(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationTransformConstant(input string) (*AuthenticationTransformConstant, error) {
	vals := map[string]AuthenticationTransformConstant{
		"aes128gcm":   AuthenticationTransformConstant_Aes128Gcm,
		"aes192gcm":   AuthenticationTransformConstant_Aes192Gcm,
		"aes256gcm":   AuthenticationTransformConstant_Aes256Gcm,
		"md5_96":      AuthenticationTransformConstant_Md596,
		"sha1_96":     AuthenticationTransformConstant_Sha196,
		"sha_256_128": AuthenticationTransformConstant_Sha256128,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationTransformConstant(input)
	return &out, nil
}
