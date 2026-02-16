package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SecurityThreatIntelligence{}

type SecurityThreatIntelligence struct {
	// Refers to indicators of threat or compromise highlighted in an article.Note: List retrieval is not yet supported.
	ArticleIndicators *[]SecurityArticleIndicator `json:"articleIndicators,omitempty"`

	// A list of article objects.
	Articles *[]SecurityArticle `json:"articles,omitempty"`

	// Retrieve details about hostComponent objects.Note: List retrieval is not yet supported.
	HostComponents *[]SecurityHostComponent `json:"hostComponents,omitempty"`

	// Retrieve details about hostCookie objects.Note: List retrieval is not yet supported.
	HostCookies *[]SecurityHostCookie `json:"hostCookies,omitempty"`

	// Retrieve details about hostTracker objects.Note: List retrieval is not yet supported.
	HostPairs *[]SecurityHostPair `json:"hostPairs,omitempty"`

	// Retrieve details about hostPort objects.Note: List retrieval is not yet supported.
	HostPorts *[]SecurityHostPort `json:"hostPorts,omitempty"`

	// Retrieve details about hostSslCertificate objects.Note: List retrieval is not yet supported.
	HostSslCertificates *[]SecurityHostSslCertificate `json:"hostSslCertificates,omitempty"`

	// Retrieve details about hostTracker objects.Note: List retrieval is not yet supported.
	HostTrackers *[]SecurityHostTracker `json:"hostTrackers,omitempty"`

	// Refers to host objects that Microsoft Threat Intelligence has observed.Note: List retrieval is not yet supported.
	Hosts *[]SecurityHost `json:"hosts,omitempty"`

	// A list of intelligenceProfile objects.
	IntelProfiles *[]SecurityIntelligenceProfile `json:"intelProfiles,omitempty"`

	IntelligenceProfileIndicators *[]SecurityIntelligenceProfileIndicator `json:"intelligenceProfileIndicators,omitempty"`

	// Retrieve details about passiveDnsRecord objects.Note: List retrieval is not yet supported.
	PassiveDnsRecords *[]SecurityPassiveDnsRecord `json:"passiveDnsRecords,omitempty"`

	// Retrieve details about sslCertificate objects.Note: List retrieval is not yet supported.
	SslCertificates *[]SecuritySslCertificate `json:"sslCertificates,omitempty"`

	// Retrieve details about the subdomain.Note: List retrieval is not yet supported.
	Subdomains *[]SecuritySubdomain `json:"subdomains,omitempty"`

	// Retrieve details about vulnerabilities.Note: List retrieval is not yet supported.
	Vulnerabilities *[]SecurityVulnerability `json:"vulnerabilities,omitempty"`

	// Retrieve details about whoisHistoryRecord objects.Note: List retrieval is not yet supported.
	WhoisHistoryRecords *[]SecurityWhoisHistoryRecord `json:"whoisHistoryRecords,omitempty"`

	// A list of whoisRecord objects.
	WhoisRecords *[]SecurityWhoisRecord `json:"whoisRecords,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SecurityThreatIntelligence) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SecurityThreatIntelligence{}

func (s SecurityThreatIntelligence) MarshalJSON() ([]byte, error) {
	type wrapper SecurityThreatIntelligence
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SecurityThreatIntelligence: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityThreatIntelligence: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.security.threatIntelligence"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SecurityThreatIntelligence: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SecurityThreatIntelligence{}

func (s *SecurityThreatIntelligence) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ArticleIndicators             *[]SecurityArticleIndicator             `json:"articleIndicators,omitempty"`
		Articles                      *[]SecurityArticle                      `json:"articles,omitempty"`
		HostComponents                *[]SecurityHostComponent                `json:"hostComponents,omitempty"`
		HostCookies                   *[]SecurityHostCookie                   `json:"hostCookies,omitempty"`
		HostPairs                     *[]SecurityHostPair                     `json:"hostPairs,omitempty"`
		HostPorts                     *[]SecurityHostPort                     `json:"hostPorts,omitempty"`
		HostSslCertificates           *[]SecurityHostSslCertificate           `json:"hostSslCertificates,omitempty"`
		HostTrackers                  *[]SecurityHostTracker                  `json:"hostTrackers,omitempty"`
		IntelProfiles                 *[]SecurityIntelligenceProfile          `json:"intelProfiles,omitempty"`
		IntelligenceProfileIndicators *[]SecurityIntelligenceProfileIndicator `json:"intelligenceProfileIndicators,omitempty"`
		PassiveDnsRecords             *[]SecurityPassiveDnsRecord             `json:"passiveDnsRecords,omitempty"`
		SslCertificates               *[]SecuritySslCertificate               `json:"sslCertificates,omitempty"`
		Subdomains                    *[]SecuritySubdomain                    `json:"subdomains,omitempty"`
		Vulnerabilities               *[]SecurityVulnerability                `json:"vulnerabilities,omitempty"`
		WhoisHistoryRecords           *[]SecurityWhoisHistoryRecord           `json:"whoisHistoryRecords,omitempty"`
		WhoisRecords                  *[]SecurityWhoisRecord                  `json:"whoisRecords,omitempty"`
		Id                            *string                                 `json:"id,omitempty"`
		ODataId                       *string                                 `json:"@odata.id,omitempty"`
		ODataType                     *string                                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ArticleIndicators = decoded.ArticleIndicators
	s.Articles = decoded.Articles
	s.HostComponents = decoded.HostComponents
	s.HostCookies = decoded.HostCookies
	s.HostPairs = decoded.HostPairs
	s.HostPorts = decoded.HostPorts
	s.HostSslCertificates = decoded.HostSslCertificates
	s.HostTrackers = decoded.HostTrackers
	s.IntelProfiles = decoded.IntelProfiles
	s.IntelligenceProfileIndicators = decoded.IntelligenceProfileIndicators
	s.PassiveDnsRecords = decoded.PassiveDnsRecords
	s.SslCertificates = decoded.SslCertificates
	s.Subdomains = decoded.Subdomains
	s.Vulnerabilities = decoded.Vulnerabilities
	s.WhoisHistoryRecords = decoded.WhoisHistoryRecords
	s.WhoisRecords = decoded.WhoisRecords
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityThreatIntelligence into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["hosts"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Hosts into list []json.RawMessage: %+v", err)
		}

		output := make([]SecurityHost, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSecurityHostImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Hosts' for 'SecurityThreatIntelligence': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Hosts = &output
	}

	return nil
}
