package say

import (
    "regexp"
)

var (
    /* regexp to extract the message from a say action. */
    re *regexp.Regexp = regexp.MustCompile("say \"?([^\"]*)\"?")
)

/* 
 * ParsedSay
 * Purpose:    Takes a string command assumed to be a say action and executes
 *             it.
 * Parameters: txt - string - the text of the command.
 * Return:     r - string - text for the response to the command. This is 
 *             empty ("") if the command is successful.
 *             c - string - the color of the response to the command.
 */
func ParsedSay(txt string) (r string, c string) {
    /* Use a regexp to capture the message text of the say action. */
    matches := re.FindStringSubmatch(txt)
    if len(matches) < 2 {
        return "To talk, type 'say \"This is my message!\"'", "Red"
    }

    return "You: " + matches[1], "Green"
}