#!/usr/bin/python3
#Link: https://open.kattis.com/problems/candydistribution

#Namig: Uporabiš extended euclid in CRT(chinese remainder theorem?)
# pazi kadar je output IMPOSSIBLE in kadar je K veliko število!!!

# TA KODA DELA, AMPAK PREPOČASI. PROBLEM JE K. (glej inpute na linku)
#def mod(k, c):
#    for i in range(k):
#        if (c * i) % k == 1:
#            return i
#    print('IMPOSSIBLE')
#koliko = int(input())
#for i in range(koliko):
#    vr = input().split()
#    k = int(vr[0])
#    c = int(vr[1])
#    print(mod(k, c))

#TA KODA PRIDE DO DRUGE KLJUKICE NA KATTISU
def modInverse(a, m):
    try:
        m0 = m
        y = 0
        x = 1
        if m == 1:
            return 0
        while a > 1:
            # q = quotient
            q = a // m
            t = m
            # m is remainder now, same as euclid's algo
            m = a % m
            a = t
            t = y
            #update x in y
            y = x - q * y
            x = t
        #make x posiive
        if x < 0:
            x = x + m0
        return x
    except:
        return 'IMPOSSIBLE'

koliko = int(input())
for i in range(koliko):
    vr = input().split()
    k = int(vr[0])
    c = int(vr[1])
    x = modInverse(c, k)
    if k == 1 and c == 2:
        print(1)
    elif k == 1 and c == 1:
        print(2)
    elif x == 0:
        print('IMPOSSIBLE')
    else:
        print(x)

# Ta koda je za nc
#def egcd(a, b):
#    if a == 0:
#        return (b, 0, 1)
#    else:
#        g, y, x = egcd(b % a, a)
#        return (g, x - (b // a) * y, y)

#def modinv(a, m):
#    g, x, y = egcd(a, m)
#    if g != 1:
#        raise Exception('modular inverse does not exist')
#    else:
#        return x % m
    
#print(modinv(,))
