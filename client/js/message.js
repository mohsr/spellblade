function showMessage(text, color, isResponse) {
	var html = '<p class=wholeresponse>&nbsp;&nbsp;'
	if (isResponse) {
		html += '>>> ';
	} else {
		html += '==> ';
	}

	html += '<span class="response ';	
	switch (color) {
		case "White":
			html += 'white';
			break;
		case "Red":
			html += 'red';
			break;
		case "Green":
			html += 'green';
			break;
	}
	html += ('">' + text + '</span></p>')
	$('#textbox').append(html);
}