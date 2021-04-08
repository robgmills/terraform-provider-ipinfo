package ipinfo

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var (
	testAccProviders         map[string]*schema.Provider
	testAccProvider          *schema.Provider
	testAccProviderFactories map[string]func() (*schema.Provider, error)
)

func init() {
	testAccProvider = Provider()
	testAccProviders = map[string]*schema.Provider{
		"ipinfo": testAccProvider,
	}
	testAccProviderFactories = map[string]func() (*schema.Provider, error){
		"ipinfo": func() (*schema.Provider, error) {
			return testAccProvider, nil
		},
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProviderImpl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("IPINFO_TOKEN"); v == "" {
		t.Log(v)
		t.Fatal("IPINFO_TOKEN must be set for acceptance tests")
	}
}
