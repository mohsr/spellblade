function submitCommand() {
	var xhr = new XMLHttpRequest();
	xhr.open('POST', '/cmd');
	xhr.onreadystatechange = function() {
		if (xhr.readyState === 4 && xhr.status === 200) {
			var returned = JSON.parse(xhr.responseText);
			showMessage(returned.Text, returned.Color, true);
		}
	}
	var box = $('#writebox');
	var data = box.val();
	box.val('');
	showMessage(data, "White", false);
	xhr.send('txt=' + data);
}