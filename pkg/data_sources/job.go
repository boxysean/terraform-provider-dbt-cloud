package data_sources

import (
	"context"
	"strconv"

	"github.com/gthesheep/terraform-provider-dbt-cloud/pkg/dbt_cloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var jobSchema = map[string]*schema.Schema{
	"project_id": &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
	},
	"environment_id": &schema.Schema{
		Type:     schema.TypeInt,
		Computed: true,
	},
	"name": &schema.Schema{
		Type:     schema.TypeString,
		Computed: true,
	},
	"job_id": &schema.Schema{
		Type:     schema.TypeInt,
		Required: true,
	},
}

func DatasourceJob() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceJobRead,
		Schema:      jobSchema,
	}
}

func datasourceJobRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dbt_cloud.Client)

	var diags diag.Diagnostics

	jobId := strconv.Itoa(d.Get("job_id").(int))

	job, err := c.GetJob(jobId)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("project_id", job.Project_Id); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("environment_id", job.Environment_Id); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", job.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("execute_steps", job.Execute_Steps); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(jobId)

	return diags
}
