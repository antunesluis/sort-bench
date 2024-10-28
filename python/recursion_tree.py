import networkx as nx
import matplotlib.pyplot as plt
from typing import List, Tuple
import sys  # Para pegar argumentos da linha de comando


class MergeSortVisualizer:
    def __init__(self):
        self.G = nx.Graph()
        self.pos = {}
        self.node_count = 0
        self.levels = {}
        self.node_labels = {}

    def merge_sort_visualize(
        self, arr: List[int], level: int = 0, x_pos: float = 0
    ) -> Tuple[List[int], int]:
        n = len(arr)
        node_id = self.node_count
        self.node_count += 1

        # Adiciona o nó atual ao grafo
        self.G.add_node(node_id)
        self.node_labels[node_id] = str(arr)
        self.pos[node_id] = (x_pos, -level)
        self.levels[level] = self.levels.get(level, 0) + 1

        if n <= 1:
            return arr, node_id

        mid = n // 2
        left_arr = arr[:mid]
        right_arr = arr[mid:]

        # Calcula as posições x para os filhos
        left_x = x_pos - (0.8 / (level + 1))
        right_x = x_pos + (0.8 / (level + 1))

        # Recursivamente ordena e visualiza as sub-arrays
        left_sorted, left_node = self.merge_sort_visualize(left_arr, level + 1, left_x)
        right_sorted, right_node = self.merge_sort_visualize(
            right_arr, level + 1, right_x
        )

        # Adiciona arestas para conectar o nó pai aos filhos
        self.G.add_edge(node_id, left_node)
        self.G.add_edge(node_id, right_node)

        # Realiza o merge
        merged = self.merge(left_sorted, right_sorted)

        return merged, node_id

    def merge(self, left: List[int], right: List[int]) -> List[int]:
        result = []
        i = j = 0

        while i < len(left) and j < len(right):
            if left[i] <= right[j]:
                result.append(left[i])
                i += 1
            else:
                result.append(right[j])
                j += 1

        result.extend(left[i:])
        result.extend(right[j:])
        return result

    def visualize(self, arr: List[int], output_file: str = "merge_sort_tree.png"):
        # Limpa o estado anterior
        self.G.clear()
        self.pos.clear()
        self.node_count = 0
        self.levels.clear()
        self.node_labels.clear()

        # Executa o merge sort e constrói a visualização
        sorted_arr, _ = self.merge_sort_visualize(arr)

        # Configura o gráfico
        plt.figure(figsize=(12, 8))
        nx.draw(
            self.G,
            pos=self.pos,
            labels=self.node_labels,
            node_color="lightblue",
            node_size=2500,
            font_size=8,
            font_weight="bold",
            width=2,
            edge_color="gray",
        )

        plt.title("Árvore de Recursão do Merge Sort")
        plt.axis("off")

        # Salva a figura em um arquivo ao invés de mostrar
        plt.savefig(output_file, format="png", bbox_inches="tight", dpi=300)
        plt.close()  # Fecha a figura para liberar memória

        return sorted_arr


def read_numbers_from_file(file_path: str) -> List[int]:
    try:
        with open(file_path, "r") as f:
            # Converte cada linha em um número inteiro
            return [int(line.strip()) for line in f if line.strip()]
    except FileNotFoundError:
        print(f"Erro: Arquivo '{file_path}' não encontrado.")
        sys.exit(1)
    except ValueError as e:
        print(f"Erro ao converter dados: {e}")
        sys.exit(1)


def demo_merge_sort_visualization(input_file: str):
    arr = read_numbers_from_file(input_file)

    print(f"Array original: {arr}")

    visualizer = MergeSortVisualizer()
    sorted_arr = visualizer.visualize(arr, "merge_sort_visualization.png")

    print(f"Array ordenado: {sorted_arr}")
    print(f"Visualização salva em 'merge_sort_visualization.png'")


if __name__ == "__main__":
    # Verifica se o arquivo de entrada foi passado como argumento
    if len(sys.argv) < 2:
        print("Uso: python <script.py> <arquivo_de_entrada>")
        sys.exit(1)

    # Pega o nome do arquivo da linha de comando
    input_file = sys.argv[1]
    demo_merge_sort_visualization(input_file)
