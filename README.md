# go-task

go-tasker é uma solução construída em Go para gerenciar e processar tarefas em fila. Ele oferece uma API RESTful para adicionar tarefas, com a opção de agendar a execução de uma tarefa em um horário específico, atualizar configurações, e muito mais.

## Sumário

- [Instalação](#instalação)
- [Uso](#uso)
- [Configuração](#configuração)
- [Contribuições](#contribuições)
- [Licença](#licença)

## Instalação

1. Instale o Go (certifique-se de ter a versão mais recente).
2. Clone este repositório:

   `git clone https://github.com/hjunior29/go-tasker`

3. Navegue até a pasta do projeto:

   `cd go-tasker`

4. Instale as dependências necessárias:

   go mod download

## Uso

(completar aqui com instruções específicas sobre como executar, usar endpoints etc.)

## Configuração

1. Configuração de Timezone:

   Para configurar a timezone para São Paulo, edite a string de conexão no arquivo de configuração.

2. Outras configurações:

   (completar aqui com detalhes adicionais sobre configuração do projeto, como configurações de banco de dados, variáveis de ambiente, etc.)

## Contribuições

Contribuições são bem-vindas! Por favor, veja o CONTRIBUTING.md para mais detalhes.

## Características

- **Agendamento de Tarefas**: Defina quando uma tarefa específica deve ser executada usando nosso sistema de agendamento integrado.
- **Flexibilidade de Payload**: O `go-task` permite que você envie qualquer tipo de corpo de requisição, tornando-o adequado para uma ampla variedade de aplicações.
- **Integração com Postgres**: Utilizamos uma conexão Postgres robusta para registrar e gerenciar tarefas, garantindo durabilidade e confiabilidade.

## Como usar

1. **Enfileirar uma tarefa**:

  Para enfileirar uma tarefa, basta enviar uma requisição com o corpo apropriado:

   ```json
   {
     "url": "https://webhook.site/4226b1b2-f7e6-47c7-a5f1-8eb280ef9d24",
     "method": "POST",
     "scheduled_at": 1,
     "payload": {
       "name": "John Doe",
       "email": "johndoe@example.com"
     }
   }
   ```

2. **Alterar número de workers** 

  Para altera a quantidade de workers sobre a fila é necessário enviar essa requisição:

  ```json
  {
    "workers": 3
  }
  ```