#!/usr/bin/env python3

import argparse
import sys
from itertools import product

ascii_art = """
　　∩.　　　∩.　　　∩.　　　∩.　　　∩.　　　∩.　　　∩.　　　∩.　　　　
  ノ×乂   ノ×乂   ノ×乂 　ノ×乂 　ノ×乂　 ノ×乂　 ノ×乂　 ノ×乂 　　
  |0┓爻|　|０爻|　|┏┓爻|　|┃┃爻|　|┏┓爻|　|┏┓爻|　|┏┓爻|　|┏┓爻| 　　
  |○┃爻|　|爻爻|　|┃┃爻|　|┗╋爻|　|┣┫爻|　|┗┓爻|　|┗┫爻|　|┣┓爻| 　　
  |0┃圭|　| * #|　|┗┛圭|　|/┃圭|　|┗┛圭|　|┗┛圭|　|┗┛圭|　|┗┛圭| 　　
  |圭圭|　|圭圭|　|圭圭|　|圭圭|　|圭圭|　|圭圭|　|圭圭|　|圭圭| 　　
 /|卅卅|廾|卅卅|廾|卅卅|廾|卅卅|廾|卅卅|廾|卅卅|廾|卅卅|廾|卅卅|且¯/|
|￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣￣||
|　　　　　　　　　- Wellcome to the Divergence -　　　　　　　　　||
|＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿＿|/
"""

general_words = ["lima","santiago","bogota","enero","febrero","marzo","abril","mayo","junio","julio","agosto","setiembre","octubre","noviembre","diciembre","january","february","march","april","may","june","july","august","september","october","november","december","verano","invierno","otoño","primavera","sistemas","enterprise","soporte"]



### Leyendo archivo
def lectura_special_words(archivo_special_words):
    # Obtener el nombre del archivo de los argumentos de línea de comandos
    archivo = archivo_special_words
    lineas = []  # Lista para almacenar las líneas del archivo
    # Leer el archivo línea por línea
    with open(archivo, 'r') as f:
        for linea in f:
            linea = linea.strip()  # Eliminar espacios en blanco al inicio y final de la línea
            lineas.append(linea)  # Agregar la línea a la lista

    return lineas

def lectura_rules(archivo_rules):
    # Obtener el nombre del archivo de los argumentos de línea de comandos
    archivo = archivo_rules
    lineas = []  # Lista para almacenar las líneas del archivo
    # Leer el archivo línea por línea
    with open(archivo, 'r') as f:
        for linea in f:
            linea = linea.strip()  # Eliminar espacios en blanco al inicio y final de la línea
            lineas.append(linea)  # Agregar la línea a la lista

    return lineas


all_s_chars = ['!', '@', '#', '$', '%', '&', '*', '(', ')', '-', '_', '+', '=', '[', ']', '{', '}', '|', '~', '`', ':', ';', '<', '>', ',', '.', '?', '/']
top_s_chars = ['!', '@', '#', '$', '%', '&', '*','-', '_', '+', '=','|', '~',':', ';', ',', '.', '?']

top_years = ['2040','2035','2030','2025','2023', '2022', '2021', '2020', '2019', '2018', '2017', '2016', '2015', '2014', '2013', '2012', '2011', '2010', '2009', '2008', '2007', '2006', '2005', '2004','2003','2002','2001','2000']
letras_minusculas = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z']
letras_mayusculas = ['A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z']
global_letters=letras_minusculas+letras_mayusculas


def camello(word):
    
    # Define las reglas de sustitución
    rules = {
        'a': ['a', 'A', '4', '@'],
        'e': ['e', 'E', '3'],
        'i': ['i', 'I', '1', '!', 'L'],
        'l': ['l', 'L', '1', '!', 'I'],
        'o': ['o', 'O', '0'],
        's': ['s', 'S', '5', '$']
    }

    # Obtén todas las letras de la palabra
    letters = [rules.get(char, [char]) for char in word]

    # Genera todas las combinaciones posibles de letras
    combinations = [''.join(prod) for prod in product(*letters)]

    # Agrega combinaciones en mayúsculas y minúsculas
    combinations.extend([combo.upper() for combo in combinations])
    combinations.extend([combo.lower() for combo in combinations])

    # Agrega la palabra original con la primera letra en mayúscula
    combinations.append(word.capitalize())

    # Genera combinaciones para la palabra con la primera letra en mayúscula
    capitalized_word = word.capitalize()
    capitalized_letters = [rules.get(char, [char]) for char in capitalized_word]
    capitalized_combinations = [''.join(prod) for prod in product(*capitalized_letters)]
    combinations.extend(capitalized_combinations)

    # Elimina duplicados y ordena la lista
    combinations = sorted(list(set(combinations)))

    return combinations


def generar_numeros():
    numeros = []
    for numero in range(10):
        numeros.append(str(numero))
    return numeros

def obtener_producto_cartesiano(matriz):
    return [''.join(combo) for combo in product(*matriz)]


def make_words(lista_special_words):
    wordlist_words=[]
    for i in general_words:
        wordlist_words=wordlist_words+camello(i)

    for i in lista_special_words:
        wordlist_words=wordlist_words+camello(i)

    wordlist_words=sorted(list(set(wordlist_words)))

    return wordlist_words

def make_top_words(lista_special_words):
    wordlist_words=[]
    for i in general_words:
        wordlist_words=wordlist_words+generate_top_word(i)

    for i in lista_special_words:
        wordlist_words=wordlist_words+generate_top_word(i)

    wordlist_words=sorted(list(set(wordlist_words)))

    return wordlist_words



def generate_list_special_words():
    archivo = "special_words.txt"
    list_sw = []  # Lista para almacenar las líneas del archivo
    # Leer el archivo línea por línea
    with open(archivo, 'r') as f:
        for linea in f:
            linea = linea.strip()  # Eliminar espacios en blanco al inicio y final de la línea
            list_sw.append(linea)  # Agregar la línea a la lista

    return list_sw    

def generate_list_top_special_words():
    archivo = "top_special_words.txt"
    list_sw = []  # Lista para almacenar las líneas del archivo
    # Leer el archivo línea por línea
    with open(archivo, 'r') as f:
        for linea in f:
            linea = linea.strip()  # Eliminar espacios en blanco al inicio y final de la línea
            list_sw.append(linea)  # Agregar la línea a la lista

    return list_sw  


def generate_top_word(palabra):
    combinaciones = [
        palabra,          # Original
        palabra.capitalize(),  # Primera letra en mayúscula
        palabra.upper()   # Todas las letras en mayúscula
    ]
    # Generar y mostrar combinaciones
    return combinaciones


def analyze_word_for_rule(word):
    list_special_words=generate_list_special_words()
    list_top_specialw=generate_list_top_special_words()
    lista_de_listas=[]
    for char in word:
        if char=='w':
            lista_de_listas.append(list_special_words)
        elif char=='y':
            lista_de_listas.append(top_years)
        elif char=='n':
            lista_de_listas.append(generar_numeros)
        elif char=='l':
            lista_de_listas.append(global_letters)
        elif char=='M':
            lista_de_listas.append(letras_mayusculas)
        elif char=='m':
            lista_de_listas.append(letras_minusculas)
        elif char=='s':
            lista_de_listas.append(all_s_chars)
        elif char=='t':
            lista_de_listas.append(list_top_specialw)
        else:
            print("Hubo un problema en el procesamiento de las reglas coloca las reglas de nuevo")

    lista_final=obtener_producto_cartesiano(lista_de_listas)

    return lista_final


def main():

    # PARSER

    descripcion_parser="Procesa argumentos -r y -sw"


    parser = MyParser(description=descripcion_parser)
    


    ### Define los argumentos
    parser.add_argument('-r', '--rule', type=str, help='Archivo con reglas')
    parser.add_argument('-sw', '--specialword', type=str, help='Archivo con palabras clave')


    args = parser.parse_args()

    if args.rule is not None and args.specialword is not None:
            ### Realiza alguna acción con los argumentos
        print(f'Reglas: {args.rule}')
        print(f'Palabras Especiales: {args.specialword}')
    else:
        print('Se requieren tanto -r como -sw. Ejecuta el script con --help para obtener más información.')
    # Crea la primera wordlist con palabras generales y del archivo de palabras speciales
    init_wordlist=make_words(lectura_special_words(args.specialword))
    init_wordlist.sort()
    wordlist_sin_repetir = list(set(init_wordlist))

    init2_wordlist=make_top_words(lectura_special_words(args.specialword))
    init2_wordlist.sort()
    wordlist2_sin_repetir = list(set(init2_wordlist))

    print(ascii_art)
    # ESCRIBIENDO EN EL ARCHIVO FINAL

    nombre_archivo = "special_words.txt"
    nombre_archivo2 = "top_special_words.txt" 

    # Escribir el contenido de la lista en el archivo línea por línea
    with open(nombre_archivo, 'w') as archivo:
        for elemento in wordlist_sin_repetir:
            archivo.write(elemento + '\n')

    # Escribir el contenido de la lista en el archivo línea por línea
    with open(nombre_archivo2, 'w') as archivo:
        for elemento in wordlist2_sin_repetir:
            archivo.write(elemento + '\n')

    # ANALIZANDO REGLAS 
    init_rules=lectura_rules(args.rule)
    lista_wordlist_final=[]
    for i in init_rules:
        lista_wordlist_final=lista_wordlist_final+analyze_word_for_rule(i)       

    lista_wordlist_final.sort()
    lista_wordlist_final = list(set(lista_wordlist_final))

    nombre_archivo = "resultado_wordlist.txt" 

    # Escribir el contenido de la lista en el archivo línea por línea
    with open(nombre_archivo, 'w') as archivo:
        for elemento in lista_wordlist_final:
            archivo.write(elemento + '\n')
   
    print("")
    print("")
    print("RESULTADO ALMACENADO EN: resultado_wordlist.txt")

if __name__ == "__main__":
    # Llamar a la función main
    main()
