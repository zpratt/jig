package main

import (
	"github.com/zpratt/jig/adapters"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
)

type JigConfig struct {
	Tools []string `yaml:"tools"`
}

func main() {
	var notInstalled []string
	desiredState := func() DesiredState {
		if runtime.GOOS == "darwin" {
			return DesiredState{PlatformAdapter: adapters.Darwin{}}
		}
		// TODO: update factory to include a windows version
		return DesiredState{PlatformAdapter: adapters.Darwin{}}
	}()

	jigConfig := parseConfig()

	for _, tool := range jigConfig.Tools {
		_, err := exec.LookPath(tool)

		if err != nil {
			notInstalled = append(notInstalled, tool)
		}
	}

	if len(notInstalled) > 0 {
		for _, notInstalledTool := range notInstalled {
			desiredState.InstallPackage(notInstalledTool)
		}
	} else {
		log.Printf("all tools installed")
	}
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
