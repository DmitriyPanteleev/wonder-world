from textual.app import App, ComposeResult
from textual.widgets import Label

class GameInterface(App):
    CSS_PATH = "GameInterface.css"

    BINDINGS = [
        ("a", "move_left", "Move left"),
        ("d", "moveright", "Move right"),
        ("w", "up", "Move up"),
        ("s", "down", "Move down"),
    ]

    def __init__(self):
        super().__init__()
        self.location = Label(id="location")
        self.location.border_title = "Location"
        self.location.text = " " * 80  # Initialize the label with spaces
        self.x_pos = 10  # Initialize the X position

    async def move_left(self):
        # Move the X symbol one position to the left
        if self.x_pos > 0:
            self.x_pos -= 1
            self.update_location()

    async def move_right(self):
        # Move the X symbol one position to the right
        if self.x_pos < 79:
            self.x_pos += 1
            self.update_location()

    async def move_up(self):
        # Move the X symbol one position up
        if self.location.text[self.x_pos] != "\n":
            self.x_pos -= 80
            self.update_location()

    async def move_down(self):
        # Move the X symbol one position down
        if self.x_pos < 80 * 23 and self.location.text[self.x_pos + 80] != "\n":
            self.x_pos += 80
            self.update_location()

    def update_location(self):
        # Update the location label with the new X position
        self.location.text = self.location.text[:self.x_pos] + "X" + self.location.text[self.x_pos + 1:]

    def compose(self) -> ComposeResult:
        yield self.location

if __name__ == "__main__":
    app = GameInterface()
    app.run()
