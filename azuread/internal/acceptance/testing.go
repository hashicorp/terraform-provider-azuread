package acceptance

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var AzureADProvider *schema.Provider
var SupportedProviders map[string]terraform.ResourceProvider

func PreCheck(t *testing.T) {
	variables := []string{
		"ARM_CLIENT_ID",
		"ARM_CLIENT_SECRET",
		"ARM_TENANT_ID",
		"ARM_TEST_LOCATION",
		"ARM_TEST_LOCATION_ALT",
	}

	for _, variable := range variables {
		value := os.Getenv(variable)
		if value == "" {
			t.Fatalf("`%s` must be set for acceptance tests!", variable)
		}
	}
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
	metadataURL := os.Getenv("ARM_METADATA_URL")
	return authentication.AzureEnvironmentByNameFromEndpoint(context.TODO(), metadataURL, envName)
}

func RequiresImportError(resourceName string) *regexp.Regexp {
	message := "to be managed via Terraform this resource needs to be imported into the State. Please see the resource documentation for %q for more information."
	return regexp.MustCompile(fmt.Sprintf(message, resourceName))
}
