package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/synadia-io/control-plane-sdk-go/controlplane"
)

func dataSourceClusters() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			clusters, err := client.Clusters.ListClusters(ctx)
			if err != nil {
				return diag.FromErr(err)
			}

			ids := make([]string, len(clusters))
			for i, cluster := range clusters {
				ids[i] = cluster.ID
			}
			d.SetId("synadia_clusters")
			d.Set("ids", ids)
			return nil
		},
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceCluster() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			id := d.Get("id").(string)
			cluster, err := client.Clusters.GetCluster(ctx, id)
			if err != nil {
				return diag.FromErr(err)
			}

			d.SetId(cluster.ID)
			d.Set("name", cluster.Name)
			d.Set("region", cluster.Region)
			d.Set("tier", cluster.Tier)
			return nil
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"region": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tier": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceOrganization() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			org, err := client.Organizations.GetMyOrganization(ctx)
			if err != nil {
				return diag.FromErr(err)
			}

			d.SetId(org.ID)
			d.Set("name", org.Name)
			return nil
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			email := d.Get("email").(string)
			user, err := client.Users.GetUserByEmail(ctx, email)
			if err != nil {
				return diag.FromErr(err)
			}

			d.SetId(user.ID)
			d.Set("name", user.Name)
			d.Set("email", user.Email)
			return nil
		},
		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceStream() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			// Placeholder: implement once control-plane exposes streams.
			d.SetId("stream-placeholder")
			return nil
		},
		Schema: map[string]*schema.Schema{},
	}
}

func dataSourceKVBucket() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			// Placeholder: implement once control-plane exposes KV bucket info.
			d.SetId("kv-placeholder")
			return nil
		},
		Schema: map[string]*schema.Schema{},
	}
}

