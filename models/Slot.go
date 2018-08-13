package models

import "time"

type Slot struct {
	DockId      int       `json:"dock_id"`
	SlotStartDT time.Time `json:"slot_start_dt"`
	SlotEndDT   time.Time `json:"slot_end_dt"`
	Capacity    int       `json:"capacity"`
	Remaining   int
}

type SlotList []*Slot
