package parse

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
    res.Text = txt
    res.Color = "White"
}