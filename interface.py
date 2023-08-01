from textual.app import App, ComposeResult
from textual.widgets import Static


class GameInterface(App):
    CSS_PATH = "GameInterface.css"

    def compose(self) -> ComposeResult:
        yield Static("Location", classes="box")
        yield Static("Instruments", classes="box")


if __name__ == "__main__":
    app = GameInterface()
    app.run()
