package newsfeed

// Item : holds rest data
type Item struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Post  string `json:"post"`
}

// List : holds list of Items
type List struct {
	Items []Item
}

// New - creates new List of Items
func New() *List {
	return &List{}
}

// Add - adds Item to List
func (l *List) Add(i Item) {
	l.Items = append(l.Items, i)
}

// GetAll - gets all items from list
func (l *List) GetAll() []Item {
	return l.Items
}
