# Web Scraping para Análise de Movimentações Bancárias

Este projeto realiza o web scraping das movimentações bancárias da empresa na qual trabalho. O objetivo é coletar dados das movimentações bancárias e processá-los para obter um melhor entendimento e análise das finanças.

## Descrição
O scraper coleta dados de uma conta bancária administrativa e transforma as informações das movimentações em um formato JSON bem estruturado. Esses dados podem ser utilizados para análises financeiras e relatórios.

## Funcionalidades
- Coleta de Dados: Raspa informações das movimentações bancárias de uma URL fornecida.
- Processamento de Dados: Converte dados brutos em um formato JSON estruturado.
- Salvamento em Arquivo: Salva os dados JSON em um arquivo na raiz do projeto.

## Tecnologias Utilizadas
- Golang: Linguagem de programação principal para o projeto.
- Colly: Biblioteca para web scraping.
- Goquery: Biblioteca para manipulação de HTML.
- Godotenv: Biblioteca para carregar variáveis de ambiente de um arquivo .env.
- JSON: Formato para salvar e estruturar os dados coletados.
- Make: Arquivo make

## Estrutura do Projeto
- cmd/: Contém o arquivo principal do programa.
- internal/: Contém as lógicas internas, como modelos e processadores.
- tests/: Contém testes unitários
