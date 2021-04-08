package ipinfo

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIPInfoDataSource_basic(t *testing.T) {
	datasourceName := "data.ipinfo.info"
	dsIP := "8.8.8.8"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIPInfoDataSourceWithIP(dsIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "ip", dsIP),
					resource.TestCheckResourceAttr(datasourceName, "hostname", "dns.google"),
				),
			},
		},
	})
}

func TestAccIPInfoDataSourceEmptyIP_basic(t *testing.T) {
	datasourceName := "data.ipinfo.info"
	dsIP := ""

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIPInfoDataSourceWithIP(dsIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "ip"),
					resource.TestCheckResourceAttrSet(datasourceName, "hostname"),
				),
			},
		},
	})
}

func TestAccIPInfoDataSourceNoIP_basic(t *testing.T) {
	datasourceName := "data.ipinfo.info"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: testAccProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccIPInfoDataSourceNoIP(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(datasourceName, "ip"),
					resource.TestCheckResourceAttrSet(datasourceName, "hostname"),
				),
			},
		},
	})
}

func testAccIPInfoDataSourceWithIP(ip string) string {
	return fmt.Sprintf(`
		data "ipinfo" "info" {
  			ip = "%s"
		}
		`, ip)
}

func testAccIPInfoDataSourceNoIP() string {
	return fmt.Sprintf(`data "ipinfo" "info" {}`)
}
