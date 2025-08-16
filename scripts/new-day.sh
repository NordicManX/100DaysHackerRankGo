#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Uso: $0 02 matriz de inteiros e soma"
    echo "Exemplo: $0 02 simpleArraySum"
    exit 1
fi

DAY_NUMBER=$(printf "%02d" $1) 
PROBLEM_NAME="$2"
FOLDER_NAME="day${DAY_NUMBER}-${PROBLEM_NAME}"


FULL_PATH="../days/${FOLDER_NAME}"

if [ -d "$FULL_PATH" ]; then
    echo "Erro: A pasta ${FULL_PATH} já existe."
    exit 1
fi

echo "Criando estrutura para o Dia ${DAY_NUMBER} em ${FULL_PATH}..."
mkdir -p "$FULL_PATH"
touch "${FULL_PATH}/main.go"

cat > "${FULL_PATH}/README.md" << EOL
# Dia ${DAY_NUMBER}: ${PROBLEM_NAME}

### Enunciado Resumido

(Cole o resumo do problema aqui)

### Abordagem

(Explique como você resolveu o problema)

### Complexidade

- **Tempo:** O(...)
- **Espaço:** O(...)
EOL

echo "✅ Estrutura criada com sucesso!"