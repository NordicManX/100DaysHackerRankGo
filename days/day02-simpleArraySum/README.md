cat > "${FULL_PATH}/README.md" << EOL
<<<<<<< HEAD
# Dia ${02}: ${matriz e soma de numeros inteiros}
=======
# Dia ${02}: ${simpleArraySum}
>>>>>>> 3984ec279c7c4826ce320b0cba7ad2ed08f5eafb



### Dificuldade
- [x] Fácil
- [ ] Médio
- [ ] Difícil

### Enunciado Resumido
Dada uma lista de números inteiros, o desafio é calcular e retornar a soma de todos os seus elementos.

### Abordagem
A solução foi implementada seguindo uma lógica iterativa simples:
1.  Foi declarada uma variável acumuladora, chamada `soma`, inicializada com o valor 0.
2.  Utilizei um laço `for...range` para percorrer cada `numero` individualmente no array de entrada.
3.  A cada iteração, o `numero` atual foi adicionado à variável `soma`.
4.  Após o término do laço, a função retornou o valor final da `soma`.

### Complexidade

- **Tempo:** O(n)
  - A explicação é que o algoritmo precisa "visitar" cada um dos `n` elementos do array exatamente uma vez. Isso significa que o tempo de execução cresce de forma linear com o tamanho do array. Se o array tiver o dobro de elementos, o tempo levará o dobro.

- **Espaço:** O(1)
  - A explicação é que o algoritmo usa uma quantidade constante de memória extra. Ele apenas cria uma variável (`soma`) para guardar o resultado. A memória utilizada não aumenta se o array de entrada tiver 10 ou 1 milhão de elementos.

EOL

echo "✅ Estrutura criada com sucesso!"
