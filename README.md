# Divergence

Divergence es una herramienta para generar diccionarios en base a una lista de palabras y archivo de reglas que indicara las combinaciones que tendra el diccionario.

EJEMPLO DE COMANDO:
```
python3 divergent.py -r rules.txt -sw palabras.txt
```
En este ejemplo "rules.txt" es el archivo donde iran las reglas
y "palabras.txt" es el archivo con las palabras especiales.
### NOTA 1: 
Este script ya viene con palabras comunes base que se combinaran y se generaran en el resultado final.

## Nomenclatura de Reglas
```
w -> combinator word (combinaciones avanzadas)
t -> top word (combinaciones base y simples)
y -> year (Top Years: 2022, 2023, 2021, etc.)
n -> number
l -> letter (Mayus y minus)
M -> mayus
m -> minus
s -> special char
```

## Demostracion de Reglas

```
tys : <PalabraTop><Year><Special char> : Empresa2022$
wys : <PalabraAvanzada><Year><Special char> : 3mpresA2022$
tnnM : <PalabraTop><Num><Num><Letter Mayus> : Empresa04E
```
## Ejemplo de Reglas
```
tys 
tys
tnnM
```

### NOTA 2: 
Lo mas recomendable es que utilize solo la letra "t" en cuanto a palabras dentro de las reglas y no la "w" ya que lo que se quiere es
conseguir un diccionario base no tan grande y luego utilizar las reglas avanzandas usando hashcat con el diccionario que obtenemos,
sin embargo este escript tambien cuenta con combinaciones mas complejas en cuanto a las palabras usando la letra "w" en su archivo 
de reglas.
