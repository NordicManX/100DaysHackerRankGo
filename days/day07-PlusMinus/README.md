# Dia 07: PlusMinus

### Enunciado Resumido

Dado um array de inteiros, o desafio é calcular as proporções (frações) de seus elementos que são positivos, negativos e zeros. Ao final, deve-se imprimir cada uma das três proporções em uma nova linha, com precisão de 6 casas decimais.

### Abordagem

A solução percorre o array de inteiros uma única vez, mantendo três contadores distintos: um para números positivos, um para negativos e um para zeros.

### Complexidade

Tempo: O(n)

O algoritmo precisa percorrer cada um dos n elementos do array exatamente uma vez. Portanto, o tempo de execução cresce de forma linear com o tamanho da entrada.

Espaço: O(1)

O espaço de memória utilizado é constante. O algoritmo usa apenas algumas variáveis para os contadores, e essa quantidade não muda, independentemente do tamanho do array de entrada. Não são criadas novas estruturas de dados que dependam de n.
