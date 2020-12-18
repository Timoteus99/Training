#!/usr/bin/python3

class Candydist():
    def __init__(self):
        self.s = 0
        self.t = 0

    def gcd_ex(self, a, b):
        r0 = a
        r1 = b
        s0 = 1
        s1 = 0
        t0 = 0
        t1 = 1

        while r1 != 0:
            q = r0 // r1
            r0 = r0 - q * r1
            # swap
            tmp = r0
            r0 = r1
            r1 = tmp

            s0 = s0 - q * s1
            #swap
            tmp = s0
            s0 = s1
            s1 = tmp

            t0 = t0 - q * t1
            #swap
            tmp = t0
            t0 = t1
            t1 = tmp
        self.s = s0
        self.t = t0

        return r0

    def mod_inverse(self, a, m):
        self.a = a
        self.m = m

        a = a % m
        if a < 0:
            a = a + m

        r0 = self.gcd_ex(a, m)
        if r0 > 1:
            return -1

        self.s = self.s % m;

        if self.s < 0:
            self.s = self.s + m;

        return self.s

koliko = int(input())
candydist = Candydist()
for i in range(koliko):
    vr = input().split()
    k = int(vr[0])
    c = int(vr[1])
    x = candydist.mod_inverse(c, k)
    if k == 1 and c == 2:
        print(1)
    elif k == 1 and c == 1:
        print(2)
    elif x == 0:
        print('IMPOSSIBLE')
    else:
        print(x)



