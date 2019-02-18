#!/usr/bin/env python3
import random
import sys
repl_match = [['0', '6', '9'], ['2', '3'], ['3', '2', '5'], ['4', '11'], ['5', '3'], ['6', '0', '9'], ['9', '0', '6'],
              ['11', '4'], ['-', '/'], ['/', '-']]
rem_match = [['6', '5'], ['7', '1'], ['8', '0', '6', '9'], ['9', '5', '3'], ['+', '-'], ['*', '/']]
add_match = [['0', '8'], ['1', '7'], ['3', '9'], ['5', '6', '9'], ['6', '8'], ['9', '8'], ['-', '+'], ['/', '*']]
tasks = []
random.seed()
if len(sys.argv)<2:
    print("INVALIDARGUMENTS")
    exit(0)
level = sys.argv[1]
while len(tasks) == 0:
    expr=''
    if level=='1':
        if random.randint(0,1)==0:
            a = random.randint(0,9)
            b = random.randint(0,9)
            c = a + b
            expr=str(a)+'+'+str(b)+'='+str(c)
        else:
            a = random.randint(0, 9)
            b = random.randint(0, a)
            c = a - b
            expr = str(a) + '-' + str(b) + '=' + str(c)
    elif level=='2':
        if random.randint(0,1)==0:
            a = random.randint(10,99)
            b = random.randint(10,99)
            c = a + b
            expr = str(a) + '+' + str(b) + '=' + str(c)
        else:
            a = random.randint(10, 99)
            b = random.randint(10, a)
            c = a - b
            expr = str(a) + '-' + str(b) + '=' + str(c)
    elif level=='3':
        n = random.randint(3, 6);
        ex=[random.randint(0,100)]
        for i in range(n-1):
            ex.append(random.randint(-100,100))
        if sum(ex) < 0:
            ex.append(random.randint(-sum(ex),-2*sum(ex)))
        for i in ex:
            if expr=='':
                expr+=str(i)
            else:
                if i>0:
                    expr += '+' + str(i)
                elif i<0:
                    expr += str(i)
                elif random.randint(0,1)==0:
                    expr += '+0'
                else:
                    expr += '-0'
        expr+='='+str(sum(ex))
    else:
        print("INVALIDARGUMENTS")
        exit(0)
    # find tasks
    for j in repl_match:
        cur_num = expr.find(j[0])
        while cur_num != -1:
            for k in range(1, len(j)):
                changed_str = expr[:cur_num] + j[k] + expr[cur_num + 1:]
                changed_str = changed_str.replace('/', '//')
                eq_parts = changed_str.split('=')
                try:
                    if eval(eq_parts[0]) != eval(eq_parts[1]):
                        tasks.append(changed_str.replace('//', '/'))
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
                                    if eval(eq_parts[0]) != eval(eq_parts[1]):
                                        tasks.append(changed_str2.replace('//', '/'))
                                except:
                                    pass
                        cur_num2 = expr.find(l[0], cur_num2 + 1)
            cur_num = expr.find(j[0], cur_num + 1)

    # print random task
    if len(tasks)!=0:
        print(tasks[random.randint(0,len(tasks)-1)])
