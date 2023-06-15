
provider "google" {
  credentials = file("/home/yaseen/suse/third_assignment/suse-technical-test-95ff394a471c.json")
  project     = "suse-technical-test"
  region      = "us-central1"
}

resource "google_compute_instance" "vm_instance" {
  name         = "suse-technical-test-machine"
  machine_type = "n1-standard-1"
  zone         = "asia-south2-b"

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-10"
    }
  }

  network_interface {
    network = "default"
    
  }
}


