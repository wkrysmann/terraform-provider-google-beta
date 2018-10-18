// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccComputeRegionAutoscaler_regionAutoscalerBasicExample(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckComputeRegionAutoscalerDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccComputeRegionAutoscaler_regionAutoscalerBasicExample(acctest.RandString(10)),
			},
			{
				ResourceName:      "google_compute_region_autoscaler.foobar",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccComputeRegionAutoscaler_regionAutoscalerBasicExample(val string) string {
	return fmt.Sprintf(`
resource "google_compute_region_autoscaler" "foobar" {
  name   = "my-region-autoscaler-%s"
  region = "us-central1"
  target = "${google_compute_region_instance_group_manager.foobar.self_link}"

  autoscaling_policy = {
    max_replicas    = 5
    min_replicas    = 1
    cooldown_period = 60

    cpu_utilization {
      target = 0.5
    }
  }
}

resource "google_compute_instance_template" "foobar" {
  name           = "my-instance-template-%s"
  machine_type   = "n1-standard-1"
  can_ip_forward = false

  tags = ["foo", "bar"]

  disk {
    source_image = "${data.google_compute_image.debian_9.self_link}"
  }

  network_interface {
    network = "default"
  }

  metadata {
    foo = "bar"
  }

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }
}

resource "google_compute_target_pool" "foobar" {
  name = "my-target-pool-%s"
}

resource "google_compute_region_instance_group_manager" "foobar" {
  name   = "my-region-igm-%s"
  region = "us-central1"

  instance_template  = "${google_compute_instance_template.foobar.self_link}"
  target_pools       = ["${google_compute_target_pool.foobar.self_link}"]
  base_instance_name = "foobar"
}

data "google_compute_image" "debian_9" {
	family  = "debian-9"
	project = "debian-cloud"
}
`, val, val, val, val,
	)
}
