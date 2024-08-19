#!/usr/bin/python3

import sys, termios, tty, os, time, curses

# The most compatible way
# rows, columns = os.popen('stty size', 'r').read().split()

# print ("Size of current terminal is %s - rows and %s - column" % (rows,columns))
# print ("Size of current terminal is %s - rows and %s - column" % (ts.lines,ts.columns))
# second commit
# stdscr = curses.initscr()

def getch():
    fd = sys.stdin.fileno()
    old_settings = termios.tcgetattr(fd)
    try:
        tty.setraw(sys.stdin.fileno())
        ch = sys.stdin.read(1)
 
    finally:
        termios.tcsetattr(fd, termios.TCSADRAIN, old_settings)
    return ch

button_delay = 0.2
# Alternative way
ts = os.get_terminal_size()
hero_ycoord = ts.lines
hero_xcoord = int(ts.columns/2)
hero_old_ycoord = hero_ycoord
hero_old_xcoord = hero_xcoord
cave = [[]]

# Initial cave
for i in range(0,(ts.lines-1)):
    temp_string = []
    for j in range(0,(ts.columns)):
        if (i == 0) or (i == ts.lines-2) :
            temp_string.append("-")
            continue
        if (j == 0) or (j == ts.columns-1) :
            temp_string.append("|")
            continue
        temp_string.append(".")
    cave.append(temp_string)
cave[hero_xcoord][hero_ycoord] = "@"

# Drawing a cave
def draw_cave() :
    os.system('clear')
    for i in range(0,(ts.lines)):
        temp_string = "".join(cave[i])
        print(temp_string)

#print ("-"*ts.columns)
#print ("|" + "."*(ts.columns-2) + "|")
#print ("|" + "."*(ts.columns-2) + "|")
#print ("|" + "."*(ts.columns-2) + "|")
#print ("-"*ts.columns)

while True:
    cave[hero_old_xcoord][hero_old_ycoord] = "."
    cave[hero_xcoord][hero_ycoord] = "@"
    draw_cave()
    hero_old_ycoord = hero_ycoord
    hero_old_xcoord = hero_xcoord

    char = getch()
 
    if (char == "p"):
        print("Stop!")
        exit(0)
 
    if (char == "a"):
        # print("Left pressed")
        hero_ycoord = hero_ycoord - 1
        time.sleep(button_delay)
 
    elif (char == "d"):
        # print("Right pressed")
        hero_ycoord = hero_ycoord + 1
        time.sleep(button_delay)
 
    elif (char == "w"):
        # print("Up pressed")
        hero_xcoord = hero_xcoord - 1
        time.sleep(button_delay)
 
    elif (char == "s"):
        # print("Down pressed")
        hero_xcoord = hero_xcoord + 1
        time.sleep(button_delay)
