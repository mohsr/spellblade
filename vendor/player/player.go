package player

/*
 * Represents a player in the game.
 * Schema: {"Name":  The name of the character.
 *          "Class": The class of the character.
 *          "Height": The height of the character.
 *          "Hair": The hair color of the character.
 *          "Level": The level of the character.}
 */
type Player struct {
    Name   string
    Class  string
    Height string
    Hair   string
    Level  int
}

/*
 * NewPlayer
 * Purpose:    Generate a new Player.
 * Parameters: name - string - the name of the character.
 *             class - string - the class of the character.
 *             height - string - the height of the character.
 *             hair - string - the hair color of the character.
 * Return:     The newly generated Player instance.
 */
func NewPlayer(name string, class string, height string, hair string) Player {
    p := Player{Name: name, Class: class, Height: height, Hair: hair, Level: 1}
    return p
}