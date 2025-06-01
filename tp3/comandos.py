import funciones
import auxiliares

MENSAJE_NO_RECORRIDO = "No se encontro recorrido"
MENSAJE_CANCIONES = "Tanto el origen como el destino deben ser canciones"
MENSAJE_NO_PERTENECE = "La cancion no pertenece"
TIPO_CANCIONES = "canciones"
TIPO_USUARIOS = "usuarios"
INICIO_USUARIO = "usuario:"

def comando_camino(origen, destino, grafo):
  """este comando imprime una lista con la cual se conecta (en la menor cantidad de pasos posibles) una canción con 
  otra, considerando los usuarios intermedios y las listas de reproducción en las que aparecen"""
  if origen not in grafo.obtener_vertices() or destino not in grafo.obtener_vertices():
    print(MENSAJE_CANCIONES)
    return
  if origen.startswith(INICIO_USUARIO) or destino.startswith(INICIO_USUARIO):
    print(MENSAJE_CANCIONES)
    return
  padres = funciones.caminos_minimos_bfs(grafo, origen, destino)
  if destino not in padres:
    print(MENSAJE_NO_RECORRIDO)
    return
  camino = funciones.reconstruir_camino(padres, destino)
  if not camino:
    print(MENSAJE_NO_RECORRIDO)
    return
  auxiliares.imprimir_camino_minimo(camino, grafo)

def comando_mas_importantes(cantidad, pageranks):
  """este comando muestra las 'n' canciones más centrales/importantes del mundo según el algoritmo de PageRank, 
  ordenadas de mayor importancia a menor importancia"""
  cantidad = int(cantidad)
  canciones = {}
  for v, rank in pageranks.items():
    if not v.startswith(INICIO_USUARIO):
      canciones[v] = rank
  canciones_ordenadas = sorted(canciones.items(), key = lambda x: x[1], reverse = True)
  auxiliares.imprimir_mas_importantes(canciones_ordenadas, cantidad)

def comando_recomendacion(grafo, tipo, cantidad, canciones):
  """este comando da una lista de 'n' usuarios o canciones para recomendar, dado el listado de canciones que 
  ya sabemos que le gustan a la persona a la cual recomedar"""
  recomendaciones = {}
  for cancion in canciones:
    funciones.pagerank_personalizado(grafo, cancion, recomendaciones)
  if tipo == TIPO_CANCIONES:
    lista = {}
    for v, rank in recomendaciones.items():
      if not v.startswith(INICIO_USUARIO):
        lista[v] = rank
  elif tipo == TIPO_USUARIOS:
    lista = {}
    for v, rank in recomendaciones.items():
      if v.startswith(INICIO_USUARIO):
        lista[v] = rank
  recomendaciones_ordenadas = sorted(lista.items(), key = lambda x: x[1], reverse = True)
  auxiliares.imprimir_recomendaciones(recomendaciones_ordenadas, cantidad, tipo)

def comando_ciclo(grafo,cantidad, cancion):
  """este comando permite obtener un ciclo de largo 'n' (dentro de la red de canciones) que comience en 
  la canción indicada"""
  if cancion not in grafo.obtener_vertices():
    print(MENSAJE_NO_PERTENECE)
    return
  ciclo = funciones.detectar_ciclo(grafo, cancion, cantidad) 
  if not ciclo : 
    print(MENSAJE_NO_RECORRIDO)
    return
  else:
    auxiliares.imprimir_ciclo(ciclo)

def comando_rango(grafo, cantidad, cancion):
  """este comando permite obtener la cantidad de canciones que se encuenten a exactamente 'n' saltos desde 
  la cancion pasada por parámetro"""
  cantidad = int(cantidad)
  if cancion not in grafo.obtener_vertices():
    print(MENSAJE_NO_PERTENECE)
    return 0
  contador = funciones.bfs(grafo, cantidad ,cancion)
  print(contador)