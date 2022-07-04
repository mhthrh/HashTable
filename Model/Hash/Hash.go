package Hash

const (
	size int = 10
)

func New() *HashTable {
	h := HashTable{}
	for i := 0; i < size; i++ {
		h.Array[i] = &LinkedList{
			value:   "",
			address: nil,
		}
	}
	return &h
}

type HashTable struct {
	Array [size]*LinkedList
}

type LinkedList struct {
	value   string
	address *LinkedList
}

func (h *HashTable) Insert(value string) error {

	h.Array[hash(value)].address = &LinkedList{
		value:   value,
		address: h.Array[hash(value)].address,
	}
	return nil
}
func (h *HashTable) Search(value string) bool {
	start := hash(value)
	l := h.Array[start].address
	for l != nil {
		if l.value == value {
			return true
		}
		l = l.address
	}
	return false
}
func (h *HashTable) Delete(value string) bool {
	start := hash(value)
	l := h.Array[start].address
	for l != nil {
		if l.value == value {
			l.address = l.address.address
			return true
		}
		l = l.address
	}
	return false
}

func hash(value string) int {
	sum := 0
	for _, c := range value {
		sum += int(c)
	}
	return sum % size
}
