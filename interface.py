from textual.app import App, ComposeResult
from textual.widgets import Static, Label

class GameInterface(App):
    CSS_PATH = "GameInterface.css"

    def compose(self) -> ComposeResult:
        location = Label(id="location")
        location.border_title = "Location"
        instruments = Label(id="instruments")
        instruments.border_title = "Instruments"
        yield location
        yield instruments


if __name__ == "__main__":
    app = GameInterface()
    app.run()
