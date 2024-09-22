import tkinter as tk

class InputFrame:
    def __init__(self, root, calcular, exportar_csv):
        self.frame = tk.Frame(root)
        self.frame.pack(side=tk.TOP, fill=tk.X, padx=10, pady=10)

        tk.Label(self.frame, text="a:").grid(row=0, column=0, padx=5, pady=5)
        self.a = tk.Entry(self.frame)
        self.a.grid(row=0, column=1, padx=5, pady=5)

        tk.Label(self.frame, text="c:").grid(row=0, column=3, padx=5, pady=5)
        self.c = tk.Entry(self.frame)
        self.c.grid(row=0, column=4, padx=5, pady=5)

        tk.Label(self.frame, text="m:").grid(row=0, column=6, padx=5, pady=5)
        self.m = tk.Entry(self.frame)
        self.m.grid(row=0, column=7, padx=5, pady=5)

        tk.Label(self.frame, text="semilla:").grid(row=1, column=0, padx=5, pady=5)
        self.seed = tk.Entry(self.frame)
        self.seed.grid(row=1, column=1, padx=5, pady=5)
        
        self.option = tk.StringVar(value="lineal")

        # Crear radio buttons
        radio_frame = tk.Frame(self.frame)
        radio_frame.grid(row=2, column=0, columnspan=2, pady=10)

        tk.Radiobutton(radio_frame, text="lineal", variable=self.option, value="lineal").pack(side=tk.LEFT, padx=10)
        tk.Radiobutton(radio_frame, text="Multiplicativo", variable=self.option, value="multiplicativo").pack(side=tk.LEFT, padx=10)

        # Botón para agregar datos
        agregar_btn = tk.Button(self.frame, text="Calcular", command=calcular)
        agregar_btn.grid(row=5, column=0, columnspan=1, pady=10)
        
        exportar_btn = tk.Button(self.frame, text="Exportar a csv", command=exportar_csv)
        exportar_btn.grid(row=5, column=1, columnspan=1, pady=10)