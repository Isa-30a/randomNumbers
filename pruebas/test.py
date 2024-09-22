import subprocess

def ejecutar_programa_y_dar_entrada(comando, entrada):
    # Ejecuta el programa dado y proporciona entrada estándar
    proceso = subprocess.Popen(comando, stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
    
    # Enviar la entrada al programa y capturar la salida
    salida, error = proceso.communicate(input=entrada)
    
    # Imprime la salida estándar
    if proceso.returncode == 0:
        print("Salida estándar del programa:\n", salida)
    else:
        print("Error del programa:\n", error)

# Ejemplo de uso
comando = ["./main"]  # Puedes cambiar "cat" por cualquier programa que desees ejecutar
entrada = "poisson;datos_a1;0.5;5"
ejecutar_programa_y_dar_entrada(comando, entrada)