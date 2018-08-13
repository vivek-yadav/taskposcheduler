package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/vivek-yadav/taskposcheduler/models"
	"github.com/vivek-yadav/taskposcheduler/utils"
)

var po string
var slot string
var output string

func main() {
	var err error
	getEnv()

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

	scheduledSlotList := new(models.ScheduledSlotList)
	scheduledSlotList.GetOptimizedSchedule(poList, slotList)

	err = scheduledSlotList.SaveTo(output)
	utils.CheckAndExit(err)
	fmt.Println("Output is saved in : ", output)

	var choice int
	fmt.Println("Want to see results based on: \n1. dock_id\n2. date\n Please enter choice no: ")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		var dockId int
		fmt.Println("Show results of dock_id: ")
		fmt.Scan(&dockId)
		err = scheduledSlotList.ShowResultsForDockId(dockId)
		utils.CheckAndExit(err)
	case 2:
		var dtStr string
		fmt.Println("Show results for date (YYYY-MM-DD): ")
		fmt.Scan(&dtStr)
		var dt time.Time
		dt, err = time.Parse(utils.FIND_DT_FORMAT, dtStr)
		utils.CheckAndExit(err)
		err = scheduledSlotList.ShowResultsForDate(dt)
		utils.CheckAndExit(err)
	default:
		fmt.Println("Invalid Choice")
	}

}

func getEnv() {
	var ok bool
	po, ok = os.LookupEnv("PO")
	if !ok {
		log.Fatal("[ERROR] Please enter PO=pos.csv as environment variable")
	}
	slot, ok = os.LookupEnv("SLOT")
	if !ok {
		log.Fatal("[ERROR] Please enter SLOT=slots.csv as environment variable")
	}
	output, ok = os.LookupEnv("OUTPUT")
	if !ok {
		log.Fatal("[ERROR] Please enter OUTPUT=output.csv as environment variable")
	}
}
