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
| Systems | Manage NATS systems | Planned |
| Teams | Manage Teams | Planned |
| Agent-tokens | Manage SCP Agent Access Tokens | Planned |
| Accounts | Manage NATS accounts | Planned |
| NATS-Users | Manage NATS users | Planned |
| JetStream | Manage JetStream assets for an Account | Planned |
| Streams | Manage JetStream streams for an Account | Planned |
| kv | Manage JetStream KV Buckets for an Account | Planned |
| Object | Manage JetStream Object Buckets for an Account | Planned |
| Mirrors | Manage JetStream mirrors for an Account | Planned |
| Consumers | Manage JetStream consumers for a Stream | Planned |
| Alerts | Manage alerts for NATS Systems and Accounts | Planned |
| App-users | Manage People with access to Control Plane | Planned |
| Authorization | Manage Authorization Policies for Control Plane users | Planned |
| signing-key-groups | Manage Account Signing Key Groups | Planned |
| NATS-User-issuance | Manage NATS user issuances | Planned |
| team-service-accounts | Manage Team Service Accounts for accessing API | Planned |
| App-service-accounts | Manage App Service Accounts for accessing API | Planned |


### Data Sources

| Name | Description | Status |
|------|-------------|--------|
| `system` | Fetches a system details. | Planned |


## Status

This provider is **under active development**. All resources and data sources listed are currently in the **planned** stage and being implemented using the official [control-plane-sdk-go](https://github.com/synadia-io/control-plane-sdk-go).

## Getting Started

Instructions for installation and usage will be published with the first tagged release.





