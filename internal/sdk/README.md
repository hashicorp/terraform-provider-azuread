## SDK for Strongly-Typed Resources

This package is a prototype for creating strongly-typed Data Sources and Resources - and in future will likely form the foundation for Terraform Data Sources and Resources in this Provider going forward.

## Should I use this package to build resources?

Yes, going forward it is strongly preferred that new resources use this SDK. For guidance, please refer to the [AzureRM Provider Contributor Guide](https://github.com/hashicorp/terraform-provider-azurerm/blob/main/contributing/topics/guide-new-resource.md), amending as appropriate for the AzureAD Provider.

---

## What's the long-term intention for this package?

Each Service Package contains the following:

* Client - giving reference to the SDK Client which should be used to interact with Azure
* ID Parsers, Formatters and/or Validators - giving a canonical ID for each Resource 
* Validation functions specific to this service package, for example for the Name

This package can be used to tie these together in a more strongly typed fashion, for example:

```go
package example

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/manicminer/hamilton/msgraph"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/terraform-provider-azuread/internal/sdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/services/example/parse"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/pluginsdk"
	"github.com/hashicorp/terraform-provider-azuread/internal/tf/validation"
)

type ExampleModel struct {
	Description string `tfschema:"description"`
	DisplayName string `tfschema:"display_name"`
	ObjectId    string `tfschema:"object_id"`
}

type ExampleResource struct{}

func (r ExampleResource) IDValidationFunc() pluginsdk.SchemaValidateFunc {
	return validation.IsUUID
}

var _ sdk.Resource = ExampleResource{}

func (r ExampleResource) ResourceType() string {
	return "azuread_example"
}

func (r ExampleResource) ModelObject() interface{} {
	return &ExampleModel{}
}

func (r ExampleResource) Arguments() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"display_name": {
			Description:      "The display name of the resource",
			Type:             pluginsdk.TypeString,
			Optional:         true,
			ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
		},

		"description": {
			Description:      "The description of the resource",
			Type:             pluginsdk.TypeString,
			Optional:         true,
			ValidateDiagFunc: validation.ValidateDiag(validation.StringIsNotEmpty),
		},
	}
}

func (r ExampleResource) Attributes() map[string]*pluginsdk.Schema {
	return map[string]*pluginsdk.Schema{
		"object_id": {
			Description: "The object ID of the resource",
			Type:        pluginsdk.TypeString,
			Computed:    true,
		},
	}
}

func (r ExampleResource) Create() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Example.ExampleClient

			var model ExampleModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			properties := msgraph.Group{
				Description: msgraph.NullableString(msgraph.StringNullWhenEmpty(model.Description)),
				DisplayName: &model.DisplayName,
			}

			result, _, err := client.Create(ctx, properties)
			if err != nil {
				return fmt.Errorf("creating %s: %+v", parse.ExampleId{}, err)
			}
			if result.ID() == nil {
				return fmt.Errorf("creating %s: returned ID was nil (API error)", parse.ExampleId{})
			}

			id := parse.NewExampleID(*result.ID())
			metadata.SetID(id)

			return nil
		},
	}
}

func (r ExampleResource) Update() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Example.ExampleClient

			id, err := parse.ExampleID(metadata.ResourceData.Id())
			if err != nil {
				return err
			}

			var model ExampleModel
			if err := metadata.Decode(&model); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			properties := msgraph.Group{
				DirectoryObject: msgraph.DirectoryObject{
					Id: pointer.To(id.ID()),
				},
				DisplayName: &model.DisplayName,
			}

			if metadata.ResourceData.HasChange("display_name") {
				properties.Description = msgraph.NullableString(msgraph.StringNullWhenEmpty(model.Description))
			}

			_, err = client.Update(ctx, properties)
			if err != nil {
				return fmt.Errorf("updating %s: %+v", id, err)
			}

			return nil
		},
	}
}

func (r ExampleResource) Read() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Example.ExampleClient

			id := parse.NewExampleID(metadata.ResourceData.Id())

			var state ExampleModel
			if err := metadata.Decode(&state); err != nil {
				return fmt.Errorf("decoding: %+v", err)
			}

			result, status, err := client.Get(ctx, id.ID(), odata.Query{})
			if err != nil {
				if status == http.StatusNotFound {
					return metadata.MarkAsGone(id)
				}
				return fmt.Errorf("retrieving %s: %+v", id, err)
			}
			if result == nil {
				return fmt.Errorf("retrieving %s: API error, result was nil", id)
			}

			state.Description = string(pointer.From(result.Description))
			state.DisplayName = pointer.From(result.DisplayName)
			state.ObjectId = pointer.From(result.ID())

			return metadata.Encode(&state)
		},
	}
}

func (r ExampleResource) Delete() sdk.ResourceFunc {
	return sdk.ResourceFunc{
		Timeout: 5 * time.Minute,
		Func: func(ctx context.Context, metadata sdk.ResourceMetaData) error {
			client := metadata.Client.Example.ExampleClient

			id := parse.NewExampleID(metadata.ResourceData.Id())

			_, err := client.Delete(ctx, id.ID())
			if err != nil {
				return fmt.Errorf("deleting %s: %+v", id, err)
			}

			return nil
		},
	}
}
```

The end result being the removal of a lot of common bugs by moving to a convention - for example:

* The Context object passed into each method _always_ has a deadline/timeout attached to it
* The Read function is automatically called at the end of a Create and Update function - meaning users don't have to do this 
* Each Resource has to have an ID Formatter and Validation Function
* The Model Object is validated via unit tests to ensure it contains the relevant struct tags (TODO: also confirming these exist in the state and are of the correct type, so no Set errors occur)

Ultimately this improves maintainability and allows bugs to be caught by the Compiler (for example if a Read function is unimplemented) - or Unit Tests (for example should the `tfschema` struct tags be missing) - rather than during Provider Initialization, which reduces the feedback loop.
