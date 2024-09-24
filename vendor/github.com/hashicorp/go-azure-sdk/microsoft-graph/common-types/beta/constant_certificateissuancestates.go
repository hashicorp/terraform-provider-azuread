package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateIssuanceStates string

const (
	CertificateIssuanceStates_ChallengeIssueFailed         CertificateIssuanceStates = "challengeIssueFailed"
	CertificateIssuanceStates_ChallengeIssued              CertificateIssuanceStates = "challengeIssued"
	CertificateIssuanceStates_ChallengeValidationFailed    CertificateIssuanceStates = "challengeValidationFailed"
	CertificateIssuanceStates_ChallengeValidationSucceeded CertificateIssuanceStates = "challengeValidationSucceeded"
	CertificateIssuanceStates_DeleteFailed                 CertificateIssuanceStates = "deleteFailed"
	CertificateIssuanceStates_Deleted                      CertificateIssuanceStates = "deleted"
	CertificateIssuanceStates_EnrollmentNotNeeded          CertificateIssuanceStates = "enrollmentNotNeeded"
	CertificateIssuanceStates_EnrollmentSucceeded          CertificateIssuanceStates = "enrollmentSucceeded"
	CertificateIssuanceStates_InstallFailed                CertificateIssuanceStates = "installFailed"
	CertificateIssuanceStates_Installed                    CertificateIssuanceStates = "installed"
	CertificateIssuanceStates_IssueFailed                  CertificateIssuanceStates = "issueFailed"
	CertificateIssuanceStates_IssuePending                 CertificateIssuanceStates = "issuePending"
	CertificateIssuanceStates_Issued                       CertificateIssuanceStates = "issued"
	CertificateIssuanceStates_RemovedFromCollection        CertificateIssuanceStates = "removedFromCollection"
	CertificateIssuanceStates_RenewVerified                CertificateIssuanceStates = "renewVerified"
	CertificateIssuanceStates_RenewalRequested             CertificateIssuanceStates = "renewalRequested"
	CertificateIssuanceStates_RequestCreationFailed        CertificateIssuanceStates = "requestCreationFailed"
	CertificateIssuanceStates_RequestSubmitFailed          CertificateIssuanceStates = "requestSubmitFailed"
	CertificateIssuanceStates_Requested                    CertificateIssuanceStates = "requested"
	CertificateIssuanceStates_ResponsePending              CertificateIssuanceStates = "responsePending"
	CertificateIssuanceStates_ResponseProcessingFailed     CertificateIssuanceStates = "responseProcessingFailed"
	CertificateIssuanceStates_Revoked                      CertificateIssuanceStates = "revoked"
	CertificateIssuanceStates_Unknown                      CertificateIssuanceStates = "unknown"
)

func PossibleValuesForCertificateIssuanceStates() []string {
	return []string{
		string(CertificateIssuanceStates_ChallengeIssueFailed),
		string(CertificateIssuanceStates_ChallengeIssued),
		string(CertificateIssuanceStates_ChallengeValidationFailed),
		string(CertificateIssuanceStates_ChallengeValidationSucceeded),
		string(CertificateIssuanceStates_DeleteFailed),
		string(CertificateIssuanceStates_Deleted),
		string(CertificateIssuanceStates_EnrollmentNotNeeded),
		string(CertificateIssuanceStates_EnrollmentSucceeded),
		string(CertificateIssuanceStates_InstallFailed),
		string(CertificateIssuanceStates_Installed),
		string(CertificateIssuanceStates_IssueFailed),
		string(CertificateIssuanceStates_IssuePending),
		string(CertificateIssuanceStates_Issued),
		string(CertificateIssuanceStates_RemovedFromCollection),
		string(CertificateIssuanceStates_RenewVerified),
		string(CertificateIssuanceStates_RenewalRequested),
		string(CertificateIssuanceStates_RequestCreationFailed),
		string(CertificateIssuanceStates_RequestSubmitFailed),
		string(CertificateIssuanceStates_Requested),
		string(CertificateIssuanceStates_ResponsePending),
		string(CertificateIssuanceStates_ResponseProcessingFailed),
		string(CertificateIssuanceStates_Revoked),
		string(CertificateIssuanceStates_Unknown),
	}
}

func (s *CertificateIssuanceStates) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCertificateIssuanceStates(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCertificateIssuanceStates(input string) (*CertificateIssuanceStates, error) {
	vals := map[string]CertificateIssuanceStates{
		"challengeissuefailed":         CertificateIssuanceStates_ChallengeIssueFailed,
		"challengeissued":              CertificateIssuanceStates_ChallengeIssued,
		"challengevalidationfailed":    CertificateIssuanceStates_ChallengeValidationFailed,
		"challengevalidationsucceeded": CertificateIssuanceStates_ChallengeValidationSucceeded,
		"deletefailed":                 CertificateIssuanceStates_DeleteFailed,
		"deleted":                      CertificateIssuanceStates_Deleted,
		"enrollmentnotneeded":          CertificateIssuanceStates_EnrollmentNotNeeded,
		"enrollmentsucceeded":          CertificateIssuanceStates_EnrollmentSucceeded,
		"installfailed":                CertificateIssuanceStates_InstallFailed,
		"installed":                    CertificateIssuanceStates_Installed,
		"issuefailed":                  CertificateIssuanceStates_IssueFailed,
		"issuepending":                 CertificateIssuanceStates_IssuePending,
		"issued":                       CertificateIssuanceStates_Issued,
		"removedfromcollection":        CertificateIssuanceStates_RemovedFromCollection,
		"renewverified":                CertificateIssuanceStates_RenewVerified,
		"renewalrequested":             CertificateIssuanceStates_RenewalRequested,
		"requestcreationfailed":        CertificateIssuanceStates_RequestCreationFailed,
		"requestsubmitfailed":          CertificateIssuanceStates_RequestSubmitFailed,
		"requested":                    CertificateIssuanceStates_Requested,
		"responsepending":              CertificateIssuanceStates_ResponsePending,
		"responseprocessingfailed":     CertificateIssuanceStates_ResponseProcessingFailed,
		"revoked":                      CertificateIssuanceStates_Revoked,
		"unknown":                      CertificateIssuanceStates_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CertificateIssuanceStates(input)
	return &out, nil
}
