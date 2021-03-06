function help() {
	showMessage("Help", "Gold", true);
	showMessage("For help with menu actions (such as logging out), type " +
		        "\"help menu\".", "White", true);
	showMessage("For help with say actions (such as talking), type " +
		        "\"help say\".", "White", true);
	showMessage("For help with game actions (such as moving and " +
		        "interacting), type \"help game\".", "White", true);
}

function helpMenu() {
	showMessage("Help - menu actions", "Gold", true);
	showMessage("Login - logs the player into the game.", "White", true);
	showMessage("Logout - logs the player out of the game.", "White", true);
}

function helpSay() {
	showMessage("Help - say actions", "Gold", true);
	showMessage("Say actions help you talk to other players in the game.",
		        "White", true);
	showMessage("To say something, type 'say \"This is my message!\"'", 
		        "White", true);
}

function helpGame() {
	showMessage("Help - game actions", "Gold", true);
	showMessage("Search - searches the current area the player is in.", 
		        "White", true);
}

function clear() {
    document.getElementById('textbox').innerHTML = '';
}