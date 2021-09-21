package odata

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

const (
	ErrorAddedObjectReferencesAlreadyExist      = "One or more added object references already exist"
	ErrorCannotDeleteOrUpdateEnabledEntitlement = "Permission (scope or role) cannot be deleted or updated unless disabled first"
	ErrorConflictingObjectPresentInDirectory    = "A conflicting object with one or more of the specified property values is present in the directory"
	ErrorResourceDoesNotExist                   = "Resource '.+' does not exist or one of its queried reference-property objects are not present"
	ErrorRemovedObjectReferencesDoNotExist      = "One or more removed object references do not exist"
	ErrorServicePrincipalInvalidAppId           = "The appId '.+' of the service principal does not reference a valid application object"
)

type Id string

func (o *Id) UnmarshalJSON(data []byte) error {
	var id string
	if err := json.Unmarshal(data, &id); err != nil {
		return err
	}
	*o = Id(regexp.MustCompile(`/v2/`).ReplaceAllString(id, `/v1.0/`))
	return nil
}

type ShortType = string

const (
	ShortTypeAdministrativeUnit                          ShortType = "administrativeUnit"
	ShortTypeApplication                                 ShortType = "application"
	ShortTypeConditionalAccessPolicy                     ShortType = "conditionalAccessPolicy"
	ShortTypeCountryNamedLocation                        ShortType = "countryNamedLocation"
	ShortTypeDevice                                      ShortType = "device"
	ShortTypeDirectoryRole                               ShortType = "directoryRole"
	ShortTypeDirectoryRoleTemplate                       ShortType = "directoryRoleTemplate"
	ShortTypeDomain                                      ShortType = "domain"
	ShortTypeEmailAuthenticationMethod                   ShortType = "emailAuthenticationMethod"
	ShortTypeFido2AuthenticationMethod                   ShortType = "fido2AuthenticationMethod"
	ShortTypeGroup                                       ShortType = "group"
	ShortTypeIpNamedLocation                             ShortType = "ipNamedLocation"
	ShortTypeNamedLocation                               ShortType = "namedLocation"
	ShortTypeMicrosoftAuthenticatorAuthenticationMethod  ShortType = "microsoftAuthenticatorAuthenticationMethod"
	ShortTypeOrganization                                ShortType = "organization"
	ShortTypePasswordAuthenticationMethod                ShortType = "passwordAuthenticationMethod"
	ShortTypePhoneAuthenticationMethod                   ShortType = "phoneAuthenticationMethod"
	ShortTypeServicePrincipal                            ShortType = "servicePrincipal"
	ShortTypeSocialIdentityProvider                      ShortType = "socialIdentityProvider"
	ShortTypeTemporaryAccessPassAuthenticationMethod     ShortType = "temporaryAccessPassAuthenticationMethod"
	ShortTypeUser                                        ShortType = "user"
	ShortTypeWindowsHelloForBusinessAuthenticationMethod ShortType = "windowsHelloForBusinessAuthenticationMethod"
)

type Type = string

const (
	TypeAdministrativeUnit                          Type = "#microsoft.graph.administrativeUnit"
	TypeApplication                                 Type = "#microsoft.graph.application"
	TypeConditionalAccessPolicy                     Type = "#microsoft.graph.conditionalAccessPolicy"
	TypeCountryNamedLocation                        Type = "#microsoft.graph.countryNamedLocation"
	TypeDevice                                      Type = "#microsoft.graph.device"
	TypeDirectoryRole                               Type = "#microsoft.graph.directoryRole"
	TypeDirectoryRoleTemplate                       Type = "#microsoft.graph.directoryRoleTemplate"
	TypeDomain                                      Type = "#microsoft.graph.domain"
	TypeEmailAuthenticationMethod                   Type = "#microsoft.graph.emailAuthenticationMethod"
	TypeFido2AuthenticationMethod                   Type = "#microsoft.graph.fido2AuthenticationMethod"
	TypeGroup                                       Type = "#microsoft.graph.group"
	TypeIpNamedLocation                             Type = "#microsoft.graph.ipNamedLocation"
	TypeNamedLocation                               Type = "#microsoft.graph.namedLocation"
	TypeMicrosoftAuthenticatorAuthenticationMethod  Type = "#microsoft.graph.microsoftAuthenticatorAuthenticationMethod"
	TypeOrganization                                Type = "#microsoft.graph.organization"
	TypePasswordAuthenticationMethod                Type = "#microsoft.graph.passwordAuthenticationMethod"
	TypePhoneAuthenticationMethod                   Type = "#microsoft.graph.phoneAuthenticationMethod"
	TypeServicePrincipal                            Type = "#microsoft.graph.servicePrincipal"
	TypeSocialIdentityProvider                      Type = "#microsoft.graph.socialIdentityProvider"
	TypeTemporaryAccessPassAuthenticationMethod     Type = "#microsoft.graph.temporaryAccessPassAuthenticationMethod"
	TypeUser                                        Type = "#microsoft.graph.user"
	TypeWindowsHelloForBusinessAuthenticationMethod Type = "#microsoft.graph.windowsHelloForBusinessAuthenticationMethod"
)

// OData is used to unmarshall OData metadata from an API response.
type OData struct {
	Context      *string `json:"@odata.context"`
	MetadataEtag *string `json:"@odata.metadataEtag"`
	Type         *Type   `json:"@odata.type"`
	Count        *int    `json:"@odata.count"`
	NextLink     *string `json:"@odata.nextLink"`
	Delta        *string `json:"@odata.delta"`
	DeltaLink    *string `json:"@odata.deltaLink"`
	Id           *Id     `json:"@odata.id"`
	Etag         *string `json:"@odata.etag"`

	Error *Error `json:"-"`

	Value interface{} `json:"value"`
}

func (o *OData) UnmarshalJSON(data []byte) error {
	// Perform unmarshalling using a local type
	type odata OData
	var o2 odata
	if err := json.Unmarshal(data, &o2); err != nil {
		return err
	}
	*o = OData(o2)

	// Look for errors in the "error" and "odata.error" fields
	var e map[string]json.RawMessage
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	for _, k := range []string{"error", "odata.error"} {
		if v, ok := e[k]; ok {
			var e2 Error
			if err := json.Unmarshal(v, &e2); err != nil {
				return err
			}
			o.Error = &e2
			break
		}
	}
	return nil
}

// Error is used to unmarshal an API error message.
type Error struct {
	Code            *string          `json:"code"`
	Date            *string          `json:"date"`
	Message         *string          `json:"-"`
	RawMessage      *json.RawMessage `json:"message"` // sometimes a string, sometimes an object :/
	ClientRequestId *string          `json:"client-request-id"`
	RequestId       *string          `json:"request-id"`

	InnerError *Error `json:"innerError"` // nested errors

	Details *[]struct {
		Code   *string `json:"code"`
		Target *string `json:"target"`
	} `json:"details"`

	Values *[]struct {
		Item  string `json:"item"`
		Value string `json:"value"`
	} `json:"values"`
}

func (e *Error) UnmarshalJSON(data []byte) error {
	// Perform unmarshalling using a local type
	type error Error
	var e2 error
	if err := json.Unmarshal(data, &e2); err != nil {
		return err
	}
	*e = Error(e2)

	// Unmarshal the message, which can be a plain string or an object wrapping a message
	if raw := e.RawMessage; raw != nil && len(*raw) > 0 {
		switch string((*raw)[0]) {
		case "\"":
			var s string
			if err := json.Unmarshal(*raw, &s); err != nil {
				return err
			}
			e.Message = &s
		case "{":
			var m map[string]interface{}
			if err := json.Unmarshal(*raw, &m); err != nil {
				return err
			}
			if v, ok := m["value"]; ok {
				if s, ok := v.(string); ok {
					e.Message = &s
				}
			}
		default:
			return fmt.Errorf("unrecognised error message: %#v", string(*raw))
		}
	}
	return nil
}

func (e Error) String() string {
	sl := make([]string, 0)
	if e.Code != nil {
		sl = append(sl, *e.Code)
	}
	if e.Message != nil {
		sl = append(sl, *e.Message)
	}
	if e.InnerError != nil {
		if is := e.InnerError.String(); is != "" {
			sl = append(sl, is)
		}
	}
	return strings.Join(sl, ": ")
}

func (e Error) Match(errorText string) bool {
	re := regexp.MustCompile(errorText)
	return re.MatchString(e.String())
}
