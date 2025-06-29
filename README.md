# Terraform Provider for Synadia

This Terraform provider allows users to manage and configure [Synadia](https://synadia.com) NATS clusters and resources via the [Synadia Control Plane](https://github.com/synadia-io/control-plane-sdk-go).

## Provider Features

- Create, configure, and destroy NATS clusters.
- Manage JetStream streams, consumers, KV stores, and object stores.
- Configure leafnodes, gateways, and service imports/exports.
- Support for Synadia Cloud and Synadia Platform (private endpoint).

### Resources

| Name | Description | Status |
|------|-------------|--------|
| `synadia_cluster` | Manages a Synadia NATS cluster. | Planned |
| `synadia_organization` | Manages an organization in the Synadia control plane. | Planned |
| `synadia_project` | Manages a project under a Synadia organization. | Planned |
| `synadia_user` | Manages a user in the Synadia control plane. | Planned |
| `synadia_jwt_claim` | Manages a user JWT claim (identity and permissions). | Planned |
| `synadia_permission` | Manages permissions for accounts or users. | Planned |
| `synadia_stream` | Manages a JetStream stream. | Planned |
| `synadia_consumer` | Manages a JetStream consumer. | Planned |
| `synadia_kv_bucket` | Manages a JetStream KV bucket. | Planned |
| `synadia_object_store` | Manages a JetStream object store. | Planned |
| `synadia_cluster_gateway` | Configures a cluster gateway connection. | Planned |
| `synadia_leafnode` | Configures a leafnode connection. | Planned |
| `synadia_service_export` | Exports a service to other accounts. | Planned |
| `synadia_service_import` | Imports a service from another account. | Planned |


### Data Sources

| Name | Description | Status |
|------|-------------|--------|
| `synadia_clusters` | Fetches a list of clusters. | Planned |
| `synadia_cluster` | Fetches details for a specific cluster. | Planned |
| `synadia_organization` | Fetches details of an organization. | Planned |
| `synadia_user` | Fetches details of a user. | Planned |
| `synadia_stream` | Fetches details of a stream. | Planned |
| `synadia_kv_bucket` | Fetches details of a KV bucket. | Planned |


## Status

This provider is **under active development**. All resources and data sources listed are currently in the **planned** stage and being implemented using the official [control-plane-sdk-go](https://github.com/synadia-io/control-plane-sdk-go).

## Getting Started

Instructions for installation and usage will be published with the first tagged release.
