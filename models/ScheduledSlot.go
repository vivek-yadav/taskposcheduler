package models

import (
	"time"
)

type ScheduledSlot struct {
	SlotStartDT time.Time `json:"slot_start_dt"`
	SlotEndDT   time.Time `json:"slot_end_dt"`
	DockId      int       `json:"dock_id"`
	PoId        int       `json:"po_id"`
	ItemId      int       `json:"item_id"`
	Quantity    int       `json:"quantity"`
}

type ScheduledSlotList []*ScheduledSlot
