
# TGame - v1.0
[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)

- ***TGame** é um bot de gacha para Telegram, feito de jogador para jogadores. Você pode criar sua própria instância do bot e modificar da forma que você quiser, adequando-o para o seu estilo de jogo, features e universo.*

## Tecnologias utilizadas
- Golang v1.23.1
- PostgreSQL (última versão)
- Redis (última versão)
- Força de vontade!

## Features
- Draws ou "giros"
- Sistema de raridade
- Ranking "global" de jogadores
- Sistema de loja
- Mini-jogos de adivinhação com prêmios
- Eventos criados pela administração
- Sistema de recompensa diária
- Sorteio de cartas diárias, desde comuns até lendárias
## Features administrativas
- Dashboard de métricas, desde usuários até uso de recursos do seu servidor
- Painel restrito para admins adicionarem cartas, editarem, excluirem etc.
- Criação de tickets resgatáveis
- Criação eventos com prêmios diários

## Como usar
#### REQUISITOS:
- VPS: 
    - Debian GNU/Linux
    - 2GB de RAM.
    - 4GB de SWAP.
- No geral você não precisa fazer muito esforço, pois os serviços irão rodar automaticamente porque serão inicializados pelo Docker.

1. Instale o Docker. Se você está usando uma distro derivada ou diferente (Ubuntu, Fedora, Rock) busque na documentação da sua distro no site oficial do Docker e busque como instalar.

## Rodando com o Docker [Branch de desenvolvimento]

Rode o seguinte comando após o Docker ser instalado:

```bash
git clone https://github.com/naftalino/tg-game
cd tg-game
sudo docker-compose up --build
```
Desta forma, o docker criará uma imagem com base no `docker-compose.yml` e subirá banco de dados, cache e o bot.

## Rodando com o Docker [Produção]
- Após fazer o que tem de ser feito, você pode rodar o seguinte comando:
```bash
sudo docker-compose build --no-cache
sudo docker-compose up -d
```

## Licença

[MIT](https://choosealicense.com/licenses/mit/)

## Autores

- [Mao](https://t.me/adorabat)

## Contribuição

A sua contribuição é mais que bem-vinda. Seja ela uma frase de apoio, um commit, qualquer coisa. Não deixe de dar a estrela no repo! Tmj!