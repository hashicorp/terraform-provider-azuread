package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRecoveryInformationType string

const (
	BitLockerRecoveryInformationType_PasswordAndKey BitLockerRecoveryInformationType = "passwordAndKey"
	BitLockerRecoveryInformationType_PasswordOnly   BitLockerRecoveryInformationType = "passwordOnly"
)

func PossibleValuesForBitLockerRecoveryInformationType() []string {
	return []string{
		string(BitLockerRecoveryInformationType_PasswordAndKey),
		string(BitLockerRecoveryInformationType_PasswordOnly),
	}
}

func (s *BitLockerRecoveryInformationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBitLockerRecoveryInformationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBitLockerRecoveryInformationType(input string) (*BitLockerRecoveryInformationType, error) {
	vals := map[string]BitLockerRecoveryInformationType{
		"passwordandkey": BitLockerRecoveryInformationType_PasswordAndKey,
		"passwordonly":   BitLockerRecoveryInformationType_PasswordOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerRecoveryInformationType(input)
	return &out, nil
}
