// simple ajax post
$(document).ready(function () {
    Notification.requestPermission().then((res) => {
        console.log(res);
    });
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
                new Notification("Desafio recibido y guardado correctamente 😎");
            });
        } else {
            alert("No se estan mandando valores válidos")
        }

        e.preventDefault();
    });
});