// 配列は固定長、スライスは可変長の配列
// 配列・スライス共に0から始まる
// 配列の長さを越えた位置を参照すると panic が発生する

package foo

// 配列の宣言（長さは変更不可）
// var a [4]int
// a[0] = 1

// スライスの宣言（長さは変更可能）
// 宣言だけの場合は nil が入るため、そのままアクセスすると panic が発生する
// var a []int

// スライスを作成するには make を使う（変数宣言を兼ねる）
// a := make([]int, 3)
// a[0] = 1
// a[1] = 2
// a[2] = 3

// スライスの初期化は初期値を代入することでも可能
// a := []int{1, 2, 3}

// 多次元の配列やスライスも扱える
// 固定長２次元の配列
// arr1 := [2][3]int {
// 	{1, 2, 3},
// 	{4, 5, 6},
// }

// 2次元のスライス
// arr2 := [][]int {
// 	{1, 2, 3},
// 	{4, 5, 6},
// }

// スライスを伸長するには append を使う
// a := []int{1, 2, 3}
// fmt.Println(a) // [1 2 3]
// a = append(a, 4, 5, 6)
// fmt.Println(a) // [1 2 3 4 5 6]
// fmt.Printf("a の長さは %d\n", len(a)) // a の長さは 6

// 配列とスライスには長さ len とは別に容量 cap がある
// 配列の場合は len と cap が同じ
// スライスの場合は append するたびにメモリを再確保しなくて良いように cap は常に len 以上
// make は以下のように len と cap を指定できる
// append を繰り返すと len が cap より大きくなるたびにメモリの再確保が発生し、パフォーマンスが低下する
// 事前に長さが判明しているのであれば、len と cap を指定して make でスライスを作成すると良い
// a := make([]int, 0, 100)
// for i := 0; i < 100; i++ {
// 	a = append(a, i)
// }

// スライスは添え字で範囲を指定して部分参照ができる
// fmt.Println(a[2:5]) // [2, 3, 4]　が表示される

// スライスから要素を削除する方法は3つ

// 1. 新しく同じ型のスライスを用意して詰め直す
// a2 := make([]int, 0, len(a)/2)
// for i := 0; i < len(a); i++ {
// 	if i % 2 == 0 {
// 		// 奇数は削除
// 		a2 = append(a2, a[i])
// 	}
// }
// a = a2

// 2. appendを使う
// n := 50
// a = append(a[:n], a[n+1:]...) // 添え字 50 の要素を削除

// 3. copyを使う
// n := 50
// a = a[:n+copy(a[n:], a[n+1:])] // 添え字 50 の要素を削除
