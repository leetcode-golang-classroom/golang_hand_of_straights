# golang_hand_of_straights

Alice has some number of cards and she wants to rearrange the cards into groups so that each group is of size `groupSize`, and consists of `groupSize` consecutive cards.

Given an integer array `hand` where `hand[i]` is the value written on the `ith` card and an integer `groupSize`, return `true` if she can rearrange the cards, or `false` otherwise.

## Examples

**Example 1:**

```
Input: hand = [1,2,3,6,2,3,4,7,8], groupSize = 3
Output: true
Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8]

```

**Example 2:**

```
Input: hand = [1,2,3,4,5], groupSize = 4
Output: false
Explanation: Alice's hand can not be rearranged into groups of 4.

```

**Constraints:**

- `1 <= hand.length <= 104`
- `0 <= hand[i] <= 109`
- `1 <= groupSize <= hand.length`

## 解析

給定一個整數陣列 hand, 還有一個整數 groupSize

hand 代表手中的卡片

每個 hand[i] 的值代表第 I 章卡片上的數字

groupSize 代表想要把所有卡片平均分堆,每堆內的卡片個數

而這 groupSize 的每個卡片上數字需要是 連續的

要求寫一個演算法來判斷是否能夠把手中的卡片照上面的分堆法來分

首先思考要能夠均勻分堆 每堆要是 groupSize

代表 len(hand) 必須要能被 groupSize 整除

所以如果 len(hand) 無法整除 groupSize 則回傳 false

而要檢查每堆都要是連續的

這可以透過把卡片數字放在 HashTable 來紀錄每種數字有多少個

然後在 HashTable 最小的數字 min 每次 找連續 groupSize 個

出來做遞減

首先當發現要找的數字不在 HashTable 代表 分堆無法具有連續數字所以回傳 false

當發現遞減的數 遞減之後變成 0 則 更新目前最小數 min 位下一個 key

而要每次找當下最小的數的方式可以透過 MinHeap 來做處理 可以達到 log(n) 的良好時間複雜度

當發現遞減的數 遞減之後變成 0 則把目前 minHeap 的 top pop 掉

因為數字有 n個 所以時間複雜度是 O(nLogn)

空間複雜度是 O(n)

![](https://i.imgur.com/WjHUIj6.png)

## 程式碼
```go
package sol

import "container/heap"

type MinHeap []int

func (h *MinHeap) Len() int {
	return len(*h)
}
func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func isNStraightHand(hand []int, groupSize int) bool {
	nHand := len(hand)
	if groupSize == 1 {
		return true
	}
	if nHand%groupSize != 0 {
		return false
	}
	count := make(map[int]int)
	for _, val := range hand {
		count[val] += 1
	}
	priorityQueue := MinHeap{}
	heap.Init(&priorityQueue)
	for key := range count {
		heap.Push(&priorityQueue, key)
	}
	for priorityQueue.Len() > 0 {
		start := priorityQueue[0]
		end := start + groupSize
		for num := start; num < end; num++ {
			_, ok := count[num]
			if !ok {
				return false
			}
			count[num] -= 1
			if count[num] == 0 {
				heap.Pop(&priorityQueue)
			}
		}
	}
	return true
}

```
## 困難點

1. 要想出如何逐步檢查分堆內有連續數字邏輯不太直覺

## Solve Point

- [x]  先判斷 len(hand) % groupSize 是否等於 0, 若否則回傳 false
- [x]  建立 HashTable count 累計每個值出現的次數
- [x]  把每個在 count 的 key 放入 minHeap
- [x]  每次從 minHeap 的 top 開始去 count 中找連續 groupSize 數字做遞減
- [x]  當發現 要找的數不在 count 回傳 false
- [x]  當遞減完 該數字的累計次數 變成 0, 則把 minHeap Pop 一個值
- [x]  重腹以上不驟直到 minHeap 為空，則代表可以分成功 回傳 true