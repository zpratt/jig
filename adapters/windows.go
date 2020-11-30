package adapters

import (
	utilsexec "k8s.io/utils/exec"
	"log"
)

type WindowsAdapter struct {
	exec utilsexec.Interface
}

func NewWindowsAdapter(exec utilsexec.Interface) WindowsAdapter {
	windows := WindowsAdapter{
		exec: exec,
	}

	return windows
}

func (w WindowsAdapter) InstallPackage(packageName string) {
	packageManagers := w.findInstalledPackageManagers()

	if len(packageManagers) == 0 {
		log.Fatalf("no supported package managers installed")
	}

	w.exec.Command(packageManagers[0], "install", packageName).CombinedOutput()

	log.Printf("installing package %s", packageName)
}

func (w WindowsAdapter) UpdatePackageList() {

}

func (w WindowsAdapter) findInstalledPackageManagers() []string {
	supportedPackageManagers := []string{"choco", "scoop"}
	var installedPackageManagers []string

	for _, packageManager := range supportedPackageManagers {
		_, err := w.exec.LookPath(packageManager)

		if err == nil {
			installedPackageManagers = append(installedPackageManagers, packageManager)
		}
	}

	return installedPackageManagers
}
