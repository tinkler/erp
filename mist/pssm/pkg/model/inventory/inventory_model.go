package inventory

import "sort"

type storeInItemArr []*StoreInItem

func (a storeInItemArr) Len() int           { return len(a) }
func (a storeInItemArr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a storeInItemArr) Less(i, j int) bool { return a[i].Index < a[j].Index }

func (m *StoreIn) sortItems() {
	sort.Sort(storeInItemArr(m.StoreInItems))
	for i := range m.StoreInItems {
		m.StoreInItems[i].Index = i + 1
	}
}
