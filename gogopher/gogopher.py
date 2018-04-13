"
    2018 Code Jam Qual Problem Go, Gopher!
    My solution creates a 2d array representing the board excluding the edges.
    The edges are excluded because they are not considered because they could produce
    a value outside of the rectangle we need to fill. Initially, the board is filled
    with 9's. This number indicates the amount of neighbors that require the gopher's
    attention (including the selected square). My program finds the maximum value
    in the neighbor array and deploys the gopher to that position. The program then
    decrements the neighbor values for the 3x3 grid around the point. The program
    continues to select the maximum value until every value is 0, indicating that
    every neighbor has been filled by the gopher.
"
# Helper function that prints the board for debugging
def prettyBoardPrint(board):
    for row in board:
        print(row)

# Decrements the neighbor count for all neighboring cells. (Only for unique cells)
def lowerNeighbors(outputset, board, x, y):
    if (x, y) not in outputset:
        outputset.add((x,y))
        for j in range(-1,2):
            for i in range(-1,2):
                if x + i >= 0 and x + i < len(board[0]) and y + j >= 0 and y + j < len(board):
                    board[y+j][x+i] -= 1

# Returns the first point with the highest neighbor value.
def maxBoard(board):
    maxv = 0
    maxPoint = (0,0)
    for j in range(len(board)):
        for i in range(len(board[0])):
            if board[j][i] > maxv:
                maxPoint = (i,j)
                maxv = board[j][i]
    return maxPoint

ntests = int(input())
for test in range(1, ntests + 1):
    a = int(input())
    if a == 20:
        board = [[9 for i in range(4-2)] for j in range(5-2)]
    if a == 200:
        board = [[9 for i in range(10-2)] for j in range(20-2)]
    outputset = set()
    x, y = 1, 1
    while (x > 0 and y > 0):
        x, y = maxBoard(board)
        print(x+2, y+2)
        x, y = [int(n) for n in input().split()]
        lowerNeighbors(outputset, board, x-2, y-2)

