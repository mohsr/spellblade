package parse

import (
    "strings"
)

/*
 * NOTES:
 * A text command is either a MENU ACTION or a GAME ACTION.
 *
 * A command is recognized as a MENU ACTION if the first word is a recognized 
 * menu action. The format for a menu action is as follows:
 *                   >>> [MENU_ACTION]
 * An example menu action is as follows:
 *                   >>> logout
 * If the command is not a recognized menu action, it is treated as a
 * game action. 
 *
 * A GAME ACTION is made up of three parts:
 *                   >>> [VERB][FILLER][TARGET]
 * The VERB is always necessary and will determine the action done.
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
 *
 */
func Parse(res *Response, txt string) {
    /* Break up the words. */
    words := strings.Fields(txt)

    /* TESTING -- print the words. */
    for i := 0; i < len(words); i++ {
        println(words[i])
    }

    res.Text = txt
    res.Color = "White"
}