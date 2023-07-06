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
    