# Dynamic programming
##题目-- [Best Time to Buy and Sell Stock - LeetCode](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/)
###1st Version
在Swift4中没有了for的i++实现，所以采用了迭代器。所以在二层迭代时需要去构建一个subArray。算法结果没有问题，但是在提交时因为性能问题被拒绝了。

```swift
class Solution {
	func maxProfit(_ prices: [Int]) -> Int {
		var profit = 0
		for (i,buy) in prices.enumerated() {
			if(i+1 == prices.count) {
				return profit
			}			
			let sub = prices[(i+1)...]
			for (_,sell) in sub.enumerated() {
				profit = max((sell-buy),profit)
			}
		}
		return profit
	}
}
```

###2nd Version
最初以为是在迭代器层面调用的性能影响，所以进一步优化了代码，其实应该不构成瓶颈。取消了迭代器的实现，代码上更精简了一些。提交，依旧因为性能问题被拒。

```swift
class Solution {
	func maxProfit(_ prices: [Int]) -> Int {
		var profit = 0
		for i in 0..<prices.count {
			for j in (i+1)..<prices.count {
				profit = max((prices[j]-prices[i]),profit)
			}			
		}
		return profit
	}
}
```


###Acceptable Version
所以最终考虑剪枝，避免无意义的循环。即发现后一买入日期比前一买入日期股票价格还要贵的话，其最终卖出利润肯定会较之前买入来的小，所以这种情况直接忽略。所以可以减少很多次循环。尤其是leetcode测试用例中的情况`[10000,9999......,0]`只要O(n)的时间效率就可以完成。

```swift
class Solution {
	func maxProfit(_ prices: [Int]) -> Int {
		var profit = 0
		var buyAt = 0
		for i in 0..<prices.count {
			if (prices[i]>prices[buyAt]) {
				continue
			}
			for j in (i+1)..<prices.count {
				let delt = prices[j]-prices[i]
				if delt > profit {
					buyAt = i
					profit = delt
				}
			}
		}
		return profit
	}
}
```

##题目--[Cheapest Flights Within K Stops - LeetCode](https://leetcode.com/problems/cheapest-flights-within-k-stops/description/)

###Acceptable Version
理解K值的意义就是这道题的关键，即K+1就是广度的层级。
为了找到最便宜的票价，我们需要将可联通的目的地串联起来，这里定义了一个Node结构，左子树表示该节点可以到达节点，右子树表示同级的节点。每一级的起始点，通过heaps数组保存，便于下一级便利。每一个节点保存从最初节点到当前节点的总价也就是本题要求的值。如果发现当前节点的总价大于目前最便宜的总价，就没有必要在当前节点继续下去。完全遍历直到找到最小的值。

如果没有合适的路径，最终返回-1，在当前场景下Int.max值9223372036854775807简直是个天文数字，可以做星际旅行了，所以用它作为是否有路径的判断即可。

这里通过链表连接是希望可以从起始点追溯终点。在本题中没有要求输入路径，将同级和下级用网状结构编织起来，而且并没有完全建立理想的结构，回溯还是需要双向链表来实现，有待改进。当然也用类似于sumPrice的变量，将这个路径字段保存。


```swift
class Solution {
	
	class Node {
		
		var flight:[Int]
		var sumPrice:Int
		var left:Node?
		var right:Node?
		init(_ f:[Int]) {
			self.flight = f
			self.sumPrice = f[2]
		}
		
	}
	
	
	func findCheapestPrice(_ n: Int, _ flights: [[Int]], _ src: Int, _ dst: Int, _ K: Int) -> Int {
		
		var heaps = (0..<(K+1)).map { _ in Node([0,0,0])}
		var cheapestPrice = Int.max
		
		for i in 0..<(K+1) {
			
			for flight in flights {
				
				if(i == 0) {
					if (flight[0] == src) {
						let node = Node(flight)
						node.right = heaps[i]
						heaps[i]=node
						if (flight[1] == dst && flight[2] < cheapestPrice) {
							cheapestPrice = flight[2]
						}
					}
				} else {
					var parent = heaps[i-1]
					while parent.right != nil {
						if(parent.flight[1] == flight[0]) {
							let child = Node(flight)
							let sum = parent.sumPrice + child.sumPrice 
							if(sum < cheapestPrice) {
								child.sumPrice = sum
								child.right = heaps[i]
								heaps[i]=child
								if (flight[1] == dst) {
									cheapestPrice = sum
								}
							}
						}
						parent = parent.right!						
					}
				}
			}			
		}
		
    	return cheapestPrice == Int.max ? -1 : cheapestPrice
	}
}

```

