package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/synadia-io/control-plane-sdk-go/controlplane"
)

func resourceCluster() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)

			req := &controlplane.CreateClusterRequest{
				Name:   d.Get("name").(string),
				Region: d.Get("region").(string),
				Tier:   d.Get("tier").(string),
			}

			orgID := d.Get("organization_id").(string)
			cluster, err := client.Clusters.CreateCluster(ctx, orgID, req)
			if err != nil {
				return diag.FromErr(err)
			}

			d.SetId(cluster.ID)
			return resourceClusterRead(ctx, d, m)
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			cluster, err := client.Clusters.GetCluster(ctx, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}

			d.Set("name", cluster.Name)
			d.Set("region", cluster.Region)
			d.Set("tier", cluster.Tier)
			d.Set("organization_id", cluster.Organization.ID)
			return nil
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			if err := client.Clusters.DeleteCluster(ctx, d.Id()); err != nil {
				return diag.FromErr(err)
			}
			d.SetId("")
			return nil
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"region": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tier": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "standard",
			},
			"organization_id": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

// ----------------------------------------------------------------

func resourceOrganization() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceOrganizationCreate,
		ReadContext:   resourceOrganizationRead,
		UpdateContext: resourceOrganizationUpdate,
		DeleteContext: resourceOrganizationDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceOrganizationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	req := &controlplane.CreateOrganizationRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	org, err := client.Organizations.CreateOrganization(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(org.ID)
	return resourceOrganizationRead(ctx, d, m)
}

func resourceOrganizationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	org, err := client.Organizations.GetOrganization(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", org.Name)
	d.Set("description", org.Description)
	return nil
}

func resourceOrganizationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	req := &controlplane.UpdateOrganizationRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	_, err := client.Organizations.UpdateOrganization(ctx, d.Id(), req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceOrganizationRead(ctx, d, m)
}

func resourceOrganizationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	err := client.Organizations.DeleteOrganization(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

// ----------------------------------------------------------------

func resourceProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceProjectCreate,
		ReadContext:   resourceProjectRead,
		UpdateContext: resourceProjectUpdate,
		DeleteContext: resourceProjectDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"organization_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceProjectCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	req := &controlplane.CreateProjectRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	orgID := d.Get("organization_id").(string)

	project, err := client.Projects.CreateProject(ctx, orgID, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(project.ID)
	return resourceProjectRead(ctx, d, m)
}

func resourceProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	orgID := d.Get("organization_id").(string)
	project, err := client.Projects.GetProject(ctx, orgID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", project.Name)
	d.Set("description", project.Description)
	return nil
}

func resourceProjectUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	req := &controlplane.UpdateProjectRequest{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
	}

	orgID := d.Get("organization_id").(string)

	_, err := client.Projects.UpdateProject(ctx, orgID, d.Id(), req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceProjectRead(ctx, d, m)
}

func resourceProjectDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	orgID := d.Get("organization_id").(string)
	err := client.Projects.DeleteProject(ctx, orgID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

// --------------------------------------------------

// synadia_user resource

func resourceUser() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUserCreate,
		ReadContext:   resourceUserRead,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Schema: map[string]*schema.Schema{
			"organization_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"email": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"roles": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	req := &controlplane.CreateUserRequest{
		Email: d.Get("email").(string),
		Name:  d.Get("name").(string),
	}

	if v, ok := d.GetOk("roles"); ok {
		for _, r := range v.([]interface{}) {
			req.Roles = append(req.Roles, r.(string))
		}
	}

	orgID := d.Get("organization_id").(string)
	projectID, _ := d.Get("project_id").(string)

	user, err := client.Users.CreateUser(ctx, orgID, projectID, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(user.ID)
	return resourceUserRead(ctx, d, m)
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	orgID := d.Get("organization_id").(string)
	projectID, _ := d.Get("project_id").(string)

	user, err := client.Users.GetUser(ctx, orgID, projectID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("email", user.Email)
	d.Set("name", user.Name)
	d.Set("roles", user.Roles)
	return nil
}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	req := &controlplane.UpdateUserRequest{
		Name: d.Get("name").(string),
	}

	if v, ok := d.GetOk("roles"); ok {
		req.Roles = make([]string, 0)
		for _, r := range v.([]interface{}) {
			req.Roles = append(req.Roles, r.(string))
		}
	}

	orgID := d.Get("organization_id").(string)
	projectID, _ := d.Get("project_id").(string)

	_, err := client.Users.UpdateUser(ctx, orgID, projectID, d.Id(), req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceUserRead(ctx, d, m)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	orgID := d.Get("organization_id").(string)
	projectID, _ := d.Get("project_id").(string)

	err := client.Users.DeleteUser(ctx, orgID, projectID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

// --------------------------------------------------
// synadia_jwt_claim resource

func resourceJWTClaim() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceJWTClaimCreate,
		ReadContext:   resourceJWTClaimRead,
		UpdateContext: resourceJWTClaimUpdate,
		DeleteContext: resourceJWTClaimDelete,
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"jwt": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"permissions": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceJWTClaimCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	req := &controlplane.CreateJWTClaimRequest{
		UserID:      d.Get("user_id").(string),
		Permissions: expandStringList(d.Get("permissions").([]interface{})),
	}

	claim, err := client.JWTClaims.CreateJWTClaim(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(claim.ID)
	d.Set("jwt", claim.JWT)
	return resourceJWTClaimRead(ctx, d, m)
}

func resourceJWTClaimRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	claim, err := client.JWTClaims.GetJWTClaim(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("user_id", claim.UserID)
	d.Set("permissions", claim.Permissions)
	d.Set("jwt", claim.JWT)
	return nil
}

func resourceJWTClaimUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	req := &controlplane.UpdateJWTClaimRequest{
		Permissions: expandStringList(d.Get("permissions").([]interface{})),
	}

	_, err := client.JWTClaims.UpdateJWTClaim(ctx, d.Id(), req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceJWTClaimRead(ctx, d, m)
}

func resourceJWTClaimDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	err := client.JWTClaims.DeleteJWTClaim(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

// --------------------------------------------------
// synadia_permission resource

func resourcePermission() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourcePermissionCreate,
		ReadContext:   resourcePermissionRead,
		UpdateContext: resourcePermissionUpdate,
		DeleteContext: resourcePermissionDelete,
		Schema: map[string]*schema.Schema{
			"subject": {
				Type:     schema.TypeString,
				Required: true,
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"allow": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

func resourcePermissionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	req := &controlplane.CreatePermissionRequest{
		Subject: d.Get("subject").(string),
		Action:  d.Get("action").(string),
		UserID:  d.Get("user_id").(string),
		Allow:   d.Get("allow").(bool),
	}

	perm, err := client.Permissions.CreatePermission(ctx, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(perm.ID)
	return resourcePermissionRead(ctx, d, m)
}

func resourcePermissionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	perm, err := client.Permissions.GetPermission(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("subject", perm.Subject)
	d.Set("action", perm.Action)
	d.Set("user_id", perm.UserID)
	d.Set("allow", perm.Allow)
	return nil
}

func resourcePermissionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	req := &controlplane.UpdatePermissionRequest{
		Subject: d.Get("subject").(string),
		Action:  d.Get("action").(string),
		Allow:   d.Get("allow").(bool),
	}

	_, err := client.Permissions.UpdatePermission(ctx, d.Id(), req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourcePermissionRead(ctx, d, m)
}

func resourcePermissionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	err := client.Permissions.DeletePermission(ctx, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

// --------------------------------------------------
// synadia_stream resource

func resourceStream() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceStreamCreate,
		ReadContext:   resourceStreamRead,
		UpdateContext: resourceStreamUpdate,
		DeleteContext:   resourceStreamDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subjects": {
				Type:     schema.TypeList,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"max_msgs": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_bytes": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"max_age_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func resourceStreamCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	req := &controlplane.CreateStreamRequest{
		Name:     d.Get("name").(string),
		Subjects: expandStringList(d.Get("subjects").([]interface{})),
	}

	if v, ok := d.GetOk("max_msgs"); ok {
		req.MaxMsgs = int32(v.(int))
	}
	if v, ok := d.GetOk("max_bytes"); ok {
		req.MaxBytes = int32(v.(int))
	}
	if v, ok := d.GetOk("max_age_seconds"); ok {
		req.MaxAgeSeconds = int32(v.(int))
	}

	clusterID := d.Get("cluster_id").(string)

	stream, err := client.Streams.CreateStream(ctx, clusterID, req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(stream.ID)
	return resourceStreamRead(ctx, d, m)
}

func resourceStreamRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	clusterID := d.Get("cluster_id").(string)
	stream, err := client.Streams.GetStream(ctx, clusterID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", stream.Name)
	d.Set("subjects", stream.Subjects)
	d.Set("max_msgs", stream.MaxMsgs)
	d.Set("max_bytes", stream.MaxBytes)
	d.Set("max_age_seconds", stream.MaxAgeSeconds)
	return nil
}

func resourceStreamUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	req := &controlplane.UpdateStreamRequest{
		Name:     d.Get("name").(string),
		Subjects: expandStringList(d.Get("subjects").([]interface{})),
	}

	if v, ok := d.GetOk("max_msgs"); ok {
		req.MaxMsgs = int32(v.(int))
	}
	if v, ok := d.GetOk("max_bytes"); ok {
		req.MaxBytes = int32(v.(int))
	}
	if v, ok := d.GetOk("max_age_seconds"); ok {
		req.MaxAgeSeconds = int32(v.(int))
	}

	clusterID := d.Get("cluster_id").(string)

	_, err := client.Streams.UpdateStream(ctx, clusterID, d.Id(), req)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceStreamRead(ctx, d, m)
}

func resourceStreamDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*controlplane.Client)

	clusterID := d.Get("cluster_id").(string)
	err := client.Streams.DeleteStream(ctx, clusterID, d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

// --------------------------------------------------

// Helper to convert []interface{} to []string
func expandStringList(list []interface{}) []string {
	res := make([]string, len(list))
	for i, v := range list {
		res[i] = v.(string)
	}
	return res
}

//-------------------------------------------------
func resourceConsumer() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			req := &controlplane.CreateConsumerRequest{
				Name:     d.Get("name").(string),
				Durable:  d.Get("durable").(string),
				StreamID: d.Get("stream_id").(string),
			}
			clusterID := d.Get("cluster_id").(string)
			consumer, err := client.Consumers.CreateConsumer(ctx, clusterID, req)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(consumer.ID)
			return resourceConsumer().ReadContext(ctx, d, m)
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			clusterID := d.Get("cluster_id").(string)
			consumer, err := client.Consumers.GetConsumer(ctx, clusterID, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			d.Set("name", consumer.Name)
			d.Set("durable", consumer.Durable)
			d.Set("stream_id", consumer.StreamID)
			return nil
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			clusterID := d.Get("cluster_id").(string)
			err := client.Consumers.DeleteConsumer(ctx, clusterID, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId("")
			return nil
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {Type: schema.TypeString, Required: true},
			"stream_id":  {Type: schema.TypeString, Required: true},
			"name":       {Type: schema.TypeString, Required: true},
			"durable":    {Type: schema.TypeString, Required: true},
		},
	}
}

//-------------------------------------------------

func resourceKVBucket() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			req := &controlplane.CreateKVBucketRequest{
				Name:      d.Get("name").(string),
				MaxValueSize: int64(d.Get("max_value_size").(int)),
			}
			clusterID := d.Get("cluster_id").(string)
			bucket, err := client.KV.CreateKVBucket(ctx, clusterID, req)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(bucket.ID)
			return resourceKVBucket().ReadContext(ctx, d, m)
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			clusterID := d.Get("cluster_id").(string)
			bucket, err := client.KV.GetKVBucket(ctx, clusterID, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			d.Set("name", bucket.Name)
			d.Set("max_value_size", int(bucket.MaxValueSize))
			return nil
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			clusterID := d.Get("cluster_id").(string)
			err := client.KV.DeleteKVBucket(ctx, clusterID, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId("")
			return nil
		},
		Schema: map[string]*schema.Schema{
			"cluster_id":      {Type: schema.TypeString, Required: true},
			"name":            {Type: schema.TypeString, Required: true},
			"max_value_size":  {Type: schema.TypeInt, Optional: true, Default: 1024},
		},
	}
}

//-------------------------------------------------

func resourceObjectStore() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			req := &controlplane.CreateObjectStoreRequest{
				Name: d.Get("name").(string),
			}
			clusterID := d.Get("cluster_id").(string)
			store, err := client.ObjectStore.CreateObjectStore(ctx, clusterID, req)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(store.ID)
			return resourceObjectStore().ReadContext(ctx, d, m)
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			clusterID := d.Get("cluster_id").(string)
			store, err := client.ObjectStore.GetObjectStore(ctx, clusterID, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			d.Set("name", store.Name)
			return nil
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			clusterID := d.Get("cluster_id").(string)
			err := client.ObjectStore.DeleteObjectStore(ctx, clusterID, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId("")
			return nil
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {Type: schema.TypeString, Required: true},
			"name":       {Type: schema.TypeString, Required: true},
		},
	}
}

//-------------------------------------------------

func resourceClusterGateway() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			req := &controlplane.CreateGatewayRequest{
				RemoteClusterID: d.Get("remote_cluster_id").(string),
				Name:            d.Get("name").(string),
			}
			clusterID := d.Get("cluster_id").(string)
			gw, err := client.Clusters.CreateGateway(ctx, clusterID, req)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(gw.ID)
			return resourceClusterGateway().ReadContext(ctx, d, m)
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			clusterID := d.Get("cluster_id").(string)
			gw, err := client.Clusters.GetGateway(ctx, clusterID, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			d.Set("remote_cluster_id", gw.RemoteClusterID)
			d.Set("name", gw.Name)
			return nil
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			clusterID := d.Get("cluster_id").(string)
			err := client.Clusters.DeleteGateway(ctx, clusterID, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId("")
			return nil
		},
		Schema: map[string]*schema.Schema{
			"cluster_id":        {Type: schema.TypeString, Required: true},
			"remote_cluster_id": {Type: schema.TypeString, Required: true},
			"name":              {Type: schema.TypeString, Required: true},
		},
	}
}

//-------------------------------------------------

func resourceLeafnode() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			req := &controlplane.CreateLeafnodeRequest{
				RemoteURL: d.Get("remote_url").(string),
				Name:      d.Get("name").(string),
			}
			clusterID := d.Get("cluster_id").(string)
			ln, err := client.Clusters.CreateLeafnode(ctx, clusterID, req)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(ln.ID)
			return resourceLeafnode().ReadContext(ctx, d, m)
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			clusterID := d.Get("cluster_id").(string)
			ln, err := client.Clusters.GetLeafnode(ctx, clusterID, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			d.Set("remote_url", ln.RemoteURL)
			d.Set("name", ln.Name)
			return nil
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			clusterID := d.Get("cluster_id").(string)
			err := client.Clusters.DeleteLeafnode(ctx, clusterID, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId("")
			return nil
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"remote_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
//-------------------------------------------------
func resourceServiceExport() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			req := &controlplane.CreateServiceExportRequest{
				Name:       d.Get("name").(string),
				Subject:    d.Get("subject").(string),
				Visibility: d.Get("visibility").(string),
			}
			clusterID := d.Get("cluster_id").(string)
			exp, err := client.Services.CreateServiceExport(ctx, clusterID, req)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(exp.ID)
			return resourceServiceExport().ReadContext(ctx, d, m)
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			clusterID := d.Get("cluster_id").(string)
			exp, err := client.Services.GetServiceExport(ctx, clusterID, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			d.Set("name", exp.Name)
			d.Set("subject", exp.Subject)
			d.Set("visibility", exp.Visibility)
			return nil
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			clusterID := d.Get("cluster_id").(string)
			err := client.Services.DeleteServiceExport(ctx, clusterID, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId("")
			return nil
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subject": {
				Type:     schema.TypeString,
				Required: true,
			},
			"visibility": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
//-------------------------------------------------

func resourceServiceImport() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			req := &controlplane.CreateServiceImportRequest{
				Name:           d.Get("name").(string),
				RemoteCluster:  d.Get("remote_cluster").(string),
				SubjectMapping: d.Get("subject_mapping").(string),
			}
			clusterID := d.Get("cluster_id").(string)
			imp, err := client.Services.CreateServiceImport(ctx, clusterID, req)
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId(imp.ID)
			return resourceServiceImport().ReadContext(ctx, d, m)
		},
		ReadContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			clusterID := d.Get("cluster_id").(string)
			imp, err := client.Services.GetServiceImport(ctx, clusterID, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			d.Set("name", imp.Name)
			d.Set("remote_cluster", imp.RemoteCluster)
			d.Set("subject_mapping", imp.SubjectMapping)
			return nil
		},
		DeleteContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
			client := m.(*controlplane.Client)
			clusterID := d.Get("cluster_id").(string)
			err := client.Services.DeleteServiceImport(ctx, clusterID, d.Id())
			if err != nil {
				return diag.FromErr(err)
			}
			d.SetId("")
			return nil
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"remote_cluster": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subject_mapping": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
