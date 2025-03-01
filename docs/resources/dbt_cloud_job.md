---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "dbt_cloud_job Resource - terraform-provider-dbt-cloud"
subcategory: ""
description: |-
  
---

# dbt_cloud_job (Resource)

## Example Usage

```terraform
resource "dbt_cloud_job" "test" {
  project_id     = 101
  environment_id = 1
  name           = "Test job"
  execute_steps  = [
    "dbt run",
    "dbt test"
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **environment_id** (Number) Environment ID to create the job in
- **execute_steps** (List of String) List of commands to execute for the job
- **name** (String) Job name
- **project_id** (Number) Project ID to create the job in

### Optional

- **dbt_version** (String) Version number of DBT to use in this job
- **generate_docs** (Boolean) Flag for whether the job should generate documentation
- **id** (String) The ID of this resource.
- **is_active** (Boolean) Flag for whether the job is marked active or deleted
- **num_threads** (Number) Number of threads to use in the job
- **run_generate_sources** (Boolean) Flag for whether the job should run generate sources
- **target_name** (String) Target name for the DBT profile
- **triggers** (Map of Boolean) Flags for which types of triggers to use, keys of github_webhook, schedule, custom_branch_only


