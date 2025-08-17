#!/bin/bash

if [ "$#" -ne 2 ]; then
    echo "Uso: $0 03 'compare os trigemeos'"
    echo "Exemplo: $0 03 compareTheTriplets"
    exit 1
fi

SCRIPT_DIR=$(dirname "$0")

DAY_NUMBER=$(printf "%02d" $1)
PROBLEM_TITLE="$2"


PROBLEM_NAME_FOR_FOLDER=$(echo "$PROBLEM_TITLE" | tr ' ' '-')

FOLDER_NAME="day${DAY_NUMBER}-${PROBLEM_NAME_FOR_FOLDER}"

FULL_PATH="${SCRIPT_DIR}/../days/${FOLDER_NAME}"



if [ -d "$FULL_PATH" ]; then
    echo "Erro: A pasta ${FULL_PATH} já existe."
    exit 1
fi

echo "Criando estrutura para o Dia ${DAY_NUMBER} em ${FULL_PATH}..."
mkdir -p "$FULL_PATH"
touch "${FULL_PATH}/main.go"

cat > "${FULL_PATH}/README.md" << EOL
# Dia ${DAY_NUMBER}: ${PROBLEM_TITLE}

### Enunciado Resumido

(Cole o resumo do problema aqui)

### Abordagem

(Explique como você resolveu o problema)

### Complexidade

- **Tempo:** O(...)
- **Espaço:** O(...)
EOL

echo "✅ Estrutura criada com sucesso!"