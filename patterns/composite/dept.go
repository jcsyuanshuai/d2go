package composite

type Countable interface {
	Count() int
}

type Employee struct {
	Name string
}

func (e Employee) Count() int {
	return 1
}

type Dept struct {
	Name  string
	Items []Countable
}

func (d Dept) Count() int {
	sum := 0
	for _, item := range d.Items {
		sum += item.Count()
	}
	return sum
}

func (d *Dept) AddSub(item Countable) {
	d.Items = append(d.Items, item)
}

func InitDept() Countable {
	root := &Dept{Name: "root"}

	for i := 0; i < 10; i++ {
		root.AddSub(&Employee{})
		root.AddSub(&Dept{
			Name: "sub",
			Items: []Countable{
				&Employee{},
			},
		})
	}

	return root
}
