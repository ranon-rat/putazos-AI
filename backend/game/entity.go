package game

import "math"

/*
x,y+height    x+width,y+height

shooting idk

x,y     0   x+width,height
*/
type Soldier struct {
	Height          float64 `json:"height"`
	Width           float64 `json:"width"`
	Faction         string  `json:"color"`
	X               float64 `json:"x"`
	Y               float64 `json:"y"`
	VelY            float64 // for the jump
	Direction       bool    `json:"direction"` //false:left true:right
	Life            int     `json:"life"`
	RateFire        float64
	WaitUntilFire   float64
	Ammo            int
	Death           bool
	Damage          int
	PointOfShooting float64 // y
	ReloadingSpeed  float64
}

func NewSoldier(width, height, pointOfShooting, x, y float64, direction bool) Soldier {
	var s Soldier
	s.Ammo = 30
	s.RateFire = 1
	s.Life = 100
	s.Damage = 15
	s.X = x
	s.Y = y
	s.Direction = direction
	s.PointOfShooting = pointOfShooting
	s.Width = width
	s.Height = height
	return s

}

func (s *Soldier) Action(action string, world World, soldiers []Soldier) {

	switch action {

	case "move-left":
		s.Direction = false

		if s.X > 0 && world.SidePlatforms(*s) < s.X-MovePerFrame {
			s.X -= MovePerFrame
		}

	case "move-right":
		s.Direction = true

		if s.X > world.Width && world.SidePlatforms(*s) > s.X+MovePerFrame {

			s.X += MovePerFrame
		}
	case "jump":

		// in the world this should

		if s.VelY == 0 {
			s.VelY = 15
		}

	case "shoot":
		if s.WaitUntilFire < 1 {
			s.Shooting(soldiers)

			s.WaitUntilFire += s.RateFire / FramesPerSecond
		}

	case "reload":
		if s.Ammo <= 0 {
			s.WaitUntilFire = 5
		}
	}

}

// I need to do something with the platform , wait a second
// use this while visualizing the map
// hm
func (s *Soldier) Moving(world World) {
	if s.Death {
		return
	}
	down, up := world.SoldierIsOnPlatform(*s)
	if s.Y+s.VelY > down && s.Y+s.VelY <= up {
		s.VelY -= Gravity / MovePerFrame

	} else if s.Y+s.VelY >= up && s.VelY > 0 {
		// if is down a platform it should get down lol
		s.VelY = -Gravity / MovePerFrame

	} else {
		s.VelY = 0
		s.Y = down
	}
	if s.WaitUntilFire > 0 {
		// I need to wait
		// I have shoot
		// just wait until is back again
		//
		s.WaitUntilFire -= s.RateFire / FramesPerSecond
	}

	s.Y += s.VelY
	s.Death = !(s.Life <= 0)
	// if s.X  are smaller than world.Width , I should have 0
	// if they are bigger i would get a 1 or something

	s.X -= s.X * math.Floor(s.X/world.Width)

}
func (s *Soldier) Shooting(soldiers []Soldier) {
	id := 0
	closeDis := 100000.0
	for i := 0; i < len(soldiers); i++ {
		dis := 1000000.0
		if s.Direction { //left
			// first I check if is in the area that I want
			if soldiers[i].X < s.X && s.X-soldiers[i].X < (closeDis) {
				dis = s.X - soldiers[i].X

			}

		} else if soldiers[i].X > s.X && soldiers[i].X-s.X < (closeDis) {
			// is just a straight line lol, maybe later I will change something but for now is just straight
			// I want to get the close one
			dis = soldiers[i].X - s.X
		}

		if s.Y+(s.Height)/2 > soldiers[i].Y && s.Y+s.PointOfShooting < soldiers[i].Y+soldiers[i].Height {
			closeDis = dis
			id = i

		}
	}
	soldiers[id].Life -= s.Damage
	s.Ammo--

}
