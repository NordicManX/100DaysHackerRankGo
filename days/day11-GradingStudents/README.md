# Dia 11: GradingStudents

### Enunciado Resumido

O problema consiste em implementar uma função para arredondar as notas dos estudantes de acordo com as regras da HackerLand University. As regras são:

Qualquer nota abaixo de 40 é considerada reprovada.

Se uma nota for menor que 38, ela não é arredondada, pois o resultado ainda seria uma reprovação.

Se a diferença entre a nota e o próximo múltiplo de 5 for menor que 3, a nota deve ser arredondada para esse próximo múltiplo de 5. Caso contrário, a nota permanece a mesma.

A função deve receber uma lista de notas e retornar uma nova lista com as notas após a aplicação dessas regras de arredondamento.

### Abordagem

Para resolver o problema, a abordagem mais direta é iterar sobre a lista de notas originais, aplicando as regras a cada uma delas. Para cada nota, o processo é o seguinte:

Verificação da Condição Mínima: Primeiro, verificamos se a nota é menor que 38. Se for, de acordo com as regras, nenhum arredondamento deve ocorrer. A nota original é mantida.

Cálculo para Arredondamento: Se a nota for 38 ou maior, precisamos determinar se ela deve ser arredondada.

Calculamos o próximo múltiplo de 5. Uma maneira eficiente de fazer isso é encontrar o resto da divisão da nota por 5 (usando o operador módulo, %).

A diferença entre a nota e o próximo múltiplo de 5 será 5 - (nota % 5).

Por exemplo, para a nota 73: 73 % 5 = 3. A diferença para o próximo múltiplo (75) é 5 - 3 = 2.

Aplicação do Arredondamento: Comparamos a diferença calculada com 3.

Se a diferença for menor que 3, somamos essa diferença à nota original para arredondá-la para cima. (Ex: 73 + 2 = 75).

Se a diferença for 3 ou maior, a nota não é alterada. (Ex: Para a nota 67, a diferença é 5 - (67 % 5) = 5 - 2 = 3, então a nota permanece 67).

Construção do Resultado: Cada nota, após ser processada (arredondada ou não), é adicionada a uma nova lista. Ao final do loop, essa nova lista, contendo todas as notas finais, é retornada.

### Complexidade

Tempo: O(n)
A solução requer que percorramos a lista de notas uma única vez. Sendo n o número de estudantes (ou o tamanho da lista de notas), o número de operações é diretamente proporcional a n. As operações dentro do loop (cálculo de módulo, comparações e soma) são de tempo constante, O(1). Portanto, a complexidade de tempo total é linear.

Espaço: O(n)
A função precisa retornar uma nova lista com as notas arredondadas. Para isso, alocamos um novo array ou lista de tamanho n para armazenar os resultados. Portanto, o espaço adicional utilizado pela solução é proporcional ao número de notas de entrada.
