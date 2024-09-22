
from tkinter import ttk
import tkinter as tk


class Table:
    def __init__(self, root):
        
        tabla_frame = tk.Frame(root)
        tabla_frame.pack(side=tk.TOP, fill=tk.BOTH, expand=True, padx=15, pady=10)

        self.tabla = ttk.Treeview(tabla_frame, columns=("Xᵢ₋₁", "Xᵢ", "rᵢ"), show="headings")
        self.tabla.heading("Xᵢ₋₁", text="Xᵢ₋₁")
        self.tabla.heading("Xᵢ", text="Xᵢ")
        self.tabla.heading("rᵢ", text="rᵢ")

        self.tabla.pack(fill=tk.BOTH, expand=True)

        scrollbar = ttk.Scrollbar(tabla_frame, orient=tk.VERTICAL, command=self.tabla.yview)
        self.tabla.configure(yscrollcommand=scrollbar.set)
        scrollbar.pack(side=tk.RIGHT, fill=tk.Y)


    def limpiar_tabla(self):
        for item in self.tabla.get_children():
            self.tabla.delete(item)
            
    