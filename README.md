# platypack

A tool to generate (and eventually) manage Pulumi packages. Currently it supports generating scaffolding for packages authored in `go`, `typescript`, and `python`. It sets up all the boilderplate needed to author a pulumi package in your languages of choice, and automatically generate SDKs in `typescript`, `go`, `python`, and `csharp`

## Usage

Create a new package with `platypack new <language> <packageName>`

```sh
# generate a new package authored in Go, with the name "myPackage"
$ platypack new go myPackage

# generate a new package authored in Typescript, with the name "vpcComponent"
$ platypack new typescript vpcComponent
```

This will create a folder `./packageName` with everything you need to build and publish your component. There is a README at `./packageName/README.md` with instructions on building, installing, and running the generated package and examples.

## Installing Platypack

Requires `go >= 1.16` due to the usage of `//go:embed` directives for the templates

To install, simply run `make` from the root of this project. 