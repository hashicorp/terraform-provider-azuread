package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentNotificationTemplateType string

const (
	EnrollmentNotificationTemplateType_Email EnrollmentNotificationTemplateType = "email"
	EnrollmentNotificationTemplateType_Push  EnrollmentNotificationTemplateType = "push"
)

func PossibleValuesForEnrollmentNotificationTemplateType() []string {
	return []string{
		string(EnrollmentNotificationTemplateType_Email),
		string(EnrollmentNotificationTemplateType_Push),
	}
}

func (s *EnrollmentNotificationTemplateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentNotificationTemplateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentNotificationTemplateType(input string) (*EnrollmentNotificationTemplateType, error) {
	vals := map[string]EnrollmentNotificationTemplateType{
		"email": EnrollmentNotificationTemplateType_Email,
		"push":  EnrollmentNotificationTemplateType_Push,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentNotificationTemplateType(input)
	return &out, nil
}
