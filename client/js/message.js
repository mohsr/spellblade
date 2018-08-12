function showMessage(text, color, isResponse) {
	var prompt;
	if (isResponse) {
		prompt = '>>>';
	} else {
		prompt = '==>';
	}
	var html = '<p class="response ';
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
	html += ('">&nbsp;&nbsp;' + prompt + ' ' + text + '</p>');
	$('#textbox').append(html);
}