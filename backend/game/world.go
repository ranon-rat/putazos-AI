package game

const (
	MovePerFrame    = 10.0 / FramesPerSecond
	FramesPerSecond = 30.0 // change it if you want but the frames per second are  fixed , and i will do nothing for change it
	Gravity         = 1    // stupid ass song
)

type Platform struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type World struct {
	Width     float64    `json:"width"`
	Height    float64    `json:"height"`
	Platforms []Platform `json:"platforms"`
}

func (w World) SoldierIsOnPlatform(soldier Soldier) (float64, float64) {

	down := 0.0
	downDis := 0.0
	lowDown := 100000.0
	lowDisDown := 100000.0

	up := w.Height
	upDis := 0.0
	lowUp := 100000.0
	lowDisUp := 100000.0
	// i hate the antichrist
	for i := 0; i < len(w.Platforms); i++ {
		// with this i get the platform that its up
		if soldier.Y <= (w.Platforms[i].Y) &&
			(soldier.X <= (w.Platforms[i].X+w.Platforms[i].Width) && soldier.X >= w.Platforms[i].X) {
			upDis = (w.Platforms[i].Y) - soldier.Y
			up = w.Platforms[i].Y
		}

		// the platform that is down
		if soldier.Y >= (w.Platforms[i].Y+w.Platforms[i].Height) &&
			(soldier.X <= (w.Platforms[i].X+w.Platforms[i].Width) && soldier.X >= w.Platforms[i].X) {
			downDis = soldier.Y - (w.Platforms[i].Y + w.Platforms[i].Height)
			down = w.Platforms[i].Y + w.Platforms[i].Height
		}
		if lowDisUp > upDis && upDis >= 0 {
			lowDisUp = upDis
			lowUp = up
		}
		if lowDisDown > downDis && downDis >= 0 {
			lowDown = down
			lowDisDown = downDis
		}

	}
	return lowDown, lowUp
}
func (w World) SidePlatforms(soldier Soldier) (x float64) {
	closeDis := 100000.0
	for i := 0; i < len(w.Platforms); i++ {
		dis := 1000000.0
		dx := 0.0
		if soldier.Direction { //left

			// first I check if is in the area that I want
			// i hate this shit
			if w.Platforms[i].X+w.Platforms[i].Width < soldier.X &&
				soldier.X-(w.Platforms[i].X+w.Platforms[i].Width) < (closeDis) {
				dis = soldier.Y - w.Platforms[i].X
				dx = w.Platforms[i].X + w.Platforms[i].Width

			}

		} else if w.Platforms[i].X > soldier.X+soldier.Width && w.Platforms[i].X-soldier.X < (closeDis) {
			// is just a straight line lol, maybe later I will change something but for now is just straight
			// I want to get the close one
			dis = w.Platforms[i].X - soldier.X
			dx = w.Platforms[i].X
		}

		if soldier.Y+soldier.Height > w.Platforms[i].Y && soldier.Y < w.Platforms[i].Y+w.Platforms[i].Height {
			closeDis = dis
			x = dx

		}
	}
	return

}
