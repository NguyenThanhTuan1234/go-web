package models

type Content struct {
	Content string
	Title   string
	Uid     int
}

func (c *Content) GetContent() string {
	return c.Content
}

func (c *Content) GetTitle() string {
	return c.Title
}

func (c *Content) GetUid() int {
	return c.Uid
}

func NewContent(content, title string, uid int) Content {
	return Content{
		Content: content,
		Title:   title,
		Uid:     uid,
	}
}
