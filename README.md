# Data Structures and Algorithms

- [Introdução](#introdução)
- [Time complexity](#time-complexity)
- [Big O](#big-o)
- [Bitwise Algorithms](#bitwise-algorithms)
- [Array](#array)
- [Two Pointer Technique](#two-pointer-technique)
- [Recusion](#recursion)
- [Linked Lists](#linked-list)
- [Queue](#queue)
- [Stack](#stack)
- [Hash](#hash)
- [Searching Algorithms](#searching-algorithms)
- [Sorting Algorithms](#sorting-algorithms)

## Introdução

Estruturas de dados e algoritmos são fundamentais na ciência da computação, pois nos ajudam a organizar e processar dados de forma eficiente.

São utilizados para resolver problemas em software, desde o gerenciamento de grandes volumes de dados até a otimização da performance de operações.

Em DSA (*Data Structures and Algorithms*), o foco principal é resolver problemas de forma eficiente e escalável. Para avaliar a eficiência de um algoritmo, analisamos dois tipos de complexidade:

- **Time complexity**: tempo de execução do algoritmo em relação ao tamanho da entrada.
- **Space complexity**: quantidade de memória utilizada pelo algoritmo durante a execução.

As estruturas de dados podem ser divididas basicamente em dois tipos: **lineares** e **não lineares**:
- **Lineares**: os elementos são organizados em sequência, um após o outro.
  - Exemplos: arrays, queues (filas) e stacks (pilhas).
  - Cada elemento (exceto o primeiro e o último) possui um predecessor e um sucessor.
- **Não lineares**: os elementos não seguem uma sequência linear, sendo organizados de forma hierárquica ou em rede.
  - Um elemento pode estar conectado a um ou mais elementos.
  - São divididas principalmente em dois grupos:
    - **Trees (árvores)**: estrutura hierárquica com um elemento raiz.
    - **Graphs (grafos)**: estrutura mais flexível, onde os elementos podem ter múltiplas conexões sem hierarquia fixa.

## Time complexity
Quantifica o tempo de execução de um algoritmo em função do tamanho da entrada de dados, e não o tempo real gasto pela máquina que o executa.

A análise é feita de forma abstrata, considerando o crescimento do algoritmo conforme a entrada aumenta, ignorando fatores como hardware, linguagem ou otimizações específicas.

Uma forma comum de estimar a complexidade é observar estruturas como loops e chamadas recursivas:
- Um loop simples geralmente indica complexidade linear (**O(n)**)
- Loops aninhados podem indicar complexidade quadrática (**O(n²)**)
- Divisão do problema em partes menores pode indicar complexidade logarítmica (**O(log n)**)

Ao utilizar **Big O**, normalmente analisamos o **pior caso (worst-case)**, ou seja, o cenário em que o algoritmo terá o maior custo possível.
- Isso garante previsibilidade de desempenho
- Existem também outros casos:
  - **Best-case**: melhor cenário possível
  - **Average-case**: comportamento médio do algoritmo

[Big O Cheat Sheet](https://www.bigocheatsheet.com/).

## Big O
É uma forma de categorizar um algoritmo com base no tempo de execução e no espaço de memória utilizados, em função do tamanho da entrada de dados.

Não é usado como uma medida exata, e sim uma forma de generalizar o crescimento e a complexidade do algoritmo.
Auxilia na tomada de decisões sobre qual algoritmo ou estrutura de dados utilizar.

### O(1)
- Algoritmo com tempo constante.
- Não depende do tamanho da entrada de dados.
- A operação é executada em tempo fixo.
- Ex: acessar um elemento de um **array** pelo índice é O(1).

### O(n)
- Algoritmo com crescimento linear.
- Depende diretamente do tamanho da entrada de dados.
- Ex: percorrer todos os elementos de um **array** é O(n).

### O(log n)
- Algoritmo com tempo logarítmico.
- O problema é reduzido a cada passo (geralmente pela metade).
- Ex: **binary search** (busca binária), onde a estrutura é dividida em partes menores a cada iteração.

### O(n^2)
- Algoritmo com crescimento quadrático.
- O tempo de execução cresce proporcionalmente ao quadrado da entrada.
- Ex: loops aninhados que percorrem a mesma estrutura.

### O(2^n)
- Algoritmo com crescimento exponencial.
- O número de operações dobra a cada novo elemento da entrada.
- Ex: cálculo de Fibonacci usando recursão sem otimização (forma ingênua).

### Regras de simplificação
- *Drop constants* é quando tempo algo como *O(n + n)*, simplificamos para *O(n)*.
  - Dois loops em sequência pode ser um exemplo disso.
- *Drop Non-Dominants* é quando tempo algo como *O(n^2 +n)*, simplificamos para *O(n^2)*.
  - Loops aninhados crescem mais rápidos que um loop simples, com isso o tempo maior seria do loop aninhado, por isso ignoramos o outro.
    
## Bitwise Algorithms

São algoritimos para fazer operação individual em **bits**.

Estes algoritimos que manipulam a representação de números binários.

Geralmente usados em sistemas de baixo nível, como criptografia, otimizações de tarefas onde manipular os *bits* de forma individual é vantajoso.

Existem várias formas de operações **bitwise**, as mais conhecidas são: *AND(&)*, *OR(|)*, *XOR(^)*, *NOT(~)*, *left shifting(<<)* e *right shifting(>>)*.

- **Bit shifting** desloca os bits de um valor para a esquerda ou direita, usando um número que define quantas posições serão movidas.
  - **Operador AND (&)**
    - Faz a comparação de dois operandos, se o valor do bit comparado na posição for 1 em ambos, o retorno é 1, senão, será 0.
    - Ex: `111 & 100` os bits resultantes serão `100`, pois só a primeira posição teve os dois operandos com bit 1. Se fosse `111 & 001` o resultado seria `001`.
    - `0 & 0` = 0
    - `0 & 1` = 0
    - `1 & 0` = 0
    - `1 & 1` = 1
  - **Operador OR (|)**
    - Faz a comparação de dois operandos, se o valor do bit comparado na posição for 0 em ambos, o retorno será 0, senão, será 1.
    - Ex: `111 | 100` os bits resultantes serão `111`, pois nenhuma posição comparada teve os bits em ambos sendo 0. Se fosse `110 | 010` o resultado seria `110`.
    - `0 | 0` = 0
    - `0 | 1` = 1
    - `1 | 0` = 1
    - `1 | 1` = 1
  - **Operador XOR (^)**
    - Chamado de ou exclusivo. Se o bit na posição comparada não for igual, o retorno é 1, no caso, opostos se tornam 1 e iguais se tornam 0.
    - Ex: `111 ^ 100` os bits resultantes serão `011`, pois quando são opostos o bit comparado vira 1. Se fosse `110 ^ 010` o resultado seria `100`.
    - `0 ^ 0` = 0
    - `0 ^ 1` = 1
    - `1 ^ 0` = 1
    - `1 ^ 1` = 0
  - **Operador NOT (~)**
    - Diferente dos outros operadores, não necessita de dois operandos para realizar a operação.
    - Faz a inversão de todos os bits do valor sendo operado.
    - Ex: `100` resultaria nos bits `011` (considerando apenas esses 3 bits).
    - `~1` = 0
    - `~0` = 1
  - **Left shifting**
    - Usa dois operandos, `a<<b`, onde `a` é o valor a sofrer a mudança nos bits e `b` é a quantidade de *bits* que queremos deslocar.
    - Os bits serão movidos para a esquerda conforme a quantidade de posições que indicamos.
    - Evite usar números negativos como parâmetro, pois o comportamento depende da linguagem e pode gerar erro ou resultado inesperado.
    - Cuidado ao usar valores maiores que a quantidade de bits suportada pelo tipo da variável. Por exemplo: `a<<33` com `a` sendo armazenado em uma variável de 32 bits.
    - Ex:
      - [Go](./golang/1-bitwise/bitwise.go)
      - [Java](./java/1-bitwise/Bitwise.java)
  - **Right shifting**
    - Semelhante ao **left shifting**, a diferença é que o deslocamento ocorre para a direita.
    - `a>>b` segue as mesmas regras do **left shifting**.

## Array

**Arrays** são estruturas de dados lineares que guardam elementos do mesmo tipo de dado.

São alocados de forma contígua em memória, o que significa, é um espaço de memória com o tamanho da estutura, onde cada elemento está alocado sequencialmente em um indice do espaço alocado.

Cada item é indexado e começa no **index** 0 na maioria das linguagens. Podemos acessar um elemento diretamente através do seu índice.

Ex:
  - [Go](./golang/2-array/array.go)
  - [Java](./java/2-array/Main.java)

## Two Pointer Technique

É uma ténica muito usada para resolver problemas usando **arrays** ou **strings**.

Usa dois *ponteiros* que percorrem os elementos, podendo ser de direções opostas(inicio e fim) ou mesma direção com passos diferentes, dependendo do problema.

Exemplos:
  ```
  // reverse array
  function reverse(arr)
    left = 0
    right = arr length -1      

    while left < right 
      // swap elements at left and right pointers
      tmp = arr[left]
      arr[left] = arr[right]
      arr[right] = tmp

      left++ // move left pointer to the right
      right-- // move right pointer to the left
    
  function isPalindrome(str):
    // str: string to check for palindrome
    left = 0
    right = arr length -1

    while left < right
      if str[left] != str[right] // if characters at the pointers do not match
          return false // return false (not a palindrome)
      left++ // move left pointer to the right
      right-- // move right pointer to the left

    return true
  ```

## Recursion

É uma técnica quando uma **função** chama ela mesmo.

Geralmente usada para resolver um problema que pode ser quebrado em problemas menores nos quais são resolvidos *recursivamente*.

A cada chamada da função, será criada uma pilha com o item atual no topo, esse processo continua até o caso base ser atingido, que é a condição que para a **recursão**. Quando o caso base ocorre, a função volta o ciclo seguindo a sequência da pilha de chamadas e retornando os seus resultados.

Geralmente temos 3 partes em **recursão**:
  - Caso base: que é a condição que para a **recursão**.
  - Caso recursivo; que é quando a função faz a chamada de sí mesma com um pedaço menor do problema.
  - Retorno/resultado: é quando a recursão termina e combina o resultado da pilha de execução.

Ex:
    ```
    function fibonacci(n)
      // base case: if n is 0 or 1, return n
      if n == 0
      return n
      else if n == 1
      return 1

      // recursive case: sum of two previous fibonacci numbers
      return fibonacci(n -1) + fibonacci(n -2)
    ```
  - [Go](./golang/3-recursion/recursion.go)
  - [Java](./java/3-recursion/Main.java)

## Linked List

São estruturas de dados lineares que armazenam os dados em nós que são conectados por ponteiros.

Diferentemente de **arrays**, não são armazenados de forma contígua em memória.

Podem ser facilmente alterada no seu tamanho, adicionando ou removendo nós.

Acessados de forma sequencial, através do inicio(head).

Tipos mais usados de **linked lists**:
  - **Singly linked list** cada nó aponta para o próximo item.
  - **Doubly linked list** cada nós aponta para o próximo nó e para o anterior.
  - **Cirular linked list** o último nó aponta para o primeiro nó, formando um circulo.

Ex:
  - [Go](./golang/4-linkedlist/linkedlist.go)
  - [Java](./java/4-linkedlist/Main.java)

## Queue

**Queue** ou fila, é uma estrutura de dados, usa o conceito **FIFO(first in first out)**.

O primeiro elemento inserido será o primeiro a sair da estrutura.

Ex:
  - [Go](./golang/5-queue/queue.go)
  - [Java](./java/5-queue/Main.java)

## Stack

**Stack** ou pilha, é uma estrutura de dados, usa o conceito **LIFO(last in first out)**.

Ex:
  - [GO](./golang/6-stack/stack.go)
  - [Java](./java/6-stack/Main.java)

## Hash

## Searching Algorithms

São essenciais para encontrar itens especificos dentro de uma coleção de dados.

Os algoritimos mais comuns são: **linear search** e **binary search**.

### Linear Search
É denifido com um algoritimo sequencial, que começa no primeiro elemento e percorre a cada elemento atẽ encontrar o elemento especifico, ou chegar ao final da estrutura.

Pode ser quebrados em passos a serem seguidos:
  - *Passo 1* Ponto de partida no primeiro elemento da coleção de dados.
  - *Passo 2* Comparar cada elemento se coincide com o elemento desejado.
  - *Passo 3* Encontrado o item desejado, interromper a busca e retorna o *index* do item ou true.
  - *Passo 4* Mover para o próximo elemento caso não encontrar o item na posição atual.
  - *Passo 5* Repetir até o final da coleção.
  - *Passo 6* Não encontrado o item, retornar informando que o item não foi encontrado.

Ex:
  ```
  function linearSearch(array, target)
    // target is the value we're searching for

    for element in array
      if element == target
        return element index
    
    return -1
  ```
  - [GO](./golang/7-linearsearch/linearsearch.go)
  - [Java](./java/7-linearsearch/Main.java)

### Binary Search
É um algoritmo de busca usado em *arrays ordenados*, caso contrário não irã funcionar da forma esperada.

Usa o principio de **divide and conquer (dividir e conquistar)**, divide o o array pela metade repetidamente até encontrar o elemento buscado ou não encontrá-lo.

Busca o elemento do meio, compara com o elemento buscado e decide qual lado após a divisão irá ser usado, faz isso repetidamente.

O passo a passo é o seguinte:
  - *Passo 1* encontrar o elemento do meio e comparar com o elemento buscado, se o item do meio ser o item buscado, retornar a posição do elemento do meio.
  - *Passo 2* se o elemento do meio não for igual, checar se ele mé maior ou menor que o elemento do meio.
  - *Passo 3* se maior, seguir para a busca no lado direito, senão, buscar no lado direito.
  - *Passo 4* repetir *passos 1, 2 e 3* até o tamanho do array final ser 1.
  - *Passo 5* se não encontrar o elemento buscado no array, retornar informando que o item não foi encontrado.

Ex:
  ```
  function binarySearch(array, target)
    // array need to be sorted
    // target is the value we're searching for
    
    left = 0
    right = array length -1

    while left <= right
      mid = left + (right - left) / 2
      
      if array[mid] == target // check if middle element is the target
        return mid
      else if array[mid] < target // if target is greater, search the right half
        left = mid + 1
      else // if target is smaller, search the left half
        right = mid - 1       
    return -1
  ```
  - [GO](./golang/8-binarysearch/binarysearch.go)
  - [Java](./java/8-binarysearch/Main.java)

## Sorting Algorithms
### Bubble Sort

## Tree

## Graph