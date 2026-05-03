package iterator

type Collection interface {
    createIterator() Iterator
}

//Collection Interface.