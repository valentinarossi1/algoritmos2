INICIO_USUARIO = "usuario:"
APARECE_EN_PLAYLIST = " --> aparece en playlist --> "
DE_USUARIO = "de --> "
TIENE_PLAYLIST = " --> tiene una playlist --> "
DONDE_APARECE = " --> donde aparece"
SEPARADOR_CAMINO = " --> "
TIPO_USUARIOS = "usuarios"
INICIO_CICLO = 0

def imprimir_camino_minimo(camino, grafo):
  """esta funcion imprime el camino mas corto entre dos canciones, teniendo en cuenta las playlists en 
  las que aparecen y los usuarios intermedios"""
  for i in range(len(camino)):
    if i != len(camino) - 1:
      playlist = grafo.peso_arista(camino[i], camino[i + 1])
      if not camino[i].startswith(INICIO_USUARIO):
        print(f'{camino[i]}{APARECE_EN_PLAYLIST}{playlist}', end = SEPARADOR_CAMINO)
      else:
        usuario = camino[i].replace(INICIO_USUARIO, "")
        print(f'{DE_USUARIO}{usuario}{TIENE_PLAYLIST}{playlist}{DONDE_APARECE}', end = SEPARADOR_CAMINO)
    else:
      print(f'{camino[i]}', end = "\n")

def imprimir_mas_importantes(lista_ordenada_pageranks, n):
  """esta funcion imprime las 'n' canciones mas importantes segun el algoritmo de PageRank"""
  lista = lista_ordenada_pageranks[:n] 
  for i in range(n - 1):
    print(f'{lista[i][0]}; ', end = "")
  print(f'{lista[n - 1][0]}')

def imprimir_recomendaciones(recomendaciones, n, tipo):
    """esta funcione imprime las canciones/usuarios a recomendar en base a las canciones recibidas"""
    n = int(n)
    if tipo == TIPO_USUARIOS:
      recomendaciones = [(recomendacion[0].replace(INICIO_USUARIO, ""), recomendacion[1]) for recomendacion in recomendaciones]
    lista = recomendaciones[:n]
    for i in range(n - 1):
      print(f'{lista[i][0]}; ', end="")
    print(f'{lista[n-1][0]}')

def imprimir_ciclo(ciclo):
  """esta funcion imprime un ciclo de largo 'n' que comienza en la cancion indicada"""
  print(ciclo[INICIO_CICLO], end = "")
  for i in range(1, len(ciclo)):
    print(f'{SEPARADOR_CAMINO}{ciclo[i]}', end = "")
  return
    