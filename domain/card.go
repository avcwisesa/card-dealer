package domain

type Card struct {
	Content string
}

func (c *Card) IsEmpty() bool {
	return c.Content == ""
}
