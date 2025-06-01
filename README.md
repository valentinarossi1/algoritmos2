# Algoritmos y Estructuras de Datos - FIUBA (Cátedra Buchwald)

Repositorio con los TDAs y TPs desarrollados en la materia.

---

### TDAS:
- Pila
- Cola
- Lista
- ABB
- Diccionario
- Cola de prioridad

---

### TP1 - Calculadora Polaca Inversa  
Programa que evalúa expresiones matemáticas en notación postfija.

---

### TP2 - Análisis de logs
Programa que procesa archivos de logs de un servidor web para extraer y mostrar estadísticas sobre visitantes y recursos solicitados.

Comandos implementados:

- `agregar_archivo <nombre_archivo>`: procesa un archivo de log para cargar datos al sistema.  
- `ver_visitantes <desde> <hasta>`: muestra todas las IPs que solicitaron algún recurso en el servidor, dentro del rango de IPs determinado.  
- `ver_mas_visitados <n>`: muestra los n recursos más solicitados.

---

### TP3 - Recomendify
Programa que modela un sistema de recomendación basado en un grafo que conecta usuarios, playlists y canciones en una plataforma de streaming.

Comandos implementados:

- `camino <origen> >>>> <destino>`: imprime una lista con la cual se conecta (en la menor cantidad de pasos posibles) una canción con otra, considerando los usuarios intermedios y las listas de reproducción en las que aparecen. 
- `mas_importantes <n>`: muestra las n canciones más centrales del mundo según el algoritmo de pagerank, ordenadas de mayor a menor importancia.
- `recomendacion <usuarios/canciones> <n> <cancion1 >>>> cancion2 >>>> ... >>>> cancionK>`: imprime una lista de n usuarios o canciones para recomendar, a partir el listado de canciones que ya sabemos que le gustan a la persona a la cual recomendar.
- `ciclo <n> <cancion>`: permite obtener un ciclo de largo n (dentro de la red de canciones) que comience en la canción indicada.  
- `rango <n> <cancion>`: cuenta la cantidad de canciones que están a exactamente n saltos desde la canción indicada.
