#Congruencial Mixto Simulacion Digital
import pandas as pd
a = 1
c = 4
m = 5
x0 = 3

def congruencial_mixto(a, c, m, x0):
    #(ax0+c)mod m
    xn = (a*x0+c)%m

    row =[[x0, a, c, m, x0/(m-1)]]
    while xn != x0:
        row.append([xn, a, c, m, xn/(m-1)])
        x1 = (a*xn+c)%m
        if(xn == x1):
            break
        xn=x1
    return row

def congruencial_multiplicativo(a,m,x0):
    #(ax0)mod m
    xn = (a*x0)%m
    row =[[x0, a, x0/(m-1)]]
    while xn != x0:
        row.append([xn, a, xn/(m-1)])
        x1 = (a*xn)%m
        if(xn == x1):
            break
        xn=x1
    return row

p = pd.DataFrame(data= congruencial_multiplicativo(a, m, x0), columns=['Xn', 'a' , 'Ri'])

print(p)
print("------------------------------------------------")
p1= pd.DataFrame(data= congruencial_mixto(a, c, m, x0), columns=['Xn', 'a', 'c', 'm', 'Ri'])
print(p1)
