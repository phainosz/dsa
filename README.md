# Data Structures and Algorithms

- [Introdução](#introdução)
- [Time complexity](#time-complexity)
- [Big O](#big-o)
- [Algoritimos Bitwise](#bitwise-algorithms)
- [Array](#array)
- [Searching Algorithms](#searching-algorithms)
- [Sorting Algorithms](#sorting-algorithms)

## Introdução
- Estrutura de dados e algoritimos são fundamentais em ciência da computação, nos ajudam a organizar e processar dados de forma eficiente.
- São usados para resolver problemas e desafios em software, desde gerenciar grandes volumes de dados até otimizar a velocidade de atividades.
- Em DSA, o foco principal é resolver problemas de forma eficiente e eficaz. Para determinar a eficiência do algoritmo, olhamos para dois tipos de complexidade:
  - **Time complexity**, tempo demorado para execução do algoritmo.
  - **Space complexity**, quantidade de memória o algoritmo usa.

## Time complexity
- Quantifica o tempo total gasto por um algoritmo executar em função do tamanho da entrada de dados, não o tempo gasto da máquina que o executou.
- A maneira mais fácil para medir complexidade seria olhando por **loops** dentro do código.
- Sempre medir o pior cenário possível ao usar **Big O**.
- [Big O Cheat Sheet](https://www.bigocheatsheet.com/).

## Big O
- É uma forma de categorizar o algoritmo em tempo e espaço gastos baseando-se na entrada de dados.
- Não é usado como uma medida exata, e sim uma forma de generalizar o crescimento e complexidade gasto no algoritmo.
- Auxilia em tomadas de deciões sobre qual algoritmo ou extrutura de dados usar.

#### O(1)
- Algoritmo com tempo constante.
- Não depende do tamanho da entrada de dados, executa apenas uma vez.
- Verificar o valor de uma posição do **array** usando indice é constante.

#### O(n)
- Algoritimo tem crescimento linear.
- Depende do tamanho da entrada de dado.
- Percorrer um **array** é O(n).

#### O(log n)
- Algoritimo com tempo logaritimico.
- Quando o tamanho da estrutura de dado diminui em cada passo executado.
- **Binary search** é um exemplo deste caso. A estrutura é dividade em pedaços menores.

#### O(n^2)
- Algoritimo com tempo elevado ao quadrado.
- A performance e diretamente relacionada com o tamanho da estrutura ao quadrado.
- Loops aninhados podem ser um exemplo para este caso.

#### O(2^n)
- Algoritimo com tempo exponencial.
- Resolução do algoritmo de fibbonaci usando recursão é um exemplo para tempo exponencial.

#### Regras de simplificação
- *Drop constants* é quando tempo algo como *O(n + n)*, simplificamos para *O(n)*.
  - Dois loops em sequência pode ser um exemplo disso.
- *Drop Non-Dominants* é quando tempo algo como *O(n^2 +n)*, simplificamos para *O(n^2)*.
  - Loops aninhados crescem mais rápidos que um loop simples, com isso o tempo maior seria do loop aninhado, por isso ignoramos o outro.
    
## Algoritimos Bitwise
- São algoritimos para fazer operação individual em **bits**.
- Estes algoritimos que manipulam a representação de números binários.
- Geralmente usados em sistemas de baixo nível, como criptografia, otimizações de tarefas onde manipular os *bits* de forma individual é vantajoso.
- Existem várias formas de operações **bitwise**, as mais conhecidas são: *AND(&)*, *OR(|)*, *XOR(^)*, *NOT(~)*, *left shifting(<<)* e *right shifting(>>)*.
- **Bit shifting** troca o **bit** de posição usando um valor de entrada como quantidade de *bits* a serem trocados.
  - **Operador AND (&)**
    - Faz a comparação de dois parâmetros, se o valor do bit comparado na posição for 1 em ambos, o retorno é 1, senão, será 0.
    - Ex: `111 & 100` os bits resultantes serão `100`, pois só a primeira posição teve os dois parâmetros o bit 1. Se fosse `111 & 001` o resultado seria `001`.
    - `0 & 0` = 0
    - `0 & 1` = 0
    - `1 & 0` = 0
    - `1 & 1` = 1
  - **Operador OR (|)**
    - Faz a comparação de dois parâmetros, se o valor do bit comparado na posição for 0 em ambos, o retorno será 0, senão, será 1.
    - Ex: `111 | 100` os bits resultantes serão `111`, pois nenhuma posição comparada teve os bits em ambos sendo 0. Se fosse `110 | 010` o resultado seria `110`.
    - `0 | 0` = 1
    - `0 | 1` = 1
    - `1 | 0` = 1
    - `1 | 1` = 0
  - **Operador XOR (^)**
    - Chamado de ou exclusivo. Se o bit na posição comparada não for igual, o retorno é 1, no caso, opostos se tornam 1 e iguais se tornam 0.
    - Ex: `111 ^ 100` os bits resultantes serão `011`, pois quando são opostos o bit comparado vira 1. Se fosse `110 ^ 010` o resultado seria `100`.
    - `0 ^ 0` = 0
    - `0 ^ 1` = 1
    - `1 ^ 0` = 1
    - `1 ^ 1` = 0
  - **Operador NOT (~)**
    - Diferente dos outros operadores, não necessita de dois parâmetros para realizar a operação.
    - Faz a troca dos bits do valor sendo operado.
    - Ex: `100` resultaria nos bits `011`
  - **Left shifting**
    - Usa dois operadores, `a<<b`, a é o valor a sofrer a mudança nos bits, b é a quantidade de *bits* que queremos trocar.
    - Pode ser representado matematicamente como: **a\*(2^b)**.
    - Nunca usar números negativos como parâmetro, terá um comportamento inesperado.
    - Cuidade ao usar parâmetros maiores que a quantidade suportada pelo tipo da variável usada. Por exemplo: `a<<33` a sendo armazenado em variável de 32 bits.
    <details>
    <summary>Ex:</summary>

    ```go
    func main() {
      // a = 5(00000101), b = 9(00001001)
      a, b := 5, 9

      fmt.Printf("a value: %d\n", a)
      fmt.Printf("a in binary: %08b\n", a)
      fmt.Printf("a in binary left shifting: %08b\n", a<<1)
      fmt.Printf("a value after left shifting: %d\n\n", a<<1)

      fmt.Printf("b value: %d\n", b)
      fmt.Printf("b in binary: %08b\n", b)
      fmt.Printf("b in binary left shifting: %08b\n", b<<1)
      fmt.Printf("b value after left shifting: %d\n", b<<1)
      
      //outputs
      //a value: 5
      //a in binary: 101
      //a in binary left shifting: 1010
      //a value after left shifting: 10

      //b value: 9
      //b in binary: 1001
      //b in binary left shifting: 10010
      //b value after left shifting: 18
    }
    ```
    </details>
  - **Right shifting**

## Array

## Searching Algorithms
### Linear Search
### Binary Search

## Sorting Algorithms
### Bubble Sort

## Recursion

## Linked List

## Queue

## Stack

## Hash

## Tree

## Graph