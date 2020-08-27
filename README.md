# cf-cli-version-plugin
A cf CLI plugin to tell you what version of the cf CLI you should be using based on your cf API

# Requirements
You need to have added the [cf-cl](github.com/cloudfoundry/cli) to your GOPATH, this works on CLI v6 or v7 

To set up a $GOPATH for the first time see [the golang docs](https://golang.org/doc/gopath_code.html)

To do this for the first time on a machine with go 1.14+ run 
  `mkdir -p $GOPATH/src/code.cloudfoundry.org` 
  `cd $GOPATH/src/code.cloudfoundry.org`
  `git clone git@github.com:cloudfoundry/cli.git`
  `cd cli`
  `make clean build`

# Installation
To install this plugin, clone the repo and run make clean build install

# Usage
cf detailed-version
