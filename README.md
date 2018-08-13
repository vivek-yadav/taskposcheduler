# taskposcheduler
optimise task po scheduler

# Run Binary
### Linux
PO=pos1.csv SLOT=slots1.csv OUTPUT=output.csv ./taskposcheduler.linux

### Windows
PO=pos1.csv SLOT=slots1.csv OUTPUT=output.csv taskposcheduler.windows.exe

### Mac OS
PO=pos1.csv SLOT=slots1.csv OUTPUT=output.csv taskposcheduler.darwin

# Run Code

This is a GoLang project.
Please ensure you have installed golang
and setup environment variables : GOPATH, GOROOT, PATH
for this setup you can refer GoLang installation doc, on golang website.

## To fetch the project code 
go get -u github.com/vivek-yadav/taskposcheduler

## To install dependencies 
- change working dir to where project code is located.
- Use this : [https://github.com/golang/dep] to install dep
OR 
just execute:
go get -u github.com/golang/dep/cmd/dep
- cd $GOPATH/src/github.com/vivek-yadav/taskposcheduler
- dep ensure

## Execute from code
- ensure you are in project code location
- cd $GOPATH/src/github.com/vivek-yadav/taskposcheduler 
- PO=pos1.csv SLOT=slots1.csv OUTPUT=output.csv go run main.go
