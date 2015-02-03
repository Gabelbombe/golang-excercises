package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Template struct {
	Cluster    string
	Datacenter string
	Host       string
	Password   string
	Username   string
	VMName     string
}

func main() {

	template := Template{
		password:   os.Getenv("PACKER_USERNAME"),
		username:   os.Getenv("PACKER_PASSWORD"),
		datacenter: "The_Datacenter",
		cluster:    "The_Cluster",
		host:       "The_Host",
		vm_name:    "VM_Name",
	}

	var target = ""
	target += os.Getenv("PACKER_OUTPUT")
	target += "/output/testbox/packer-oracle-6.6-x86_64.vmx"

	// This needs to be quoted to support spaces, NOT encoded...
	vi := fmt.Sprintf("vi://%s:%s@%s/%s/host/%s/",
		template.username,
		template.password,
		template.host,
		template.datacenter,
		template.cluster)

	args := []string{
		"--acceptAllEulas",
		"--powerOn",

		"--name=GoToolTest",
		"--diskMode=thin",
		"--compress=9",

		"--datastore=PXDEVCLVMW01_DS01",
		"--ipAllocationPolicy=transientPolicy",
		"--network=PG_PXDEV01_10.158.6.0",

		"--noSSLVerify=false",
	}

	// Add target and location
	args = append(args, target)
	args = append(args, vi)

	cmd := exec.Command("ovftool", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()

}
