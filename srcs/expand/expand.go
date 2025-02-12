package expand

import "os"

func Expand(s string) string {
	// var result []rune
	// var inQuotes bool
	// var escaped bool
	// var varName []rune

	s = expandTilde(s)
	//os.Expandの使い方を調べる
	s = os.Expand(s, os.Getenv)

	return s

	// for _, r := range s {
	// 	if escaped {
	// 		result = append(result, r)
	// 		escaped = false
	// 		continue
	// 	}

	// 	switch r {
	// 	case '\'':
	// 		inQuotes = !inQuotes
	// 	case '"':
	// 		inQuotes = !inQuotes
	// 	case '\\':
	// 		escaped = true
	// 	case '$':
	// 		if !inQuotes {
	// 			// $の後に続く環境変数を処理
	// 			varName = nil // 変数名のリセット
	// 			continue
	// 		}
	// 	default:
	// 		// 環境変数名が続いているとき
	// 		if r == '{' {
	// 			// 環境変数名が波括弧で囲まれている場合
	// 			varName = append(varName, r)
	// 		} else if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
	// 			varName = append(varName, r)
	// 		} else {
	// 			// 環境変数名が終わった
	// 			if len(varName) > 0 {
	// 				// $の展開
	// 				varNameStr := string(varName)
	// 				// 変数を環境変数として展開
	// 				if value, exists := os.LookupEnv(varNameStr); exists {
	// 					result = append(result, []rune(value)...)
	// 				} else {
	// 					result = append(result, []rune("$"+varNameStr)...)
	// 				}
	// 				varName = nil // 変数名リセット
	// 			}
	// 			// 普通の文字として結果に追加
	// 			result = append(result, r)
	// 		}
	// 	}
	// }

	// // 最後の環境変数名が未展開の場合
	// if len(varName) > 0 {
	// 	varNameStr := string(varName)
	// 	if value, exists := os.LookupEnv(varNameStr); exists {
	// 		result = append(result, []rune(value)...)
	// 	} else {
	// 		result = append(result, []rune("$"+varNameStr)...)
	// 	}
	// }

	// return string(result)
}
