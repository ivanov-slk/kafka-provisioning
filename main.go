package main

import (
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
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

		// Strimzi operator

		return nil
	})
}
