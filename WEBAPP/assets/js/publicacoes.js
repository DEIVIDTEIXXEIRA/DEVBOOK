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
    evento.preventDefault();

    const elementoClicado = $(evento.target);
    const publicacaoId = elementoClicado.closest('div').data('publicacao-id');

    elementoClicado.prop('disable', true);
    $.ajax({
        url:  `/publicacoes/${publicacaoId}/curtir`,
        method: "POST",
    }).done(function() {
        const contadorDeCutidas = elementoClicado.next('span');
        const quantidadeDeCutidas = parseInt(contadorDeCutidas.text());
        
        contadorDeCutidas.text(quantidadeDeCutidas + 1);
        
    }).fail(function() {
        alert("Erro ao curtir a published..");
    }).always(function() {
        elementoClicado.prop('disable', false);
    })

}

