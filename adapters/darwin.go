package adapters

import (
	"log"
	"os/exec"

	utilsexec "k8s.io/utils/exec"
)

type Darwin struct {
	exec utilsexec.Interface
}

func NewDarwinAdapter(exec utilsexec.Interface) Darwin {
	return Darwin{
		exec: exec,
	}
}

func (d Darwin) InstallPackage(packageName string) {
	packageManager := "brew"
	d.UpdatePackageList(packageManager)

	log.Printf("installing package %s", packageName)
	_, _ = d.exec.Command(packageManager, "install", packageName).CombinedOutput()
	log.Printf("installed package %s", packageName)
}

func (d Darwin) UpdatePackageList(packageManager string) {
	abortIfBrewIsNotInstalled()

	log.Printf("updating package list")
	d.exec.Command(packageManager, "update")
	log.Printf("finished updating package list")
}

func abortIfBrewIsNotInstalled() {
	_, err := exec.LookPath("brew")

	if err != nil {
		log.Fatalf("brew is not installed")
	}
}
