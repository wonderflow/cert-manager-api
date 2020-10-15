package main

import (
	"context"
	"log"

	certmanager "github.com/wonderflow/cert-manager-api/pkg/apis/certmanager/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var scheme = runtime.NewScheme()

func init() {
	_ = certmanager.AddToScheme(scheme)
}

func main() {
	k8sClient, err := client.New(ctrl.GetConfigOrDie(), client.Options{Scheme: scheme})
	if err != nil {
		log.Fatal(err)
	}
	err = k8sClient.Create(context.Background(), &certmanager.Issuer{
		ObjectMeta: metav1.ObjectMeta{Name: "test", Namespace: "default"},
		Spec: certmanager.IssuerSpec{
			IssuerConfig: certmanager.IssuerConfig{
				SelfSigned: &certmanager.SelfSignedIssuer{},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}
