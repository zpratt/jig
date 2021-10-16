package main

import (
	"github.com/zpratt/jig/adapters"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	utilsexec "k8s.io/utils/exec"
	"log"
	"os"
	"os/exec"
	"runtime"
)

type JigConfig struct {
	Tools []string `yaml:"tools"`
}

func main() {
	desiredState := desiredStateFactory()
	jigConfig := parseConfig()
	notInstalled := findMissingPackages(jigConfig)

	if len(notInstalled) > 0 {
		for _, notInstalledTool := range notInstalled {
			desiredState.InstallPackage(notInstalledTool)
		}
	} else {
		log.Printf("all tools installed")
	}
}

func findMissingPackages(jigConfig JigConfig) []string {
	var notInstalled []string

	for _, tool := range jigConfig.Tools {
		_, err := exec.LookPath(tool)

		if err != nil {
			notInstalled = append(notInstalled, tool)
		}
	}
	return notInstalled
}

func parseConfig() JigConfig {
	config := JigConfig{}

	jigConfigFile := "jig.yaml"
	_, err := os.Stat(jigConfigFile)

	if err != nil {
		log.Fatalf("%s config does not exist", jigConfigFile)
	}

	file, _ := ioutil.ReadFile(jigConfigFile)
	err = yaml.Unmarshal(file, &config)

	if err != nil {
		log.Fatalf("failed to parse %s", jigConfigFile)
	}

	return config
}

func desiredStateFactory() *DesiredState {
	execInst := utilsexec.New()

	factories := map[string]func() *DesiredState{
		"darwin": func() *DesiredState {
			return &DesiredState{PlatformAdapter: adapters.NewDarwinAdapter(execInst)}
		},
		"windows": func() *DesiredState {
			return &DesiredState{PlatformAdapter: adapters.NewWindowsAdapter(execInst)}
		},
	}

	return factories[runtime.GOOS]()
}
