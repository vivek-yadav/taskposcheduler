package main

import (
	"fmt"

	"github.com/vivek-yadav/taskposcheduler/models"
	"github.com/vivek-yadav/taskposcheduler/utils"
)

func main() {
	var err error
	po := "pos.csv"
	//fmt.Println("Enter PO Data file:")
	//fmt.Scan(&po)

	slot := "slots.csv"
	//fmt.Println("Enter Slots Data file:")
	//fmt.Scan(&slot)

	var poRecords [][]string
	poRecords, err = utils.LoadCsv(po)
	utils.CheckAndExit(err)

	var slotRecords [][]string
	slotRecords, err = utils.LoadCsv(slot)
	utils.CheckAndExit(err)

	//2006-01-02T15:04:05

	slotList := new(models.SlotList)
	poList := new(models.POList)

	err = slotList.Parse(slotRecords, true)
	utils.CheckAndExit(err)

	err = poList.Parse(poRecords, true)
	utils.CheckAndExit(err)

	fmt.Print(slotList.ToCSV())
	fmt.Print(poList.ToCSV())

}
