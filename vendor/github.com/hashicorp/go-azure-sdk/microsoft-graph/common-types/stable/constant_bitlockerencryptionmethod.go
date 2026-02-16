package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerEncryptionMethod string

const (
	BitLockerEncryptionMethod_AesCbc128 BitLockerEncryptionMethod = "aesCbc128"
	BitLockerEncryptionMethod_AesCbc256 BitLockerEncryptionMethod = "aesCbc256"
	BitLockerEncryptionMethod_XtsAes128 BitLockerEncryptionMethod = "xtsAes128"
	BitLockerEncryptionMethod_XtsAes256 BitLockerEncryptionMethod = "xtsAes256"
)

func PossibleValuesForBitLockerEncryptionMethod() []string {
	return []string{
		string(BitLockerEncryptionMethod_AesCbc128),
		string(BitLockerEncryptionMethod_AesCbc256),
		string(BitLockerEncryptionMethod_XtsAes128),
		string(BitLockerEncryptionMethod_XtsAes256),
	}
}

func (s *BitLockerEncryptionMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBitLockerEncryptionMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBitLockerEncryptionMethod(input string) (*BitLockerEncryptionMethod, error) {
	vals := map[string]BitLockerEncryptionMethod{
		"aescbc128": BitLockerEncryptionMethod_AesCbc128,
		"aescbc256": BitLockerEncryptionMethod_AesCbc256,
		"xtsaes128": BitLockerEncryptionMethod_XtsAes128,
		"xtsaes256": BitLockerEncryptionMethod_XtsAes256,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerEncryptionMethod(input)
	return &out, nil
}
