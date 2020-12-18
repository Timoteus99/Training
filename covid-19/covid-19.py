import matplotlib.pyplot as plt
import pandas as pd
import matplotlib
import glob, os

# uporaben link: https://datatofish.com/use-pandas-to-calculate-stats-from-an-imported-csv-file/
# datasets link: https://github.com/CSSEGISandData/COVID-19

def dnevna_statistika(csvfile):
    '''vrne SLOVAR okuzenih, umrlih, ozdravelih in smrtnost'''
    df = pd.read_csv(csvfile, error_bad_lines=False)
    uradno_okuzeni_svet = df['Confirmed'].sum()
    umrli = df['Deaths'].sum()
    okrevani = df['Recovered'].sum()
    svetovna_smrtnost = (umrli / uradno_okuzeni_svet) * 100
    dnevna_stat = {"Okuzeni" : uradno_okuzeni_svet, "Umrli" : umrli, "Okrevani" : okrevani, "Smrtnost" : round(svetovna_smrtnost, 5), "Datum" : csvfile[-14 : -4]}
    return dnevna_stat

def tekoca_svetovna_statistika():
    '''vrne seznam slovarjev z podatki dnevnih situacij'''
    mycsvdir = '/home/legion/Desktop/COVID-19-master/csse_covid_19_data/csse_covid_19_daily_reports/' # pot do csv datotek
    csvfiles = glob.glob(os.path.join(mycsvdir, '*.csv')) #dostopaj do vseh csv datotek (predpostavim da je koncnica .csv)
    dataframes = []  #seznam slovarjev
    for csvfile in csvfiles:
        dataframes.append(dnevna_statistika(csvfile))
    situacija = sorted(dataframes, key = lambda x:x["Datum"]) #slovarji so med sabo POMESANI (ne vem zakaj)
    return situacija #seznam slovarjev

# dela! dodamo se datume, in preprosto sortiramo po datumih namesto po st okuzenih.
# Tako lahko učinkovito analiziramo podatke tudi ce je prirastek negativen.

#zrisem se preprost graf z krivuljami po dnevih 
#prikazem krivuljo rasti okuzenih, umrlih in okrevanih

def izrisi_situacijo():
    '''prvo pretvori v sezname nato plotira in vrne graf'''
    dnevi, seznam_okuzenih, seznam_umrlih, seznam_okrevanih = [], [], [], [] #filam iz tekoce_svetovne_statistike() v sezname
    situacija = tekoca_svetovna_statistika()
    index = 0
    while index < len(situacija):
        seznam_okuzenih.append(int(situacija[index]['Okuzeni'])) 
        seznam_umrlih.append(int(situacija[index]['Umrli']))
        seznam_okrevanih.append(int(situacija[index]['Okrevani']))
        dnevi.append(index)
        index += 1
    fig, ax = plt.subplots()
    ax.plot(dnevi, seznam_okuzenih, "-b", label = 'Okuženi')
    ax.plot(dnevi, seznam_okrevanih, "-g", label = 'Okrevani')
    ax.plot(dnevi, seznam_umrlih, "-r",label = 'Umrli')
    ax.set(xlabel = 'Čas [dnevi]', ylabel = 'Število ljudi', title = 'COVID-19 - spremljanje svetovne situacije')
    ax.grid()
    plt.legend(loc="upper left")
    plt.ylim(0, seznam_okuzenih[-1])
    fig.savefig("graf.png")
    return plt.show()
print(izrisi_situacijo())

