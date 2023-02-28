package main

import (
	"fmt"
	"sort"
)

func main() {
	//ages:=make(map[string]int)
	ages := map[string]int{
		"alice":   34,
		"charlir": 31,
	}
	ages["alice"] = 32
	fmt.Println(ages["alice"])
	//使用内置的delete可以删除元素
	delete(ages, "alice")
	fmt.Println(ages)
	ages["bob"] += 35 //即使这些元素不在map中也没有关系
	fmt.Println(ages)
	for name, value := range ages {
		fmt.Printf("%s\t%d\n", name, value)
	}
	fmt.Println()
	//Map的迭代顺序是不确定的
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}
