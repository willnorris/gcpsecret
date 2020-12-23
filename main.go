// The gcpsecret tool fetches a secret from GCP Secret Manager.
//
// Authenticate by settings the GOOGLE_APPLICATION_CREDENTIALS environment
// variable.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	smpb "google.golang.org/genproto/googleapis/cloud/secretmanager/v1"
)

func main() {
	project := flag.String("project", "", "GCP Project")
	secret := flag.String("secret", "", "Secret name")
	flag.Parse()

	if *project == "" {
		log.Fatalf("must specify -project")
	}
	if *secret == "" {
		log.Fatalf("must specify -secret")
	}
	secretName := fmt.Sprintf("projects/%s/secrets/%s/versions/latest", *project, *secret)

	ctx := context.Background()
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to setup client: %v", err)
	}

	req := &smpb.AccessSecretVersionRequest{
		Name: secretName,
	}
	resp, err := client.AccessSecretVersion(ctx, req)
	if err != nil {
		log.Fatalf("failed to access secret version: %v", err)
	}

	if len(resp.Payload.Data) == 0 {
		log.Fatalf("secret %q is empty", secretName)
	}
	fmt.Print(string(resp.Payload.Data))
}
