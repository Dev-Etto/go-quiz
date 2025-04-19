# Jogo de Quiz em Go

Este é um simples jogo de quiz no terminal escrito em Go. O jogo lê perguntas de um arquivo CSV, embaralha-as e as apresenta ao jogador. O jogador tem um tempo limitado para responder a todas as perguntas.

## Funcionalidades

- Lê perguntas e respostas de um arquivo CSV.
- Embaralha a ordem das perguntas.
- Fornece um limite de tempo de 5 minutos para o jogo.
- Concede 10 pontos para cada resposta correta.
- Exibe regras e instruções antes de iniciar o jogo.

## Como Jogar

1. Clone este repositório ou baixe o código-fonte.
2. Certifique-se de que o Go está instalado no seu sistema.
3. Coloque um arquivo CSV chamado `question_go.csv` no mesmo diretório do programa. O arquivo CSV deve ter o seguinte formato:
   ```
   Pergunta,Opção1,Opção2,Opção3,Opção4,ÍndiceRespostaCorreta
   ```
   Exemplo:
   ```
   Qual é a capital da França?,Paris,Londres,Berlim,Madrid,1
   ```
4. Execute o programa com o seguinte comando:
   ```bash
   go run main.go
   ```
5. Siga as instruções exibidas na tela para jogar.

## Regras

1. Você será questionado com 10 perguntas aleatórias.
2. Cada resposta correta vale 10 pontos.
3. Você tem 5 minutos para responder a todas as perguntas.
4. Insira o número correspondente à sua resposta.
5. O jogo termina quando o tempo acabar ou todas as perguntas forem respondidas.

## Estrutura do Projeto

- `main.go`: O arquivo principal do programa contendo a lógica do jogo.
- `question_go.csv`: O arquivo CSV contendo as perguntas do quiz (deve ser criado pelo usuário).

## Dependências

Este projeto utiliza as seguintes bibliotecas padrão do Go:

- `bufio`: Para leitura de entrada do usuário.
- `csv`: Para leitura de perguntas do arquivo CSV.
- `fmt`: Para entrada e saída formatada.
- `math/rand`: Para embaralhar as perguntas.
- `os`: Para operações de arquivo e entrada/saída.
- `strconv`: Para conversão de strings em inteiros.
- `time`: Para gerenciar o temporizador do jogo.

## Melhorias Futuras

- Adicionar suporte para múltiplos arquivos CSV ou categorias.
- Implementar um placar de líderes.
- Permitir personalização do limite de tempo e do número de perguntas.

## Licença

Este projeto é open-source e está disponível sob a licença MIT.
