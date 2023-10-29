# go-task

`go-task` é uma solução de fila de push desenvolvida para agilizar e otimizar o gerenciamento de tarefas em sistemas distribuídos. Com o `go-task`, você pode enfileirar tarefas para serem executadas por meio de requisições POST ou PUT, com a opção de agendar a execução de uma tarefa em um horário específico.

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
