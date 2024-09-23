package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdministratorConfiguredDeviceComplianceState string

const (
	AdministratorConfiguredDeviceComplianceState_BasedOnDeviceCompliancePolicy AdministratorConfiguredDeviceComplianceState = "basedOnDeviceCompliancePolicy"
	AdministratorConfiguredDeviceComplianceState_NonCompliant                  AdministratorConfiguredDeviceComplianceState = "nonCompliant"
)

func PossibleValuesForAdministratorConfiguredDeviceComplianceState() []string {
	return []string{
		string(AdministratorConfiguredDeviceComplianceState_BasedOnDeviceCompliancePolicy),
		string(AdministratorConfiguredDeviceComplianceState_NonCompliant),
	}
}

func (s *AdministratorConfiguredDeviceComplianceState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAdministratorConfiguredDeviceComplianceState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAdministratorConfiguredDeviceComplianceState(input string) (*AdministratorConfiguredDeviceComplianceState, error) {
	vals := map[string]AdministratorConfiguredDeviceComplianceState{
		"basedondevicecompliancepolicy": AdministratorConfiguredDeviceComplianceState_BasedOnDeviceCompliancePolicy,
		"noncompliant":                  AdministratorConfiguredDeviceComplianceState_NonCompliant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdministratorConfiguredDeviceComplianceState(input)
	return &out, nil
}
