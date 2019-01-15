Terraform Provider for Azure Active Directory
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

General Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.11.x (to build the provider plugin)

Windows Specific Requirements
-----------------------------
- [Make for Windows](http://gnuwin32.sourceforge.net/packages/make.htm)
- [Git Bash for Windows](https://git-scm.com/download/win)

For *GNU32 Make*, make sure its bin path is added to PATH environment variable.*

For *Git Bash for Windows*, at the step of "Adjusting your PATH environment", please choose "Use Git and optional Unix tools from Windows Command Prompt".*

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terraform-providers/terraform-provider-azuread`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:terraform-providers/terraform-provider-azuread
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-azuread
$ make build
```

Using the provider
----------------------

```
# Configure the Microsoft Azure AD Provider
provider "azuread" {
  # NOTE: Environment Variables can also be used for Service Principal authentication
  # Terraform also supports authenticating via the Azure CLI too.
  # see here for more info: http://terraform.io/docs/providers/azuread/index.html

  # subscription_id = "..."
  # client_id       = "..."
  # client_secret   = "..."
  # tenant_id       = "..."
}

# Create an application
resource "azuread_application" "example" {
  name = "ExampleApp"
}

# Create a service principal
resource "azuread_service_principal" "example" {
  application_id = "${azuread_application.example.application_id}"
}
```

Further [usage documentation is available on the Terraform website](https://www.terraform.io/docs/providers/azuread/index.html).

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-azuread
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

The following ENV variables must be set in your shell prior to running acceptance tests:
- ARM_CLIENT_ID
- ARM_CLIENT_SECRET
- ARM_SUBSCRIPTION_ID
- ARM_TENANT_ID
- ARM_TEST_LOCATION
- ARM_TEST_LOCATION_ALT

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
