# Code-In-Go-Lang
Coding in Go Lang

## Notes

### Maps
 1. Create a map:
    ``` 
    myMap := make(map[int] int)
    ```
 2. check if map contains key:
    ```
     if val, isPresent := myMap[key]; isPresent {
            //do something
     }
    ```
 3. Add/Update map
    ```
    myMap[key] = val
    ```
4. To iterate over a map
   ```
   for key, val := range myMap {
      //do something
   }
   ```

### Set
1. Create a set: 
   ```
   mySet := make(map[int] bool)
   ```

2. Check if set contains key: 
   ```
   if val, isPresent := mySet[key]; isPresent {
      //do something
   }
   ```
3. Add value to set
   ```
   mySet[key] = true
   ```

### Character / rune
1. to create a char/character, rune is the keyword
   ```
    var myChar rune
   ```
2. To convert rune to string
   ```
   val := string(myChar)
   ```
3. To convert string to rune
   ```
   myRune := []rune(myString)
   ```
4. If you want to replace a char at index i in string. 
   a) Convert it to array of rune
   b) update the char at index i
      ```
      myRune[i] = 'c'
      ```
   c) Convert the rune array to string (follow #2)

### String
1. Find length of string
  ```
  len(str)
  ```
2. Substring of string
  ```
  s[: len(s) - 3]
  ```
3. To check if 2 strings are equal
   ```
   if s1 == s2 {
      // do something
   }
   ```
4. To remove one or multiple whitespaces from string, use strings.Fields
   It returns array of string without any whitespaces
  ```
  words := strings.Fields(str)
  ```
5. To convert array to string, use strings.Join()
   ```
   strings.Join(words, " ")
   ```

### String Builder
If we are updating strig frequently, a new string will be created in the memory. 
Reason - strings are immutable in Go, so if we are adding content to it, we are creating a new string every time. 
Use strings.Builder instead
1. To create string builder
   ```
   var sb strings.Builder
   ```
2. To add string to string builder
   ```
   sb.WriteString(s)
   ```
3. To convert it to string
   ```
   sb.String()
   ```

### Int
1. Convert int to string
   strNum := strconv.Itoa(num)
2. Max int value
   max := math.MaxInt32

### Stack Array

1. Create a stack array of type rune/char
    ```
    var stack[] rune
    ```
2. To push to the stack
    ```
    stack = append(stack, c)
    ```
3. Pop from stack
    ```
    stack = stack[: len(stack) -1]
    ```
4. Peek from stack
   ```
   prev := stack[len(stack) -1]
   ```
5. Check stack length
   ```
   len := len(stack)
   ```


### To create a slice
1. Slice is a variable sized array. We can use it when we don't know the exact size of array we want. 
   ```
   mySlice := make([]int, size)
   ```
2. While passing slice as an argument, we don't need to pass them as pointer. In Go, everything is passed by value, slices too. But a slice variable is a pointer to the value already.  
3. To check if two slice are equal
   ```
   reflect.DeepEqual(sl1, sl2)
   ```

### Array
1. To create an array of fixed size
   ```
   arr := make([]int, 10)
   ```
2. To sort array element
   ```
   sort.Ints(arr)
   ```
3. To create a 2-D array of nxn
   ```
   arr := make([][]int, n)
   ```
### For loop in go lang

```
for (i := 0; i< len(word); i++) {
   // do something
}
```

or 

```
for index, val := range arr {
   // do something
}
```
    