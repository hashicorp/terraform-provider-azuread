package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCResizeValidationCode string

const (
	CloudPCResizeValidationCode_CloudPCNotFound          CloudPCResizeValidationCode = "cloudPcNotFound"
	CloudPCResizeValidationCode_InternalServerError      CloudPCResizeValidationCode = "internalServerError"
	CloudPCResizeValidationCode_OperationConflict        CloudPCResizeValidationCode = "operationConflict"
	CloudPCResizeValidationCode_OperationNotSupported    CloudPCResizeValidationCode = "operationNotSupported"
	CloudPCResizeValidationCode_Success                  CloudPCResizeValidationCode = "success"
	CloudPCResizeValidationCode_TargetLicenseHasAssigned CloudPCResizeValidationCode = "targetLicenseHasAssigned"
)

func PossibleValuesForCloudPCResizeValidationCode() []string {
	return []string{
		string(CloudPCResizeValidationCode_CloudPCNotFound),
		string(CloudPCResizeValidationCode_InternalServerError),
		string(CloudPCResizeValidationCode_OperationConflict),
		string(CloudPCResizeValidationCode_OperationNotSupported),
		string(CloudPCResizeValidationCode_Success),
		string(CloudPCResizeValidationCode_TargetLicenseHasAssigned),
	}
}

func (s *CloudPCResizeValidationCode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCResizeValidationCode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCResizeValidationCode(input string) (*CloudPCResizeValidationCode, error) {
	vals := map[string]CloudPCResizeValidationCode{
		"cloudpcnotfound":          CloudPCResizeValidationCode_CloudPCNotFound,
		"internalservererror":      CloudPCResizeValidationCode_InternalServerError,
		"operationconflict":        CloudPCResizeValidationCode_OperationConflict,
		"operationnotsupported":    CloudPCResizeValidationCode_OperationNotSupported,
		"success":                  CloudPCResizeValidationCode_Success,
		"targetlicensehasassigned": CloudPCResizeValidationCode_TargetLicenseHasAssigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCResizeValidationCode(input)
	return &out, nil
}
