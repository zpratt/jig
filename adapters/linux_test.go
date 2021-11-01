package adapters

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/zpratt/jig/internal"
)

func TestLinux_InstallPackage(t *testing.T) {
	installedPackageManager := "brew"
	packageToInstall := gofakeit.LetterN(10)
	shellOutput := "someoutput"

	command := internal.MakeFakeCommand(shellOutput)
	fakeExec := internal.MakeFakeExec(installedPackageManager, &command)

	darwinAdapter := NewLinuxAdapter(&fakeExec)

	darwinAdapter.InstallPackage(packageToInstall)

	actualBrewOutput := command.CombinedOutputLog[0][0]
	if actualBrewOutput != installedPackageManager {
		t.Fatalf("%s not called", installedPackageManager)
	}

	brewAction := command.CombinedOutputLog[0][1]
	if brewAction != "install" {
		t.Fatalf("%s install not called", installedPackageManager)
	}

	if command.CombinedOutputCalls != 1 {
		t.Fatalf("%s install not called", installedPackageManager)
	}

}

func TestLinux_UpdatePackageList(t *testing.T) {
	t.Log("TODO: add tests")
}
