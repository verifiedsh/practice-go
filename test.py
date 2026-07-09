def duplicate_func(string):
    string = "".join(c.lower() for c in string if c.isalnum()) 
    # rev = string[::-1]
    start , end = 0, len(string) - 1
    while start < end:
        # if not string[start]:
        #     start += 1
        #     continue
        # if not string[end]:
        #     end -= 1
        #     continue
        if string[start] == string[end]:
            return True, start, end 
        start += 1
        end -= 1    
    return False, None 

print(duplicate_func(input("Enter sample text:\n")))

def reverse_string(string):
  rev = string[::-1]
  return(rev)

print(reverse_string("Hello"))