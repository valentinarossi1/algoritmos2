from collections import deque
import random

ITERACIONES_PAGERANK = 40
MAX_ITER = 30
COEF_AMORTIGUACION = 0.85
TOLERANCIA_PAGERANK = 0.0000001
LARGO_RANDOM_WALK = 10

def caminos_minimos_bfs(grafo, origen, destino):
  """recorrido BFS para determinar el camino mas corto entre dos canciones"""
  padres = {}
  visitados = set()
  cola = deque()

  padres[origen] = None
  visitados.add(origen)
  cola.append(origen)
  while cola:
    v = cola.popleft()
    if v == destino:
      break
    for w in grafo.adyacentes(v):
      if w not in visitados:
        padres[w] = v
        visitados.add(w)
        cola.append(w)
  return padres

def reconstruir_camino(padres, destino):
  """esta funcion reconstruye el camino hasta una cancion determinada"""
  camino = []
  while destino is not None:
    camino.append(destino)
    destino = padres[destino]
  return camino[::-1]

def pagerank(grafo):
  """algoritmo de PageRank para determinar la importancia de cada cancion en el grafo"""
  N = len(grafo.obtener_vertices())
  pageranks = {}
  adyacentes = {}
  grados = {}
  for v in grafo.obtener_vertices():
    pageranks[v] = 1.0/N
    adyacentes[v] = grafo.adyacentes(v)
    grados[v] = len(adyacentes[v])

  for _ in range(MAX_ITER):
    nuevo_pagerank = {}
    convergencia = True 
    for vertice in grafo.obtener_vertices():
      suma = 0
      for vecino in adyacentes[vertice]:
        if grados[vecino] > 0:  
          suma += pageranks[vecino] / grados[vecino]
      nuevo_pagerank[vertice] = (1 - COEF_AMORTIGUACION) / N + COEF_AMORTIGUACION * suma
      if abs(nuevo_pagerank[vertice] - pageranks[vertice]) > TOLERANCIA_PAGERANK:
        convergencia = False
    pageranks = nuevo_pagerank
    if convergencia:
      break
  return pageranks

def random_walk(grafo, inicio):
  """esta funcion realiza un random walk desde un vertice inicial, eligiendo aleatoriamente los vertices 
  adyacentes a visitar en cada paso"""
  camino = [inicio]
  vertice = inicio
  for _ in range(LARGO_RANDOM_WALK):
    adyacentes = grafo.adyacentes(vertice)
    if not adyacentes:
      break
    vertice = random.choice(adyacentes)
    camino.append(vertice)
  return camino

def pagerank_personalizado(grafo, cancion, recomendaciones):
  """esta función calcula recomendaciones personalizadas para para una cancion del grafo, segun la cantidad de 
  veces que fue visitada al realizar las iteraciones de random walk"""
  for _ in range(ITERACIONES_PAGERANK):
    camino = random_walk(grafo, cancion)
    for i in range(len(camino) - 1):
      vertice_origen = camino[i]
      vertice_destino = camino[i + 1]
      if vertice_destino not in recomendaciones:
        recomendaciones[vertice_destino] = 0
      recomendaciones[vertice_destino] += 1 / len(grafo.adyacentes(vertice_origen))

def dfs(grafo, origen, vertice, cantidad, ciclo, visitados):
  """recorrido DFS para detectar ciclos de una longitud especifica en un grafo, partiendo desde una cancion inicial"""
  visitados.add(vertice)
  if len(ciclo) == cantidad:
    if origen in grafo.adyacentes(vertice):
      ciclo.append(origen)
      return ciclo
    else:
      return None
  for w in grafo.adyacentes(vertice):
    if w in visitados: 
      continue
    ciclo.append(w)
    solucion = dfs(grafo, origen, w, cantidad, ciclo, visitados)
    if solucion is not None:
      return solucion
    else:
      return None
  visitados.remove(vertice)
  return None

def detectar_ciclo(grafo, origen, n):
  ciclo = []
  visitados = set()
  ciclo.append(origen)
  camino = dfs(grafo, origen, origen, n, ciclo,visitados)
  return camino


def bfs(grafo, cantidad, origen):
  """recorrido BFS para obtener la cantidad de canciones que se encuentran a determinada distancia desde la cancion inicial """
  contador = 0
  visitados = set()
  orden = {}
  cola = deque()
  
  orden[origen] = 0
  visitados.add(origen)
  cola.append(origen)
  if cantidad <= 0:
    return 0
  if cantidad == 1:
    return len(grafo.adyacentes(origen))
  while len(cola) > 0:
    desencolado = cola.popleft()
    for w in grafo.adyacentes(desencolado):
      if w not in visitados:
        visitados.add(w)
        orden[w] = orden[desencolado] + 1
        if orden[w] == cantidad:
          contador += 1
        if orden[w] < cantidad: 
          cola.append(w)          
  return contador