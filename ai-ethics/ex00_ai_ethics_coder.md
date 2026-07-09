#   PART A: Self-assessment.

## 1a. How have you used AI for coding so far?

    Initially, i started off with using AI as my first option to attempt tasks and questions and then memorizing the resulting code from the answers provided by AI without actually understanding the scope and logic of the code. After being advised and discovering that, it wasn't an efficient way to learn programming I switched to using AI to study and approach the questions differently by giving it the tasks but setting and configuring it to study mode in order to properly develop and harness my problem solving skills with code.

## 1b.Do you ask AI for solutions before trying yourself?

    I try questions myself first but not long enough because, immediately I  meet a blocker I result to using AI to solve the issue I face.

## 1c. Can you explain code you've submitted without AI's help?

    Actually, to the best of my knowledge I can explain code I have submitted without AI's help.

## 1d. What would happen if AI was suddenly unavailable during an exam or interview?

    If AI is unavailable during an exam or interview, I would be able to attempt and solve some tasks or questions I encounter based on my level of dept and understading of programming on a general basis.

## 2. Identify your current pattern: Which learner are you now? 
- Learner B(AI is my learning amplifier)

- I try writing the code myself based on the fundamentals I know. 
- Ask AI questions like "Why is this approach efficient and are there better approaches? What are the edge cases and trade-offs of this approach?".
- critically vet my level of understanding by explaining concepts and knowledge to peers.
- Use AI to dive deeper, not to replace thinking and brainstorming.


## PART B: Palindrome solution + reflection
## The Palindrome Challenge

## 1. Write pseudocode for a palindrome check:
- START by defining a function name for the FUNCTION that takes a STRING as an argument(INPUT).
- Initialize a VARIABLE which is the REVERSE of a SLICE of the user INPUT with index ignored.
- Write a condition expression IF the two(2) variables are same, RETURN TRUE. ELSE, then RETURN FALSE. Then, END the function.
- Call on the function to test examples.
            
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
# Test function with examples.
```
## 3.  Test with examples like "racecar", "hello", and "A man a plan a canal Panama":

```py
1. "racecar"
OUTPUT: TRUE

2. "hello"
OUTPUT: FALSE

3. "A man a plan a canal panama"
OUTPUT: FALSE
```
## 4. Issues I debugged:
- No issues found

# Code without bug:
```py
def palindrome_func(string):
#Building a function entails, firstly defining a function_name for the funtion that receives input of string data type as an argument.
    invert = string[::-1]
    #Since a palindrome is a word or phrase or collection of characters whose reversal is same as it's forward approach, declare a variable that is a reversed version of the slice of the input recieved by the function from the user.
    return string == invert
    #Make comparisons between the initial input and the reversed version of it, if they turn out to be the same, validate the string to be a palindrome. Otherwise, nullify the string not being a palindrome.
print(palindrome_func("racecar"))
print(palindrome_func("hello"))
print(palindrome_func("A man a plan a canal panama"))
# Test function with examples.
```

## 5. Comments explaining my reasoning:
-  Included in the code above.


## Step 2 
## After using AI to go deeper:
1. Overall time complexity is O(n).

2. Edge cases I missed;
- Case Sensitivity
- Spaces
- Punctuation
- Unicode normalization
- Empty string
- Non-string inputs


3. Alternatives and trade-offs;
- Clean + Pointers.
- Recursion.
- Slice Recursion.
- Two Pointer approach:
```py
def palindrome_func(s):
    s = ''.join(c.lower() for c in s if c.isalnum())
    return s == s[::-1]
```
4. How does it perform on very long strings?"
Let’s say 100 million characters.

My approach:
- Creates another duplicate 100 million characters string
- Uses ~2× memory temporarily
- Still O(n)

Issues:
- Memory overhead
- Slower due to copying


## Step 3 - Reflection:

- What did you learn by struggling first?

    By struggling first, I learnt to be more self-reliant and depend on my own initiative there by testing the dept of my knowledge of the code I'm to write.

- How is your understanding different than if you'd just asked for the solution?

    If I just asked for the solution, I would not know for sure if I have a proper understanding on the code written. But now I wrote it on my own before asking AI I'm more confident in my own understqanding of the logic and scope of the code.

- Can you now implement similar functions (reverse a string, find duplicates) without AI?

    Now, I can write similar functions (e.g., reverse a string, find duplicates) without AI.

- What mental model did you build?

    The mental model I built is a trial first approach before requesting help to help of any sort to advance and fine-tune my answers.

## Part C: Testing Understanding
## Variation:
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

## Edge cases:
- Non-string inputs
- Strings with only non-alphanumeric characters
- Unicode normalization issues
- Extremely long strings (memory)

## How to make it more efficient:
- Avoid building the cleaned string entirely(by two-pointer + skip logic).
```py
def palindrome_func(s):
    if not isinstance(s, str):
        raise TypeError("Input must be a string")

    start, end = 0, len(s) - 1

    while start < end:
        while start < end and not s[start].isalnum():
            start += 1
        while start < end and not s[end].isalnum():
            end -= 1

        if s[start].lower() != s[end].lower():
            return False, start, end

        start += 1
        end -= 1

    return True, None
```


## Part D: Personal fairness contract
##                Personal rules for ethical AI use in learning.

## * I will use AI when:

-   I have attempted rectifying building the code to solve the problem on my own for at least 15 minutes.

-   I require deeper understanding about the logic and scope of the code I'm working on.

-   I need other better ways to modify my code or solution.

-   I have successfully built a substantial code that can function to some reasonable extent.

-   I need to summarize and explain bulky documentations to strengthen my knowledge about a dataset I need tot work.

## * I will NOT use AI when:

- I have not attempted building the code or finding the solution on my own for atleast 15 focused minutes.

- I have an official interview or assessment of any nature.

- I have not learnt the fundamentals of the programming language I am working with.

- I have no gotten a code that work reasonably.

- I have no idea or have not done proper research on the project at hand.

## * I know I'm using AI fairly when:

- I can give a valid and proper explanation of the code I have written and submitted without assistance from AI.

- I can write codes that solve similar issues on my own.

- I possess a high level of confidence in myself as a programmer.

- I can think and reason logically to produce steps that help build functional systems.

- I know the first correct step and subsequent steps to take on my own when faced with any challenge or code.

## Signature and date of contract:
OGBANJE ABRAHAM OKWUTE

Fri.20/02/2026

## Part E: Scenario analysis
## Interview: "Explain how you'd implement a caching system." If you always relied on AI, can you answer?

    I can not answer because I cannot explain code I did not write, though I will attempt to answer.

## Production bug at 2 AM: AI is unavailable. Can you debug code you don't fully understand?

    I can not debug a code I don't fully understand because I have insufficient knowledgeabout the logic/scope of the code. If I try, I might end up putting more bugs.

## New tech with little documentation: If you never learned to read docs and experiment, what happens?

    If I never learned to read docs and experiment, it will be almost impossible to learn to use the New tech properly.


## Part F: Skills assessment + action plan
## Rate yourself (1–5) in each skill and plan improvements:

|Skill	                    |Rating	|Improvement Plan                                                    |
|---------------------------|-------|--------------------------------------------------------------------|
|Problem Decomposition	    |4/5	|Critically observing the problem at hand before attempting to solve.|
|Systems Thinking	        |3/5	|Dive deeper into the core concepts of the programming language used.|
|Critical Evaluation	  	|3/5	|Make proper analysis of the code and properly vet each line of code.|
|Debugging Mindset	        |2/5	|Ensure my code executes and gives the expected output by fishing out unnecessary parts of the code.|
|Conceptual Understanding   |3/5	|Make proper and more detailed studies on programming languages I work with.|

## Choose your lowest-rated skill and write 3 specific actions you'll take this week to improve it without relying on AI:
### Skill: 
Debbuging Mindset
### Actions: 
- Pay more attention to error messages that pop up when the code does not run.
- Proof-read my code to ensure that the appropriate logic is followed in the building of the code before execution.
- Study more on the fundamentals of the programming language I use, so I gain more knowledge on the proper struture of the codes I write.