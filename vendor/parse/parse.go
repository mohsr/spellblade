package parse

import (
    "game"
    "menu"
    "strings"
    "say"
)

/*
 * NOTES:
 * A text command is either a MENU ACTION, a SAY ACTION, or a GAME ACTION.
 *
 * A command is recognized as a MENU ACTION if the first word is a recognized 
 * menu action. The format for a menu action is as follows:
 *                   >>> [MENU_ACTION]
 * An example menu action is as follows:
 *                   >>> logout
 * If the command is not a recognized menu action, it is treated as a
 * say action or a game action depending on the first word.
 *
 * A SAY ACTION is made up of two parts:
 *                   >>> say "[WORDS]"
 * A say action always starts with "say" and is followed by a message in 
 * quotes. This message is then displayed to other nearby players. Any 
 * content between "say" and the beginning of the quotes is ignored.
 * Example say actions are as follows:
 *                   >>> say "hi there!"
 *                   >>> say "what's up?"
 * If the command does not start with say and is not a recognized menu action, 
 * then it is treated as a game action.
 *
 * A GAME ACTION is made up of three parts:
 *                   >>> [VERB][FILLER][TARGET]
 * The VERB is always necessary and will determine the action done. It is
 * always precisely the first word, and is never more than that.
 * The FILLER is never necessary and consists of words like "the" and "a", and
 * can be a variable amount of words.
 * The TARGET is usually necessary and will determine the target of the verb.
 * Example game actions are as follows:
 *                   >>> attack the dragon
 *                   >>> move north
 *                   >>> search
 */

/*
 * Holds the response of parsing a text command.
 * Schema: {"Text":  The response text to send back.
 *          "Color": The color of the response text to display.}
 */
type Response struct {
    Text  string
    Color string
}

/*
 * Parse
 * Purpose:    Parses a string command from a user and performs an action 
 *             accordingly.
 * Parameters: res - *Response - pointer to the response to store the 
 *                               action result in.
 *             txt - string - the text to be parsed.
 * Return:     n/a
 */
func Parse(res *Response, txt string) {
    /* Break up the words. */
    words := strings.Fields(txt)

    if len(words) == 0 || strings.Contains(strings.ToLower(txt), "script") {
        BadCommand(res)
        return
    }

    /* Check the first word of the command and do the appropriate action. */
    switch strings.ToLower(words[0]) {
        /* First, check for menu actions. */
        case "login":
            res.Text, res.Color = menu.LogIn();
        case "logout":
            res.Text, res.Color = menu.Logout();
        /* Next, check if it's a say action. */
        case "say":
            res.Text, res.Color = say.ParsedSay(txt);
        /* Lastly, check for game actions. */
        case "search":
            res.Text, res.Color = game.Search();
        default:
            res.Text = "Sorry, command not recognized. Type \"help\" for help!"
            res.Color = "Red"
    }
}

/*
 * BadCommand
 * Purpose:    Modifies a given pointed-to response to say that the 
 *             command was invalid.
 * Parameters: res - *Response - pointer to the response to modify.
 * Return:     n/a
 */
func BadCommand(res *Response) {
    res.Text = "Sorry, invalid command!"
    res.Color = "Red"
}