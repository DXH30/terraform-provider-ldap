package example

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
    return &schema.Provider{
        Schema: map[string]*schema.Schema{
            "api_key": &schema.Schema{
                Type: schema.TypeString,
                Required: true,
                DefaultFunc: schemaEnvDefaultFunc("EXAMPLE_API_KEY", nil),
            },
        },
        ResourcesMap: map[string]*schema.Resource{
            "example_resource": resourceExample(),
        },
    }
}
