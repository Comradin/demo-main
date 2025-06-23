# Overview

The purpose of this repository is to create a small demo tool with a sub package
in the pkg/foo directory. While the package works perfectly fine, this will change
if you have a look into the companion repository [demo-main-import](https://github.com/Comradin/demo-main-import).

# Unimportability

Even though the code in the `pkg/foo` directory could in theory be imported, as it is
done by the main package, the module name in the `go.mod` file prevents the import.

Go packages that need to be imported, must have a name, that allows the Go tool to
clone the package. But in this case, the module is named `demo-main` which isn't
resolved into a version control system, that the Go tool can download the code from.
