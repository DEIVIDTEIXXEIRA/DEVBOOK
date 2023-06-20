$('#nova-publicacao').on('submit', criarPublicacao);

function criarPublicacao(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/publicacoes",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
        }
    }).done(function() {
        alert("Publicação criada com sucesso!!");
    }).fail(function() {
        alert("Erro ao criar publicação!!");
    });
}
