## kubernetes-kafka-provisioning

A repository for automatic provisioning of a [Strimzi](https://strimzi.io) Kafka cluster in Kubernetes.

## Pre-requisites

The repository leverages [Pulumi](https://www.pulumi.com/) for managing the provisioning. Installing Pulumi on Linux can be done via `curl -fsSL https://get.pulumi.com | sh`.

A working Kubernetes cluster with Prometheus CRDs is also needed. Testing was performed on the configuration [here](https://github.com/ivanov-slk/virtualization-set-up).

## Usage

Run `pulumi up` to provision the Kafka cluster with default settings. Run `pulumi destroy` to destroy it.

## Monitoring

The Kafka cluster exposes metrics for usage by Prometheus. There are custom configurations for monitoring [Strimzi](https://github.com/strimzi/strimzi-kafka-operator) Kafka clusters. Strimzi's Grafana dashboards are included and Prometheus is configured to supply metrics for them.
