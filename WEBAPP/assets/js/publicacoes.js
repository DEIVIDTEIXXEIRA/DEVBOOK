$('#nova-publicacao').on('submit', criarPublicacao);
$('.curtir-publicacao').on('click', curtirPublicacao);

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
        window.location = "/home"
    }).fail(function() {
        alert("Erro ao criar publicação!!");
    });
}

function curtirPublicacao(evento) {
    console.log("curtindo published");

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    $.ajax({
        url:  `/publicacoes/${publicacaoId}/curtir`,
        method: "POST",
    }).done(function() {
        alert("Publicação curtida")
    }).fail(function() {
        alert("Erro ao curtir a published..")
    })
}

