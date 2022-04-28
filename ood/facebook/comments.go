package facebook

type Comment struct {
	Id         int64
	Content    string
	TotalLikes int64
	Owner      *Member
}
