package main

// import "fmt"

// func main()  {
// 	// var card string ="Ace of spades"
// 	card := "Ace of spades" //-> for new variable
// 	card = "five of diamonds"
// 	fmt.Println(card)
// }

// func main()  {
// 	// var card string ="Ace of spades"
// 	cards :=[]string{newCard(), "Ace of diamonds"}
// 	cards = append(cards, "Six of spades")

// 	for i, card :=range cards{
// 		fmt.Println(i,card)
// 	}
// 	fmt.Println(cards)

// }

// func main()  {
// 	// var card string ="Ace of spades"
// 	cards :=deck{newCard(), "Ace of diamonds"}
// 	cards = append(cards, "Six of spades")

// 	cards.print()

// }
// func newCard() string{
// 	return "Five of diamonds"
// }
// func main()  {
// 	// var card string ="Ace of spades"
// 	cards := newDeck()
// 	cards.print()
// 	hand, remainingCards := deal(cards,5)
// 	hand.print()
// 	remainingCards.print()
// }

// func main()  {
// 	// var card string ="Ace of spades"
// 	cards := newDeck()
// 	// fmt.Println([]byte(cards.toString()))
// 	fmt.Println(cards.saveToFile("deck_of_cards"))
// }

// func main()  {
// 	// var card string ="Ace of spades"
// 	cards := newDeckFromFile("deck_of_cards")
// 	cards.print()
// }

func main()  {
	// var card string ="Ace of spades"
	cards := newDeck()
	cards.shuffle()
	cards.print()
}