package twoSum


func main(nums []int, target int) []int {
	for k1, _ := range nums {
		for k2 := k1 + 1; k2 < len(nums); k2++ {
			if target == nums[k1] + nums[k2] {
				return []int{k1, k2}
			}
		}
	}
	return []int{}
}
