# Terraform Provider for Synadia

This Terraform provider allows users to manage and configure [Synadia](https://synadia.com) NATS clusters and resources via the [Synadia Control Plane](https://github.com/synadia-io/control-plane-sdk-go).

## Status

This provider is **under active development**. All resources and data sources listed are currently in the **planned** stage and being implemented using the official [control-plane-sdk-go](https://github.com/synadia-io/control-plane-sdk-go).

## Getting Started

Instructions for installation and usage will be published with the first tagged release.

### Resources

| Name | Description | Status |
|------|-------------|--------|
| account_signing_key_group | Manages Account Signing Key Groups | Planned |
| alert_rule | Manages alert rule| Planned |
| kv_bucket | Manages key value bucket | Planned |
| mirror | Manages mirror between streams / bucets / object stores | Planned |
| object_bucket | Manages object bucket | Planned |
| nats_user_revocation | Manages nats user revocation | Planned |
| stream | Manages jetstream stream | Planned |
| stream_export | Manages stream export entity | Planned |
| stream_import | Manages stream import entity | Planned |
| subject_export | Manages subject export entity | Planned |
| subject_import | Manages subject import entity | Planned |
| user | Manages control plane user | Planned |
| app_service_account_token | Manages application service account token | Planned |
| kv_pull_consumer | Manages key value store pull consumer | Planned |
| kv_push_consumer | Manages key value store push consumer | Planned |
| mirror_pull_consumer | Manages mirror pull consumer | Planned |
| mirror_push_consumer | Manages mirror push consumer | Planned |
| object_pull_consumer | Manages object store pull consumer | Planned |
| object_push_consumer | Manages object store push consumer| Planned |
| app_service_account | Manages application service account | Planned |
| app_user | Manages application user | Planned |
| personal_access_token | Manages personal access token | Planned |
| team | Manages team | Planned |
| pull_consumer | Manages stream pull consumer | Planned |
| push_consumer | Manages stream push consumer | Planned |
| stream_shares | Manages stream share configuration between accounts | Planned |
| subject_shares | Manages subject share configuration between accounts | Planned |
| account | Manages account | Planned |
| system_alert_rule | Manages system alert rule | Planned |
| system | Manages system configuration | Planned |
| team_service_account | Manages team service account | Planned |
| team_service_account_token | Manages service account token | Planned |

### Data Sources

| Name | Description | Status |
|------|-------------|--------|
| account | Fetches configuration | Planned |
| alert_rule | Fetches configuration | Planned |
| jetstream_placement_options | Fetches jetstream placement options | Planned |
| nats_user_revocation | Fetches nats user recovation configuration | Planned |
| account_signing_key_groups | Fetches list of account signing key groups | Planned |
| account_team_app_users | Fetches list of account team application users | Planned |
| alert_rules | Fetches list of configured alert rules | Planned |
| jetstream_assets | Fetches list of jetstream assets | Planned |
| kv_buckets | Fetches list of key value buckets | Planned |
| mirrors | Fetches list of mirrors | Planned |
| object_buckets | Fetches list of object buckets | Planned |
| stream_exports | Fetches list of stream exports | Planned |
| stream_exports_shared | Fetches list of stream exports that are shared | Planned |
| stream-imports | Fetches list of stream imports | Planned |
| streams | Fetches list of streams | Planned |
| subject_exports | Fetches list of subject exports | Planned |
| subject_imports | Fetches list of subject imports | Planned |
| users | Fetches list of users | Planned |
| app_service_account | Fetches application service account configurations | Planned |
| app_service_account_tokens | Fetches list of application service account tokens | Planned |
| app_user | Fetches application user configuration | Planned |
| app_user_roles | Fetches list of applicaiton user roles | Planned |
| policies | Fetches list of policies | Planned |
| roles | Fetches list of roles | Planned |
| nats_user_issuance | Fetches nats user issuance configuration | Planned |
| kv_bucket | Fetches key value bucket configuration | Planned |
| kv_consumers | Fetches key value store consumers | Planned |
| mirror | Fetches mirror configuration | Planned |
| mirror_consumers | Fetches list of mirror consumers | Planned |
| nats_user_bearer_jwt | Fetches nats user bearer jwt | Planned |
| nats_user_creds | Fetches nats user creds file | Planned |
| nats_user_http_gw_token | Fetches nats user http gateway token | Planned |
| nats_user | Fetches nats user configuration | Planned |
| nats_user_issuances | Fetches user issuance configuration | Planned |
| nats_user_team_app_users | Fetches list of nats application users by team | Planned |
| object_bucket | Fetches object bucket configuration | Planned |
| object_consumers | Fetches list of object bucket consumers | Planned |
| personal_access_token | Fetches personal access token | Planned |
| pull_consumer | Fetches pull consumer configuration | Planned |
| push_consumer | Fetches push consumer configuration | Planned |
| app_service_accounts | Fetches list of application service accounts | Planned |
| app_users | Fetches list of application users | Planned |
| personal_access_tokens | Fetches list of personal access tokens | Planned |
| teams | Fetches list of teams | Planned |
| account_signing_key | Fetches account signing key configuration | Planned |
| account_signing_key_group | Fetches account signing key group configuration | Planned |
| account_signing_key_group_keys | Fetches list of account signing key groups | Planned |
| stream | Fetches stream configuration | Planned |
| consumers | Fetches list of consumers | Planned |
| stream_export | Fetches stream export configuration | Planned |
| stream_shares | Fetches list if stream shares | Planned |
| stream_import | Fetches stream import configuration | Planned |
| subject_export | Fetches subject export configuration | Planned |
| subject_shares | Fetches list of subject shares | Planned |
| subject_import | Fetches subject import configuration | Planned |
| current_agent_token | Fetches current agent token | Planned |
| system | Fetches system configuration | Planned |
| system_alert_rule | Fetches system alert rule configuration | Planned |
| system_limits | Fetches system limits configuration | Planned |
| accounts | Fetches list of accounts | Planned |
| agent_tokens | Fetches list of agent tokens | Planned |
| clusters | Fetches list of clusters | Planned |
| serviers | Fetches list of servers | Planned |
| system_alert_rules | Fetches list of system alert rules | Planned |
| system_accounts | Fetches list of system accounts | Planned |
| system_servers | Fetches list of system servers | Planned |
| system_team_app_users | Fetches list of system team application users | Planned |
| team | Fetches team configuration | Planned |
| team_limits | Fetches team limits configuration | Planned |
| team_accounts | Fetches list of team accounts | Planned |
| team_app_users | Fetches list of team application users | Planned |
| team_nats_users | Fetches list of team nats users | Planned |
| team_service_accounts | Fetches list of team service accounts | Planned |
| team_systems | Fetches list of team systems | Planned |
| team_service_account | Fetches team service account configuration | Planned |
| team_service_account_token | Fetches team service account token | Planned |

