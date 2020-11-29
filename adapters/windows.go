package adapters

import (
	"log"
	"os/exec"
)

type Windows struct{}

func (w Windows) InstallPackage(packageName string) {
	packageManagers := findInstalledPackageManagers()

	if len(packageManagers) == 0 {
		log.Fatalf("no supported package managers installed")
	}

	log.Printf("installing package %s", packageName)
}

func (w Windows) UpdatePackageList() {

}

func findInstalledPackageManagers() []string {
	supportedPackageManagers := []string{"choco", "scoop"}
	var installedPackageManagers []string

	for _, packageManager := range supportedPackageManagers {
		_, err := exec.LookPath(packageManager)

		if err != nil {
			installedPackageManagers = append(installedPackageManagers, packageManager)
		}
	}

	return installedPackageManagers
}
