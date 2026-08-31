# TODO

## Query post_publisher
- [x] arrumar queries do post_publisher
    - [x] CreatePublisher
    - [x] GetPublisherByPostID
    - [x] UpdatePublisherAsCompleted
    - [x] UpdatePublisherAsFailed

## Post

* [x] Atualizar payload de `POST /v1/post` para receber `social_connections_ids` (array)
* [x] `PostService` cria o post normalmente, sem `social_connections`
* [ ] Validar que cada `social_connections_id` pertence ao usuário autenticado pelo `ctx`
* [ ] Para cada `social_connections_id`, criar um registro em `post_publishers`

## Social_Connections

* [ ] Criar endpoint GET /v1/socials
* [ ] Buscar no banco as social_connections do usuário autenticado pelo ctx
* [ ] Retornar para o frontend apenas os dados necessários: id, provider
* [ ] Não retornar access_token nem refresh_token

## Publisher

* [ ] Criar interface `Publisher`
* [ ] Criar `TwitterPublisher`
  * [ ] Buscar `post.content`
  * [ ] Buscar `social_connection.provider`
  * [ ] Buscar `social_connection.access_token`
  * [ ] Chamar a API do Twitter usando `access_token`
  * [ ] Enviar `content` como `text` no body da requisição
  * [ ] Não enviar `title`

## Finalização

* [ ] Se Twitter publicar com sucesso, atualizar `post_publisher` para `completed`
* [ ] Se Twitter retornar erro, atualizar `post_publisher` para `failed`
* [ ] Testar o fluxo completo: `POST /v1/post` → `post_publisher` → `TwitterPublisher` → Twitter
