package main

import "fmt"

func main() {

	// 01. 珂珂吃香蕉。这里右N堆香蕉第i堆右nums[i]根香蕉。警卫离开了，将在H小时后回来。珂珂可以决定吃香蕉的速度K 根/小时，
	// 每个小时，珂珂会从香蕉堆里选择一堆吃掉K根，如果这堆香蕉小于K根，它将会把这一堆全部吃掉，但是在这一小时内不会吃另一堆
	// 珂珂喜欢慢慢吃，但是也想在警卫回来之前，把香蕉吃完，求珂珂在警卫回来之前，吃完香蕉的最小速度K。
	//minEatingSpeedTest()

	// 02. 求某个数的平方根
	//squareRootTest()

	// 03. 运送货物。传送带上的包裹必须在D天内运走。传送带上第i个包裹的重量为weight[i]。每天，我们都会按给出重量的顺序
	// 往传送带上装在包裹。但是装载的重量不会超过船的最大运载。 求能在D天内将货物全部运走的船的最低运载能力。
	//getDaysWithCarryTest()
	//shipWithinDaysTest()

	// 04. 分割数组的最大值。 给定要给非负整数数组和一个整数m，你需要将这个数组分成m个非空的连续子数组。
	// 设计一个算法使得这m个子数组各自和的最大值最小。
	// nums=[7,2,5,10,8], m = 2. 输出18，最好的分割方式是[7,2,5]和[10,8]
	splitArrayTest()
}

func minEatingSpeedTest() {
	nums := []int{30, 11, 23, 4, 20}
	fmt.Println(minEatingSpeed(nums, 8))

	nums2 := []int{3, 6, 7, 11}
	fmt.Println(minEatingSpeed(nums2, 8))
}

// 核心问题是得发现这道题是需要通过二分来解。最值问题一般是动态规划、二分（后续补充）
// 二分相对更简单能看出来。比如说，如果要求的值与题目中的某个条件有单调不减或者单调不增关系，那么就可以用二分解
// 比如这道题，珂珂的吃香蕉速度K与吃完消耗的时间t有单调不增的关系。即珂珂吃香蕉的速度越快，那么消耗的时间就越短
// 题目要求的就是当香蕉数目一定时，吃完这些香蕉所耗时间为H时，K的值。 即对于函数 t := f(nums, K), 求t=H时，K的值
// 这个函数很明显是个单调不增函数，也就是说K的取值，是单调不增的，也即是数据结构中称为的有序结构。这个时候就可以用二分了
// 对于二分法，需要确定题目要求的是左边界还是右边界，还是找到值就行，本题明显是左边界，因为要求满足条件的K的最小值
func minEatingSpeed(nums []int, h int) int {
	if len(nums) > h {
		return -1
	}

	var min = 0
	var max = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}

	for min <= max {
		var mid = min + (max-min)/2
		t := getTimeForSpeed(nums, mid)
		fmt.Println("min=", min, "max=", max, "mid=", mid, "t=", t)
		if t == h {
			max = mid - 1
		} else if t > h {
			min = mid + 1
		} else if t < h {
			max = mid - 1
		}
	}

	return min
}

// 以速度k吃完所有香蕉需要的时间是t小时
func getTimeForSpeed(nums []int, k int) int {

	var t int

	for i := 0; i < len(nums); i++ {
		if nums[i] <= k {
			t += 1
		} else {
			if nums[i]%k == 0 {
				t += nums[i] / k
			} else {
				t += nums[i]/k + 1
			}
		}
	}

	return t
}

func squareRootTest() {
	fmt.Println(squareRoot(100))
}

func squareRoot(num int) int {
	var min = 1
	var max = num

	for min <= max {
		var mid = min + (max-min)/2
		fmt.Println("min=", min, "max=", max, "mid=", mid)
		if square(mid) > num {
			max = mid - 1
		} else if square(mid) < num {
			min = mid + 1
		} else if square(mid) == num {
			return min
		}
	}
	return -1
}

func square(x int) int {
	return x * x
}

func shipWithinDaysTest() {
	var weight = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(shipWithinDays(weight, 5))
}

func shipWithinDays(weights []int, d int) int {

	var min = 0 // 最小值定义为货物中的最大重量，因为船的最低运载量肯定要与最重的货物重量相当
	var max = 0 // 最大值定义为货物重量的总和
	for i := 0; i < len(weights); i++ {
		if weights[i] > min {
			min = weights[i]
		}
		max += weights[i]
	}

	for min <= max {
		var mid = min + (max-min)/2
		if getDaysWithCarry(weights, mid) == d {
			max = mid - 1
		} else if getDaysWithCarry(weights, mid) > d {
			min = mid + 1
		} else if getDaysWithCarry(weights, mid) < d {
			max = mid - 1
		}
	}

	return min
}

func getDaysWithCarryTest() {
	weights := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(getDaysWithCarry(weights, 5))
}

// 此方法返回，运载能力为K时，将货物全部运走所需的天数
func getDaysWithCarry(weights []int, k int) int {

	var d int
	var t = k
	for i := 0; i < len(weights); {
		t -= weights[i]
		if t < 0 {
			d++
			t = k
		} else {
			i++
		}
	}
	d++
	return d
}

func splitArrayTest(){
	var nums = []int{7,2,5,10,8}
	fmt.Println(splitArray(nums, 2))
}

// 这道题比较难看出是使用二分来做。
// 假设求得的结果是x，x代表数组的各个分组，和最大的那个分组的和。
// 因为数组是固定的，那么能看出x越大，m应该会越小；x越小，m应该越大。这就有了一个单调关系
func splitArray(nums []int, m int) int {

	var min = 0
	var max = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > min {
			min = nums[i]
		}
		max += nums[i]
	}

	for min <= max {
		var mid = min + (max-min)/2
		if getDaysWithCarry(nums, mid) == m {
			max = mid - 1
		} else if getDaysWithCarry(nums, mid) > m {
			min = mid + 1
		} else if getDaysWithCarry(nums, mid) < m {
			max = mid - 1
		}
	}
	return min
}
