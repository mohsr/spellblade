package menu

/*
 * LogIn
 * Purpose:    Logs a player into the game.
 * Parameters: n/a
 * Return:     r - string - the text to give as a response to the player.
 *             c - string - the color of the response to give the player.
 */
func LogIn() (r string, c string) {
    return "You've logged in. Have fun!", "Green"
}

/*
 * Logout
 * Purpose:    Logs a player out from the game.
 * Parameters: n/a
 * Return:     r - string - the text to give as a response to the player.
 *             c - string - the color of the response to give the player.
 */
func Logout() (r string, c string) {
    return "You've successfully logged out. See you later!", "Green"
}