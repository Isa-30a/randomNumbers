from input_frame import InputFrame
from table import Table
from tkinter import END, filedialog, messagebox, simpledialog, Toplevel, Text
import csv
import subprocess
import platform

class App:
    def __init__(self, root):

        self.root = root
        self.root.title("Simulación de numeros aleatorios")
        self.root.geometry("600x400")
        
        self.input_frame = InputFrame(self.root, self.calcular, self.exportar_csv, self.probar)
        self.table = Table(self.root)

    # Función para agregar datos a la tabla
    def calcular(self):
        self.table.limpiar_tabla()
        
        a = int(self.input_frame.a.get())
        m = int(self.input_frame.m.get())
        seed = int(self.input_frame.seed.get())
        aux = seed
        option = self.input_frame.option.get()
        
        if(option == "multiplicativo"):
            xi = (a*aux)%m
            numbers = {}    
            while xi != seed :
                xi = (a*aux)%m
                if xi in numbers:
                    break
                ri=xi/(m-1)
                self.table.tabla.insert("", "end", values=(aux,xi,ri))
                numbers[xi] = True
                aux = xi
        else:
            c = int(self.input_frame.c.get())
            xi = ((a*aux)+c)%m
            numbers = {}
            while xi != seed:
                xi = (a*aux+c)%m 
                if xi in numbers:
                    break
                ri=xi/(m-1)
                self.table.tabla.insert("", "end", values=(aux, xi, ri))
                numbers[xi] = True
                aux = xi
    
    def exportar_csv(self):
        file_path = "../exported_csv/" + simpledialog.askstring('Ventana', "Ingrese el nombre del archivo")+ ".csv"

        if not file_path:
            return  # Si el usuario cancela, no hacer nada

        items = self.table.tabla.get_children()

        with open(file_path, mode='w', newline='', encoding='utf-8') as file:
            writer = csv.writer(file, delimiter=",")  # Asegurar que el delimitador sea una coma
            # Escribir los encabezados

            # Escribir las filas
            for item in items:
                row = self.table.tabla.item(item)['values']
                column_value = row [2]
                writer.writerow([column_value])

    def probar(self):
        #[opcion prueba];csv;
        opcion = self.input_frame.pruebas.get()
        #"CORRIDAS", "MEDIA", "VARIANZA", , "KOLGOMOROV", "POISSON", "POKER"
        #if(opcion == "BERNOULLI" or opcion == "BERNOULLI" or opcion == )
        extra = extra2 =""
        texto = opcion+";"+self.cargar_csv()+";"
        
        match opcion:
            case "BERNOULLI":
                extra = simpledialog.askfloat('ventana',"La probabilidad de exito (decimal)")
            case "POISSON":
                extra = simpledialog.askfloat('Ventana', "Ingrese la media")
                extra2 = simpledialog.askinteger('Ventana', "Ingrese el numero de intervalos")
                
            case "INTERVALOS":
                extra = simpledialog.askinteger('Ventana', "Ingrese el numero de intervalos")
            case _:
                pass
        
        texto += str(extra)+";"+str(extra2)
        
        print(texto)
        sistema_operativo = platform.system()
        if sistema_operativo == "Windows":
            comando = "main"
        else:
            comando = "./main"
        resultado = self.ejecutar_programa_y_dar_entrada(comando, texto)
        self.abrir_dialogo(resultado)
        
    
    def cargar_csv(self):
        # Establecer la carpeta inicial del filedialog
        initial_dir = "../exported_csv"
        file_path = filedialog.askopenfilename(initialdir=initial_dir, filetypes=[("CSV files", "*.csv")])
#sfile_path = filedialog.askopenfilename(filetypes=[("CSV files", "*.csv")])

        if not file_path:
            return  # Si el usuario cancela, no hacer nada
        
        return "../exported_csv/"+file_path.split("/")[-1]
    
    def ejecutar_programa_y_dar_entrada(self, comando, entrada):
        # Ejecuta el programa dado y proporciona entrada estándar
        proceso = subprocess.Popen(comando, stdin=subprocess.PIPE, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
        
        # Enviar la entrada al programa y capturar la salida
        salida, error = proceso.communicate(input=entrada)
        print(salida)
        # Imprime la salida estándar
        if proceso.returncode == 0:
            return salida
        else:
            return error
        
    def abrir_dialogo(self, resultado):
        
        # Crear un cuadro de diálogo (ventana secundaria)
        dialogo = Toplevel(self.root) # Título del cuadro de diálogo
        dialogo.geometry("600x400")  # Tamaño del diálogo

        # Crear un área de texto dentro del diálogo
        text_area = Text(dialogo, wrap='word')
        text_area.pack(expand=True, fill='both')
        
        text_area.insert(END, resultado)
