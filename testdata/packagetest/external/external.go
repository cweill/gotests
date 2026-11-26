package external

type Item struct{ ID int }

func Existing() int { return 0 }

func Find(id int) *Item { return nil }
