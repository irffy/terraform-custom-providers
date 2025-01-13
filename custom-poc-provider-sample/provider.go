// filepath: /c:/Projects/terraform-custom-providers/custom-poc-provider-sample/provider.go
package main

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// func Provider() *schema.Provider {
//     return &schema.Provider{
//         ResourcesMap: map[string]*schema.Resource{},
//     }
// }

func Provider() *schema.Provider {
    return &schema.Provider{
        ResourcesMap: map[string]*schema.Resource{
            "example_server": resourceServer(),
        },
    }
}