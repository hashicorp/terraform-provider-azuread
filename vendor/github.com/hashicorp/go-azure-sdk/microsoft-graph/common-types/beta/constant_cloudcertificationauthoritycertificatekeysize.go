package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudCertificationAuthorityCertificateKeySize string

const (
	CloudCertificationAuthorityCertificateKeySize_ECP256  CloudCertificationAuthorityCertificateKeySize = "eCP256"
	CloudCertificationAuthorityCertificateKeySize_ECP256k CloudCertificationAuthorityCertificateKeySize = "eCP256k"
	CloudCertificationAuthorityCertificateKeySize_ECP384  CloudCertificationAuthorityCertificateKeySize = "eCP384"
	CloudCertificationAuthorityCertificateKeySize_ECP521  CloudCertificationAuthorityCertificateKeySize = "eCP521"
	CloudCertificationAuthorityCertificateKeySize_Rsa2048 CloudCertificationAuthorityCertificateKeySize = "rsa2048"
	CloudCertificationAuthorityCertificateKeySize_Rsa3072 CloudCertificationAuthorityCertificateKeySize = "rsa3072"
	CloudCertificationAuthorityCertificateKeySize_Rsa4096 CloudCertificationAuthorityCertificateKeySize = "rsa4096"
	CloudCertificationAuthorityCertificateKeySize_Unknown CloudCertificationAuthorityCertificateKeySize = "unknown"
)

func PossibleValuesForCloudCertificationAuthorityCertificateKeySize() []string {
	return []string{
		string(CloudCertificationAuthorityCertificateKeySize_ECP256),
		string(CloudCertificationAuthorityCertificateKeySize_ECP256k),
		string(CloudCertificationAuthorityCertificateKeySize_ECP384),
		string(CloudCertificationAuthorityCertificateKeySize_ECP521),
		string(CloudCertificationAuthorityCertificateKeySize_Rsa2048),
		string(CloudCertificationAuthorityCertificateKeySize_Rsa3072),
		string(CloudCertificationAuthorityCertificateKeySize_Rsa4096),
		string(CloudCertificationAuthorityCertificateKeySize_Unknown),
	}
}

func (s *CloudCertificationAuthorityCertificateKeySize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudCertificationAuthorityCertificateKeySize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudCertificationAuthorityCertificateKeySize(input string) (*CloudCertificationAuthorityCertificateKeySize, error) {
	vals := map[string]CloudCertificationAuthorityCertificateKeySize{
		"ecp256":  CloudCertificationAuthorityCertificateKeySize_ECP256,
		"ecp256k": CloudCertificationAuthorityCertificateKeySize_ECP256k,
		"ecp384":  CloudCertificationAuthorityCertificateKeySize_ECP384,
		"ecp521":  CloudCertificationAuthorityCertificateKeySize_ECP521,
		"rsa2048": CloudCertificationAuthorityCertificateKeySize_Rsa2048,
		"rsa3072": CloudCertificationAuthorityCertificateKeySize_Rsa3072,
		"rsa4096": CloudCertificationAuthorityCertificateKeySize_Rsa4096,
		"unknown": CloudCertificationAuthorityCertificateKeySize_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudCertificationAuthorityCertificateKeySize(input)
	return &out, nil
}
