#Number
##题目-- [Two Sum - LeetCode](https://leetcode.com/problems/two-sum/description/)
旧实现的重新整理,Swift2版本
###Brute force
```swift
func towSum(numbers:Array<Int>,target:Int)->Array<Int>{
    for (var i=0;i<numbers.count;i++){
        for(var j=i+1;j<numbers.count;j++){
            if(numbers[i]+numbers[j]==target){
                return [i,j]
            }
        }
    }
    return []
}
```
###Hash Table
```swift
func towSum(numbers:Array<Int>,target:Int)->Array<Int>{
    var hashMap:Dictionary<Int,Int> = Dictionary<Int,Int>()
    for (var i=0;i<numbers.count;i++){
        let key = target-numbers[i]
        if(key==0){
            continue
        }
        if let index1 = hashMap[key] {
            return [index1,i]
        }
        hashMap[numbers[i]]=i
    }
    return []
}
```
## 题目 -- [Two Sum II - Input array is sorted - LeetCode](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/description/)
### Binary Search
```swift
func bsearch(array:Array<Int>,target:Int,start:Int)->Int{
    var left = start;
    var right = array.count-1;
    while (left<right){
        let mid = (left+right)/2
        if(array[mid]<target){
            left = mid+1
        }else{
            right = mid
        }
    }
    if(left==right&&array[left]==target){
        return left
    }
    return -1
}

func twoSum(array:Array<Int>,target:Int)->Array<Int>{
    for(var i=0;i<array.count;i++){
        let index2 = bsearch(array, target: target-array[i], start: i+1);
        if(index2>0){
            return [i,index2]
        }
    }
    return []
}
```

### Two Pointers
```swift
func twoSum(array:Array<Int>,target:Int)->Array<Int>{
    var left = 0
    var right = array.count-1
    while (left<right){
        if(array[left]+array[right]>target){
            right--
        }else if(array[left]+array[right]<target){
            left++
        }else{
            return [left,right]
        }
    }
    return []
}
```
## 题目-- [Two Sum III - Data structure design](https://leetcode.com/problems/two-sum-iii-data-structure-design)
```swift
class TwoSum {
    var hashMap:Dictionary<Int,Int> = Dictionary()
    func add(num:Int){
        var newCount = 1
        if let count = hashMap[num]{
            newCount = count+1
        }
        hashMap[num]=newCount
    }
    func find(target:Int)->Bool{
        for num in hashMap.keys{
            if let another=hashMap[target-num]{
                if(target-num==num){
                    if (another>1){
                        return true
                    }
                }else{
                    return true
                }
            }
        }
        return false
    }
}

```
## 题目 -- [Valid Number - LeetCode](https://leetcode.com/problems/valid-number/description/)
```swift
class Solution {
    func isNumber(s: String) -> Bool {
        
        var isNumric = false
        var scanIndex = s.characters.startIndex
        var scancount = 0
        
        while(scanIndex != s.characters.endIndex && String(s.characters[scanIndex]) == " "){
            scancount++
            scanIndex = scanIndex.successor()
        }
        
        if(scanIndex != s.characters.endIndex && (String(s.characters[scanIndex]) == "+" || String(s.characters[scanIndex]) == "-")){
            scancount++
            scanIndex = scanIndex.successor()
        }
        
        while(scanIndex != s.characters.endIndex && Int(String(s.characters[scanIndex])) != nil){
            scancount++
            isNumric = true
            scanIndex = scanIndex.successor()
        }
        
        if( scanIndex != s.characters.endIndex && String(s.characters[scanIndex]) == "."){
            scanIndex = scanIndex.successor()
            scancount++
            
            while(scanIndex != s.characters.endIndex && Int(String(s.characters[scanIndex])) != nil){
                isNumric = true
                scanIndex = scanIndex.successor()
                scancount++
            }
        }
        
        // 注意这里的isNumric判断
        if(isNumric && scanIndex != s.characters.endIndex && String(s.characters[scanIndex]) == "e"){
            scanIndex = scanIndex.successor()
            scancount++
            if(scanIndex != s.characters.endIndex && (String(s.characters[scanIndex]) == "+"
                || String(s.characters[scanIndex]) == "-")){
                    scanIndex = scanIndex.successor()
                    scancount++
            }
            
            isNumric = false
            while(scanIndex != s.characters.endIndex && Int(String(s.characters[scanIndex])) != nil){
                isNumric = true
                scanIndex = scanIndex.successor()
                scancount++
            }
        }
    
        while(scanIndex != s.characters.endIndex && String(s.characters[scanIndex]) == " "){
            scancount++
            scanIndex = scanIndex.successor()
        }
        
        return isNumric && scancount == s.characters.count
    }
}

```