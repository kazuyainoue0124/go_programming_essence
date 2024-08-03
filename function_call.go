// Go では型名を変数名や関数名の後に書く

package foo

// import (
// 	"fmt"
// 	"log"
// )

// Go では多値を返すことができる
// 以下の例ではユーザーが見つかった場合はerrorをnilにして返す
// 見つからなかった場合はerrorを返す
// func FindUser(name string) (*User, error) {
// 	user, err := findUserFromList(users, name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return user, nil
// }

// func main() {
// 	user, err := FindUser("Bob")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(user.Name)
// }

// Go ではエラーが発生しうる関数を呼び出した際はまずerrをチェックすることが推奨される
// エラーがnilでない場合はエラー処理を行う
