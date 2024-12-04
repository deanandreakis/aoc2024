
left=[]
right=[]
total=0
similarity=0

with open("inputd1.txt", "r") as inputlines:
    lines = inputlines.readlines()

for line in lines:
    as_list = line.split("   ")
    left.append(as_list[0])
    right.append(as_list[1].replace("\n",""))


left.sort()
right.sort()

for x, y in zip(left, right):
    total = total + abs(int(x)-int(y))

print(total)

# lets find the similarity score

for value in left:
    count = right.count(value)
    similarity = similarity + (int(value) * count)

print(similarity)

