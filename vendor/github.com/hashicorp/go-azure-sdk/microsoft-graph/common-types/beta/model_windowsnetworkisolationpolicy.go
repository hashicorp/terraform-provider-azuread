package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsNetworkIsolationPolicy struct {
	// Contains a list of enterprise resource domains hosted in the cloud that need to be protected. Connections to these
	// resources are considered enterprise data. If a proxy is paired with a cloud resource, traffic to the cloud resource
	// will be routed through the enterprise network via the denoted proxy server (on Port 80). A proxy server used for this
	// purpose must also be configured using the EnterpriseInternalProxyServers policy. This collection can contain a
	// maximum of 500 elements.
	EnterpriseCloudResources *[]ProxiedDomain `json:"enterpriseCloudResources,omitempty"`

	// Sets the enterprise IP ranges that define the computers in the enterprise network. Data that comes from those
	// computers will be considered part of the enterprise and protected. These locations will be considered a safe
	// destination for enterprise data to be shared to. This collection can contain a maximum of 500 elements.
	EnterpriseIPRanges *[]IPRange `json:"enterpriseIPRanges,omitempty"`

	// Boolean value that tells the client to accept the configured list and not to use heuristics to attempt to find other
	// subnets. Default is false.
	EnterpriseIPRangesAreAuthoritative *bool `json:"enterpriseIPRangesAreAuthoritative,omitempty"`

	// This is the comma-separated list of internal proxy servers. For example, '157.54.14.28, 157.54.11.118, 10.202.14.167,
	// 157.53.14.163, 157.69.210.59'. These proxies have been configured by the admin to connect to specific resources on
	// the Internet. They are considered to be enterprise network locations. The proxies are only leveraged in configuring
	// the EnterpriseCloudResources policy to force traffic to the matched cloud resources through these proxies.
	EnterpriseInternalProxyServers *[]string `json:"enterpriseInternalProxyServers,omitempty"`

	// This is the list of domains that comprise the boundaries of the enterprise. Data from one of these domains that is
	// sent to a device will be considered enterprise data and protected. These locations will be considered a safe
	// destination for enterprise data to be shared to.
	EnterpriseNetworkDomainNames *[]string `json:"enterpriseNetworkDomainNames,omitempty"`

	// This is a list of proxy servers. Any server not on this list is considered non-enterprise.
	EnterpriseProxyServers *[]string `json:"enterpriseProxyServers,omitempty"`

	// Boolean value that tells the client to accept the configured list of proxies and not try to detect other work
	// proxies. Default is false
	EnterpriseProxyServersAreAuthoritative *bool `json:"enterpriseProxyServersAreAuthoritative,omitempty"`

	// List of domain names that can used for work or personal resource.
	NeutralDomainResources *[]string `json:"neutralDomainResources,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &WindowsNetworkIsolationPolicy{}

func (s *WindowsNetworkIsolationPolicy) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		EnterpriseCloudResources               *[]ProxiedDomain `json:"enterpriseCloudResources,omitempty"`
		EnterpriseIPRangesAreAuthoritative     *bool            `json:"enterpriseIPRangesAreAuthoritative,omitempty"`
		EnterpriseInternalProxyServers         *[]string        `json:"enterpriseInternalProxyServers,omitempty"`
		EnterpriseNetworkDomainNames           *[]string        `json:"enterpriseNetworkDomainNames,omitempty"`
		EnterpriseProxyServers                 *[]string        `json:"enterpriseProxyServers,omitempty"`
		EnterpriseProxyServersAreAuthoritative *bool            `json:"enterpriseProxyServersAreAuthoritative,omitempty"`
		NeutralDomainResources                 *[]string        `json:"neutralDomainResources,omitempty"`
		ODataId                                *string          `json:"@odata.id,omitempty"`
		ODataType                              *string          `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.EnterpriseCloudResources = decoded.EnterpriseCloudResources
	s.EnterpriseIPRangesAreAuthoritative = decoded.EnterpriseIPRangesAreAuthoritative
	s.EnterpriseInternalProxyServers = decoded.EnterpriseInternalProxyServers
	s.EnterpriseNetworkDomainNames = decoded.EnterpriseNetworkDomainNames
	s.EnterpriseProxyServers = decoded.EnterpriseProxyServers
	s.EnterpriseProxyServersAreAuthoritative = decoded.EnterpriseProxyServersAreAuthoritative
	s.NeutralDomainResources = decoded.NeutralDomainResources
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WindowsNetworkIsolationPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["enterpriseIPRanges"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling EnterpriseIPRanges into list []json.RawMessage: %+v", err)
		}

		output := make([]IPRange, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalIPRangeImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'EnterpriseIPRanges' for 'WindowsNetworkIsolationPolicy': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.EnterpriseIPRanges = &output
	}

	return nil
}
