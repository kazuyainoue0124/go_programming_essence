package foo

// for 文には4種類ある

// 1. C言語形式
// for i := 0; i < len(s); i++ {
// 	// do something
// }

// 2. C言語のwhile文の形式
// for i < 3 {
// 	// do something
// }

// 3. 無限ループ
// for {
// 	// do something
// }

// 4. map, スライスをイテレート
// map
// for k, v := range obj {
// 	// k は map のキー、　v　は値
// }

// for i, v := range arr {
// 	// i は arr の添え字、v は値
// }

// C言語と同じように break でループを中断し continue でループを継続できる
// また Go　には Labeled Break という記法があり、多重ループの内側から外側のループまで一気に抜けることができる
// finished:
// 	for i := 0; i < 100; i:++ {
// 		for j := 0; j < 100; j++ {
// 			if check(i, j) {
// 				break finished
// 			}
// 		}
// 	}

// for ループとは別に channel の読み取りを for ループで書ける
// この場合、channel がクローズされるまでループが続く
// for v := range ch {
// 	// do something
// }
