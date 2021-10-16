package adapters

import (
	utilsexec "k8s.io/utils/exec"
	"log"
	"os/exec"
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
	d.UpdatePackageList()
	log.Printf("installing package %s", packageName)
	_, _ = d.exec.Command("brew", "install", packageName).CombinedOutput()
	log.Printf("installed package %s", packageName)
}

func (d Darwin) UpdatePackageList() {
	abortIfBrewIsNotInstalled()

	log.Printf("updating package list")
	exec.Command("brew", "update")
	log.Printf("finished updating package list")
}

func abortIfBrewIsNotInstalled() {
	_, err := exec.LookPath("brew")

	if err != nil {
		log.Fatalf("brew is not installed")
	}
}
