import os

def preberi_izracunaj_izpisi(base_file, results_file):

    try:
        f = open(base_file, 'r')
        seznam = f.readlines()
        f.close()
    except:
        raise ValueError('Datoteka ' + base_file + ' ne obstaja!')

    #preverim datoteko
    i, j, k, l, m = 0, 1, 2, 3, 4
    while m <= len(seznam):
        if len(seznam[i].rstrip()) != 19:
            raise ValueError('Napaka v ' + str(i + 1) + ' vrstici!')
        if seznam[j].rstrip() != 'Buy' and seznam[j].rstrip() != 'Sell':
            raise ValueError('Napaka v ' + str(j + 1) + ' vrstici!')
        try:
            float(seznam[k].rstrip())
            float(seznam[l].rstrip())
            float(seznam[m].rstrip())
        except:
            raise ValueError('Napaka nekje v vrsticah od ' + str(k + 1) + ' do ' + str(m + 1))
        i += 5
        j += 5
        k += 5
        l += 5
        m += 5

    #drugace vse ok, grem racunat
    nov_seznam = []
    for element in seznam:
        nov_seznam.append(element.rstrip())

    i = 0
    koncni_seznam = []
    while i < len(nov_seznam):
        koncni_seznam.append(nov_seznam[i : i + 5])
        i += 5

    sez_rezultatov = []

    for element in koncni_seznam:
        datum = [element[0][:10]]
        if datum not in sez_rezultatov:
            sez_rezultatov.append(datum)
    pomozni_sez = []
    for element in sez_rezultatov:
        for elt in element:
            pomozni_sez.append(str(elt))

    f = open(results_file, 'w') #pazi na 'a' ce dodajas nove podatke

    i = 0
    for datum in pomozni_sez:
        buy = 0
        sell = 0
        for podseznam in koncni_seznam:
            if 'Buy' in podseznam and datum == podseznam[0][:10]:
                buy += float(podseznam[-1])
            if 'Sell' in podseznam and datum == podseznam[0][:10]:
                sell += float(podseznam[-1])
        f.write('----------' + sez_rezultatov[i][0] + '----------' + '\n' + 
                'Kupljeno: ' + str(buy) + '\n' + 
                'Prodano: ' + str(sell) + '\n')
        i += 1
    f.close()

    return sez_rezultatov

preberi_izracunaj_izpisi("bazaena.txt", "rezultati.txt")