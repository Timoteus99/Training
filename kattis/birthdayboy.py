# birthday boy kattis problem - link: https://open.kattis.com/problems/birthdayboy
# math way :D --> dela!

def kateriDan(dan):
    '''zapise dan v nn-nn slogu'''
    dneviVmesecu = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
    mesec = 1
    for st in dneviVmesecu:
        if dan > st:
            dan = dan - st
            mesec += 1
        else: break
        
    dan = str(dan)
    mesec = str(mesec)

    if len(str(dan)) == 1:
        dan = '0' + dan
    if len(str(mesec)) == 1:
        mesec = '0' + mesec
    
    return mesec + '-' + dan
        

def dolociDatum(datumi):
    '''doloci datum ki ustrez pogojem'''
    st_dni = {'01':0,
              '02':31,
              '03':59,
              '04':90,
              '05':120,
              '06':151,
              '07':181,
              '08':212,
              '09':243,
              '10':273,
              '11':304,
              '12':334}

    dan_v_mesecu = {'01':1,
                    '02':2,
                    '03':3,
                    '04':4,
                    '05':5,
                    '06':6,
                    '07':7,
                    '08':8,
                    '09':9,
                    '10':10,
                    '11':11,
                    '12':12,
                    '13':13,
                    '14':14,
                    '15':15,
                    '16':16,
                    '17':17,
                    '18':18,
                    '19':19,
                    '20':20,
                    '21':21,
                    '22':22,
                    '23':23,
                    '24':24,
                    '25':25,
                    '26':26,
                    '27':27,
                    '28':28,
                    '29':29,
                    '30':30,
                    '31':31}
    
    danes = st_dni['10'] + dan_v_mesecu['27']
    tab_dni = []
    for datum in datumi:
        dan = datum.split('-')[1]
        mesec = datum.split('-')[0]
        dan_v_letu = st_dni[mesec] + dan_v_mesecu[dan]
        tab_dni.append(dan_v_letu)
    
    tab_dni = sorted(tab_dni)
    
    
    
    tab_razlik = []
    naj_razlika = 0
    kje = 0
    while kje < len(tab_dni):
        if kje == len(tab_dni) - 1:
            razlika = 365 - tab_dni[kje] + tab_dni[0]
            if razlika > naj_razlika:
                tab_razlik = [tab_dni[0]]
                naj_razlika = razlika
            elif razlika == naj_razlika:
                tab_razlik.append(tab_dni[0])
        
        else:
            elt = tab_dni[kje]
            razlika = tab_dni[kje + 1] - elt
            if razlika > naj_razlika:
                tab_razlik = [tab_dni[kje + 1]]
                naj_razlika = razlika
            elif razlika == naj_razlika:
                tab_razlik.append(tab_dni[kje + 1])
        kje += 1
            
    if len(tab_razlik) == 1:
        iskani_dan = tab_razlik[0] - 1
    
    else:    
        najblizji_datum = 0
        min_razlika = 365
        for dan in tab_razlik:
            if dan - 1 > danes:
                razlika = dan - danes
                if razlika < min_razlika:
                    min_razlika = razlika
                    najblizji_datum = dan
            else:
                razlika = 365 - danes + dan
                if razlika < min_razlika:
                    min_razlika = razlika
                    najblizji_datum = dan
            
        iskani_dan = najblizji_datum - 1
  
    if iskani_dan == 0:
        iskani_dan = 365
        
    rez = kateriDan(iskani_dan)
    return rez

    
datumi = []
st_delavcev = int(input())
for delavec in range(0,st_delavcev):
    delavec = input()
    datumi.append(delavec.split()[-1])

print(dolociDatum(datumi))

#date library way --> napaka pri osmem? testu

#from datetime import date, timedelta
#import sys
#import datetime

#datumi = ['01-09','12-31','03-22','09-20','11-11','02-28','09-25','12-15','11-17','05-01','08-02','04-25','06-12','03-26','01-20','01-20']
#datumi = ['01-09','09-20','11-11']
#datumi = ['04-28','10-28','04-29']
#datumi = ['04-30','10-29','04-29']

#danes_date = date(2018, 10, 27)
#danes_str = "10-27"

#def pripravi_datume(datumi):
#    datumi.sort()
#    print(datumi)
#    #print(datumi[0][0])
#    #print(datumi)
#    novi_datumi = []
#    a = ''
#    b = ''
#    for element in datumi:
#        if element[0] == '0':
#            a = element.replace('0', "",1)
#            if a[-2] == '0' or a[2] == '0':
#                b = a.replace('0',"",1)
#
#                novi_datumi.append(b.split("-"))
#            else:
#                novi_datumi.append(a.split("-"))
#        else:
#            novi_datumi.append(element.split("-"))
#    print(novi_datumi)
#    if len(datumi) < 2:
#        extra_datum = date(2018, int(novi_datumi[0][0]), int(novi_datumi[0][1])) + timedelta(1)
#        novi_datumi.append(extra_datum.strftime("%m-%d").split("-"))
#    return novi_datumi

#print(koncni_datumi)

#def najdi_datum(datumi):
#    datum = date(2018, 10, 28)
#    najvecja_raz = 0
#    i = 0
#    if len(datumi) == 0:
#        return datum
#    while i < len(datumi):
#       j = i + 1
#        if i == len(datumi) - 1:
#            j = 0
#            datum1 = date(2018, int(novi_datumi[i][0]), int(novi_datumi[i][1]))
#            datum2 = date(2019, int(novi_datumi[j][0]), int(novi_datumi[j][1]))
#            razlika = datum2 - datum1
#            print("element: ", i, "datum1: ", datum1, ", datum2: ", datum2, ", razlika: ", razlika)
#        else:
#            datum1 = date(2018, int(novi_datumi[i][0]), int(novi_datumi[i][1]))
#            datum2 = date(2018, int(novi_datumi[j][0]), int(novi_datumi[j][1]))
#            razlika = datum2 - datum1
#            print("element: ", i, "datum1: ", datum1, ", datum2: ", datum2, ", razlika: ", razlika)
#            
#        if najvecja_raz < abs(razlika.days):
#            najvecja_raz = abs(razlika.days)
#            datum = datum2
#        elif najvecja_raz == abs(razlika.days):
#            if abs((datum2 - danes_date).days) < abs((datum - danes_date).days):
#                datum = datum2
#        i += 1
#
#    datum = datum - timedelta(1)
#    
#    print(datum)
#    if datum == danes_date:
#        print("Ne more bit danes!")
#        print(datumi)
#        datumi.append(['10', '27'])
#        datum = najdi_datum(datumi)
#    datumi.append(["10", "27"])
#    print(datumi)
#    while datum in datumi:
#        print("loop")
#        datum = najdi_datum(datumi)
#    
#    if datum in datumi:
#        return datum - timedelta(1)
#    return datum
#
#def sprejmi_argumente():
#    stevilo_delavcev = int(input("Vnesi: "))
#    stevilo_delavcev = int(input())
#    if stevilo_delavcev == 0:
#        return ["10-28"]
#    datumi = []
#    for element in range(0, stevilo_delavcev):
#        datumi.append(input("Vnesi rojstni dan delavca st. " + str(element + 1) + ": ").split()[1])
#        datumi.append(input().split()[1])
#    return datumi

#datumi = []
#stevilo_delavcev = 0
#for i in sys.stdin:
#    i = i.strip()
#    if len(i) == len(i.replace('-', "",1)):
#        stevilo_delavcev = int(i)
#        continue
#    datumi.append(i.split()[1])
#if stevilo_delavcev == 0:
#    print("10-28")
#else:
#    novi_datumi = pripravi_datume(datumi)
#    najden_datum = najdi_datum(novi_datumi)
#    idealni_datum = najden_datum.strftime("%m-%d")
#    print(idealni_datum)


    

    