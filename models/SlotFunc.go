package models

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/vivek-yadav/taskposcheduler/utils"
)

func (t *Slot) ToCSV() string {
	return fmt.Sprintf("%d,%s,%s,%d,%d", t.DockId, t.SlotStartDT.Format(utils.DT_FORMAT), t.SlotEndDT.Format(utils.DT_FORMAT), t.Capacity, t.Remaining)
}
func (slotList *SlotList) ToCSV() string {
	var b strings.Builder
	b.WriteString("dock_id,slot_start_dt,slot_end_dt,capacity,remaining\n")
	for _, v := range *slotList {
		b.WriteString(v.ToCSV())
		b.WriteString("\n")
	}
	return b.String()
}

func (t *Slot) Parse(in []string) {
	var err error
	t.DockId, err = strconv.Atoi(in[0])
	utils.CheckAndExit(err)
	t.SlotStartDT, err = time.Parse(utils.DT_FORMAT, in[1])
	utils.CheckAndExit(err)
	t.SlotEndDT, err = time.Parse(utils.DT_FORMAT, in[2])
	utils.CheckAndExit(err)
	t.Capacity, err = strconv.Atoi(in[3])
	utils.CheckAndExit(err)
	t.Remaining = t.Capacity
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

func (slotList *SlotList) FindClosest(val int) (index int, err error) {
	if len(*slotList) == 0 {
		err = errors.New("Array is empty")
		return
	}
	dif := math.MaxInt32
	index = 0
	for i, v := range *slotList {
		if v.Remaining-val < dif && v.Remaining >= val {
			dif = v.Remaining - val
			index = i
		}
	}
	return
}
