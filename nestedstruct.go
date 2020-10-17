package main

import "fmt"

type car struct { // this is the main struct
	plate     string
	trademark carTrademark // from child struct
	owner     carOwner     //from child struct
	price     carPrice     //from child struct
	year      string
}

type carTrademark struct {
	trademark string
	model     string
}

type carOwner struct {
	id      string
	name    string
	surname string
}

type carPrice struct {
	price      int
	discounted int
}

func (p carPrice) newOffer(a int) int {
	return p.price - a
}

func main() {
	newCar := car{
		plate:     "34fab4515",
		trademark: carTrademark{trademark: "Honda", model: "Accord"},
		owner:     carOwner{id: "23524251012", name: "Ali", surname: "Can"},
		price:     carPrice{price: 200000, discounted: 175000},
	}

	fmt.Println("plate  : " + newCar.plate)
	fmt.Println("trademark  : " + newCar.trademark.trademark + " " + newCar.trademark.model)
	fmt.Println("owner   : " + newCar.owner.name + " " + newCar.owner.surname)
	fmt.Println("price(tl)  :  ", newCar.price.price)
	fmt.Println("discounted(tl) : ", newCar.price.discounted)
	fmt.Println("new offer(tl)  : ", newCar.price.newOffer(65000))
}
