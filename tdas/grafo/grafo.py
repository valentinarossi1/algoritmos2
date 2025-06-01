class Grafo:
    def __init__(self, es_dirigido = False, lista_vertices = None):
        self.vertices = {}
        self.dirigido = es_dirigido
        if lista_vertices != None :
            for vertice in lista_vertices:
                self.agregar_vertice(vertice)
    
    def agregar_vertice(self, v):
        """agrega el vertice v al grafo si no existe"""
        if v not in self.vertices:
            self.vertices[v] = {}

    def eliminar_vertice(self, v):
        """elimina el vertice v del grafo y todas sus aristas, si existe"""
        if v not in self.vertices:
            return
        for w in self.vertices[v]:
            self.vertices[w].pop(v)
        self.vertices.pop(v)
            
    def agregar_arista(self, v, w, peso = 1):
        """agrega una arista en el grafo entre v y w con el peso recibido, si ambos vertices existen"""
        if v not in self.vertices or w not in self.vertices:
            return
        self.vertices[v][w] = peso
        if self.dirigido == False:
            self.vertices[w][v] = peso

    def eliminar_arista(self, v, w):
        """elimina la arista del grafo entre v y w, verificando que ambos vertices existan"""    
        if v not in self.vertices or w not in self.vertices:
            return 
        if w not in self.vertices[v]:
            return
        self.vertices[v].pop(w)
        if self.dirigido == False:
            self.vertices[w].pop(v)

    def estan_unidos(self, v, w):
        """devuelve True si los vertices v y w estan unidos, False si no lo estan"""
        if v not in self.vertices or w not in self.vertices:
            return
        return w in self.vertices[v]
    
    def peso_arista(self, v, w):
        """devuelve el peso de la arista entre los vertices v y w, si existe"""
        if v in self.vertices and w in self.vertices[v]:
            return self.vertices[v][w]
        return None    
    
    def obtener_vertices(self):
        """devuelve una lista con todos los vertices del grafo"""
        lista_vertices = []
        for v in self.vertices:
            lista_vertices.append(v)
        return lista_vertices
    
    def adyacentes(self, v):
        """devuelve una lista con todos los adyacentes al vertice v, si no existe devuelve una lista vacia"""
        if v not in self.vertices:
            return []
        lista_adyacentes = []
        for w in self.vertices[v]:
            lista_adyacentes.append(w)
        return lista_adyacentes