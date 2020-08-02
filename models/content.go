package models

type Content struct {
	Title   string
	Content string
	Uid     int
}

func (c *Content) GetTitle() string {
	return c.Title
}

func (c *Content) GetContent() string {
	return c.Content
}

func (c *Content) GetUid() int {
	return c.Uid
}

func NewContent(title, content string, uid int) Content {
	return Content{
		Title:   title,
		Content: content,
		Uid:     uid,
	}
}
