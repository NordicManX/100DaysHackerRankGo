# Dia 05: DiagonalDifference

### Enunciado Resumido

Dada uma matriz quadrada, o desafio é calcular a diferença absoluta entre as somas de suas duas diagonais: a principal (do canto superior esquerdo ao inferior direito) e a secundária (do canto superior direito ao inferior esquerdo).

### Abordagem

A abordagem mais eficiente para resolver este problema é otimizar a coleta de dados, percorrendo a matriz uma única vez.

Durante esta iteração, realizamos o cálculo das somas de ambas as diagonais de forma simultânea. Para cada linha i da matriz, o elemento correspondente da diagonal principal (matriz[i][i]) e o da diagonal secundária (matriz[i][n-1-i]) são identificados e adicionados a duas variáveis acumuladoras distintas.

Ao concluir a iteração, essas variáveis contêm as somas totais de cada diagonal. O passo final é calcular a diferença entre esses dois totais e retornar seu valor absoluto, conforme solicitado pelo problema.

### Complexidade

Tempo: O(n)

A complexidade de tempo é linear em relação ao número de linhas (ou colunas) da matriz, n. Isso ocorre porque percorremos a matriz apenas uma vez com um único laço que executa n iterações.

Espaço: O(1)

A complexidade de espaço é constante. A quantidade de memória extra utilizada não depende do tamanho da matriz de entrada. Usamos apenas um número fixo de variáveis (DiagonalPrincipal, DiagonalSecundaria, n, i) para realizar o cálculo.