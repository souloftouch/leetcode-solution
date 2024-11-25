package tools

func display(nums []int) {
	print("{")
	for i, x := range nums {
		print(x)
		if i < len(nums)-1 {
			print(", ")
		}
	}
	print("}")
	println()
}
