## SDK for Strongly-Typed Resources

This package is a prototype for creating strongly-typed Data Sources and Resources - and in future will likely form the foundation for Terraform Data Sources and Resources in this Provider going forward.

## Should I use this package to build resources?

Yes, going forward it is strongly preferred that new resources use this SDK. For guidance, please refer to the [AzureRM Provider Contributor Guide](https://github.com/hashicorp/terraform-provider-azurerm/blob/main/contributing/topics/guide-new-resource.md), amending as appropriate for the AzureAD Provider.

## What's the long-term intention for this package?

Each Service Package contains the following:

* Client - giving reference to the SDK Client which should be used to interact with Azure
* ID Parsers, Formatters and/or Validators - giving a canonical ID for each Resource 
* Validation functions specific to this service package, for example for the Name

This package can be used to tie these together in a more strongly typed fashion.
The end result being the removal of a lot of common bugs by moving to a convention:

* The Context object passed into each method _always_ has a deadline/timeout attached to it
* The Read function is automatically called at the end of a Create and Update function - meaning users don't have to do this 
* Each Resource has to have an ID Formatter and Validation Function
* The Model Object is validated via unit tests to ensure it contains the relevant struct tags (TODO: also confirming these exist in the state and are of the correct type, so no Set errors occur)

Ultimately this improves maintainability and allows bugs to be caught by the Compiler (for example if a Read function is unimplemented) - or Unit Tests (for example should the `tfschema` struct tags be missing) - rather than during Provider Initialization, which reduces the feedback loop.
