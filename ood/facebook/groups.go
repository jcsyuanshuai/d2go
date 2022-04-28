package facebook

type Group struct {
	Name    string
	Id      int64
	Desc    string
	Members []*Member
}
