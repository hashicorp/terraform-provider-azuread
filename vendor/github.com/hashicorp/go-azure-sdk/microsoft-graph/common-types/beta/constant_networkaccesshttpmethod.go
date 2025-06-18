package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessHttpMethod string

const (
	NetworkaccessHttpMethod_Connect NetworkaccessHttpMethod = "connect"
	NetworkaccessHttpMethod_Delete  NetworkaccessHttpMethod = "delete"
	NetworkaccessHttpMethod_Get     NetworkaccessHttpMethod = "get"
	NetworkaccessHttpMethod_Head    NetworkaccessHttpMethod = "head"
	NetworkaccessHttpMethod_Options NetworkaccessHttpMethod = "options"
	NetworkaccessHttpMethod_Patch   NetworkaccessHttpMethod = "patch"
	NetworkaccessHttpMethod_Post    NetworkaccessHttpMethod = "post"
	NetworkaccessHttpMethod_Put     NetworkaccessHttpMethod = "put"
	NetworkaccessHttpMethod_Trace   NetworkaccessHttpMethod = "trace"
)

func PossibleValuesForNetworkaccessHttpMethod() []string {
	return []string{
		string(NetworkaccessHttpMethod_Connect),
		string(NetworkaccessHttpMethod_Delete),
		string(NetworkaccessHttpMethod_Get),
		string(NetworkaccessHttpMethod_Head),
		string(NetworkaccessHttpMethod_Options),
		string(NetworkaccessHttpMethod_Patch),
		string(NetworkaccessHttpMethod_Post),
		string(NetworkaccessHttpMethod_Put),
		string(NetworkaccessHttpMethod_Trace),
	}
}

func (s *NetworkaccessHttpMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessHttpMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessHttpMethod(input string) (*NetworkaccessHttpMethod, error) {
	vals := map[string]NetworkaccessHttpMethod{
		"connect": NetworkaccessHttpMethod_Connect,
		"delete":  NetworkaccessHttpMethod_Delete,
		"get":     NetworkaccessHttpMethod_Get,
		"head":    NetworkaccessHttpMethod_Head,
		"options": NetworkaccessHttpMethod_Options,
		"patch":   NetworkaccessHttpMethod_Patch,
		"post":    NetworkaccessHttpMethod_Post,
		"put":     NetworkaccessHttpMethod_Put,
		"trace":   NetworkaccessHttpMethod_Trace,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessHttpMethod(input)
	return &out, nil
}
