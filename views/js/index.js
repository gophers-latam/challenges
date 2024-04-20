// simple ajax post
$(document).ready(function () {
    let path = window.location.protocol + '//' + window.location.hostname + ':' + window.location.port;
    $("form").submit(function (e) {
        let description = $("#description").val();
        let level = $("#level").val();
        let challengeType = $("#challengetype").val();

        if (description !== "" && level !== "" && challengeType !== "") {
            let formData = {
                description,
                level,
                challengeType
            };
            $.ajax({
                type: "POST",
                url: `${path}/challenges`,
                data: formData,
                dataType: "json",
                encode: true,
            }).done(function (res) {
                console.log(res);
                // new Notificacion() needs permission, by default are disabled
                // instead use alert dialog
                alert("Desafio registrado para aprobar 👀");
            });
        } else {
            alert("No se estan mandando valores válidos");
        }

        e.preventDefault();
    });
});