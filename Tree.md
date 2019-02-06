#String
旧实现的重新整理,Swift2版本
## 题目 -- [Path Sum - LeetCode](https://leetcode.com/problems/path-sum/description/)

作为作为笔试题提供给面试者。自己在做的时候，因为看到题目的数值都是正式，所以考虑在sum大于目标的时候及时停止递归。后来发现想多了，数字是任意的，不是持续递增，没有必要加限制。此外对guard的理解需要明确，guard里面是满足继续执行的条件。

```swift
func hasPathSum(_ root: TreeNode?, _ sum: Int) -> Bool {
    
    var found = false
    
    func calSum(_ node: TreeNode?, _ total: Int) {
        
        guard found != true , let n = node else {
            return
        }
        
        let nSum = n.val+total
        
        if(nSum == sum && n.left == nil && n.right == nil){
            found = true
            return
        }
        
        calSum(n.left,nSum)
        calSum(n.right,nSum)

    
    }
    
    calSum(root,0)
    
    return found
    
}
```
