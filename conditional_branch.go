package foo

// Go では if 文の括弧は不要
// if x == 2 && y == 3 {
// 	doSomething()
// }

// セミコロンを用いて次のようにも書ける
// if user, err := userName(); err == nil {
// 	fmt.Println(user.Name)
// } else {
// 	log.Println(err)
// }

// Go の switch 文では break を書く必要がない
// switch i {
// case 1:
// 	fmt.Println("one")
// case 2:
// 	fmt.Println("two")
// case 3, 4:
// 	fmt.Println("three or four")
// default:
// 	fmt.Println("other")
// }

// 明示的に次の case に降下したい場合は fallthrough を使う
// switch i {
// case 1:
// 	fmt.Println("one")
// case 2:
// 	fmt.Print("two or")
// 	fallthrough // two or three or four が表示される
// case 3, 4:
// 	fmt.Println("three or four")
// default:
// 	fmt.Println("other")
// }

// switch case には固定値ではなく式も書ける
// switch {
// case i > hit:
// 	fmt.Println("bigger than", hit)
// case i < hit:
// 	fmt.Println("less than", hit)
// case i == hit:
// 	fmt.Println("equal to", hit)
// }
