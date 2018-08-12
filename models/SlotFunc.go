package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/vivek-yadav/taskposcheduler/utils"
)

func (t *Slot) ToCSV() string {
	format := "2006-01-02T15:04:05"
	return fmt.Sprintf("%d,%s,%s,%d", t.DockId, t.SlotStartDT.Format(format), t.SlotEndDT.Format(format), t.Capacity)
}
func (slotList *SlotList) ToCSV() string {
	var b strings.Builder
	b.WriteString("dock_id,slot_start_dt,slot_end_dt,capacity\n")
	for _, v := range *slotList {
		b.WriteString(v.ToCSV())
		b.WriteString("\n")
	}
	return b.String()
}

func (t *Slot) Parse(in []string) {
	var err error
	format := "2006-01-02T15:04:05"
	t.DockId, err = strconv.Atoi(in[0])
	utils.CheckAndExit(err)
	t.SlotStartDT, err = time.Parse(format, in[1])
	utils.CheckAndExit(err)
	t.SlotEndDT, err = time.Parse(format, in[2])
	utils.CheckAndExit(err)
	t.Capacity, err = strconv.Atoi(in[3])
	utils.CheckAndExit(err)
}

func (slotList *SlotList) Parse(records [][]string, hasHeader bool) (err error) {
	for i, v := range records {
		if hasHeader && i == 0 {
			continue
		}
		slot := new(Slot)
		slot.Parse(v)
		*slotList = append(*slotList, slot)
	}
	return
}
