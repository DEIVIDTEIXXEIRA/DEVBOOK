$(document).ready(function() {
    $('#formulario-cadastro').on('submit', function(event) {
        event.preventDefault();
        console.log("Dentro da função usuario!");

        if ($('#senha').val() !== $('#confirmar-senha').val()) {
            Swal.fire("Ops...", "As senhas não coincidem", "error");
            return;
        }

        $.ajax({
            url: "/usuarios",
            method: "POST",
            data: {
                nome: $('#nome').val(),
                email: $('#email').val(),
                nick: $('#nick').val(),
                senha: $('#senha').val()
            },
        }).done(function() {
            Swal.fire("Sucesso", "Usuário cadastrado com sucesso!", "success")
            .then(function() {
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $('#email').val(),
                        senha: $('#senha').val()
                    }
                }).done(function() {
                    window.location = "/home";
                }).fail(function() {
                    Swal.fire("Ops...", "Erro ao autenticar!", "error");
                })
            })
        }).fail(function(erro) {
            Swal.fire("Ops...", "Erro ao cadastrar o usuário!", "error");
        });
    });
});