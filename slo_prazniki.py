import datetime
from dateutil.easter import *
leto = int(input("Vnesi letnico: ")) #input je letnica, output pa dnevi tedna slovenskih dela prostih praznikov
dnevi_tedna = ("ponedeljek","torek","sreda","četrtek","petek","sobota","nedelja")
def kateri_dan(leto, mesec, dan):
    '''vrne str ime dneva v tednu glede na dani datum format Y, M, D'''
    datum = datetime.date(leto, mesec ,dan)
    dan_v_tednu = datum.weekday()
    dan_v_nizu = dnevi_tedna[dan_v_tednu]
    return(dan_v_nizu)
def velika_noc(leto):
    e = [easter(leto)] #seznam z 1 elementom ki je datum velike noci glede na leto Y-M-D
    dat = e[0]
    dan_v_tedn = dat.weekday()
    dan_v_niz = dnevi_tedna[dan_v_tedn]
    return (dan_v_niz)
#datum_velikonocnega_pon = mogoce kdaj potrebujes datum velikonocnega ponedeljka
slo_prazniki = [(leto,1,1),(leto,1,2),(leto,2,8),(leto,4,27),(leto,5,1),(leto,5,2),(leto,5,31),(leto,6,25),(leto,8,15),(leto,10,31),(leto,11,1),(leto,12,25),(leto,12,26)]
#slo_prazniki - v seznamu ni velike noci in velikonocnega ponedeljka
dela_prosti_prazniki = []
for element in range(0, 13):
    dela_prosti_prazniki.append(kateri_dan(slo_prazniki[element][0], slo_prazniki[element][1], slo_prazniki[element][2]))
dela_prosti_prazniki.append(velika_noc(leto)) #dodamo se veliko noc
dela_prosti_prazniki.append(dnevi_tedna[0]) #dodamo se velikonocni ponedeljek
#print("Ta dan je", kateri_dan(leto,mesec,dan) + ".")
vikend = 0
for element in dela_prosti_prazniki:
    if element == 'sobota' or element == 'nedelja':
        vikend += 1
cez_teden = (15 - vikend)
#print(dela_prosti_prazniki)
#print(datum_velike_noci)
print("V tem letu je število praznikov, ki padejo čez teden:", str(cez_teden), "od 15" + ".")