// +build !local

package gcp

import (
	"os"

	"cloud.google.com/go/compute/metadata"
)

// FetchToken returns ...
func FetchToken() (string, error) {
	return metadata.Get("/instance/service-accounts/default/identity?audience=" + os.Getenv("GOOGLE_CLOUD_PROJECT"))
}
