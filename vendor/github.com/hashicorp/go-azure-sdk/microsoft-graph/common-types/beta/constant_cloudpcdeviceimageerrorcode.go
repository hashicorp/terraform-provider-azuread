package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDeviceImageErrorCode string

const (
	CloudPCDeviceImageErrorCode_InternalServerError                          CloudPCDeviceImageErrorCode = "internalServerError"
	CloudPCDeviceImageErrorCode_OsVersionNotSupported                        CloudPCDeviceImageErrorCode = "osVersionNotSupported"
	CloudPCDeviceImageErrorCode_PaidSourceImageNotSupport                    CloudPCDeviceImageErrorCode = "paidSourceImageNotSupport"
	CloudPCDeviceImageErrorCode_SourceImageInvalid                           CloudPCDeviceImageErrorCode = "sourceImageInvalid"
	CloudPCDeviceImageErrorCode_SourceImageNotFound                          CloudPCDeviceImageErrorCode = "sourceImageNotFound"
	CloudPCDeviceImageErrorCode_SourceImageNotGeneralized                    CloudPCDeviceImageErrorCode = "sourceImageNotGeneralized"
	CloudPCDeviceImageErrorCode_SourceImageNotSupportCustomizeVMName         CloudPCDeviceImageErrorCode = "sourceImageNotSupportCustomizeVMName"
	CloudPCDeviceImageErrorCode_SourceImageSizeExceedsLimitation             CloudPCDeviceImageErrorCode = "sourceImageSizeExceedsLimitation"
	CloudPCDeviceImageErrorCode_SourceImageWithDataDiskNotSupported          CloudPCDeviceImageErrorCode = "sourceImageWithDataDiskNotSupported"
	CloudPCDeviceImageErrorCode_SourceImageWithDiskEncryptionSetNotSupported CloudPCDeviceImageErrorCode = "sourceImageWithDiskEncryptionSetNotSupported"
	CloudPCDeviceImageErrorCode_VmAlreadyAzureAdjoined                       CloudPCDeviceImageErrorCode = "vmAlreadyAzureAdjoined"
)

func PossibleValuesForCloudPCDeviceImageErrorCode() []string {
	return []string{
		string(CloudPCDeviceImageErrorCode_InternalServerError),
		string(CloudPCDeviceImageErrorCode_OsVersionNotSupported),
		string(CloudPCDeviceImageErrorCode_PaidSourceImageNotSupport),
		string(CloudPCDeviceImageErrorCode_SourceImageInvalid),
		string(CloudPCDeviceImageErrorCode_SourceImageNotFound),
		string(CloudPCDeviceImageErrorCode_SourceImageNotGeneralized),
		string(CloudPCDeviceImageErrorCode_SourceImageNotSupportCustomizeVMName),
		string(CloudPCDeviceImageErrorCode_SourceImageSizeExceedsLimitation),
		string(CloudPCDeviceImageErrorCode_SourceImageWithDataDiskNotSupported),
		string(CloudPCDeviceImageErrorCode_SourceImageWithDiskEncryptionSetNotSupported),
		string(CloudPCDeviceImageErrorCode_VmAlreadyAzureAdjoined),
	}
}

func (s *CloudPCDeviceImageErrorCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCDeviceImageErrorCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCDeviceImageErrorCode(input string) (*CloudPCDeviceImageErrorCode, error) {
	vals := map[string]CloudPCDeviceImageErrorCode{
		"internalservererror":                          CloudPCDeviceImageErrorCode_InternalServerError,
		"osversionnotsupported":                        CloudPCDeviceImageErrorCode_OsVersionNotSupported,
		"paidsourceimagenotsupport":                    CloudPCDeviceImageErrorCode_PaidSourceImageNotSupport,
		"sourceimageinvalid":                           CloudPCDeviceImageErrorCode_SourceImageInvalid,
		"sourceimagenotfound":                          CloudPCDeviceImageErrorCode_SourceImageNotFound,
		"sourceimagenotgeneralized":                    CloudPCDeviceImageErrorCode_SourceImageNotGeneralized,
		"sourceimagenotsupportcustomizevmname":         CloudPCDeviceImageErrorCode_SourceImageNotSupportCustomizeVMName,
		"sourceimagesizeexceedslimitation":             CloudPCDeviceImageErrorCode_SourceImageSizeExceedsLimitation,
		"sourceimagewithdatadisknotsupported":          CloudPCDeviceImageErrorCode_SourceImageWithDataDiskNotSupported,
		"sourceimagewithdiskencryptionsetnotsupported": CloudPCDeviceImageErrorCode_SourceImageWithDiskEncryptionSetNotSupported,
		"vmalreadyazureadjoined":                       CloudPCDeviceImageErrorCode_VmAlreadyAzureAdjoined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCDeviceImageErrorCode(input)
	return &out, nil
}
