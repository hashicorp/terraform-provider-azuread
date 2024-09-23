package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationApprovalPolicyPlatform string

const (
	OperationApprovalPolicyPlatform_AndroidDeviceAdministrator OperationApprovalPolicyPlatform = "androidDeviceAdministrator"
	OperationApprovalPolicyPlatform_AndroidEnterprise          OperationApprovalPolicyPlatform = "androidEnterprise"
	OperationApprovalPolicyPlatform_IOSiPadOS                  OperationApprovalPolicyPlatform = "iOSiPadOS"
	OperationApprovalPolicyPlatform_MacOS                      OperationApprovalPolicyPlatform = "macOS"
	OperationApprovalPolicyPlatform_NotApplicable              OperationApprovalPolicyPlatform = "notApplicable"
	OperationApprovalPolicyPlatform_Windows10AndLater          OperationApprovalPolicyPlatform = "windows10AndLater"
	OperationApprovalPolicyPlatform_Windows10X                 OperationApprovalPolicyPlatform = "windows10X"
	OperationApprovalPolicyPlatform_Windows81AndLater          OperationApprovalPolicyPlatform = "windows81AndLater"
)

func PossibleValuesForOperationApprovalPolicyPlatform() []string {
	return []string{
		string(OperationApprovalPolicyPlatform_AndroidDeviceAdministrator),
		string(OperationApprovalPolicyPlatform_AndroidEnterprise),
		string(OperationApprovalPolicyPlatform_IOSiPadOS),
		string(OperationApprovalPolicyPlatform_MacOS),
		string(OperationApprovalPolicyPlatform_NotApplicable),
		string(OperationApprovalPolicyPlatform_Windows10AndLater),
		string(OperationApprovalPolicyPlatform_Windows10X),
		string(OperationApprovalPolicyPlatform_Windows81AndLater),
	}
}

func (s *OperationApprovalPolicyPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOperationApprovalPolicyPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOperationApprovalPolicyPlatform(input string) (*OperationApprovalPolicyPlatform, error) {
	vals := map[string]OperationApprovalPolicyPlatform{
		"androiddeviceadministrator": OperationApprovalPolicyPlatform_AndroidDeviceAdministrator,
		"androidenterprise":          OperationApprovalPolicyPlatform_AndroidEnterprise,
		"iosipados":                  OperationApprovalPolicyPlatform_IOSiPadOS,
		"macos":                      OperationApprovalPolicyPlatform_MacOS,
		"notapplicable":              OperationApprovalPolicyPlatform_NotApplicable,
		"windows10andlater":          OperationApprovalPolicyPlatform_Windows10AndLater,
		"windows10x":                 OperationApprovalPolicyPlatform_Windows10X,
		"windows81andlater":          OperationApprovalPolicyPlatform_Windows81AndLater,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationApprovalPolicyPlatform(input)
	return &out, nil
}
