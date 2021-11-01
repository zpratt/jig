package adapters

import (
	"log"

	utilsexec "k8s.io/utils/exec"
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

	packageManager := packageManagers[0]

	w.UpdatePackageList(packageManager)

	if len(packageManagers) == 0 {
		log.Fatalf("no supported package managers installed")
	}

	_, _ = w.exec.Command(packageManager, "install", packageName).CombinedOutput()

	log.Printf("installing package %s", packageName)
}

func (w WindowsAdapter) UpdatePackageList(packageManager string) {
	log.Printf("updating package list")
	w.exec.Command(packageManager, "update")
	log.Printf("finished updating package list")
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
