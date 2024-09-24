package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDeviceImageStatusDetails string

const (
	CloudPCDeviceImageStatusDetails_InternalServerError                  CloudPCDeviceImageStatusDetails = "internalServerError"
	CloudPCDeviceImageStatusDetails_OsVersionNotSupported                CloudPCDeviceImageStatusDetails = "osVersionNotSupported"
	CloudPCDeviceImageStatusDetails_PaidSourceImageNotSupport            CloudPCDeviceImageStatusDetails = "paidSourceImageNotSupport"
	CloudPCDeviceImageStatusDetails_SourceImageInvalid                   CloudPCDeviceImageStatusDetails = "sourceImageInvalid"
	CloudPCDeviceImageStatusDetails_SourceImageNotFound                  CloudPCDeviceImageStatusDetails = "sourceImageNotFound"
	CloudPCDeviceImageStatusDetails_SourceImageNotGeneralized            CloudPCDeviceImageStatusDetails = "sourceImageNotGeneralized"
	CloudPCDeviceImageStatusDetails_SourceImageNotSupportCustomizeVMName CloudPCDeviceImageStatusDetails = "sourceImageNotSupportCustomizeVMName"
	CloudPCDeviceImageStatusDetails_SourceImageSizeExceedsLimitation     CloudPCDeviceImageStatusDetails = "sourceImageSizeExceedsLimitation"
	CloudPCDeviceImageStatusDetails_VmAlreadyAzureAdjoined               CloudPCDeviceImageStatusDetails = "vmAlreadyAzureAdjoined"
)

func PossibleValuesForCloudPCDeviceImageStatusDetails() []string {
	return []string{
		string(CloudPCDeviceImageStatusDetails_InternalServerError),
		string(CloudPCDeviceImageStatusDetails_OsVersionNotSupported),
		string(CloudPCDeviceImageStatusDetails_PaidSourceImageNotSupport),
		string(CloudPCDeviceImageStatusDetails_SourceImageInvalid),
		string(CloudPCDeviceImageStatusDetails_SourceImageNotFound),
		string(CloudPCDeviceImageStatusDetails_SourceImageNotGeneralized),
		string(CloudPCDeviceImageStatusDetails_SourceImageNotSupportCustomizeVMName),
		string(CloudPCDeviceImageStatusDetails_SourceImageSizeExceedsLimitation),
		string(CloudPCDeviceImageStatusDetails_VmAlreadyAzureAdjoined),
	}
}

func (s *CloudPCDeviceImageStatusDetails) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCDeviceImageStatusDetails(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCDeviceImageStatusDetails(input string) (*CloudPCDeviceImageStatusDetails, error) {
	vals := map[string]CloudPCDeviceImageStatusDetails{
		"internalservererror":                  CloudPCDeviceImageStatusDetails_InternalServerError,
		"osversionnotsupported":                CloudPCDeviceImageStatusDetails_OsVersionNotSupported,
		"paidsourceimagenotsupport":            CloudPCDeviceImageStatusDetails_PaidSourceImageNotSupport,
		"sourceimageinvalid":                   CloudPCDeviceImageStatusDetails_SourceImageInvalid,
		"sourceimagenotfound":                  CloudPCDeviceImageStatusDetails_SourceImageNotFound,
		"sourceimagenotgeneralized":            CloudPCDeviceImageStatusDetails_SourceImageNotGeneralized,
		"sourceimagenotsupportcustomizevmname": CloudPCDeviceImageStatusDetails_SourceImageNotSupportCustomizeVMName,
		"sourceimagesizeexceedslimitation":     CloudPCDeviceImageStatusDetails_SourceImageSizeExceedsLimitation,
		"vmalreadyazureadjoined":               CloudPCDeviceImageStatusDetails_VmAlreadyAzureAdjoined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDeviceImageStatusDetails(input)
	return &out, nil
}
