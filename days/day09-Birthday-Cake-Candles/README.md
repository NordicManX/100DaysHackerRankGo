# Dia 09: Birthday Cake Candles

### Enunciado Resumido

Dado um array de inteiros que representam as alturas de velas de aniversário, a tarefa é contar e retornar quantas velas são as mais altas. Por exemplo, para as alturas [3, 2, 1, 3], a altura máxima é 3 e ela aparece 2 vezes, então o resultado deve ser 2.

### Abordagem

A solução mais eficiente envolve percorrer a lista de alturas (candles) uma única vez, mantendo o controle da altura da vela mais alta encontrada até o momento e a contagem de quantas vezes essa altura apareceu.

Inicializamos duas variáveis: alturaMaxima para guardar a maior altura encontrada (começando com 0) e contador para o número de ocorrências (começando com 0).

Iteramos por cada altura no array de velas.

Dentro do loop, comparamos a altura atual com a alturaMaxima:

Se a altura atual for maior que a alturaMaxima, encontramos um novo recorde. Atualizamos a alturaMaxima para o valor da altura atual e reiniciamos o contador para 1.

Se a altura atual for igual à alturaMaxima, significa que encontramos outra vela com a mesma altura máxima. Apenas incrementamos o contador.

Após o loop terminar, a variável contador conterá o número exato de velas mais altas, que é o valor retornado.

Essa abordagem evita a necessidade de percorrer a lista duas vezes (uma para achar o máximo e outra para contar), tornando-a mais otimizada.

### Complexidade

Tempo: O(n)
Onde n é o número de velas no array. Isso ocorre porque o algoritmo percorre o array de velas apenas uma vez para encontrar a altura máxima e contar suas ocorrências.

Espaço: O(1)
O algoritmo utiliza apenas um número fixo de variáveis (alturaMaxima e contador) para armazenar o estado, independentemente do tamanho do array de entrada. Nenhum espaço adicional que dependa de n é alocado.
