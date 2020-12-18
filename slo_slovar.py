import urllib.request
zacetek = str(input("Vnesi koren besede: "))
def same_razlicne(zacetek):
    '''vrne mnozico vseh tistih slo besed ki se zacno z nizom zacetek in imajo največ različnih črk'''
    naslov = "http://bos.zrc-sazu.si/sbsj.html"
    vir = urllib.request.urlopen(naslov)
    vse = vir.read().decode()
    besede = vse.split('\n')[15:-13]
    ls = []
    for element in besede:
        if zacetek == element[0 : len(zacetek)]:
            ls.append(element[0 : -5])
    return set(ls)
print(same_razlicne(zacetek))