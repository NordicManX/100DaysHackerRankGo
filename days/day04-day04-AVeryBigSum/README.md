# Dia 04: day04-AVeryBigSum

### Enunciado Resumido

O desafio consiste em calcular a soma de todos os elementos de um array. A principal particularidade é que os números nesse array podem ser muito grandes, excedendo a capacidade de um tipo de inteiro padrão de 32 bits, por isso foi utilizado o padrão de 64 bits.

### Abordagem

A solução foi implementada em Go. A estratégia principal para lidar com os números grandes foi utilizar o tipo de dado int64 em todas as etapas: ao declarar a variável da soma, ao ler os números da entrada e ao definir o tipo de retorno da função.

O algoritmo funciona da seguinte forma:

Uma variável sum do tipo int64 é inicializada com o valor 0.

O código itera sobre cada elemento do array de entrada.

A cada iteração, o valor do elemento atual é adicionado à variável sum.

Ao final do loop, a função retorna o valor total acumulado em sum.

Essa abordagem garante que a soma não sofra "overflow" (estouro de capacidade), mesmo que o resultado seja um número muito grande.

### Complexidade

Tempo: O(n)

A complexidade de tempo é linear porque o algoritmo precisa percorrer cada um dos n elementos do array uma única vez para somá-los.

Espaço: O(1)

A complexidade de espaço é constante. Além do espaço já ocupado pelo array de entrada, o algoritmo utiliza apenas uma variável extra (sum) para armazenar a soma, independentemente do tamanho do array.
