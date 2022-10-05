package _0_基础排序算法

// HeapSort 堆排序 ==> 以未排序范围数组来构建大顶堆，然后将堆顶元素移置未排序数组末尾；重复执行，直至未排序数组元素个数为0
func HeapSort(sli []int){

	i:= len(sli) -1
	for i >= 0 {
		// 构建大顶堆
		buildHeap(sli, i)

		// 将堆顶元素调换到数组尾部
		swap(sli, 0, i)

		// 未排序范围减一
		i--
	}
}

// 构建大顶堆，i的意义在于完全二叉树的边界，构建大顶堆时得确认哪个数组下标范围是要构建的二叉树
func buildHeap(sli []int, i int){
	// 找到最后一个非叶子节点
	parent := i/2 - 1
	// 从最后一个非叶子节点开始，自下而上自右向左 对每个非叶子节点进行堆调整
	var j = parent
	for j >= 0 {
		headAdjust(sli, j, i)
		j--
	}
}

// 堆调整 ==> 判断当前节点的子节点的值有没有比自己值大的，如果有则调换位置且向下递归调整，如果没有则忽略
// i 的意义在于确定完全二叉树的边界，调整时肯定不能超过这个边界，所以i的值必须作为参数传入
func headAdjust(sli []int, j, i int){

	// 判断当前节点的子节点有没有比当前节点值大的，如果有，则调换位置
	left := 2 * j + 1
	right := 2 * j + 2

	// maxIndex 代表当前节点、左子节点、右子节点 中最大值的下标
	var maxIndex = j
	if left <= i && sli[left] > sli[j] {
		maxIndex = left
	}
	if right <= i && sli[right] > sli[maxIndex] {
		maxIndex = right
	}

	// 如果真的存在子节点比当前节点值大的情况，则将大的子节点值和父节点交换，且继续向下调整
	if maxIndex != j {
		swap(sli, maxIndex, j)
		headAdjust(sli, maxIndex, i)
	}

}

// 调换位置
func swap(sli []int, i, j int){
	sli[i], sli[j] =  sli[j] , sli[i]
}