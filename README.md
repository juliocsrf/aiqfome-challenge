# 🚀 Aiqfome Challenge - API REST

Olá! 👋 Este projeto foi desenvolvido como um desafio técnico do Aiqfome, onde apliquei conceitos modernos de desenvolvimento em Go, sempre pensando em **código limpo**, **manutenibilidade** e **escalabilidade**.

## 🎯 O que esta API faz?

- **Gerencia clientes** com validações robustas
- **Lista produtos** integrados via API externa (FakeStore API)
- **Sistema de favoritos** para clientes salvarem produtos preferidos
- **Autenticação JWT** completa com refresh tokens
- **Documentação automática** via Swagger

## 🛠️ Principais Tecnologias e Por Que Escolhi

### **Go 1.24.5**

Escolhi Go pela sua simplicidade, performance e excelente suporte para APIs. É uma linguagem que permite escrever código expressivo sem sacrificar velocidade. E também, atualmente é minha principal linguagem e a que mais tenho domínio.

### **Clean Architecture + DDD**

Implementei Clean Architecture para separar claramente as responsabilidades:

- **Domain**: Regras de negócio puras
- **Use Cases**: Orquestração das operações
- **Adapters**: Implementações específicas (HTTP, Database)

Isso facilita testes, manutenção e permite trocar componentes sem impactar o core da aplicação.

### **Chi Router**

Preferi o Chi pela sua simplicidade e performance. É minimalista mas poderoso, ideal para APIs REST bem estruturadas.

### **PostgreSQL + SQLC**

- **PostgreSQL**: Banco sugerido pelo próprio desafio. Mas não teria problemas em utilizar qualquer outro banco.
- **SQLC**: Gera código Go type-safe a partir de SQL. Zero reflection, máxima performance. É um biblioteca que não deixa o projeto dependente dela. Se amanhã eu não quiser utiliza-la mais, posso remover que o código continua funcionando.

### **Google Wire**

Implementei injeção de dependências com Wire para:

- Eliminar acoplamento entre camadas
- Facilitar testes unitários
- Garantir que dependências sejam resolvidas em compile-time

### **JWT + bcrypt**

- **JWT**: Tokens stateless para autenticação
- **bcrypt**: Hash seguro de senhas
- **Refresh tokens**: Renovação automática sem re-login

### **Docker + Docker Compose**

Setup completo para facilitar desenvolvimento e deploy. Qualquer pessoa pode rodar o projeto com um comando.

### **Swagger/OpenAPI**

Documentação automática da API. Facilita integração e testes para outros desenvolvedores.

## 🔐 Como Funciona a Autenticação

O sistema de auth foi pensado para ser **seguro** e **user-friendly**:

1. **Login**: `POST /api/auth/login` com email/senha
2. **Resposta**: Access token (15min) + Refresh token (7 dias)
3. **Uso**: Header `Authorization: Bearer {token}` em todas as rotas protegidas
4. **Renovação**: `POST /api/auth/refresh` quando access token expira

```bash
# 1. Fazer login
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "user@example.com", "password": "password123"}'

# 2. Usar o token retornado
curl -X GET http://localhost:8080/api/customers/123 \
  -H "Authorization: Bearer SEU_TOKEN_AQUI"
```

## 🚦 Como Rodar o Projeto

**Pré-requisito**: Docker e Docker Compose instalados

```bash
# Clone o repositório
git clone <seu-repo>
cd aiqfome-challenge

# Suba tudo com um comando
docker-compose up --build

# Pronto! API rodando em http://localhost:8080
# Documentação em http://localhost:8080/swagger/index.html
```

Simples assim! O Docker Compose cuida de:

- ✅ Subir PostgreSQL
- ✅ Aguardar banco ficar pronto
- ✅ Executar migrações automaticamente
- ✅ Iniciar a API

## 📋 Principais Endpoints

| Método | Endpoint                                    | Descrição                      |
| ------ | ------------------------------------------- | ------------------------------ |
| `POST` | `/api/auth/login`                           | Login                          |
| `POST` | `/api/auth/refresh`                         | Renovar token                  |
| `POST` | `/api/customers`                            | Criar cliente                  |
| `GET`  | `/api/customers/{id}`                       | Buscar cliente (com favoritos) |
| `GET`  | `/api/products`                             | Listar produtos                |
| `POST` | `/api/customers/{id}/favorites/{productId}` | Adicionar favorito             |

> **💡 Dica**: Use a documentação Swagger em `/swagger/index.html` para testar interativamente!

## 🏗️ Estrutura do Projeto

```
├── cmd/server/          # Entry point
├── internal/
│   ├── domain/          # Entidades e regras de negócio
│   ├── usecase/         # Casos de uso da aplicação
│   ├── adapter/         # Implementações (HTTP, DB, etc)
│   └── wire/            # Injeção de dependências
├── database/            # Migrações SQL
└── docker-compose.yml   # Setup completo
```

## 💭 Reflexões Técnicas

### **Por que esta arquitetura?**

Escolhi Clean Architecture porque facilita:

- **Testes**: Cada camada pode ser testada isoladamente
- **Manutenção**: Mudanças em uma camada não afetam outras
- **Evolução**: Fácil adicionar novas features sem quebrar existentes

### **Decisões de Design**

- **DTOs separados**: Validação e serialização controladas
- **Middleware JWT**: Proteção automática de rotas
- **Repository Pattern**: Abstração do banco de dados
- **Use Cases**: Regras de negócio centralizadas

## 🎯 Resultados

- ✅ **API REST completa** e funcional
- ✅ **Autenticação robusta** com JWT
- ✅ **Arquitetura escalável** e testável
- ✅ **Documentação automática**
- ✅ **Setup zero-friction** com Docker
- ✅ **Código limpo** e bem estruturado
