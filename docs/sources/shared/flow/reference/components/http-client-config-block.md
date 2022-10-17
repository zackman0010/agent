---
aliases:
- /docs/agent/shared/flow/reference/components/http-client-config-block/
headless: true
---

The `http_client_config` block configures settings used to connect to the
Kubernetes API server.

Name | Type | Description | Default | Required
---- | ---- | ----------- | ------- | --------
`bearer_token` | `secret` | Bearer token to authenticate with. | | no
`bearer_token_file` | `string` | File containing a bearer token to authenticate with. | | no
`proxy_url` | `string` | HTTP proxy to proxy requests through. | | no
`follow_redirects` | `bool` | Whether redirects returned by the server should be followed. | `true` | no
`enable_http_2` | `bool` | Whether HTTP2 is supported for requests. | `true` | no

`bearer_token`, `bearer_token_file`, `basic_auth`, `authorization`, and
`oauth2` are mutually exclusive and only one can be provided inside of a
`http_client_config` block.

The following sub-blocks are supported for `http_client_config`:

Name | Description | Required
---- | ----------- | --------
[`basic_auth`](#basic_auth-block) | Configure basic_auth for authenticating against Kubernetes. | no
[`authorization`](#authorization-block) | Configure generic authorization against Kubernetes. | no
[`oauth2`](#oauth2-block) | Configure OAuth2 for authenticating against Kubernetes. | no
[`tls_config`](#tls_config-block) | Configure TLS settings for connecting to Kubernetes. | no