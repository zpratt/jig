package main

import "github.com/zpratt/jig/adapters"

type DesiredState struct {
	PlatformAdapter adapters.PlatformAdapterInterface
}

func (ds DesiredState) InstallPackage(packageName string) {
	ds.PlatformAdapter.InstallPackage(packageName)
}
