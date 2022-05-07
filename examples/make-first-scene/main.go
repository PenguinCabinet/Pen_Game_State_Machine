package main

import (
	"image/color"
	"log"

	"github.com/PenguinCabinet/pgfsm"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//This is the title scene state
type TitleGameState struct {
	mplusNormalFont font.Face
}

//This is the function that is called when the state is first executed.
func (sm *TitleGameState) Init(
	stackdeep int, /*Here is the index of where this state is stacked on the stack*/
	delta float64, /*Here is the time that has elapsed between the previous frame and the current frame.*/
) {

	/*Here is the start of the font initialization process of Ebiten*/
	const dpi = 72

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)

	if err != nil {
		panic(err)
	}

	sm.mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	if err != nil {
		panic(err)
	}
	/*Here is the end of the font initialization process of Ebiten*/
}

//This is the function that is called every frame.
//Called only when this state is running.
func (sm *TitleGameState) Update(
	screen *ebiten.Image, /*Screen of ebiten, but it is deprecated to describe it in Update*/
	stackdeep int, delta float64,
) pgfsm.Result {
	/*Continue loop by returning an empty pgfsm.Result.
	Change the current running state to the new state by rewriting the returned pgfsm.Result or
	New states can be placed on top of the stack.*/
	return pgfsm.Result{}
}

//This is the function for drawing that is called every frame.
//Even if this state is not running, it will be called if it is on the stack.
func (sm *TitleGameState) Draw(screen *ebiten.Image, stackdeep int, delta float64) {
	text.Draw(screen, "Game Title", sm.mplusNormalFont, 200, 100, color.White)
}

func main() {

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pen_Game_State_Machine")

	gms := new(pgfsm.Machine)

	gms.LayoutWidth = 640
	gms.LayoutHeight = 480

	TitleSm := new(TitleGameState)

	/*Add the title scene state to the stack*/
	gms.StateAdd(TitleSm)

	if err := ebiten.RunGame(gms); err != nil {
		log.Fatal(err)
	}
}
