---
canonical: https://grafana.com/docs/alloy/latest/reference/components/prometheus/prometheus.exporter.elasticsearch/
aliases:
  - ../prometheus.exporter.elasticsearch/ # /docs/alloy/latest/reference/components/prometheus.exporter.elasticsearch/
description: Learn about prometheus.exporter.elasticsearch
labels:
  stage: general-availability
  products:
    - oss
title: prometheus.exporter.elasticsearch
---

# `prometheus.exporter.elasticsearch`

The `prometheus.exporter.elasticsearch` component embeds the [`elasticsearch_exporter`](https://github.com/prometheus-community/elasticsearch_exporter) for the collection of metrics from ElasticSearch servers.

{{< admonition type="note" >}}
Currently, {{< param "PRODUCT_NAME" >}} can only collect metrics from a single ElasticSearch server.
However, the exporter can collect the metrics from all nodes through that server configured.
{{< /admonition >}}

We strongly recommend that you configure a separate user for {{< param "PRODUCT_NAME" >}}, and give it only the strictly mandatory security privileges necessary for monitoring your node.
Refer to the [Elasticsearch security privileges](https://github.com/prometheus-community/elasticsearch_exporter#elasticsearch-7x-security-privileges) documentation for more information.

## Usage

```alloy
prometheus.exporter.elasticsearch "<LABEL>" {
    address = "<ELASTICSEARCH_ADDRESS>"
}
```

## Arguments

You can use the following arguments with `prometheus.exporter.elasticsearch`:

| Name                   | Type       | Description                                                                                            | Default                   | Required |
| ---------------------- | ---------- | ------------------------------------------------------------------------------------------------------ | ------------------------- | -------- |
| `address`              | `string`   | HTTP API address of an Elasticsearch node.                                                             | `"http://localhost:9200"` | no       |
| `aliases`              | `bool`     | Include informational aliases metrics.                                                                 |                           | no       |
| `all`                  | `bool`     | Export stats for all nodes in the cluster. If used, this flag overrides the flag `node`.               |                           | no       |
| `ca`                   | `string`   | Path to PEM file that contains trusted Certificate Authorities for the Elasticsearch connection.       |                           | no       |
| `client_cert`          | `string`   | Path to PEM file that contains the corresponding cert for the private key to connect to Elasticsearch. |                           | no       |
| `client_private_key`   | `string`   | Path to PEM file that contains the private key for client auth when connecting to Elasticsearch.       |                           | no       |
| `cluster_settings`     | `bool`     | Export stats for cluster settings.                                                                     |                           | no       |
| `clusterinfo_interval` | `duration` | Cluster info update interval for the cluster label.                                                    | `"5m"`                    | no       |
| `data_stream`          | `bool`     | Export stats for Data Streams.                                                                         |                           | no       |
| `indices_settings`     | `bool`     | Export stats for settings of all indices of the cluster.                                               |                           | no       |
| `indices`              | `bool`     | Export stats for indices in the cluster.                                                               |                           | no       |
| `node`                 | `string`   | Node's name of which metrics should be exposed                                                         |                           | no       |
| `shards`               | `bool`     | Export stats for shards in the cluster (implies indices).                                              |                           | no       |
| `slm`                  | `bool`     | Export stats for SLM (Snapshot Lifecycle Management).                                                  |                           | no       |
| `snapshots`            | `bool`     | Export stats for the cluster snapshots.                                                                |                           | no       |
| `ssl_skip_verify`      | `bool`     | Skip SSL verification when connecting to Elasticsearch.                                                |                           | no       |
| `timeout`              | `duration` | Timeout for trying to get stats from Elasticsearch.                                                    | `"5s"`                    | no       |

## Blocks

You can use the following block with `prometheus.exporter.elasticsearch`:

| Block                      | Description                                                | Required |
| -------------------------- | ---------------------------------------------------------- | -------- |
| [`basic_auth`][basic_auth] | Configure `basic_auth` for authenticating to the endpoint. | no       |

[basic_auth]: #basic_auth

### `basic_auth`

{{< docs/shared lookup="reference/components/basic-auth-block.md" source="alloy" version="<ALLOY_VERSION>" >}}

## Exported fields

{{< docs/shared lookup="reference/components/exporter-component-exports.md" source="alloy" version="<ALLOY_VERSION>" >}}

## Component health

`prometheus.exporter.elasticsearch` is only reported as unhealthy if given an invalid configuration.
In those cases, exported fields retain their last healthy values.

## Debug information

`prometheus.exporter.elasticsearch` doesn't expose any component-specific debug information.

## Debug metrics

`prometheus.exporter.elasticsearch` doesn't expose any component-specific debug metrics.

## Example

This example uses a [`prometheus.scrape` component][scrape] to collect metrics from `prometheus.exporter.elasticsearch`:

```alloy
prometheus.exporter.elasticsearch "example" {
  address = "http://localhost:9200"
  basic_auth {
    username = "<USERNAME>"
    password = "<PASSWORD>"
  }
}

// Configure a prometheus.scrape component to collect Elasticsearch metrics.
prometheus.scrape "demo" {
  targets    = prometheus.exporter.elasticsearch.example.targets
  forward_to = [prometheus.remote_write.demo.receiver]
}

prometheus.remote_write "demo" {
  endpoint {
    url = "<PROMETHEUS_REMOTE_WRITE_URL>"

    basic_auth {
      username = "<USERNAME>"
      password = "<PASSWORD>"
    }
  }
}
```

Replace the following:

- _`<PROMETHEUS_REMOTE_WRITE_URL>`_: The URL of the Prometheus `remote_write` compatible server to send metrics to.
- _`<USERNAME>`_: The username to use for authentication to the `remote_write` API.
- _`<PASSWORD>`_: The password to use for authentication to the `remote_write` API.

[scrape]: ../prometheus.scrape/

<!-- START GENERATED COMPATIBLE COMPONENTS -->

## Compatible components

`prometheus.exporter.elasticsearch` has exports that can be consumed by the following components:

- Components that consume [Targets](../../../compatibility/#targets-consumers)

{{< admonition type="note" >}}
Connecting some components may not be sensible or components may require further configuration to make the connection work correctly.
Refer to the linked documentation for more details.
{{< /admonition >}}

<!-- END GENERATED COMPATIBLE COMPONENTS -->
