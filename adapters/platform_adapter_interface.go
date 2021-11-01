package adapters

type PlatformAdapterInterface interface {
	InstallPackage(packageName string)
	UpdatePackageList(packageManager string)
}
