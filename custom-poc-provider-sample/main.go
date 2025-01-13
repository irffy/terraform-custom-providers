// filepath: /c:/Projects/terraform-custom-providers/custom-poc-provider-sample/main.go
package main

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: Provider,
    })
}