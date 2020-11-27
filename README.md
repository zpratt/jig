# jig
Tired of making sure that everyone has the correct tools installed on their machine. `jig` will read the jig.yaml file and check your `$PATH` environment variable for installed tools. 

## How To Use

Note: Currently, there are no releases, so you have to build this from scratch.

1. clone the repo
1. `make build`
1. `./jig`

You can update the jig.yaml file to include a list of tools that you expect to be installed.

## Roadmap

* [x] - check the `$PATH` for installed tools. Print a list of required programs that are not installed.
* [ ] - for programs that are not installed, attempt to install them using a package manager (ex: apt, brew, choco, scoop).
* [ ] - better document the schema of the jig file.
* [ ] - allow for project specific jig files instead of looking for one in the current directory
* [ ] - check for required configuration files (which could be expressed in the jig.yaml file)
* [ ] - check the version of tools and see if an update is available.

## Contributing

Not sure yet. I need to add tests and a roadmap of some kind. Feel free to submit issues if you have ideas.