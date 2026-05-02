package decorator


// This is the base object that can be wrapped with decorators. It provides the default behavior.

type VeggieMania struct {
}

func (p *VeggieMania) getPrice() int {
    return 15
}