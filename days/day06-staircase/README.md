# Dia 06: staircase

### Enunciado Resumido

O desafio consiste em escrever uma função que recebe um inteiro n e imprime no console uma escada, alinhada à direita, com altura e base de tamanho n. A escada deve ser construída usando o caractere # e espaços. A última linha não deve ter nenhum espaço antes dos caracteres #

### Abordagem

A solução se baseia em um laço de repetição (for) que itera n vezes, onde cada iteração corresponde a uma linha da escada. A variável do laço, que podemos chamar de i, representará o número da linha atual, indo de 1 até n.

### Complexidade

# Tempo: O(n²)

A complexidade de tempo é quadrática porque temos um laço principal que executa n vezes (para cada linha). Dentro desse laço, a operação de construir a string (seja com laços aninhados ou com funções como repeat) leva um tempo proporcional a n no pior caso (para a última linha, que tem n caracteres). Portanto, a complexidade total é n

timesn=O(n²).

# Espaço: O(n)

A complexidade de espaço é linear. Em cada iteração do laço, uma nova string é criada e armazenada na memória para representar a linha atual. A maior string criada terá n caracteres (a última linha). Como as linhas são impressas e descartadas a cada passo, o espaço máximo necessário em qualquer momento é proporcional a n.