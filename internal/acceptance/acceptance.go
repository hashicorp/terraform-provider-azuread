package acceptance

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-azuread/internal/provider"
)

var AzureADProvider *schema.Provider
var once sync.Once

func init() {
	if os.Getenv("TF_ACC") == "" {
		return
	}
	EnsureProvidersAreInitialised()
}

func EnsureProvidersAreInitialised() {
	once.Do(func() {
		AzureADProvider = provider.AzureADProvider()
	})
}

func EnvironmentName() string {
	envName, exists := os.LookupEnv("ARM_ENVIRONMENT")
	if !exists {
		envName = "public"
	}
	return envName
}

func Environment() (*azure.Environment, error) {
	envName := EnvironmentName()
	metadataHost := os.Getenv("ARM_METADATA_HOST")
	return authentication.AzureEnvironmentByNameFromEndpoint(context.TODO(), metadataHost, envName)
}

func RequiresImportError(resourceName string) *regexp.Regexp {
	message := "To be managed via Terraform, this resource needs to be imported into the State. Please see the resource documentation for %q for more information."
	message = strings.Replace(message, " ", "\\s+", -1)
	return regexp.MustCompile(fmt.Sprintf(message, resourceName))
}
