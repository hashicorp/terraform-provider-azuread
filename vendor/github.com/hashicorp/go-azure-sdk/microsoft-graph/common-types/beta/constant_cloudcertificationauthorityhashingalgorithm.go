package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudCertificationAuthorityHashingAlgorithm string

const (
	CloudCertificationAuthorityHashingAlgorithm_Sha256  CloudCertificationAuthorityHashingAlgorithm = "sha256"
	CloudCertificationAuthorityHashingAlgorithm_Sha384  CloudCertificationAuthorityHashingAlgorithm = "sha384"
	CloudCertificationAuthorityHashingAlgorithm_Sha512  CloudCertificationAuthorityHashingAlgorithm = "sha512"
	CloudCertificationAuthorityHashingAlgorithm_Unknown CloudCertificationAuthorityHashingAlgorithm = "unknown"
)

func PossibleValuesForCloudCertificationAuthorityHashingAlgorithm() []string {
	return []string{
		string(CloudCertificationAuthorityHashingAlgorithm_Sha256),
		string(CloudCertificationAuthorityHashingAlgorithm_Sha384),
		string(CloudCertificationAuthorityHashingAlgorithm_Sha512),
		string(CloudCertificationAuthorityHashingAlgorithm_Unknown),
	}
}

func (s *CloudCertificationAuthorityHashingAlgorithm) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudCertificationAuthorityHashingAlgorithm(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudCertificationAuthorityHashingAlgorithm(input string) (*CloudCertificationAuthorityHashingAlgorithm, error) {
	vals := map[string]CloudCertificationAuthorityHashingAlgorithm{
		"sha256":  CloudCertificationAuthorityHashingAlgorithm_Sha256,
		"sha384":  CloudCertificationAuthorityHashingAlgorithm_Sha384,
		"sha512":  CloudCertificationAuthorityHashingAlgorithm_Sha512,
		"unknown": CloudCertificationAuthorityHashingAlgorithm_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudCertificationAuthorityHashingAlgorithm(input)
	return &out, nil
}
