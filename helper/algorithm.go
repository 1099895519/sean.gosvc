package helper

import "math"

// @title InsertionSort
// @des   插入排序,下标从1开始
// 插入排序的思想：拿一个长度大于2的数组，从数组的第二位开始
// 取出第二位作为一个临时值（此时的位置为空），依次与前一位做比较，如果比前一位小，则前一位后挪一个位置（前一位的位置空出），
// 最后这个零时的值放到空位置上。循环对比直到最后一个值
// 算法时间复杂度的计算
// T(n) = an.n+bn+c
func InsertionSort(A map[int]int) map[int]int {
	ALen :=len(A)                       //k1
	if ALen < 2 {                       //k2
		return A
	}
	var j int                           //k3
	var i int                           //k4
	for j = 2;j<=ALen;j++ {             //n.c1
		key := A[j]                      //(n-1).c2
		i = j-1                          //(n-1).c3
		for {
			if i>0 && A[i]>key{          //（t2+t1+..tn)
				A[i+1] = A[i]			 //（t2+t1+..tn-1)
				i = i-1                  //（t2+t1+..tn-1)
			}else{
				break
			}
		}
		A[i+1] = key                    //(n-1).c4
	}
	return A
}

// @desc 桶排序
func BucketSort(A map[int]int) map[int]int {



	return A
}

//@desc 计数排序
// 先将数组至为0
// T(n) = an
func CountSort(A map[int]int,max int) map[int]int{
	//将零时数组置为0
	var numTemp = make(map[int]int)
	for i:= 0;i<=max;i++{
		numTemp[i] = 0
	}

	//计算出A中每个值的个数
	//var numTemp []int
	ALen := len(A)

	for i :=1;i<=ALen;i++{
		numTemp[A[i]] = numTemp[A[i]]+1
	}

	//计算出A中小于等于每个值的个数
	for j := 1;j<=max;j++ {
		numTemp[j] = numTemp[j]+numTemp[j-1]
	}

	//将A中元素放在合适的位置
	var returnArr = make(map[int]int)
	//for n :=ALen;n>0;n--{
	for n :=1;n<ALen;n++{
		returnArr[numTemp[A[n]]] = A[n]
		numTemp[A[n]]--
	}

	return returnArr
}

//归并排序
//利用了递归的思想，将一个数组分成若干小份，然后进行排序合并，因为合并后的子数组都说有序的，所以合并起来是一个线性时间的比较操作
func MergeSort(A map[int]int,p int,r int){
	if p<r{
		center := (r+p)/2
		q :=math.Floor(float64(center))
		MergeSort(A,p,int(q))
		MergeSort(A,int(q)+1,r)
		Merge(A,p,int(q),r)
	}
}

//进行合并操作
func Merge(A map[int]int,p int,q int,r int){
	//计算2个子数组的长度
	//lengthL :=q-p+1
	//lengthR :=r-q

	//将2个子数组复制到新的数组中

}


//快速排序
//@desc

func QuickSort()  {

}
