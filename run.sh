#!/bin/bash

# Entrar a la carpeta del ejercicio
cd "./$1" || exit

# Compilar el archivo TypeScript (asumiendo que solo hay un .ts en la carpeta)
tsc *.ts

# Ejecutar el archivo generado en JavaScript
node *.js