package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerRecoveryPasswordRotationType string

const (
	BitLockerRecoveryPasswordRotationType_Disabled                   BitLockerRecoveryPasswordRotationType = "disabled"
	BitLockerRecoveryPasswordRotationType_EnabledForAzureAd          BitLockerRecoveryPasswordRotationType = "enabledForAzureAd"
	BitLockerRecoveryPasswordRotationType_EnabledForAzureAdAndHybrid BitLockerRecoveryPasswordRotationType = "enabledForAzureAdAndHybrid"
	BitLockerRecoveryPasswordRotationType_NotConfigured              BitLockerRecoveryPasswordRotationType = "notConfigured"
)

func PossibleValuesForBitLockerRecoveryPasswordRotationType() []string {
	return []string{
		string(BitLockerRecoveryPasswordRotationType_Disabled),
		string(BitLockerRecoveryPasswordRotationType_EnabledForAzureAd),
		string(BitLockerRecoveryPasswordRotationType_EnabledForAzureAdAndHybrid),
		string(BitLockerRecoveryPasswordRotationType_NotConfigured),
	}
}

func (s *BitLockerRecoveryPasswordRotationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBitLockerRecoveryPasswordRotationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBitLockerRecoveryPasswordRotationType(input string) (*BitLockerRecoveryPasswordRotationType, error) {
	vals := map[string]BitLockerRecoveryPasswordRotationType{
		"disabled":                   BitLockerRecoveryPasswordRotationType_Disabled,
		"enabledforazuread":          BitLockerRecoveryPasswordRotationType_EnabledForAzureAd,
		"enabledforazureadandhybrid": BitLockerRecoveryPasswordRotationType_EnabledForAzureAdAndHybrid,
		"notconfigured":              BitLockerRecoveryPasswordRotationType_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BitLockerRecoveryPasswordRotationType(input)
	return &out, nil
}
