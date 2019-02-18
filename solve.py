#!/usr/bin/env python3
import json
import sys
repl_match = [['0', '6', '9'], ['2', '3'], ['3', '2', '5'], ['4', '11'], ['5', '3'], ['6', '0', '9'], ['9', '0', '6'],
              ['11', '4'], ['-', '/'], ['/', '-']]
rem_match = [['6', '5'], ['7', '1'], ['8', '0', '6', '9'], ['9', '5', '3'], ['+', '-'], ['*', '/']]
add_match = [['0', '8'], ['1', '7'], ['3', '9'], ['5', '6', '9'], ['6', '8'], ['9', '8'], ['-', '+'], ['/', '*']]
solutions = []
if len(sys.argv)!=2:
    print("INVALIDARGUMENTS")
    exit(0)
expr = sys.argv[1]

# input check
allowed_ch = list('0123456789+=-/*')
for ch in expr:
    if ch not in allowed_ch:
        print('WRONGINPUT')
        exit()
eq_parts = expr.split('=')
if len(eq_parts) != 2:
    print('WRONGINPUT')
    exit()
try:
    eval(eq_parts[0])
    eval(eq_parts[1])
except:
    print('WRONGINPUT')
    exit()
spec_symbols = list('+=-/*')
for i in range(len(expr) - 1):
    check = expr[i] + expr[i + 1]
    if (expr[i] in spec_symbols) and (expr[i + 1] in spec_symbols):
        print('WRONGINPUT')
        exit()

# find solutions
for j in repl_match:
    cur_num = expr.find(j[0])
    while cur_num != -1:
        for k in range(1, len(j)):
            changed_str = expr[:cur_num] + j[k] + expr[cur_num + 1:]
            changed_str = changed_str.replace('/', '//')
            eq_parts = changed_str.split('=')
            try:
                if eval(eq_parts[0]) == eval(eq_parts[1]):
                    solutions.append(changed_str.replace('//', '/'))
            except:
                pass
        cur_num = expr.find(j[0], cur_num + 1)
for j in rem_match:
    cur_num = expr.find(j[0])
    while cur_num != -1:
        for k in range(1, len(j)):
            changed_str = expr[:cur_num] + j[k] + expr[cur_num + 1:]
            for l in add_match:
                cur_num2 = expr.find(l[0])
                while cur_num2 != -1:
                    for m in range(1, len(l)):
                        if cur_num2 != cur_num:
                            changed_str2 = changed_str[:cur_num2] + l[m] + changed_str[cur_num2 + 1:]
                            changed_str2 = changed_str2.replace('/', '//')
                            eq_parts = changed_str2.split('=')
                            try:
                                if eval(eq_parts[0]) == eval(eq_parts[1]):
                                    solutions.append(changed_str2.replace('//', '/'))
                            except:
                                pass
                    cur_num2 = expr.find(l[0], cur_num2 + 1)
        cur_num = expr.find(j[0], cur_num + 1)

# print solutions
print(json.dumps(solutions))
