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
- [Searching Algorithms](#searching-algorithms)
- [Sorting Algorithms](#sorting-algorithms)

## Introdução
- Estrutura de dados e algoritimos são fundamentais em ciência da computação, nos ajudam a organizar e processar dados de forma eficiente.
- São usados para resolver problemas e desafios em software, desde gerenciar grandes volumes de dados até otimizar a velocidade de atividades.
- Em DSA, o foco principal é resolver problemas de forma eficiente e eficaz. Para determinar a eficiência do algoritmo, olhamos para dois tipos de complexidade:
  - **Time complexity**, tempo de execução do algoritmo.
  - **Space complexity**, quantidade de memória usada do algoritmo.
- As estruturas de dados, podem ser divididas em dois tipos basicamente: *linear* e *não linear*:
  - *linear* os elementos são organizados em sequência, um após o outro. Por exemplo, *arrays*, *queues* e *stacks*.
  - *não linear* os elementos não possuem sequência alguma, ao invés disso, são organizados em hierarquia, um elemento pode estar conectado em um ou mais elementos, são dividios em dois grupos, *graph* e *tree*.

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
    
## Bitwise Algorithms
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
    - `~1` = 0
    - `~0` = 1
  - **Left shifting**
    - Usa dois operadores, `a<<b`, a é o valor a sofrer a mudança nos bits, b é a quantidade de *bits* que queremos trocar.
    - Os bits serão movidos para equerda conforme a quantidade de posições que indicamos.
    - Nunca usar números negativos como parâmetro, terá um comportamento inesperado.
    - Cuidade ao usar parâmetros maiores que a quantidade suportada pelo tipo da variável usada. Por exemplo: `a<<33` a sendo armazenado em variável de 32 bits.
    - Ex:
      - [Go](./golang/1-bitwise/bitwise.go)
      - [Java](./java/1-bitwise/Bitwise.java)
  - **Right shifting**
    - Semelhante ao **left shifting**, diferença é o sentido que vai para a direita.
    - `a>>b` segue as mesmas regras do **left shifting**.

## Array
- **Arrays** são estruturas de dados lineares que guardam elementos do mesmo tipo de dado.
- São alocados de forma contígua em memória, o que significa, é um espaço de memória com o tamanho da estutura, onde cada elemento está alocado sequencialmente em um indice do espaço alocado.
- Cada item é indexado e começa no **index** 0 na maioria das linguagens. Podemos acessar um elemento diretamente através do seu índice.
- Ex:
  - [Go](./golang/2-array/array.go)
  - [Java](./java/2-array/Main.java)

## Two Pointer Technique
- É uma ténica muito usada para resolver problemas usando **arrays** ou **strings**.
- Usa dois *ponteiros* que percorrem os elementos, podendo ser de direções opostas(inicio e fim) ou mesma direção com passos diferentes, dependendo do problema.
- Exemplos:
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
- É uma técnica quando uma **função** chama ela mesmo.
- Geralmente usada para resolver um problema que pode ser quebrado em problemas menores nos quais são resolvidos *recursivamente*.
- A cada chamada da função, será criada uma pilha com o item atual no topo, esse processo continua até o caso base ser atingido, que é a condição que para a **recursão**. Quando o caso base ocorre, a função volta o ciclo seguindo a sequência da pilha de chamadas e retornando os seus resultados.
- Geralmente temos 3 partes em **recursão**:
  - Caso base: que é a condição que para a **recursão**.
  - Caso recursivo; que é quando a função faz a chamada de sí mesma com um pedaço menor do problema.
  - Retorno/resultado: é quando a recursão termina e combina o resultado da pilha de execução.
- Ex:
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
- São estruturas de dados lineares que armazenam os dados em nós que são conectados por ponteiros.
- Diferentemente de **arrays**, não são armazenados de forma contígua em memória.
- Podem ser facilmente alterada no seu tamanho, adicionando ou removendo nós.
- Acessados de forma sequencial, através do inicio(head).
- Tipos mais usados de **linked lists**:
  - **Singly linked list** cada nó aponta para o próximo item.
  - **Doubly linked list** cada nós aponta para o próximo nó e para o anterior.
  - **Cirular linked list** o último nó aponta para o primeiro nó, formando um circulo.
- Ex:
  - [Go](./golang/4-linkedlist/linkedlist.go)
  - [Java](./java/4-linkedlist/Main.java)

## Queue
- **Queue** ou fila, é uma estrutura de dados, usa o conceito **FIFO(first in first out)**.
- O primeiro elemento inserido será o primeiro a sair da estrutura.
- Ex:
  - [Go](./golang/5-queue/queue.go)
  - [Java](./java/5-queue/Main.java)

## Stack
- **Stack** ou pilha, é uma estrutura de dados, usa o conceito **LIFO(last in first out)**.
- Ex:
  - [GO](./golang/6-stack/stack.go)
  - [Java](./java/6-stack/Main.java)

## Searching Algorithms
- São essenciais para encontrar itens especificos dentro de uma coleção de dados.
- Os algoritimos mais comuns são: **linear search** e **binary search**.

### Linear Search
- É denifido com um algoritimo sequencial, que começa no primeiro elemento e percorre a cada elemento atẽ encontrar o elemento especifico, ou chegar ao final da estrutura.
- Pode ser quebrados em passos a serem seguidos:
  - *Passo 1* Ponto de partida no primeiro elemento da coleção de dados.
  - *Passo 2* Comparar cada elemento se coincide com o elemento desejado.
  - *Passo 3* Encontrado o item desejado, interromper a busca e retorna o *index* do item ou true.
  - *Passo 4* Mover para o próximo elemento caso não encontrar o item na posição atual.
  - *Passo 5* Repetir até o final da coleção.
  - *Passo 6* Não encontrado o item, retornar informando que o item não foi encontrado.
- Ex:
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
- É um algoritmo de busca usado em *arrays ordenados*, caso contrário não irã funcionar da forma esperada.
- Usa o principio de **divide and conquer (dividir e conquistar)**, divide o o array pela metade repetidamente até encontrar o elemento buscado ou não encontrá-lo.
- Busca o elemento do meio, compara com o elemento buscado e decide qual lado após a divisão irá ser usado, faz isso repetidamente.
- O passo a passo é o seguinte:
  - *Passo 1* encontrar o elemento do meio e comparar com o elemento buscado, se o item do meio ser o item buscado, retornar a posição do elemento do meio.
  - *Passo 2* se o elemento do meio não for igual, checar se ele mé maior ou menor que o elemento do meio.
  - *Passo 3* se maior, seguir para a busca no lado direito, senão, buscar no lado direito.
  - *Passo 4* repetir *passos 1, 2 e 3* até o tamanho do array final ser 1.
  - *Passo 5* se não encontrar o elemento buscado no array, retornar informando que o item não foi encontrado.
- Ex:
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


## Hash

## Tree

## Graph