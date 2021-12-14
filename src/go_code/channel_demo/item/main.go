package main

type container struct {
	items []string
}

func (c *container) Len() int {
	return len(c.items)
}

func (c *container) Iter() <-chan string {
	ch := make(chan string)

	go func() {
		for i := 0; i < c.Len(); i++ {

			ch <- c.items[i]
		}
		close(ch)
	}()
	return ch
}

func main() {
	c := container{items: []string{"韩立", "南宫婉"}}

	for s := range c.Iter() {
		println(s)
	}

}
