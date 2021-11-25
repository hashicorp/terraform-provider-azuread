package providerjson

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"net/http"
	"strings"
)

const (
	DataSourcesList = "/schema-data/v1/data-sources"
	ResourcesList   = "/schema-data/v1/resources"
	DataSourcesPath = "/schema-data/v1/data-sources/"
	ResourcesPath   = "/schema-data/v1/resources/"
)

func (p ProviderJSON) DataSourcesHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dsRaw := strings.Split(req.URL.RequestURI(), DataSourcesPath)
	ds := strings.Split(dsRaw[1], "/")[0]
	data, err := fromRaw(p.DataSourcesMap[ds])
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Could not process schema for %q from provider: %+v", ds, err)))
	}
	if len(ds) > 0 {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.Write([]byte(fmt.Sprintf("Marshall error: %+v", err)))
		}
	}
}

func (p ProviderJSON) ResourcesHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	dsRaw := strings.Split(req.URL.RequestURI(), ResourcesPath)
	ds := strings.Split(dsRaw[1], "/")[0]
	data, err := fromRaw(p.ResourcesMap[ds])
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Could not process schema for %q from provider: %+v", ds, err)))
	}
	if len(ds) > 0 {
		if err := json.NewEncoder(w).Encode(data); err != nil {
			w.Write([]byte(fmt.Sprintf("Marshall error: %+v", err)))
		}
	}
}

func (p *ProviderJSON) ListResources(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(p.Resources()); err != nil {
		panic(err)
	}
}

func (p *ProviderJSON) ListDataSources(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(p.DataSources()); err != nil {
		panic(err)
	}
}

func (p *ProviderJSON) DataSources() []terraform.DataSource {
	s := schema.Provider(*p)
	return s.DataSources()
}

func (p *ProviderJSON) Resources() []terraform.ResourceType {
	s := schema.Provider(*p)
	return s.Resources()
}
