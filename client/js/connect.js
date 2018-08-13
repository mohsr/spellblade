$(document).ready(function() {
    $('#writebox').keyup(function(event) {
        if (event.keyCode === 13) {
            submitCommand();
        }
    });
});

function submitCommand() {
    var box = $('#writebox');
    var data = box.val().trim();
    if (data == '') {
        return;
    }
    if (clientCommands(data)) {
        box.val('');
        return;
    }

    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/cmd');
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4 && xhr.status === 200) {
            var returned = JSON.parse(xhr.responseText);
            showMessage(returned.Text, returned.Color, true);
        }
    }
    box.val('');
    showMessage(data, "White", false);
    xhr.send('txt=' + data);
}

function clientCommands(cmd) {
    switch(cmd) {
        case "help":
            help();
            return true;
        case "clear":
            clear();
            return true;
    }
    return false;
}