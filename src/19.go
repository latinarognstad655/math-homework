x = 5
y = 3
z = 7
if x > y and x > z:
    print("The largest number is", x)
else:
    if y < x or y == z and z > x:
        print("The largest number is", y)
    else:
        print("The largest number is", z)
