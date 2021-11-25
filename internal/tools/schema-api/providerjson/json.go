package providerjson

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-azuread/internal/provider"
)

// type omit *struct{}

//type ProviderJSON struct {
//	*schema.Provider
//	//Schema map[string]*SchemaJSON `json:"schema"`
//}

type ProviderJSON schema.Provider

type SchemaJSON struct {
	Type        string      `json:"type,omitempty"`
	ConfigMode  string      `json:"config_mode,omitempty"`
	Optional    bool        `json:"optional,omitempty"`
	Required    bool        `json:"required,omitempty"`
	Default     interface{} `json:"default,omitempty"`
	Description string      `json:"description,omitempty"`
	Computed    bool        `json:"computed,omitempty"`
	ForceNew    bool        `json:"force_new,omitempty"`
	Elem        interface{} `json:"elem,omitempty"`
	MaxItems    int         `json:"max_items,omitempty"`
	MinItems    int         `json:"min_items,omitempty"`
}

type ResourceJSON struct {
	Schema map[string]SchemaJSON `json:"schema"`
	// MigrateState         omit                   `json:",omitempty"`
	// StateUpgraders       omit                   `json:",omitempty"`
	// Create               omit                   `json:",omitempty"`
	// Read                 omit                   `json:",omitempty"`
	// Update               omit                   `json:",omitempty"`
	// Delete               omit                   `json:",omitempty"`
	// Exists               omit                   `json:",omitempty"`
	// CreateContext        omit                   `json:",omitempty"`
	// ReadContext          omit                   `json:",omitempty"`
	// UpdateContext        omit                   `json:",omitempty"`
	// DeleteContext        omit                   `json:",omitempty"`
	// CreateWithoutTimeout omit                   `json:",omitempty"`
	// ReadWithoutTimeout   omit                   `json:",omitempty"`
	// UpdateWithoutTimeout omit                   `json:",omitempty"`
	// DeleteWithoutTimeout omit                   `json:",omitempty"`
	// CustomizeDiff        omit                   `json:",omitempty"`
	// Importer             omit                   `json:",omitempty"`
	Timeouts *ResourceTimeoutJSON `json:"-"`
}

type ResourceTimeoutJSON struct {
	*schema.ResourceTimeout
	Create *time.Duration `json:"create,omitempty,int"`
	Read   *time.Duration `json:"read,omitempty,int"`
	Delete *time.Duration `json:"delete,omitempty,int"`
	Update *time.Duration `json:"update,omitempty,int"`
}

//func (p *Provider) MarshalJSON() ([]byte, error) {
//	return json.Marshal(&struct {
//		Schema        interface{} `json:"schema"`
//		ResourceMap   interface{} `json:"resource_map,omitempty"`
//		DataSourceMap interface{} `json:"data_source_map,omitempty"`
//	}{
//		Schema:        p.Schema,
//		ResourceMap:   p.ResourcesMap,
//		DataSourceMap: p.DataSourcesMap,
//	})
//}
//
//func (s *Schema) MarshalJSON() ([]byte, error) {
//	return json.Marshal(&struct {
//		Type          schema.ValueType        `json:"type"`
//		ConfigMode    schema.SchemaConfigMode `json:"config_mode,omitempty"`
//		Optional      bool                    `json:"optional,omitempty"`
//		Required      bool                    `json:"required,omitempty"`
//		Default       interface{}             `json:"default,omitempty"`
//		Description   string                  `json:"description,omitempty"`
//		InputDefault  string                  `json:"-"`
//		Computed      bool                    `json:"computed,omitempty"`
//		ForceNew      bool                    `json:"force_new,omitempty"`
//		Elem          interface{}             `json:"elem,omitempty"`
//		MaxItems      int                     `json:"max_items,omitempty"`
//		MinItems      int                     `json:"min_items,omitempty"`
//		Set           interface{}             `json:"-"`
//		ComputedWhen  []string                `json:"computed_when,omitempty"`
//		ConflictsWith []string                `json:"conflicts_with,omitempty"`
//		ExactlyOneOf  []string                `json:"exactly_one_of,omitempty"`
//		AtLeastOneOf  []string                `json:"at_least_one_of,omitempty"`
//		RequiredWith  []string                `json:"required_with,omitempty"`
//		Deprecated    string                  `json:"deprecated,omitempty"`
//		Sensitive     bool                    `json:"sensitive,omitempty"`
//	}{
//		Type:          s.Type,
//		ConfigMode:    s.ConfigMode,
//		Optional:      s.Optional,
//		Required:      s.Required,
//		Default:       s.Default,
//		Description:   s.Description,
//		Computed:      s.Computed,
//		ForceNew:      s.ForceNew,
//		Elem:          s.Elem,
//		MaxItems:      s.MaxItems,
//		MinItems:      s.MinItems,
//		ComputedWhen:  s.ComputedWhen,
//		ConflictsWith: s.ConflictsWith,
//		ExactlyOneOf:  s.ExactlyOneOf,
//		AtLeastOneOf:  s.AtLeastOneOf,
//		RequiredWith:  s.RequiredWith,
//		Deprecated:    s.Deprecated,
//		Sensitive:     s.Sensitive,
//	})
//}

//func (r *Resource) MarshalJSON() ([]byte, error) {
//	return json.Marshal(&struct {
//		Schema        map[string]*schema.Schema `json:"schema,omitempty"`
//		SchemaVersion int                       `json:"schema_version,omitempty"`
//	}{
//		Schema:        r.Schema,
//		SchemaVersion: r.SchemaVersion,
//	})
//}
//
//type JsonSchema struct {
//	Type          interface{} `json:"type"`
//	ConfigMode    interface{} `json:"config_mode,omitempty"`
//	Optional      bool        `json:"optional,omitempty"`
//	Required      bool        `json:"required,omitempty"`
//	Default       interface{} `json:"default,omitempty"`
//	Description   string      `json:"description,omitempty"`
//	InputDefault  string      `json:"-"`
//	Computed      bool        `json:"computed,omitempty"`
//	ForceNew      bool        `json:"force_new,omitempty"`
//	Elem          interface{} `json:"elem,omitempty"`
//	MaxItems      int         `json:"max_items,omitempty"`
//	MinItems      int         `json:"min_items,omitempty"`
//	Set           interface{} `json:"-"`
//	ComputedWhen  []string    `json:"computed_when,omitempty"`
//	ConflictsWith []string    `json:"conflicts_with,omitempty"`
//	ExactlyOneOf  []string    `json:"exactly_one_of,omitempty"`
//	AtLeastOneOf  []string    `json:"at_least_one_of,omitempty"`
//	RequiredWith  []string    `json:"required_with,omitempty"`
//	Deprecated    string      `json:"deprecated,omitempty"`
//	Sensitive     bool        `json:"sensitive,omitempty"`
//}

//type ResourceData struct {
//	Schema        map[string]JsonSchema `json:"schema,omitempty"`
//	SchemaVersion int                   `json:"schema_version,omitempty"`
//	//MigrateState         interface{}        `json:"-"`
//	//StateUpgraders       interface{}        `json:"-"`
//	//Create               interface{}        `json:"-"`
//	//Read                 interface{}        `json:"-"`
//	//Update               interface{}        `json:"-"`
//	//Delete               interface{}        `json:"-"`
//	//Exists               interface{}        `json:"-"`
//	//CreateContext        interface{}        `json:"-"`
//	//ReadContext          interface{}        `json:"-"`
//	//UpdateContext        interface{}        `json:"-"`
//	//DeleteContext        interface{}        `json:"-"`
//	//CreateWithoutTimeout interface{}        `json:"-"`
//	//ReadWithoutTimeout   interface{}        `json:"-"`
//	//UpdateWithoutTimeout interface{}        `json:"-"`
//	//DeleteWithoutTimeout interface{}        `json:"-"`
//	// CustomizeDiff        interface{}        `json:"-"`
//}
//
//func (r ResourceData) MarshallJSON() ([]byte, error) {
//	return json.Marshal(r)
//}

//func ResourceCopy(input *schema.Resource) (r ResourceData) {
//	if input == nil {
//		return r
//	}
//	r.Schema = schemaCopy(input.Schema)
//	r.SchemaVersion = input.SchemaVersion
//
//	return r
//}

//func schemaCopy(input map[string]*schema.Schema) map[string]JsonSchema {
//	s := make(map[string]JsonSchema, 0)
//	for k, p := range input {
//		v := *p
//		s[k] = JsonSchema{
//			Type:          v.Type,
//			ConfigMode:    v.ConfigMode,
//			Optional:      v.Optional,
//			Required:      v.Required,
//			Default:       v.Default,
//			Description:   v.Description,
//			InputDefault:  v.InputDefault,
//			Computed:      v.Computed,
//			ForceNew:      v.ForceNew,
//			Elem:          v.Elem,
//			MaxItems:      v.MaxItems,
//			MinItems:      v.MinItems,
//			Set:           v.Set,
//			ComputedWhen:  v.ComputedWhen,
//			ConflictsWith: v.ConflictsWith,
//			ExactlyOneOf:  v.ExactlyOneOf,
//			AtLeastOneOf:  v.AtLeastOneOf,
//			RequiredWith:  v.RequiredWith,
//			Deprecated:    v.Deprecated,
//			Sensitive:     v.Sensitive,
//		}
//	}
//
//	return s
//}
//

func LoadData() *ProviderJSON {
	p := provider.AzureADProvider()
	return (*ProviderJSON)(p)
}

//
//func (p *Provider) ShowResource(w http.ResponseWriter, _ *http.Request) {
//
//}
