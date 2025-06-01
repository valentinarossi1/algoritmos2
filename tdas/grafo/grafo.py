import random

class Grafo:
    def __init__(self, es_dirigido=False):
        self.vertices={}
        self.dirigido=es_dirigido
    
    def agregar_vertice(self, v):
        if v not in self.vertices:
            self.vertices[v]={}

    def eliminar_vertice(self, v):
        if v not in self.vertices:
            return # error vertice no encontrado
        for w in self.vertices[v]:
            self.vertices[w].pop(v)
        self.vertices.pop(v)
            

    def agregar_arista(self, v, w, peso=1):
        if v not in self.vertices or w not in self.vertices:
            return # error vertice no encontrado
        self.vertices[v][w] = peso
        if self.dirigido == False:
            self.vertices[w][v] = peso

    def eliminar_arista(self, v, w):    
        if v not in self.vertices or w not in self.vertices:
            return #error no son vertices
        if w not in self.vertices[v]:
            return  #error arista no encontrada
        self.vertices[v].pop(w)
        if self.dirigido==False:
            self.vertices[w].pop(v)

    def estan_unidos(self, v, w):
        if v not in self.vertices or w not in self.vertices:
            return #error
        return w in self.vertices[v]
    
    def peso_arista(self, v, w):
        if v in self.vertices and w in self.vertices[v]:
            return self.vertices[v][w]
        return None    
    
    def obtener_vertices(self):
        lista_vertices=[]
        for v in self.vertices:
            lista_vertices.append(v)
        return lista_vertices
    
    def adyacentes(self, v):
        lista_adyacentes=[]
        for w in self.vertices[v]:
            lista_adyacentes.append(w)
        return lista_adyacentes
    
    def vertice_aleatorio(self):
        vertices = self.obtener_vertices()
        return random.choice(vertices)


def main():
    print("Bienvenido al programa de prueba de la clase Grafo.")
    print("Creando un grafo no dirigido...")
    grafo = Grafo(es_dirigido=False)
    
    while True:
        print("\nOpciones:")
        print("1. Agregar vértice")
        print("2. Eliminar vértice")
        print("3. Agregar arista")
        print("4. Eliminar arista")
        print("5. Verificar si dos vértices están unidos")
        print("6. Obtener el peso de una arista")
        print("7. Listar todos los vértices")
        print("8. Listar adyacentes de un vértice")
        print("9. Seleccionar un vértice aleatorio")
        print("0. Salir")
        
        opcion = input("\nSelecciona una opción: ")

        if opcion == "1":
            v = input("Introduce el nombre del vértice: ")
            grafo.agregar_vertice(v)
            print(f"Vértice '{v}' agregado.")
        
        elif opcion == "2":
            v = input("Introduce el nombre del vértice a eliminar: ")
            grafo.eliminar_vertice(v)
            print(f"Vértice '{v}' eliminado (si existía).")
        
        elif opcion == "3":
            v = input("Introduce el primer vértice: ")
            w = input("Introduce el segundo vértice: ")
            peso = input("Introduce el peso de la arista (por defecto es 1): ")
            if peso.isdigit():
                peso = int(peso)
            else:
                peso = 1
            grafo.agregar_arista(v, w, peso)
            print(f"Arista entre '{v}' y '{w}' con peso {peso} agregada.")
        
        elif opcion == "4":
            v = input("Introduce el primer vértice: ")
            w = input("Introduce el segundo vértice: ")
            grafo.eliminar_arista(v, w)
            print(f"Arista entre '{v}' y '{w}' eliminada (si existía).")
        
        elif opcion == "5":
            v = input("Introduce el primer vértice: ")
            w = input("Introduce el segundo vértice: ")
            if grafo.estan_unidos(v, w):
                print(f"Los vértices '{v}' y '{w}' están unidos.")
            else:
                print(f"Los vértices '{v}' y '{w}' no están unidos.")
        
        elif opcion == "6":
            v = input("Introduce el primer vértice: ")
            w = input("Introduce el segundo vértice: ")
            peso = grafo.peso_arista(v, w)
            if peso is not None:
                print(f"El peso de la arista entre '{v}' y '{w}' es {peso}.")
            else:
                print(f"No existe una arista entre '{v}' y '{w}'.")
        
        elif opcion == "7":
            vertices = grafo.obtener_vertices()
            print("Lista de vértices:", vertices)
        
        elif opcion == "8":
            v = input("Introduce el nombre del vértice: ")
            adyacentes = grafo.adyacentes(v)
            print(f"Vértices adyacentes a '{v}':", adyacentes)
        
        elif opcion == "9":
            try:
                aleatorio = grafo.vertice_aleatorio()
                print(f"Vértice aleatorio seleccionado: {aleatorio}")
            except ValueError as e:
                print(e)
        
        elif opcion == "0":
            print("Saliendo del programa...")
            break
        
        else:
            print("Opción no válida. Por favor, selecciona una opción válida.")

if __name__ == "__main__":
    main()