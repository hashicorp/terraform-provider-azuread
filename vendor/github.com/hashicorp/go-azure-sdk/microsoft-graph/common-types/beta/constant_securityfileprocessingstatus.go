package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileProcessingStatus string

const (
	SecurityFileProcessingStatus_ExtractionException         SecurityFileProcessingStatus = "extractionException"
	SecurityFileProcessingStatus_FileBodyIsTooLong           SecurityFileProcessingStatus = "fileBodyIsTooLong"
	SecurityFileProcessingStatus_FileDepthLimitExceeded      SecurityFileProcessingStatus = "fileDepthLimitExceeded"
	SecurityFileProcessingStatus_FileSizeIsTooLarge          SecurityFileProcessingStatus = "fileSizeIsTooLarge"
	SecurityFileProcessingStatus_FileSizeIsZero              SecurityFileProcessingStatus = "fileSizeIsZero"
	SecurityFileProcessingStatus_FileTypeIsNotSupported      SecurityFileProcessingStatus = "fileTypeIsNotSupported"
	SecurityFileProcessingStatus_FileTypeIsUnknown           SecurityFileProcessingStatus = "fileTypeIsUnknown"
	SecurityFileProcessingStatus_InternalError               SecurityFileProcessingStatus = "internalError"
	SecurityFileProcessingStatus_InvalidFileId               SecurityFileProcessingStatus = "invalidFileId"
	SecurityFileProcessingStatus_MalformedFile               SecurityFileProcessingStatus = "malformedFile"
	SecurityFileProcessingStatus_NoReviewSetSummaryGenerated SecurityFileProcessingStatus = "noReviewSetSummaryGenerated"
	SecurityFileProcessingStatus_OcrFileSizeExceedsLimit     SecurityFileProcessingStatus = "ocrFileSizeExceedsLimit"
	SecurityFileProcessingStatus_OcrProcessingTimeout        SecurityFileProcessingStatus = "ocrProcessingTimeout"
	SecurityFileProcessingStatus_PoisonFile                  SecurityFileProcessingStatus = "poisonFile"
	SecurityFileProcessingStatus_ProcessingTimeout           SecurityFileProcessingStatus = "processingTimeout"
	SecurityFileProcessingStatus_ProtectedFile               SecurityFileProcessingStatus = "protectedFile"
	SecurityFileProcessingStatus_Success                     SecurityFileProcessingStatus = "success"
	SecurityFileProcessingStatus_UnknownError                SecurityFileProcessingStatus = "unknownError"
)

func PossibleValuesForSecurityFileProcessingStatus() []string {
	return []string{
		string(SecurityFileProcessingStatus_ExtractionException),
		string(SecurityFileProcessingStatus_FileBodyIsTooLong),
		string(SecurityFileProcessingStatus_FileDepthLimitExceeded),
		string(SecurityFileProcessingStatus_FileSizeIsTooLarge),
		string(SecurityFileProcessingStatus_FileSizeIsZero),
		string(SecurityFileProcessingStatus_FileTypeIsNotSupported),
		string(SecurityFileProcessingStatus_FileTypeIsUnknown),
		string(SecurityFileProcessingStatus_InternalError),
		string(SecurityFileProcessingStatus_InvalidFileId),
		string(SecurityFileProcessingStatus_MalformedFile),
		string(SecurityFileProcessingStatus_NoReviewSetSummaryGenerated),
		string(SecurityFileProcessingStatus_OcrFileSizeExceedsLimit),
		string(SecurityFileProcessingStatus_OcrProcessingTimeout),
		string(SecurityFileProcessingStatus_PoisonFile),
		string(SecurityFileProcessingStatus_ProcessingTimeout),
		string(SecurityFileProcessingStatus_ProtectedFile),
		string(SecurityFileProcessingStatus_Success),
		string(SecurityFileProcessingStatus_UnknownError),
	}
}

func (s *SecurityFileProcessingStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileProcessingStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileProcessingStatus(input string) (*SecurityFileProcessingStatus, error) {
	vals := map[string]SecurityFileProcessingStatus{
		"extractionexception":         SecurityFileProcessingStatus_ExtractionException,
		"filebodyistoolong":           SecurityFileProcessingStatus_FileBodyIsTooLong,
		"filedepthlimitexceeded":      SecurityFileProcessingStatus_FileDepthLimitExceeded,
		"filesizeistoolarge":          SecurityFileProcessingStatus_FileSizeIsTooLarge,
		"filesizeiszero":              SecurityFileProcessingStatus_FileSizeIsZero,
		"filetypeisnotsupported":      SecurityFileProcessingStatus_FileTypeIsNotSupported,
		"filetypeisunknown":           SecurityFileProcessingStatus_FileTypeIsUnknown,
		"internalerror":               SecurityFileProcessingStatus_InternalError,
		"invalidfileid":               SecurityFileProcessingStatus_InvalidFileId,
		"malformedfile":               SecurityFileProcessingStatus_MalformedFile,
		"noreviewsetsummarygenerated": SecurityFileProcessingStatus_NoReviewSetSummaryGenerated,
		"ocrfilesizeexceedslimit":     SecurityFileProcessingStatus_OcrFileSizeExceedsLimit,
		"ocrprocessingtimeout":        SecurityFileProcessingStatus_OcrProcessingTimeout,
		"poisonfile":                  SecurityFileProcessingStatus_PoisonFile,
		"processingtimeout":           SecurityFileProcessingStatus_ProcessingTimeout,
		"protectedfile":               SecurityFileProcessingStatus_ProtectedFile,
		"success":                     SecurityFileProcessingStatus_Success,
		"unknownerror":                SecurityFileProcessingStatus_UnknownError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileProcessingStatus(input)
	return &out, nil
}
