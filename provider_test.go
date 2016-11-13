package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"godaddy": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	verifyEnvExists(t, "GD_KEY")
	verifyEnvExists(t, "GD_SECRET")
	verifyEnvExists(t, "GD_DOMAIN")
}

func verifyEnvExists(t *testing.T, key string) {
	if v := os.Getenv(key); v == "" {
		t.Fatal(fmt.Sprintf("%s must be set for acceptance tests.", key))
	}
}
