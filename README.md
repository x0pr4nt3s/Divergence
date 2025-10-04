# Divergence

Divergence es un generador de diccionarios optimizado en Go. A partir de palabras base y reglas descriptivas produce tres archivos orientados a ataques de diccionario o análisis de contraseñas:

- `special_words.txt`: variaciones amplias ("leet" y capitalización) de palabras comunes más las palabras especiales que proporciones.
- `top_special_words.txt`: versiones simples (original, capitalizada, mayúsculas) de las mismas palabras, pensadas para usarse como diccionario base.
- `resultado_wordlist.txt`: resultado final tras aplicar las reglas combinatorias sobre las listas anteriores.

El programa está diseñado para funcionar en streaming, limitar la memoria usada y evitar generar archivos inmanejables mediante parámetros de corte.

## Requisitos

- Go 1.22 o superior instalado y disponible en el `PATH`.

## Instalación rápida

```
go build -o divergence
```

También puedes ejecutarlo sin compilar explícitamente:

```
go run . -rule reglas.txt -specialword palabras.txt [opciones]
```

## Uso

```
./divergence -rule reglas.txt -specialword palabras.txt [opciones]
```

- `-rule, --rule` define la ruta al archivo con reglas (obligatorio).
- `-specialword, --specialword` apunta al archivo con palabras especiales (obligatorio).
- `-out` permite indicar un directorio de salida (por defecto, el actual).
- El programa escribe el banner de bienvenida a menos que utilices `--no-banner`.

Los archivos de entrada son texto plano, una palabra/regla por línea. El motor ya incluye un listado de palabras generales en español e inglés; tus palabras se suman a esa base.

### Archivos opcionales

- `--general` permite sumar un archivo con palabras generales adicionales; se combinan con las predefinidas.
- `--years` acepta un archivo con anios extra para extender la lista de anos predefinidos.
- `--special-chars` carga caracteres especiales adicionales para ampliar las combinaciones de `s`.


## Nomenclatura de reglas

Cada letra describe qué lista se incorpora en la posición correspondiente del resultado final:

```
w -> combinator word (variantes amplias generadas por camel/leet)
t -> top word (variantes básicas de capitalización)
y -> year (años predefinidos: 2040, 2035, 2030, 2025, 2023, ...)
n -> number (dígitos 0-9)
l -> letter (todas las letras en mayúscula y minúscula)
M -> mayúsculas A-Z
m -> minúsculas a-z
s -> special char (caracteres como !@#$%&*-_=, etc.)
```

### Ejemplos de reglas

```
tys   => <TopWord><Year><SpecialChar>    (Ej: Empresa2023$)
wyt   => <WordAvanzada><Year><TopWord>
tnnM => <TopWord><Num><Num><LetraMayus> (Ej: Empresa04E)
```

## Opciones avanzadas y límites de seguridad

Todas las opciones aceptan `0` para indicar "sin límite". Úsalas para adaptar la ejecución a la memoria disponible y al tamaño de los archivos que deseas producir.

- `--max-special <int>` limita la cantidad de variantes únicas que se escriben en `special_words.txt`. El colector detiene la generación cuando alcanza el valor indicado (`main.go:170`, `main.go:212`, `main.go:303`). Ideal si quieres un diccionario de semillas más acotado.

- `--max-top <int>` cumple la misma función para `top_special_words.txt` (`main.go:171`, `main.go:233`, `main.go:306`). Úsalo cuando sólo necesites un subconjunto representativo para ataques con reglas externas.

- `--max-rule-product <int>` restringe el número de combinaciones generadas por cada regla antes de detenerla (`main.go:172`, `main.go:280`, `main.go:309`). El valor por defecto es 250 000, lo que previene explosiones combinatorias; incrementa o reduce según tu hardware.

- `--max-result <int>` fija el máximo de líneas que se escribirán en `resultado_wordlist.txt` (`main.go:173`, `main.go:254`, `main.go:312`). Útil para ejecuciones de muestra o cuando sólo quieres producir un bloque manejable de combinaciones.

- `--allow-duplicates` desactiva la deduplicación en el archivo final (`main.go:175`, `main.go:254`, `main.go:405`). Con ello evitas mantener un conjunto en memoria y reduces uso de RAM, aunque los duplicados consumirán líneas del límite `--max-result` y pueden requerir filtrado posterior.

- `--no-banner` omite el arte ASCII de presentación para salidas más limpias en scripts (`main.go:174`).

## Ejemplo de ejecución controlada

```
./divergence \
  --rule reglas.txt \
  --specialword palabras.txt \
  --out ./salida \
  --max-special 5000 \
  --max-top 2000 \
  --max-rule-product 100000 \
  --max-result 250000 \
  --allow-duplicates
```

Este comando genera los archivos en `./salida`, corta `special_words.txt` tras 5000 variantes, `top_special_words.txt` tras 2000, limita cada regla a 100 000 combinaciones y detiene el resultado general a las primeras 250 000 líneas. Además, permite duplicados para ahorrar memoria.

## Consejos de uso

- Empieza usando reglas con `t` antes que `w` si buscas diccionarios base manejables. Las variantes avanzadas (`w`) combinan sustituciones "leet" y pueden multiplicar el tamaño rápidamente.
- Ajusta los límites en función de los recursos disponibles. Si notas uso intensivo de memoria, habilita `--allow-duplicates` o baja `--max-rule-product`.
- Los archivos generados pueden alimentarse directamente a herramientas como Hashcat, John the Ripper o pipelines de cracking personalizados.

## Ejemplo de archivos de entrada

`palabras.txt`
```
empresa
soporte
seguridad
```

`reglas.txt`
```
tys
wyt
tnnM
```

## Créditos

Proyecto original escrito en Python por x0pr4nt3s. Reimplementación optimizada en Go por Codex (GPT-5) para ofrecer un flujo más eficiente en memoria y configurable mediante flags.
