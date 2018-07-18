package algorithm

import "fmt"

func makeAdjMatrix(n int) [][]int {
	M := make([][]int, n)
	A := make([]int, n*n)
	fmt.Printf("%T, %v\n", M, M)
	for i := range M {
		M[i] = A[0:n]
		A = A[n:]
	}
	return M
}

type status struct {
	in   bool
	dist int
}

func dijkstra(start int, G [][]int) ([]status, []int) {
	from := make([]int, len(G))
	for i := range from {
		from[i] = -1
	}
	D := make([]status, len(G))
	for i := range D {
		D[i].in = false
		D[i].dist = 1<<32 - 1
	}

	D[start] = status{true, 0}
	x := start
	for i := 1; i < len(G); i++ {
		adj := G[x]
		for v, w := range adj {
			if w > 0 {
				if D[x].dist+w < D[v].dist {
					D[v].dist = D[x].dist + w
				}
			}
		}
		min := int(1<<32 - 1)
		nextx := -1
		for j := range D {
			if !D[j].in && D[j].dist < min {
				nextx = j
				min = D[j].dist
			}
		}
		D[nextx] = status{true, min}
		from[nextx] = x
		x = nextx
	}

	return D, from
}

func showPath(from []int, v int) {
	P := make([]byte, 0, len(from))
	for v != -1 {
		P = append(P, byte('0'+v))
		v = from[v]
	}
	for i := len(P) - 1; i >= 0; i-- {
		fmt.Printf("%c ", P[i])
	}
	fmt.Println("")
}

func TestDijkstra() {
	M := makeAdjMatrix(5)
	M[0][1] = 3
	M[0][2] = 8
	M[0][3] = 5
	M[1][2] = 4
	M[1][3] = 1
	M[2][4] = 10
	M[3][2] = 1
	for i := range M {
		fmt.Println(i, M[i])
	}

	status, from := dijkstra(0, M)
	fmt.Println(status)
	fmt.Println(from)

	showPath(from, 3)
}
