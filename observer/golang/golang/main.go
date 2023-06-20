package main

func main() {
	shirtItem := newItem("Nick Shirt")
	observerFirst := &Customer{
		id: "abc@gmail.com",
	}

	observerSecond := &Customer{
		id: "def@gmail.com",
	}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	shirtItem.updateAvailability()
}
