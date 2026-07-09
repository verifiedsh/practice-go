# 1. Modify your palindrome function to:
##   - Ignore spaces and punctuation.
##   - Be case-insensitive.
##   - Return the position where the string stops being a palindrome (if not one).
```py
def palindrome_func(string):
  string = "".join(c.lower() for c in string if c.isalnum())
  start = 0
  end = len(string) - 1
  while start < end:
      if string[start] != string[end]:
        return False, start, end
      start += 1
      end -= 1
  return True, None
print(palindrome_func("racecar"))
print(palindrome_func("hello"))
print(palindrome_func("A man a plan a canal panama"))
```

# 2. Feedback from AI after first attempt: 

## Things I miss:
   - Not effcient for empty strings passed as arguments to the function. 
   - Unicode edge cases
   - String with only symbols
## Making it be more efficient:
```py
def palindrome_func(string):
    start = 0
    end = len(string) - 1

    while start < end:
        while start < end and not string[start].isalnum():
            start += 1
        while start < end and not string[end].isalnum():
            end -= 1

        if string[start].casefold() != string[end].casefold():
            return False, start, end

        start += 1
        end -= 1

    return True, None
print(palindrome_func("racecar"))
print(palindrome_func("hello"))
print(palindrome_func("A man a plan a canal panama"))
```

# 3. Reflection on what AI added that I didn't consider initially:
    
   - Initially, I had did not cover for empty strings, special characters and non-string characters.
   - After using AI, I got a better understanding of the logic of loops, how to  avoid unnecessary too many variables, handle Unicode better and cover for situations where an empty string or a string without characters is passed to the function as an argument.