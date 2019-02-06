#String
旧实现的重新整理,Swift2版本
## 题目 -- [Valid Palindrome - LeetCode](https://leetcode.com/problems/valid-palindrome/description/)
```swift
func isPalindrome(input:String)->Bool{
    
    func isLetterOrDigital(char:Character)->Bool{
        return char>="a"&&char<="z" || char>="A"&&char<="Z" || char>="0"&&char<="9"
    }

    let str = input.lowercaseString
    var left = str.startIndex
    //`endIndex` is not a valid argument to `subscript`
    var right = str.endIndex.predecessor()
    
    while (left<right){
        while(!isLetterOrDigital(str[left])){
            left = left.successor()
        }
    
        while(!isLetterOrDigital(str[right])){
            right = right.predecessor()
        }
        
        if(str[left] != str[right]){
            return false
        }

        left = left.successor()
        right = right.predecessor()
    }
    return true;
}
```

## 题目 -- [Implement strStr() - LeetCode](https://leetcode.com/problems/implement-strstr/description/)
```swift 
func strStr(haystack:String,needle:String)->Int{
    
    if(needle.characters.count==0 || haystack.characters.count==0){
        return -1
    }
    
    var needleIndex = needle.characters.startIndex
    for ( var index = haystack.characters.startIndex;
        index != haystack.endIndex;index=index.successor()){
            
            var haystackIndex = index
            while(haystack.characters[haystackIndex] == needle.characters[needleIndex]){
                haystackIndex = haystackIndex.successor()
                needleIndex = needleIndex.successor()
                if(needleIndex == needle.endIndex){
                    return haystack.characters.startIndex.distanceTo(haystackIndex)-needle.characters.count
                }
                
                if(haystackIndex.successor() == haystack.endIndex){
                    return -1
                }
            }
            
            needleIndex = needle.characters.startIndex
            
    }
    return -1;
}
```

## 题目-- [Reverse Words in a String - LeetCode](https://leetcode.com/problems/reverse-words-in-a-string/description/)
```swift
func reverseWordsHighOrder(input:String)->String{
    return input.characters.split(" ").reverse().reduce("", combine: { (result, char) -> String in
        if(result == ""){
            return String(char)
        }else{
            return result + " " + String(char)
        }

    })
}

func reverseWords(input:String)->String{
    
    var output = ""
    
    let startIndex = input.startIndex
    var rightIndex = input.endIndex.predecessor()
    var leftIndex = rightIndex
    let characters = input.characters;
    
    while(leftIndex != startIndex){
        if(characters[leftIndex] == " " && leftIndex == rightIndex){
            rightIndex = rightIndex.predecessor()
            leftIndex = leftIndex.predecessor()
        }else if (characters[leftIndex] != " "){
            leftIndex = leftIndex.predecessor()
        }else{
            let slice = str[leftIndex.advancedBy(1)...rightIndex]
            output = output + String(slice) + " "
            leftIndex = leftIndex.predecessor()
            rightIndex = leftIndex
        }
    }
    
    if(leftIndex != rightIndex){
        output = output + String(str[leftIndex...rightIndex])
    }
    
    return output
}

```

## 题目-- [Reverse Words in a String II](https://leetcode.com/problems/reverse-words-in-a-string-ii)

```swift
func reverseWords(var input:String)->String{
    
    func reverse(var startIndex:String.CharacterView.Index,var endIndex:String.CharacterView.Index){
        let characters = input.characters
        while(startIndex < endIndex){
            let startRange = Range(start: startIndex, end: startIndex.advancedBy(1))
            let endRange = Range(start: endIndex, end: endIndex.advancedBy(1))
            let startStr = String(characters[startRange])
            let endStr = String(characters[endRange])
            
            input.replaceRange(startRange, with: endStr)
            input.replaceRange(endRange, with: startStr)
            startIndex = startIndex.successor()
            endIndex = endIndex.predecessor()
        }
    }

    reverse(input.startIndex, endIndex: input.endIndex.predecessor())

    var leftIndex = input.startIndex
    let characters = input.characters
    var count = 0
    while(leftIndex != input.endIndex){
        if(characters[leftIndex] == " " ){
            reverse(leftIndex.advancedBy(-count), endIndex: leftIndex.advancedBy(-1))
            leftIndex = leftIndex.successor()
            count = 0
        }else if(leftIndex.successor() == input.endIndex){
            reverse(leftIndex.advancedBy(-count), endIndex: leftIndex)
        }
        
        leftIndex = leftIndex.successor()
        count++
    }
    return input
}
```

## 题目-- [String to Integer (atoi) - LeetCode](https://leetcode.com/problems/string-to-integer-atoi/description/)
```swift
func atoi(input:String)->Int{

    var sign = 1
    let preMax = Int.max/10
    let lastNum = Int.max%10
    
    let result = input.characters.reduce(0) { (output, char) -> Int in
        
        if(output == Int.max || output == Int.min){
            return output
        }
        
        if (char == "-"){
            sign = -1
        }else if let i = Int(String(char)){
            if (output > preMax || (output == preMax && i>=lastNum)){ //这个判断省了不少逻辑
                return sign == -1 ? Int.min :Int.max
            }
            return output*10+i
        }
        
        return output
    }
    

    if (sign < 0 && result > 0){
        return result*sign
    }
    
    return result
}
```

