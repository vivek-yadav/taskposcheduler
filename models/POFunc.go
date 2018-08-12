package models

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/vivek-yadav/taskposcheduler/utils"
)

func (t *PO) ToCSV() string {
	return fmt.Sprintf("%d,%d,%d", t.PoId, t.ItemId, t.Quantity)
}
func (poList *POList) ToCSV() string {
	var b strings.Builder
	b.WriteString("po_id,item_id,quantity\n")
	for _, v := range *poList {
		b.WriteString(v.ToCSV())
		b.WriteString("\n")
	}
	return b.String()
}

func (t *PO) Parse(in []string) {
	var err error
	t.PoId, err = strconv.Atoi(in[0])
	utils.CheckAndExit(err)
	t.ItemId, err = strconv.Atoi(in[1])
	utils.CheckAndExit(err)
	t.Quantity, err = strconv.Atoi(in[2])
	utils.CheckAndExit(err)
}

func (poList *POList) Parse(records [][]string, hasHeader bool) (err error) {
	for i, v := range records {
		if hasHeader && i == 0 {
			continue
		}
		po := new(PO)
		po.Parse(v)
		*poList = append(*poList, po)
	}
	return
}
