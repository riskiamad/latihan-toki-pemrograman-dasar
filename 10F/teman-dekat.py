str_in = input().split(' ')

n = int(str_in[0])
d = int(str_in[1])

def f(xi, yi, xj, yj):
    return abs(xj-xi)**d + abs(yj-yi)**d

arr = []

for i in range(n):
    arr_in = input().split(' ')
    x = int(arr_in[0])
    y = int(arr_in[1])

    arr.append((x,y))

i=0
j=1

ans=[]

while i < len(arr) - 1:
    ans.append(f(arr[i][0], arr[i][1], arr[j][0], arr[j][1]))

    if j < len(arr) - 1:
        j += 1
    else:
        i += 1
        j = i + 1

print(min(ans), max(ans))
