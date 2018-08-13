package models

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/vivek-yadav/taskposcheduler/utils"
)

func (t *ScheduledSlot) ToCSV() string {
	return fmt.Sprintf("%s,%s,%d,%d,%d,%d", t.SlotStartDT.Format(utils.DT_FORMAT), t.SlotEndDT.Format(utils.DT_FORMAT), t.DockId, t.PoId, t.ItemId, t.Quantity)
}
func (list *ScheduledSlotList) ToCSV() string {
	var b strings.Builder
	b.WriteString("slot_start_dt,slot_end_dt,dock_id,po_id,item_id,quantity\n")
	for _, v := range *list {
		b.WriteString(v.ToCSV())
		b.WriteString("\n")
	}
	return b.String()
}

func (list *ScheduledSlotList) GetOptimizedSchedule(poList *POList, slotList *SlotList) (err error) {
	// Sort PO List
	sort.Slice((*poList), func(i, j int) bool {
		return (*poList)[i].Quantity > (*poList)[j].Quantity
	})

	for _, po := range *poList {

		sort.Slice((*slotList), func(i, j int) bool {
			if (*slotList)[i].Remaining == (*slotList)[j].Remaining &&
				(*slotList)[j].Remaining == 0 {
				return (*slotList)[i].SlotStartDT.Sub((*slotList)[j].SlotStartDT) < 0
			} else if (*slotList)[j].Remaining == 0 {
				return true
			} else if (*slotList)[i].Remaining == 0 {
				return false
			}

			if (*slotList)[i].SlotStartDT.Sub((*slotList)[j].SlotStartDT) == 0 {
				return (*slotList)[i].Remaining > (*slotList)[j].Remaining
			} else {
				return (*slotList)[i].SlotStartDT.Sub((*slotList)[j].SlotStartDT) < 0
			}
		})

		dockMap := map[int]SlotList{}
		for _, v := range *slotList {
			if _, ok := dockMap[v.DockId]; !ok {
				dockMap[v.DockId] = SlotList{}
			}
			dockMap[v.DockId] = append(dockMap[v.DockId], v)
		}

		var index int
		index, err = slotList.FindClosest(po.Quantity)
		if err != nil {
			break
		}
		if (*slotList)[index].Remaining == 0 {
			break
		}
		sslot := new(ScheduledSlot)
		if (*slotList)[index].Remaining >= po.Quantity {
			(*slotList)[index].Remaining -= po.Quantity

			sslot.DockId = (*slotList)[index].DockId
			sslot.SlotStartDT = (*slotList)[index].SlotStartDT
			sslot.SlotEndDT = (*slotList)[index].SlotEndDT
			sslot.PoId = po.PoId
			sslot.ItemId = po.ItemId
			sslot.Quantity = po.Quantity
			*list = append(*list, sslot)
		} else if (*slotList)[index].Remaining > 0 {

			poLeft := po.Quantity
			allSlotsForADock := dockMap[(*slotList)[index].DockId]
			for _, s := range allSlotsForADock {
				if s.Remaining >= poLeft {
					s.Remaining -= poLeft
					sslot = new(ScheduledSlot)
					sslot.DockId = s.DockId
					sslot.SlotStartDT = s.SlotStartDT
					sslot.SlotEndDT = s.SlotEndDT
					sslot.PoId = po.PoId
					sslot.ItemId = po.ItemId
					sslot.Quantity = poLeft
					*list = append(*list, sslot)
					break
				} else if s.Remaining > 0 {
					sslot = new(ScheduledSlot)
					sslot.DockId = s.DockId
					sslot.SlotStartDT = s.SlotStartDT
					sslot.SlotEndDT = s.SlotEndDT
					sslot.PoId = po.PoId
					sslot.ItemId = po.ItemId
					sslot.Quantity = s.Remaining
					poLeft -= s.Remaining
					s.Remaining = 0
					*list = append(*list, sslot)
				} else {
					break
				}
			}
		} else {
			break
		}

	}
	return
}

func (list *ScheduledSlotList) SaveTo(fileName string) (err error) {
	var file *os.File
	file, err = os.Create(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	var n int
	n, err = writer.WriteString(list.ToCSV())
	fmt.Printf("wrote %d bytes\n", n)
	return
}

func (list *ScheduledSlotList) ShowResultsForDockId(dockId int) (err error) {
	selectedList := new(ScheduledSlotList)
	for _, v := range *list {
		if v.DockId == dockId {
			*selectedList = append(*selectedList, v)
		}
	}
	if len(*selectedList) == 0 {
		return errors.New("dock_id not found")
	}
	fmt.Println(selectedList.ToCSV())
	return
}

func (list *ScheduledSlotList) ShowResultsForDate(dt time.Time) (err error) {
	selectedList := new(ScheduledSlotList)
	for _, v := range *list {
		if v.SlotStartDT.Format(utils.FIND_DT_FORMAT) == dt.Format(utils.FIND_DT_FORMAT) {
			*selectedList = append(*selectedList, v)
		}
	}
	if len(*selectedList) == 0 {
		return errors.New("Date not found")
	}
	fmt.Println(selectedList.ToCSV())
	return
}
