package aehcl

import (
	"os"

	"cloud.google.com/go/compute/metadata"
)

func fetchIDToken() (string, error) {
	return metadata.Get("/instance/service-accounts/default/identity?audience=" + os.Getenv("GOOGLE_CLOUD_PROJECT"))
}
