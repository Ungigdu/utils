package orderedMap

type Item struct {
	Key interface{}
	Value interface{}
}

type OrderedMap struct {
	Items map[interface{}] *Item
	list []interface{}
}

func New() *OrderedMap {
	return &OrderedMap{
		Items: make(map[interface{}]*Item),
		list:  []interface{}{},
	}
}

func (om *OrderedMap) Get(key interface{}) (int, interface{}) {
	if item, has := om.Items[key];has{
		for i,v := range om.list{
			if v == key {
				return i,item.Value
			}
		}
		return -1,nil
	}
	return -1,nil
}

func (om *OrderedMap) Set(key interface{}, value interface{}) bool {
	if index, _ := om.Get(key);index!=-1{
		om.Items[key].Value = value
		return true
	}
	om.Items[key] = &Item{
		Key:   key,
		Value: value,
	}
	om.list = append(om.list, key)
	return false
}

func (om *OrderedMap) Delete(key interface{}) bool {
	if index,_ := om.Get(key);index!=-1{
		delete(om.Items,key)
		om.list = append(om.list[:index],om.list[index+1:]...)
		return true
	}
	return false
}

func (om *OrderedMap) ToList () []interface{} {
	var l []interface{}
	for _,key := range om.list{
		if item, has := om.Items[key];has{
			l = append(l, item.Value)
		}
	}
	return l
}
