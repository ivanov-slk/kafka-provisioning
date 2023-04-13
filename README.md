## kubernetes-kafka-provisioning

A repository for automatic provisioning of a [Strimzi](https://strimzi.io) Kafka cluster in Kubernetes.

## Pre-requisites

The repository leverages [Pulumi](https://www.pulumi.com/) for managing the provisioning. Installing Pulumi on Linux can be done via `curl -fsSL https://get.pulumi.com | sh`.

A working Kubernetes cluster is also needed. Testing was performed on the configuration [here](https://github.com/ivanov-slk/virtualization-set-up).

## Usage

Run `pulumi up` to provision the Kafka cluster with default settings. Run `pulumi destroy` to destroy it.

## Monitoring

The Kafka cluster exposes metrics for usage by Prometheus. An example of how Prometheus itself is configured and how the associated Grafana dashboards are configured can be found [here](https://github.com/ivanov-slk/virtualization-set-up/tree/master/prometheus-stack). The example follows the official documentation for [Strimzi](https://strimzi.io/docs/operators/latest/deploying.html#assembly-metrics-prometheus-str).
