package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		// Namespace
		namespaceName := "kafka-system"

		namespace, err := corev1.NewNamespace(ctx, namespaceName, &corev1.NamespaceArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:        pulumi.String(namespaceName),
				Annotations: pulumi.StringMap{"linkerd.io/inject": pulumi.String("enabled")},
			},
		})
		if err != nil {
			return err
		}

		ctx.Export("Namespace", namespace.Metadata.Elem().Name())
		// Namespace

		// Strimzi operator
		strimziKafkaOperator, err := helmv3.NewRelease(ctx, "strimzikafkaoperator", &helmv3.ReleaseArgs{
			Chart:     pulumi.String("strimzi-kafka-operator"),
			Namespace: namespace.Metadata.Name(),
			Name:      pulumi.String("strimzi-kafka-operator"),
			RepositoryOpts: &helmv3.RepositoryOptsArgs{
				Repo: pulumi.String("https://strimzi.io/charts/"),
			},
		})
		if err != nil {
			return err
		}

		// Export some values for use elsewhere,
		ctx.Export("Operator", strimziKafkaOperator.Name)
		// Strimzi operator

		// Kafka and Zookeeper
		kafka, err := yaml.NewConfigFile(ctx, "kafka-cluster", &yaml.ConfigFileArgs{
			File: "strimzi-kafka-cluster.yaml",
		})

		if err != nil {
			return err
		}

		kafka_cluster := kafka.GetResource("kafka.strimzi.io/v1beta2/Kafka", "kafka-cluster", "kafka-system") //.(*apiext.CustomResource)
		ctx.Export("Kafka Cluster", kafka_cluster.URN().ToStringOutput())
		// Kafka and Zookeeper

		return nil
	})
}
