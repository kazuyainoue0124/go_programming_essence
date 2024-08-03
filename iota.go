package main

// import "fmt"

// iota は連続した整数を生成するための識別子
// enum のような使い方ができる
// iota は const の中で使う
// 以下のように使うと、Apple は 0、Orenge は 1、Banana は 2 になる
// const (
// 	Apple = iota
// 	Orange
// 	Banana
// )

// 以下のように使うと、Apple は 0、Orange は　2、Banana は 4 になる
// const (
// 	Apple = iota + iota
// 	Orange
// 	Banana
// )

// iota は登場するたびにインクリメントされる
// const (
// 	Apple = iota + iota
// 	Orange
// 	Banana = iota + 3
// )

// func main() {
// 	fmt.Println(Apple)  // 0
// 	fmt.Println(Orange) // 1 + 1 = 2
// 	fmt.Println(Banana) // 2 + 3 = 5
// }

// iota は型を指定して宣言することもできる
// type Fruit int
// type Animal int

// const (
// 	Apple Fruit = iota // Fruit(0)
// 	Orange
// 	Banana
// )

// const (
// 	Monkey Animal = iota // Animal(0)
// 	Elephant
// 	Pig
// )

// 型ごとにconst宣言を別にする必要がある
// type Fruit int
// type Animal int

// const (
// 	Apple  Fruit = iota // Fruit(0)
// 	Orange              // Fruit(1)
// 	Banana              // Fruit(2)
// )

// const (
// 	Monkey   Animal = iota // Animal(0)
// 	Elephant               // Animal(1)
// 	Pig                    // Animal(2)
// )

// func main() {
// 	var fruit Fruit = Apple
// 	fmt.Println(fruit) // 0

// 	fruit = Elephant  // コンパイルエラー
// 	fmt.Println(fruit)
// }
