package environments

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type metadataEndpoints struct {
	GalleryEndpoint  string `json:"galleryEndpoint"`
	AadGraphEndpoint string `json:"graphEndpoint"`
	AadGraphAudience string `json:"graphAudience"`
	PortalEndpoint   string `json:"portalEndpoint"`
	Authentication   struct {
		LoginEndpoint       string   `json:"loginEndpoint"`
		ManagementAudiences []string `json:"audiences"`
		Tenant              string   `json:"tenant"`
	} `json:"authentication"`
}

func EnvironmentFromMetadata(endpoint string) (*Environment, error) {
	uri := fmt.Sprintf("%s/%s", strings.TrimSuffix(endpoint, "/"), "/metadata/endpoints?api-version=1.0")
	resp, err := http.Get(uri)
	if err != nil {
		return nil, fmt.Errorf("unable to load metadata from %q: %v", uri, err)
	}

	defer resp.Body.Close()
	jsonResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response from %q: %v", uri, err)
	}

	var endpoints metadataEndpoints
	err = json.Unmarshal(jsonResponse, &endpoints)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling response from %q: %v", uri, err)
	}

	resourceManagerEndpoint := strings.TrimSuffix(endpoint, "/")

	env := Environment{
		AzureADEndpoint: AzureADEndpoint(strings.TrimSuffix(endpoints.Authentication.LoginEndpoint, "/")),
		AadGraph: Api{
			AppId:    PublishedApis["AzureActiveDirectoryGraph"],
			Endpoint: ApiEndpoint(strings.TrimSuffix(endpoints.AadGraphEndpoint, "/")),
		},
		ResourceManager: Api{
			AppId:    PublishedApis["AzureServiceManagement"],
			Endpoint: ApiEndpoint(resourceManagerEndpoint),
		},
	}

	// The following logic borrowed from go-autorest in order to meet user expectations
	if parts := strings.Split(resourceManagerEndpoint, "."); len(parts) > 2 {
		dnsSuffix := fmt.Sprintf("%s.%s", parts[len(parts)-2], parts[len(parts)-1])
		env.KeyVault = Api{
			AppId:    PublishedApis["AzureKeyVault"],
			Endpoint: ApiEndpoint(fmt.Sprintf("https://%s.%s", "vault", dnsSuffix)),
		}
		env.Storage = Api{
			AppId:    PublishedApis["AzureStorage"],
			Endpoint: ApiEndpoint(fmt.Sprintf("https://%s.%s", "storage", dnsSuffix)),
		}
	}

	return &env, nil
}
