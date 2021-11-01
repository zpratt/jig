package adapters

import (
	"log"
	"os/exec"

	utilsexec "k8s.io/utils/exec"
)

type Linux struct {
	exec utilsexec.Interface
}

func NewLinuxAdapter(exec utilsexec.Interface) Linux {
	return Linux{
		exec: exec,
	}
}

func (d Linux) InstallPackage(packageName string) {
	d.UpdatePackageList()
	log.Printf("installing package %s", packageName)
	_, _ = d.exec.Command("brew", "install", packageName).CombinedOutput()
	log.Printf("installed package %s", packageName)
}

func (Linux) UpdatePackageList() {
	abortIfBrewIsNotInstalled()

	log.Printf("updating package list")
	exec.Command("brew", "update")
	log.Printf("finished updating package list")
}
