from input_frame import InputFrame
from table import Table
from tkinter import filedialog
import csv

class App:
    def __init__(self, root):

        self.root = root
        self.root.title("Simulación de numeros aleatorios")
        self.root.geometry("600x400")
        
        self.input_frame = InputFrame(self.root, self.calcular, self.exportar_csv)
        self.table = Table(self.root)

    # Función para agregar datos a la tabla
    def calcular(self):
        self.table.limpiar_tabla()
        
        a = int(self.input_frame.a.get())
        c = int(self.input_frame.c.get())
        m = int(self.input_frame.m.get())
        seed = int(self.input_frame.seed.get())
        aux = seed
        option = self.input_frame.option.get()
        
        if(option == "multiplicativo"):
            xi = (a*aux)%m
            while xi != seed:
                xi = (a*aux)%m
                ri=xi/(m-1)
                self.table.tabla.insert("", "end", values=(aux,xi,ri))
                aux = xi
        else:
            xi = ((a*aux)+c)%m
            while xi != seed:
                xi = (a*aux+c)%m
                ri=xi/(m-1)
                self.table.tabla.insert("", "end", values=(aux, xi, ri))
                aux = xi
    
    def exportar_csv(self):
        file_path = filedialog.asksaveasfilename(defaultextension=".csv", filetypes=[("CSV files", "*.csv")])

        if not file_path:
            return  # Si el usuario cancela, no hacer nada

        items = self.table.tabla.get_children()

        with open(file_path, mode='w', newline='', encoding='utf-8') as file:
            writer = csv.writer(file, delimiter=",")  # Asegurar que el delimitador sea una coma
            # Escribir los encabezados

            # Escribir las filas
            for item in items:
                row = self.table.tabla.item(item)['values']
                writer.writerow(row)
