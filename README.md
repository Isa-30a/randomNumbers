# RandomNumbers
Esta aplicación genera numeros aleatorios por medio de generadores lineales:
 * Congruencial mixto
 * Congruencial multiplicativo

# Features:
1. Genera numeros aleatorios
2. Exporta los numeros en archivo .csv
3. Realiza pruebas estadisticas:
    * Media
    * Varianza
    * Chi-cuadrado -> Intervalos
    * Poker
    * Kolmogorov-Smirnof
    * Corridas arriba y abajo
4. Realiza variables aleatorias con distribución uniforme
    * Poisson
    * Bernoulli

# ¿Como funciona?

## Requisitos
1. Tener instalado python con tkinter
2. Tener instalado golang en el sistema

## Ejecutar el archivo 
1. Ejecuta el archivo main.py que se encuentra en la carpeta "Simulación"
    ``` python3 main.py ```
2. Empieza a utilizar la aplicación

## En la app...
Cuando ejecutes la aplicación correctamente te encontraras con esta pantalla. Esta te permite elegir los valores a, c, m y semilla (X0) para la generación de los numeros aleatorios. Ten en cuenta que si deseas utilizar el generador lineal congruencial multiplicativo, el valor de c no se utilizara, por lo tanto es opcional ingresarlo.
![imagen](https://github.com/user-attachments/assets/c34a2ab6-f288-4d46-a1f9-48949c0304e4)

Despues de colocar los numeros y darle click en calcular en la tabla se colocaran 3 columnas correspondientes a los numeros generados en Xi-1  (empezara siempre con el valor de la semilla que colocaste), Xi (empezara con el primer numero pseudoaleatorio generado a partir de los valores para la formula del generador elegido) y ri que sera el numero pseudoaleatorio entre 0 y 1 el cual se ha obtenido de la formula ``ri=xi/(m-1)`` siendo m = a la cantidad de numeros generados.
![imagen](https://github.com/user-attachments/assets/984f8022-f63a-4b47-9380-7a2fcbdac449)

### Exporta tus datos
Al darle click en el boton "Exportar a csv" estos numeros pueden ser exportados en dicho formato para poder ser utilizados en las pruebas estadisticas. Deberas ingresar un nombre con el que guardar el archivo .csv que se creara en la carpeta "exported_csv".
![imagen](https://github.com/user-attachments/assets/833880a4-aad6-42bf-98d8-ce1ff0883ad4)

### Usar variables o pruebas estadisticas
Para usar las variables o pruebas estadisticas deberas elegir primero cual deseas aplicar. Esto se logra dando click en el boton que por defecto muestra Bernoulli, alli te apareceran todas las opciones que hay.
![imagen](https://github.com/user-attachments/assets/72dcc374-d485-473a-815b-9745655208bf)

Cuando elijas la opción deseada deberas darle en el boton probar el cual desplegara una nueva ventana que te mostrara toda la información obtenida de la prueba.
Algunas pruebas piden valores adicionales. Ten en cuenta que cuando el programa te pida un decimal o flotante deberas ingresarlo con un *.* y no con una *,*
![imagen](https://github.com/user-attachments/assets/85ca9865-afa4-4df9-b615-498922e78d41)

 
