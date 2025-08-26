# Dia 10: TimeConversion

### Enunciado Resumido

O desafio consiste em converter uma string de hora no formato de 12 horas (AM/PM) para o formato militar (24 horas). Por exemplo, "07:05:45PM" deve ser convertido para "19:05:45". Casos especiais como "12:00:00AM" e "12:00:00PM" devem ser tratados corretamente, resultando em "00:00:00" e "12:00:00", respectivamente.

### Abordagem

A solução implementada em Go aborda o problema da seguinte forma:

Extração de Componentes: A string de entrada é dividida para isolar o período ("AM" ou "PM") e as partes do horário (hora, minuto e segundo).

Conversão da Hora: A parte da hora é convertida de string para um número inteiro para facilitar os cálculos matemáticos.

Lógica de Conversão:

Se o período for "PM" e a hora for diferente de 12 (de 1 a 11), somamos 12 à hora.

Se o período for "AM" e a hora for 12, a hora é alterada para 0 (meia-noite).

Nos demais casos (AM de 1 a 11 e PM às 12), o valor da hora não precisa ser alterado.

Recomposição da String: Por fim, a nova hora é formatada para garantir que tenha dois dígitos (por exemplo, "00" ou "07") e é concatenada com os minutos e segundos originais para formar a string de saída no formato de 24 horas.

### Complexidade

Tempo: O(1)

As operações realizadas (dividir a string, converter para inteiro, somar e formatar a saída) levam um tempo constante, pois o tamanho da string de entrada é sempre fixo e pequeno. A complexidade não depende do valor da hora, mas sim da estrutura fixa do formato de entrada.

Espaço: O(1)

O espaço de armazenamento utilizado para as variáveis (período, partes do horário, nova hora) também é constante e não escala com o tamanho da entrada.
