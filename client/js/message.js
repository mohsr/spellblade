function showMessage(text, color, isResponse) {
	var html = '<p class=wholeresponse>&nbsp;&nbsp;'
	if (isResponse) {
		html += '>>> ';
	} else {
		html += '==> ';
	}

	html += '<span class="response ';
	html += color.toLowerCase();
	html += ('">' + text + '</span></p>')
	$('#textbox').append(html);
	$('#textbox').scrollTop(document.getElementById('textbox').scrollHeight);
}