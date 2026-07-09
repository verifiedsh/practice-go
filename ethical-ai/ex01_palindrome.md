## The Palindrome Challenge

## 1. Pseudocode for a function that checks if a string is a palindrome:
- START by defining a function name for the FUNCTION that takes a STRING as an argument(INPUT).
- Initialize a VARIABLE which is the REVERSE of a SLICE of the user INPUT with index ignored.
- Write a condition expression IF the two(2) variables are same, RETURN TRUE. ELSE, RETURN FALSE. Then, END the function.

            
## 2. Solution in Python:
``` py
def palindrome_func(string):
#Building a function entails, firstly defining a function_name for the funtion that receives input of string data type as an argument.
    invert = string[::-1]
    #Since a palindrome is a word or phrase or collection of characters whose reversal is same as it's forward approach, declare a variable that is a reversed version of the slice of the input recieved by the function from the user.
    return string == invert
    #Make comparisons between the initial input and the reversed version of it, if they turn out to be the same, validate the string to be a palindrome. Otherwise, nullify the string not being a palindrome.
print(palindrome_func("racecar"))
print(palindrome_func("hello"))
print(palindrome_func("A man a plan a canal panama"))
```
## 3.  Test with examples like "racecar", "hello", and "A man a plan a canal Panama":

1. "racecar"
OUTPUT: TRUE

2. "hello"
OUTPUT: FALSE

3. "A man a plan a canal panama"
OUTPUT: FALSE

## 4. Comments explaining my reasoning:
-  Included in the code above.


## Step 2 
## After using AI to go deeper:
1. Overall time complexity is O(n).

2. Edge cases I missed; 
- Case sensitivity.
- Spaces or punctuations.
- Empty string.
- Non-string input.

3. Better approach(es);
- Reverse.
- Clean + Pointers.
- Recursion.
- Slice Recursion.
- Two Pointer approach:
```py
def palindrome_func(string):
    # Normalize: lowercase and keep only alphanumeric characters
    cleaned = "".join(char.lower() for char in string if char.isalnum())
    
    # Reverse and compare
    return cleaned == cleaned[::-1]

print(palindrome_func("racecar"))
print(palindrome_func("hello"))
print(palindrome_func("A man a plan a canal panama"))
```
1. "racecar" OUTPUT: TRUE
2. "hello" OUTPUT: FALSE
3. "A man a plan a canal panama"OUTPUT: TRUE

## Step 3 - Reflection:

- Before asking AI, I basically learnt to declare functions, variables and write condition expressions(using IF, ELSE and ELIF) in python manually and also the syntax after the condition statement is made. 

- Now, I learnt about O(n) time complexity and got to understand that my code initially did not cover for case sensitivity, spaces, punctuations and numbers in between the phrase received as an argument(string) which technically might be present and the string will still be a palindrome. Now the revised version covers that and improves my understanding of palindrome in python and how to ask AI better and intuitive questions to enhance my learning.             

- Now, I can write similar functions (e.g., reverse a string) without help from AI.


