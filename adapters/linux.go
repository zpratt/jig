package adapters

import (
	"log"

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

func (l Linux) InstallPackage(packageName string) {
	packageManager := "brew"
	l.UpdatePackageList(packageManager)

	log.Printf("installing package %s", packageName)
	_, _ = l.exec.Command(packageManager, "install", packageName).CombinedOutput()
	log.Printf("installed package %s", packageName)
}

func (l Linux) UpdatePackageList(packageManager string) {
	abortIfBrewIsNotInstalled()

	log.Printf("updating package list")
	l.exec.Command(packageManager, "update")
	log.Printf("finished updating package list")
}
