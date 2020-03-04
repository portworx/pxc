/*
Initialization and startup files

  main.go: Called from main.go in the root of the source tree. It handles the
    initialization of the program.
  root.go: Root cobra object, which is the program itself. All main commands
    in the handler directory register with this object.
  gendoc.go: Adds a hidden command to generate markdown files
  version.go: Adds a version command to the program

*/
package component
