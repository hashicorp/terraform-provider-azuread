package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentNotificationBrandingOptions string

const (
	EnrollmentNotificationBrandingOptions_IncludeCompanyLogo        EnrollmentNotificationBrandingOptions = "includeCompanyLogo"
	EnrollmentNotificationBrandingOptions_IncludeCompanyName        EnrollmentNotificationBrandingOptions = "includeCompanyName"
	EnrollmentNotificationBrandingOptions_IncludeCompanyPortalLink  EnrollmentNotificationBrandingOptions = "includeCompanyPortalLink"
	EnrollmentNotificationBrandingOptions_IncludeContactInformation EnrollmentNotificationBrandingOptions = "includeContactInformation"
	EnrollmentNotificationBrandingOptions_IncludeDeviceDetails      EnrollmentNotificationBrandingOptions = "includeDeviceDetails"
	EnrollmentNotificationBrandingOptions_None                      EnrollmentNotificationBrandingOptions = "none"
)

func PossibleValuesForEnrollmentNotificationBrandingOptions() []string {
	return []string{
		string(EnrollmentNotificationBrandingOptions_IncludeCompanyLogo),
		string(EnrollmentNotificationBrandingOptions_IncludeCompanyName),
		string(EnrollmentNotificationBrandingOptions_IncludeCompanyPortalLink),
		string(EnrollmentNotificationBrandingOptions_IncludeContactInformation),
		string(EnrollmentNotificationBrandingOptions_IncludeDeviceDetails),
		string(EnrollmentNotificationBrandingOptions_None),
	}
}

func (s *EnrollmentNotificationBrandingOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnrollmentNotificationBrandingOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnrollmentNotificationBrandingOptions(input string) (*EnrollmentNotificationBrandingOptions, error) {
	vals := map[string]EnrollmentNotificationBrandingOptions{
		"includecompanylogo":        EnrollmentNotificationBrandingOptions_IncludeCompanyLogo,
		"includecompanyname":        EnrollmentNotificationBrandingOptions_IncludeCompanyName,
		"includecompanyportallink":  EnrollmentNotificationBrandingOptions_IncludeCompanyPortalLink,
		"includecontactinformation": EnrollmentNotificationBrandingOptions_IncludeContactInformation,
		"includedevicedetails":      EnrollmentNotificationBrandingOptions_IncludeDeviceDetails,
		"none":                      EnrollmentNotificationBrandingOptions_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnrollmentNotificationBrandingOptions(input)
	return &out, nil
}
