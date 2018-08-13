package models

import (
	"fmt"
	"sort"
	"strings"

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

	fmt.Println(poList.ToCSV())

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

		if (*slotList)[0].Remaining == 0 {
			break
		}
		sslot := new(ScheduledSlot)
		if (*slotList)[0].Remaining >= po.Quantity {
			(*slotList)[0].Remaining -= po.Quantity

			sslot.DockId = (*slotList)[0].DockId
			sslot.SlotStartDT = (*slotList)[0].SlotStartDT
			sslot.SlotEndDT = (*slotList)[0].SlotEndDT
			sslot.PoId = po.PoId
			sslot.ItemId = po.ItemId
			sslot.Quantity = po.Quantity
			*list = append(*list, sslot)
		} else if (*slotList)[0].Remaining > 0 {

			poLeft := po.Quantity
			allSlotsForADock := dockMap[(*slotList)[0].DockId]
			for _, s := range allSlotsForADock {
				if s.Remaining >= poLeft {
					s.Remaining -= poLeft
					sslot.DockId = s.DockId
					sslot.SlotStartDT = s.SlotStartDT
					sslot.SlotEndDT = s.SlotEndDT
					sslot.PoId = po.PoId
					sslot.ItemId = po.ItemId
					sslot.Quantity = poLeft
					*list = append(*list, sslot)
				} else if s.Remaining > 0 {
					poLeft -= s.Remaining
					s.Remaining = 0
					sslot.DockId = s.DockId
					sslot.SlotStartDT = s.SlotStartDT
					sslot.SlotEndDT = s.SlotEndDT
					sslot.PoId = po.PoId
					sslot.ItemId = po.ItemId
					sslot.Quantity = s.Remaining
					*list = append(*list, sslot)
				} else {
					break
				}
			}
		} else {
			break
		}

	}

	fmt.Println(list.ToCSV())
	return
}
