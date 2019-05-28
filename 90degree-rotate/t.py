
a = None
COUNT = 3

def make_grid():
    a = []
    for i in range(COUNT):
        tmp = []
        for j in range(COUNT):
            tmp.append(0)
        a.append(tmp)
    return a

def pr(a):
    """
    print a grid
    """
    for i in range(COUNT):
        ts = ""
        for j in range(COUNT):
            ts = ts + str(a[i][j])
        print(ts)

"""
123
456
789

741
852
963

put row 0 in col len -1
put row 1 in col len -2
put row 2 in col len -3
"""
def rot90(a):
    res = make_grid()
    for i in range(COUNT):
        for j in range(COUNT):
            res[j][COUNT-i-1] = a[i][j]
    return res

def transpose(a):
    res = make_grid()
    for i in range(COUNT):
        for j in range(COUNT):
            res[j][i] = a[i][j]
    return res

def t_inpl(a):
    for n in range(COUNT-2):
        for m in range(n+1,COUNT-1):
            tmp = a[n][m]
            a[n][m] = a[m][n]
            a[m][n] = tmp
            
# not working
def transpose_inplace(a):
    for i in range(COUNT):
        for j in range(COUNT):
            tmp = a[i][j] # remember first element
            print("saving tmp",tmp)
            a[i][j] = a[j][i] # move other to first
            a[j][i] = tmp # put first in other

def test_rot90():
    a = make_grid()
    for i in range(COUNT):
        for j in range(COUNT):
            a[i][j] = i
    pr(a)
    print("calling rot90")
    b = rot90(a)
    pr(b)

def test_transpose():
    a = make_grid()
    for i in range(COUNT):
        for j in range(COUNT):
            a[i][j] = i
    pr(a)
    b = transpose(a)
    pr(b)


def reverse_row(a,row):
    i = 0
    while True:
        tmp = a[row][i]
        a[row][i] = a[row][COUNT-i-1]
        a[row][COUNT-i-1] = tmp
        i = i + 1
        if i > COUNT/2:
            break

def anotherwayto90rot():
    """
    transpose first
    reverse each row
    """
    a = make_grid()
    for i in range(COUNT):
        for j in range(COUNT):
            a[i][j] = i
    pr(a)
    b = transpose(a)
    for i in range(COUNT):
        reverse_row(b,i)
    pr(b)

def inplace90rot():
    """
    transpose first
    reverse each row
    """
    a = make_grid()
    crud = 0
    for i in range(COUNT):
        for j in range(COUNT):
            a[i][j] = crud
            crud = crud+1
    pr(a)
    # transpose_inplace(a)
    t_inpl(a)
    pr(a)

test_rot90()
