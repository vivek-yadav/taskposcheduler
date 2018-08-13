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


## Execution :
Sample Execution Output:
```bash
PO=pos1.csv SLOT=slots1.csv OUTPUT=output.csv ./taskposcheduler.linux
wrote 24681 bytes
Output is saved in :  output.csv
Want to see results based on:
1 dock_id
2 date
 Please enter choice no:
1
Show results of dock_id:
2
slot_start_dt,slot_end_dt,dock_id,po_id,item_id,quantity
2018-08-01T00:00:00,2018-08-01T01:00:00,2,195687,10015876,210
2018-08-01T01:00:00,2018-08-01T02:00:00,2,195687,10015876,214
2018-08-01T02:00:00,2018-08-01T03:00:00,2,195687,10015876,159
.
.
.
```
