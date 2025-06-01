#!/usr/bin/python3

import csv
import sys
from grafo import Grafo
from auxiliares import *
import funciones
import comandos

POS_COMANDO = 0
POS_ORIGEN = 0
POS_DESTINO = 1
FIN_COMANDO = 7
POS_N_RECOMENDACION = 2
POS_TIPO_RECOMENDACION = 1
POS_N = 1
CAMINO = "camino"
MAS_IMPORTANTES = "mas_importantes"
RECOMENDACION = "recomendacion"
CANT_ESPACIOS_RECOMENDACION = 3
CICLO = "ciclo"
RANGO = "rango"
CANT_ESPACIOS_CICLO = 2
CANT_ESPACIOS_RANGO = 2
SEPARADOR_CANCIONES = " >>>> "

def cargar_usuarios_canciones(grafo, usuario, cancion, artista, playlist):
  """esta funcion carga un usuario y una cancion en el grafo, cuya relacion indica que la cancion se 
  encuentra en una playlist del usuario. el peso de la arista es el nombre de la playlist"""
  nombre_cancion = cancion + " - " + artista
  nombre_usuario = "usuario:" + usuario
  grafo.agregar_vertice(nombre_usuario)
  grafo.agregar_vertice(nombre_cancion)
  grafo.agregar_arista(nombre_cancion, nombre_usuario, playlist)

def crear_grafo_usuarios(ruta):
  """esta funcion devuelve un grafo bipartito de usuarios y canciones"""
  grafo_bipartito = Grafo()
  with open(ruta, 'r') as archivo:
    lector = csv.reader(archivo, delimiter = '\t')
    lista_archivo = list(lector)
    for _, usuario, cancion, artista, _, playlist, _ in lista_archivo[1:]:
      cargar_usuarios_canciones(grafo_bipartito, usuario, cancion, artista, playlist)
  return grafo_bipartito

def cargar_canciones_por_usuario(canciones_por_usuario, cancion, artista, usuario):
  """ esta funcion carga una cancion a la lista de canciones de cada usuario si dicha cancion aparece 
  en alguna playlist del usuario"""
  nombre_cancion = cancion + " - " + artista
  if usuario not in canciones_por_usuario:
    canciones_por_usuario[usuario] = []
  canciones_por_usuario[usuario].append(nombre_cancion)

def cargar_diccionario_canciones(grafo, canciones_por_usuario):
  """esta funcion carga canciones al grafo y conecta las que estan en playlists del mismo usuario.
  el peso de la arista es el nombre del usuario"""
  for usuario in canciones_por_usuario:
    for cancion in canciones_por_usuario[usuario]:
      grafo.agregar_vertice(cancion)
      for otra_cancion in canciones_por_usuario[usuario]:
        if cancion == otra_cancion:
          continue
        grafo.agregar_vertice(otra_cancion)
        if not grafo.estan_unidos(cancion, otra_cancion):
          grafo.agregar_arista(cancion,otra_cancion,usuario)

def crear_grafo_canciones(ruta):
  """esta funcion devuelve un grafo de canciones"""
  grafo_canciones = Grafo()
  canciones_por_usuario = {}
  with open(ruta, 'r') as archivo:
    lector = csv.reader(archivo, delimiter = '\t')
    lista_archivo = list(lector)
    for _, usuario, cancion, artista, _, _, _ in lista_archivo[1:]:
      cargar_canciones_por_usuario(canciones_por_usuario, cancion, artista,usuario)
    cargar_diccionario_canciones(grafo_canciones, canciones_por_usuario)
  return grafo_canciones

def identificar_comando(comando, grafo_bipartito, grafo_canciones, pageranks, ruta):
  """esta funcion recibe un comando y lo identifica, llamando a su funcion correspondiente
  si el comando no existe, imprime por pantalla 'Comando no reconocido'"""
  entrada = comando.split(" ")
  if entrada[POS_COMANDO] == CAMINO:
    parametros = comando[FIN_COMANDO:].split(SEPARADOR_CANCIONES)
    comandos.comando_camino(parametros[POS_ORIGEN], parametros[POS_DESTINO], grafo_bipartito)

  elif entrada[POS_COMANDO] == MAS_IMPORTANTES:
    if len(pageranks) == 0:
      pageranks = funciones.pagerank(grafo_bipartito)
    comandos.comando_mas_importantes(entrada[POS_N], pageranks)

  elif entrada[POS_COMANDO] == RECOMENDACION:
    parametros = comando[len(entrada[POS_COMANDO]) + len(entrada[POS_TIPO_RECOMENDACION]) + len(entrada[POS_N_RECOMENDACION]) + CANT_ESPACIOS_RECOMENDACION:].split(SEPARADOR_CANCIONES)
    comandos.comando_recomendacion(grafo_bipartito, entrada[POS_TIPO_RECOMENDACION], entrada[POS_N_RECOMENDACION], parametros)

  elif entrada[POS_COMANDO] == CICLO:
    if len(grafo_canciones) == 0:
      grafo_canciones = crear_grafo_canciones(ruta)
    cantidad = entrada[POS_N]
    cantidad = int(cantidad)
    cancion = comando[len(entrada[POS_COMANDO]) + len(entrada[POS_N]) + CANT_ESPACIOS_CICLO:]
    comandos.comando_ciclo(grafo_canciones, cantidad, cancion)

  elif entrada[POS_COMANDO] == RANGO:
    if len(grafo_canciones) == 0:
      grafo_canciones = crear_grafo_canciones(ruta)
    cantidad = entrada[POS_N]
    cancion = comando[len(entrada[POS_COMANDO]) + len(entrada[POS_N]) + CANT_ESPACIOS_RANGO:]
    comandos.comando_rango(grafo_canciones, cantidad, cancion)
  
  else:
    print("Comando no reconocido")

def main():
    ruta = sys.argv[1]
    grafo_canciones = {}
    pageranks = {}
    grafo_bipartito = crear_grafo_usuarios(ruta)
    entrada = sys.stdin.readline().strip("\n")
    while entrada != "":
      identificar_comando(entrada, grafo_bipartito, grafo_canciones, pageranks, ruta)
      entrada = sys.stdin.readline().strip("\n")
 
if __name__ == "__main__":
    main()