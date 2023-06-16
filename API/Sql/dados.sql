
insert into usuarios (nome, nick, email, senha)
values
("usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$UjhhIzhP4Ty/ppLnNL9KwuNg5uqynwNKjvHJfEMMPRUNi0DsYaPQa"),
("usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$UjhhIzhP4Ty/ppLnNL9KwuNg5uqynwNKjvHJfEMMPRUNi0DsYaPQa"),
("usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$UjhhIzhP4Ty/ppLnNL9KwuNg5uqynwNKjvHJfEMMPRUNi0DsYaPQa");

insert into seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3),
(2, 3);

INSERT into publicacoes(titulo, conteudo, autor_id)
values
("publicação do usuário 1", "essa é a publicação do usuario 1", 1),
("publicação do usuário 2", "essa é a publicação do usuario 2", 2),
("publicação do usuário 3", "essa é a publicação do usuario 3", 3);
