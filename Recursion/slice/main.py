

def slice_reverse(s):
    if len(s) == 1:
        return s 
    return slice_reverse(s[1:]) + [s[0]]


def slice_in_reverse(s):
    if s == [] :
        return [] 
    if type(s[0]) == list:
        return slice_in_reverse(s[1:]) + [slice_in_reverse(s[0])]
    return slice_in_reverse(s[1:]) + [s[0]]

def main():
    s = [1, 2, 3, 4, 34]
   
    x = [1, [22, 43], 2, 3, [4, 34]]
    print(slice_reverse(s))
    print(slice_in_reverse(x))

if __name__== "__main__":
    main()
