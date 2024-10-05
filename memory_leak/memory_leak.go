package memory_leak

var (
	head *link
	tail *link
)

type link struct {
	next *link
}

func setup() {
	head = &link{}
	tail = head
}

func leak() {
	next := &link{}
	tail.next = next
	tail = tail.next
}
