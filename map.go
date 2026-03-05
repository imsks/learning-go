package main

import "fmt"

func learnMaps() {
	// // 1. Create and use a map
	// ages := make(map[string]int)
	// ages["Sachin"] = 27

	// fmt.Println(ages)

	// // 2. Map literal
	// scores := map[string]int{}

	// for i := 0; i < 10; i++ {
	// 	scores[fmt.Sprint(i)] = i
	// }

	// fmt.Println(scores)

	// // 3. The “comma ok” idiom (check if key exists)
	// ages := map[string]int{
	// 	"Sachin": 27,
	// }

	// value, ok := ages["Sachin"]

	// if ok {
	// 	fmt.Println("Working", value)
	// } else {
	// 	fmt.Println("Not Working")
	// }

	// // 4. Delete a key
	// m := map[string]int{
	// 	"a": 1,
	// 	"b": 2,
	// 	"c": 3,
	// }

	// delete(m, "ss")
	// fmt.Println(m)

	// 8. Map of slices (or other types)
	teams := make(map[string][]string)
	teams["eng"] = []string{"Alice", "Bob"}
	teams["ind"] = []string{"Sachin", "Virat"}
	fmt.Println(teams)
}
