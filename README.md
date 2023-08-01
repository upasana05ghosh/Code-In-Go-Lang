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
     if val, ok := myMap[key]; ok {
            //do something
     }
    ```
 3. Add/Update map
    ```
    myMap[key] = val
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


### For loop in go lang

```
for (i := 0; i< len(word); i++) {
   
}
```
    