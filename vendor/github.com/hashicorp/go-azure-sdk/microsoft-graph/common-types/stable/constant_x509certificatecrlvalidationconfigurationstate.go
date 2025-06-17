package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type X509CertificateCRLValidationConfigurationState string

const (
	X509CertificateCRLValidationConfigurationState_Disabled X509CertificateCRLValidationConfigurationState = "disabled"
	X509CertificateCRLValidationConfigurationState_Enabled  X509CertificateCRLValidationConfigurationState = "enabled"
)

func PossibleValuesForX509CertificateCRLValidationConfigurationState() []string {
	return []string{
		string(X509CertificateCRLValidationConfigurationState_Disabled),
		string(X509CertificateCRLValidationConfigurationState_Enabled),
	}
}

func (s *X509CertificateCRLValidationConfigurationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseX509CertificateCRLValidationConfigurationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseX509CertificateCRLValidationConfigurationState(input string) (*X509CertificateCRLValidationConfigurationState, error) {
	vals := map[string]X509CertificateCRLValidationConfigurationState{
		"disabled": X509CertificateCRLValidationConfigurationState_Disabled,
		"enabled":  X509CertificateCRLValidationConfigurationState_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := X509CertificateCRLValidationConfigurationState(input)
	return &out, nil
}
