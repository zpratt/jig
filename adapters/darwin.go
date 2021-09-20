package adapters

import (
	"log"
	"os/exec"
)

type Darwin struct{}

func (d Darwin) InstallPackage(packageName string) {
	d.UpdatePackageList()
	log.Printf("installing package %s", packageName
	exec.Command("brew", "install", packageName)
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
