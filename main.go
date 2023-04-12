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
			// SkipCrds: pulumi.Bool(true),
			// Values: pulumi.Map{
			// 	"controller": pulumi.Map{
			// 		"enableCustomResources": pulumi.Bool(false),
			// 		"appprotect": pulumi.Map{
			// 			"enable": pulumi.Bool(false),
			// 		},
			// 		"appprotectdos": pulumi.Map{
			// 			"enable": pulumi.Bool(false),
			// 		},
			// 		"service": pulumi.Map{
			// 			"extraLabels": appLabels,
			// 		},
			// 	},
			// },
			// Version: pulumi.String("latest"),
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

		// // Ugly... why not just parse plain YAML?
		// kafka, err := apiext.NewCustomResource(ctx, "kafka-cluster", &apiext.CustomResourceArgs{
		// 	ApiVersion: pulumi.String("kafka.strimzi.io/v1beta2"),
		// 	Kind:       pulumi.String("Kafka"),
		// 	Metadata: &metav1.ObjectMetaArgs{
		// 		Name:      pulumi.String("kafka-cluster"),
		// 		Namespace: namespace.Metadata.Name(),
		// 	},
		// 	OtherFields: k8s.UntypedArgs{
		// 		"spec": k8s.UntypedArgs{
		// 			"kafka": k8s.UntypedArgs{
		// 				"version":  "3.4.0",
		// 				"replicas": 1,
		// 				"listeners": []k8s.UntypedArgs{
		// 					{
		// 						"name": "plain",
		// 						"port": 9092,
		// 						"type": "internal",
		// 						"tls":  false,
		// 					}, {
		// 						"name": "tls",
		// 						"port": 9093,
		// 						"type": "internal",
		// 						"tls":  false,
		// 					},
		// 				},
		// 				"config": k8s.UntypedArgs{
		// 					"offsets.topic.replication.factor":         1,
		// 					"transaction.state.log.replication.factor": 1,
		// 					"transaction.state.log.min.isr":            1,
		// 					"default.replication.factor":               1,
		// 					"min.insync.replicas":                      1,
		// 					"inter.broker.protocol.version":            "3.4",
		// 				},
		// 				"storage": map[string]string{"type": "ephemeral"},
		// 			},
		// 			"zookeeper": k8s.UntypedArgs{
		// 				"replicas": 1,
		// 				"storage":  map[string]string{"type": "ephemeral"},
		// 			},
		// 			"entityOperator": k8s.UntypedArgs{
		// 				"topicOperator": map[string]string{},
		// 				"userOperator":  map[string]string{},
		// 			},
		// 		},
		// 	},
		// })

		if err != nil {
			return err
		}

		kafka_cluster := kafka.GetResource("kafka.strimzi.io/v1beta2/Kafka", "kafka-cluster", "kafka-system") //.(*apiext.CustomResource)
		// kafka_cluster.URN().OutputState
		ctx.Export("Kafka Cluster", kafka_cluster.URN().ToStringOutput())
		// ctx.Export("Kafka", kafka_cluster.Metadata.Name())
		// ctx.Export("API Version", kafka_cluster.ApiVersion)
		// ctx.Export("Custom Resource Type", kafka_cluster.Kind)
		// Kafka and Zookeeper

		return nil
	})
}
