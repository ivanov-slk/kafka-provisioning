package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	helmv3 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
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

		ctx.Export("name", namespace.Metadata.Elem().Name())
		// Namespace

		// Strimzi operator
		strimziKafkaOperator, err := helmv3.NewRelease(ctx, "strimzikafkaoperator", &helmv3.ReleaseArgs{
			Chart:     pulumi.String("strimzi-kafka-operator"),
			Namespace: namespace.Metadata.Name(),
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

		// Export some values for use elsewhere
		ctx.Export("name", strimziKafkaOperator.Name)
		// Strimzi operator

		return nil
	})
}
