package adapters

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/zpratt/jig/internal"
	"log"
	"testing"
)

func TestDarwin_InstallPackage(t *testing.T) {
	installedPackageManager := "brew"
	packageToInstall := gofakeit.LetterN(10)
	shellOutput := "someoutput"

	command := internal.MakeFakeCommand(shellOutput)
	fakeExec := internal.MakeFakeExec(installedPackageManager, &command)

	darwinAdapter := NewDarwinAdapter(&fakeExec)

	darwinAdapter.InstallPackage(packageToInstall)

	actualBrewOutput := command.CombinedOutputLog[0][0]
	if actualBrewOutput != installedPackageManager {
		log.Fatalf("%s not called", installedPackageManager)
	}

	brewAction := command.CombinedOutputLog[0][1]
	if brewAction != "install" {
		log.Fatalf("%s install not called", installedPackageManager)
	}

	if command.CombinedOutputCalls != 1 {
		log.Fatalf("%s install not called", installedPackageManager)
	}

}

func TestDarwin_UpdatePackageList(t *testing.T) {
}
