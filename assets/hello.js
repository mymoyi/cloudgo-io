$(document).ready(function() {
    $.ajax({
        url: "/api",
        dataType: "text"
    }).then(function(data) {
        data = JSON.parse(data)
        console.log(data.ID)
        console.log(data.Content)
    });
});
