# Dia 03: compareTheTriplets

### Enunciado Resumido

O problema consiste em comparar as pontuações de duas pessoas, Alice e Bob, que foram avaliadas em três categorias distintas. As pontuações são representadas por dois arrays (ou "trincas") de três elementos cada, a para Alice e b para Bob. A comparação é feita elemento por elemento:

Se a[i] > b[i], Alice ganha 1 ponto.

Se a[i] < b[i], Bob ganha 1 ponto.

Se a[i] = b[i], ninguém ganha ponto.

O objetivo é calcular a pontuação total de cada um e retorná-la como um array de dois elementos: [pontuação_de_alice, pontuação_de_bob].

### Abordagem

Para resolver o problema, a abordagem foi a seguinte:

Inicialização: Duas variáveis, aliceScore e bobScore, foram inicializadas com o valor 0 para armazenar a pontuação de cada participante.

Iteração: Foi utilizado um laço for que itera três vezes, correspondendo aos três elementos dos arrays de pontuação (índices de 0 a 2).

Comparação: Dentro do laço, em cada iteração i, o elemento a[i] foi comparado com b[i].

Uma estrutura condicional if-else if foi usada para verificar qual dos dois valores era maior.

A variável de pontuação correspondente (aliceScore ou bobScore) era incrementada em 1 caso a condição fosse satisfeita. Nenhuma ação era tomada se os valores fossem iguais.

Resultado: Após o término do laço, as pontuações finais armazenadas nas variáveis foram agrupadas em um novo array de inteiros, que foi então retornado pela função.

### Complexidade

Tempo: O(1) - A complexidade de tempo é constante porque o laço de repetição é executado um número fixo de vezes (3), independentemente do valor das pontuações de entrada. A quantidade de operações não cresce com o tamanho da entrada, pois a entrada é sempre de tamanho fixo.

Espaço: O(1) - A complexidade de espaço também é constante. A função utiliza apenas um número fixo de variáveis (aliceScore, bobScore, i) e um array de resultado de tamanho fixo (2). A memória utilizada não depende do tamanho da entrada.
