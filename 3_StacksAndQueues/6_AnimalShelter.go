package main

import "fmt"

/*
Animal Shelter
An animal shelter, which holds only dogs and cats, operates on a strictly "first in, first out" basis.
People must adopt either the "oldest" (based on arrival time) of all animals at the shelter, or they can 
select whether they would prefer a dog or a cat (and will receive the oldest animal of that type). 
They cannot select which specific animal they would like. 
Create the data structures to maintain this system and implement operations such as enqueue, 
dequeueAny, dequeueDog, and dequeueCat.
*/
type AnimalQueue struct {
	dogs []Animal
	cats []Animal
	order int
}


func (q *AnimalQueue) Enqueue(a Animal) {
	q.order++
	a.setOrder(q.order)	
	if _, ok := a.(*Cat); ok {
		q.cats = append(q.cats, a)	
		return
	}
	q.dogs = append(q.dogs, a)
}

func (q *AnimalQueue) Dequeue() Animal {
	if len(q.cats) == 0 && len(q.dogs) == 0 {
		return nil
	}
	if len(q.dogs) == 0 {
		return q.DequeueCat()
	}
	if len(q.cats) == 0 {
		return q.DequeueDog()
	}
	
	cat, dog := q.cats[0], q.dogs[0]
	if cat.isOlderThan(dog) {
		return q.DequeueCat()
	} 
	return q.DequeueDog()
}

func (q *AnimalQueue) DequeueDog() Animal {
	top := q.dogs[0]
	q.dogs = q.dogs[1:]
	return top
}

func (q *AnimalQueue) DequeueCat() Animal {
	top := q.cats[0]
	q.cats = q.cats[1:]
	return top
}
type Cat struct {
	*AbstractAnimal
}

func NewCat(name string) *Cat {
	return &Cat{AbstractAnimal: &AbstractAnimal{name: name}}
}

type Dog struct {
	*AbstractAnimal
}

func NewDog(name string) *Dog {
	return &Dog{AbstractAnimal: &AbstractAnimal{name: name}}
}

type Animal interface {
	getName() string
	getOrder() int
	setOrder(order int)
	isOlderThan(that Animal) bool

}

type AbstractAnimal struct {
	order int
	name string
}

func (this *AbstractAnimal) setOrder(order int) {
	this.order = order
}

func (this *AbstractAnimal) getOrder() int {
	return this.order
}


func (this *AbstractAnimal) getName() string {
	return this.name
}

func (this AbstractAnimal) isOlderThan(that Animal) bool {
	return this.order < that.getOrder()
}

func main() {
	q := &AnimalQueue{}
	fmt.Println(q.Dequeue())
	cat, dog := NewCat("cat"), NewDog("dog")
	q.Enqueue(cat)
	fmt.Println(q.Dequeue().getName())

	q.Enqueue(dog)
	fmt.Println(q.Dequeue().getName())

	q.Enqueue(dog)
	q.Enqueue(cat)
	fmt.Println(q.Dequeue().getName())
}