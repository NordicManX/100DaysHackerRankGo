# Dia 08: MiniMax-Sum

### Enunciado Resumido

Dado um array com cinco números inteiros positivos, o objetivo é encontrar os valores mínimo e máximo que podem ser obtidos somando-se exatamente quatro desses cinco números. Ao final, esses dois valores (a soma mínima e a soma máxima) devem ser impressos em uma única linha, separados por um espaço.

### Abordagem

A maneira mais eficiente de resolver o problema é perceber que:

A soma mínima de quatro dos cinco números é obtida somando todos os números, exceto o maior número do array.

A soma máxima de quatro dos cinco números é obtida somando todos os números, exceto o menor número do array.

O algoritmo, portanto, consiste em:

Calcular a soma total de todos os cinco elementos do array. É importante usar um tipo de dado que suporte números grandes (como um inteiro de 64 bits), conforme a dica do problema, para evitar estouro (overflow).

Encontrar o menor e o maior valor no array.

Calcular a soma mínima: soma_mínima = soma_total - maior_valor.

Calcular a soma máxima: soma_máxima = soma_total - menor_valor.

Imprimir os dois resultados.

Uma abordagem alternativa seria ordenar o array. Após a ordenação, a soma mínima seria a soma dos quatro primeiros elementos e a soma máxima seria a soma dos quatro últimos.

### Complexidade

Analisando a abordagem principal (sem ordenação):

Tempo: O(n)

O algoritmo percorre a lista uma única vez para calcular a soma total e encontrar os valores mínimo e máximo. Como o número de elementos (n) é fixo em 5, a complexidade é, na prática, constante, ou O(1).

Espaço: O(1)

A solução requer apenas algumas variáveis para armazenar a soma total, o valor mínimo e o valor máximo, não importando o tamanho do array de entrada. Portanto, o uso de memória é constante.
