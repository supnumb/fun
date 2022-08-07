package member

type Level struct {
	value int
	name  string
}

type Contact struct {
	name      string
	telephone string
	address   string
}

type Member struct {
	name      string
	gender    int
	telephone string
	level     int
	address   string
	contacts  []*Contact
}
