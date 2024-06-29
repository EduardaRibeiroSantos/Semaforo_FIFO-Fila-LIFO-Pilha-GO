package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Nodo struct {
	dado    string
	proximo *Nodo
}

type Fila struct {
	primeiro *Nodo
	ultimo   *Nodo
}

type Pilha struct {
	topo *Nodo
}

func NovaFila() *Fila {
	return &Fila{}
}

func NovaPilha() *Pilha {
	return &Pilha{}
}

// Insere um novo dado na fila
func (f *Fila) Insere(novoDado string) {
	novoNodo := &Nodo{dado: novoDado}
	if f.primeiro == nil {
		f.primeiro = novoNodo
		f.ultimo = novoNodo
	} else {
		f.ultimo.proximo = novoNodo
		f.ultimo = novoNodo
	}
}

// Remove um dado da fila
func (f *Fila) Remove() string {
	if f.primeiro == nil {
		fmt.Println("Fila já está vazia.")
		return ""
	}
	removido := f.primeiro.dado
	f.primeiro = f.primeiro.proximo
	if f.primeiro == nil {
		f.ultimo = nil
	}
	return removido
}

// Insere um novo dado na pilha
func (p *Pilha) Insere(novoDado string) {
	novoNodo := &Nodo{dado: novoDado, proximo: p.topo}
	p.topo = novoNodo
}

// Remove um dado da pilha
func (p *Pilha) Remove() string {
	if p.topo == nil {
		fmt.Println("Impossível remover valor de pilha vazia.")
		return ""
	}
	removido := p.topo.dado
	p.topo = p.topo.proximo
	return removido
}

// String retorna uma representação em string da fila
func (f *Fila) String() string {
	var result string
	for nodo := f.primeiro; nodo != nil; nodo = nodo.proximo {
		result += fmt.Sprintf("%s -> ", nodo.dado)
	}
	result += "nil"
	return result
}

// String retorna uma representação em string da pilha
func (p *Pilha) String() string {
	var result string
	for nodo := p.topo; nodo != nil; nodo = nodo.proximo {
		result += fmt.Sprintf("%s -> ", nodo.dado)
	}
	result += "nil"
	return result
}

func main() {

	pista1 := NovaFila()
	pista2 := NovaFila()

	// Distribuição de 10 carros de maneira aleatória entre as pistas 1 e 2
	for i := 1; i <= 10; i++ {
		x := rand.Intn(2) + 1
		if x == 1 {
			pista1.Insere(fmt.Sprintf("\033[1;96mcarro %d\033[0m", i))
			fmt.Printf("Carro %d entrando na pista %d: %s\n", i, x, pista1)
		} else {
			pista2.Insere(fmt.Sprintf("\033[1;95mcarro %d\033[0m", i))
			fmt.Printf("Carro %d entrando na pista %d: %s\n", i, x, pista2)
		}
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	}

	fmt.Println("\n")

	// Loop principal que controla os semáforos das pistas 1 e 2
	for {
		count := 0
		for count < 10 {
			if count < 5 {
				fmt.Println("\033[42;1m--green light on (Semáforo 1)--\033[0m")
				if pista1.primeiro != nil {
					time.Sleep(2 * time.Second)
					pista1.Remove()
					fmt.Printf("Carro saiu da pista 1: %s\n", pista1)
				}
			} else {
				fmt.Println("\033[41;1m--red light on (Semáforo 1)--\033[0m")
				break
			}
			time.Sleep(1 * time.Second)
			count++
		}

		count = 0
		for count < 10 {
			if count < 5 {
				fmt.Println("\033[42;1m--green light on (Semáforo 2)--\033[0m")
				if pista2.primeiro != nil {
					time.Sleep(2 * time.Second)
					pista2.Remove()
					fmt.Printf("Carro saiu da pista 2: %s\n", pista2)
				}
			} else {
				fmt.Println("\033[41;1m--red light on (Semáforo 2)--\033[0m")
				break
			}
			time.Sleep(1 * time.Second)
			count++
		}

		if pista1.primeiro == nil && pista2.primeiro == nil {
			fmt.Println("\033[1;33m--Todos os carros foram em direção ao estacionamento--\033[0m\n")
			break
		}
	}

	estacionamento := NovaPilha()

	for i := 1; i <= 5; i++ {
		estacionamento.Insere(fmt.Sprintf("\033[1;34mcarro %d\033[0m", i))
		fmt.Printf("Carro %d entrando no estacionamento: %s\n", i, estacionamento)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("\n")

	fmt.Println("\033[1;33m--O estacionamento está lotado.--\033[0m")
	fmt.Println(estacionamento)

	// Escolhe um carro aleatório para ser removido do estacionamento
	y := rand.Intn(5) + 1
	fmt.Printf("O carro a ser removido é o carro %d\n", y)

	// Remove carros da pilha e devolve-os para a pista 1 de acordo com o valor aleatório gerado
	switch y {
	case 1:
		fmt.Println("Os carros a serem removidos antes do carro 1 são [carro2 -> carro3 -> carro4 -> carro5].")
		for i := 5; i >= 1; i-- {
			estacionamento.Remove()
			time.Sleep(2 * time.Second)
			pista1.Insere(fmt.Sprintf("\033[1;96mcarro %d\033[0m", i))
			fmt.Printf("Carro %d voltou para a pista 1: %s\n", i, pista1)
		}
	case 2:
		fmt.Println("Os carros a serem removidos antes do carro 2 são [carro3 -> carro4 -> carro5].")
		for i := 5; i >= 2; i-- {
			estacionamento.Remove()
			time.Sleep(2 * time.Second)
			pista1.Insere(fmt.Sprintf("\033[1;96mcarro %d\033[0m", i))
			fmt.Printf("Carro %d voltou para a pista 1: %s\n", i, pista1)
		}
	case 3:
		fmt.Println("Os carros a serem removidos antes do carro 3 são [carro4 -> carro5].")
		for i := 5; i >= 3; i-- {
			estacionamento.Remove()
			time.Sleep(2 * time.Second)
			pista1.Insere(fmt.Sprintf("\033[1;96mcarro %d\033[0m", i))
			fmt.Printf("Carro %d voltou para a pista 1: %s\n", i, pista1)
		}
	case 4:
		fmt.Println("O carro a ser removido antes do carro 4 é o [carro5].")
		for i := 5; i >= 4; i-- {
			estacionamento.Remove()
			time.Sleep(2 * time.Second)
			pista1.Insere(fmt.Sprintf("\033[1;96mcarro %d\033[0m", i))
			fmt.Printf("Carro %d voltou para a pista 1: %s\n", i, pista1)
		}
	case 5:
		fmt.Println("O carro 5 é o último então pode sair direto.")
		estacionamento.Remove()
		time.Sleep(2 * time.Second)
		pista1.Insere(fmt.Sprintf("\033[1;96mcarro %d\033[0m", 5))
		fmt.Printf("Carro %d voltou para a pista 1: %s\n", 5, pista1)
	}
}
