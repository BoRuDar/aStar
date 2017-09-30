package main

import (
	"math"
	"strconv"
	"github.com/fogleman/gg"
	"io/ioutil"
	"log"
	"encoding/json"
)

const (
	FREE = iota
	BLOCKED
	START
	END
	PATH
)

var (
	START_POINT, END_POINT Point
	WIDTH, HEIGHT int
	matrix [][]Point
)

type Point struct {
	x, y, state int
	H, G, F int
	parent *Point
}

type Points map[string]Point

type Coord struct {
	X,Y int
}

type Config struct {
	Width int
	Height int
	Start Coord
	End Coord
	Obstacles []Coord
}

func main()  {

	log.Println("Reading config file...")
	conf, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("Cannot open config.json")
	}

	var j Config
	err = json.Unmarshal(conf,&j)
	if err != nil {
		log.Fatal("Cannot decode JSON from file ")
	}
	log.Println("Config file is read!")

	//init configs
	log.Println("Initializing...")
	START_POINT := Point{x: j.Start.X, y: j.Start.Y, state:START}
	END_POINT := Point{x: j.End.X, y: j.End.Y, state:END}
	WIDTH = j.Width
	HEIGHT = j.Height

	//make array of array of Point
	matrix := make([][]Point,WIDTH, WIDTH*HEIGHT)
	for i:=0; i<len(matrix);i ++ {
		matrix[i] = make([]Point, HEIGHT)
	}

	//init array [][]Point
	for x := 0; x < WIDTH; x++ {
		for y := 0; y < HEIGHT; y++ {
			matrix[x][y] = Point{x:x, y:y, state:FREE}
		}
	}

	open, close := Points{}, Points{}
	open[START_POINT.Key()] = START_POINT

	//set start & finish
	matrix[START_POINT.x][START_POINT.y] = START_POINT
	matrix[END_POINT.x][END_POINT.y] = END_POINT

	//set obstacles
	for _, o := range j.Obstacles {
		matrix[o.X][o.Y].state = BLOCKED
	}

	log.Println("Initialization is completed!")

	for {
		//get new Point with min F value
		current := *MinF(open)

		//when we find the end
		if current.Equal(END_POINT) {
			log.Println("Path is found")

			log.Print("Points in path: ")
			for !current.Equal(START_POINT) {
				current = *current.parent
				if !current.Equal(START_POINT){
					matrix[current.x][current.y].state = PATH
					log.Print(matrix[current.x][current.y], " ")
				}
			}
			break
		}
		parseNeighbours(current, &matrix, &open, &close)
	}

	drawResult(&matrix)
	log.Println("Exited!")
//end
}

func parseNeighbours(curr Point, m *[][]Point, open, close *Points) {
	delete(*open, curr.Key())
	(*close)[curr.Key()] = curr

	nCoord := generateNeighboursCoord(curr)

	for _, c := range nCoord{
		tmpPoint := (*m)[c.x][c.y]

		if _, inClose := (*close)[tmpPoint.Key()]; inClose || tmpPoint.state == BLOCKED {
			continue
		}

		if _, inOpen := (*open)[tmpPoint.Key()]; inOpen{
			continue
		}

		tmpPoint.G = curr.GetG(tmpPoint)
		tmpPoint.H = GetH(tmpPoint, END_POINT)
		tmpPoint.F = tmpPoint.GetF()
		tmpPoint.parent = &curr //ref is needed?

		(*open)[tmpPoint.Key()] = tmpPoint
	}
}

func GetH(a, b Point) int {
	tmp := math.Abs(float64(a.x - b.x))
	tmp += math.Abs(float64(a.y - b.y))

	return int(tmp)
}

func (current Point) GetG(target Point) int {
	if target.x != current.x &&
		target.y != current.y {
		return current.G + 14
	}

	return current.G + 10
}

func (p Point) GetF() int {
	return p.G + p.H
}

func (p Point) Key() string {
	return strconv.Itoa(p.x) + ":" + strconv.Itoa(p.y)
}

func (a Point) Equal(b Point) bool {
	return a.x == b.x && a.y == b.y
}

func MinF(points Points) (min *Point){
	min = &Point{F:WIDTH*HEIGHT*10+1}

	for _, p := range points{
		if p.F < min.F {
			*min = p
		}
	}
	return
}

func addCoordIfValid(coords *[]Point, x,y  int){

	if x >= 0  && y >= 0 &&
		x < WIDTH && y < HEIGHT{
		*coords = append(*coords, Point{x:x, y:y})
	}
}

func generateNeighboursCoord(curr Point) (res []Point)  {

	//top left
	addCoordIfValid(&res, curr.x -1, curr.y +1)
	//top middle
	addCoordIfValid(&res, curr.x, curr.y +1)
	//top right
	addCoordIfValid(&res, curr.x +1, curr.y +1)

	//left
	addCoordIfValid(&res, curr.x-1, curr.y)
	//right
	addCoordIfValid(&res, curr.x+1, curr.y)

	//bottom left
	addCoordIfValid(&res, curr.x -1, curr.y -1)
	//bottom middle
	addCoordIfValid(&res, curr.x, curr.y -1)
	//bottom right
	addCoordIfValid(&res, curr.x +1, curr.y -1)

	return
}

func drawResult(m *[][]Point)  {
	log.Println("Drawing started")
	fileName := "out.png"
	//radius
	r := WIDTH*HEIGHT/4
	//init image
	dc := gg.NewContext(r*2*10 + 10, r*2*10 + 10)

	//set color & draw
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			dc.DrawCircle( float64(10 + x*r*2 + WIDTH*2), float64(10 + y*r*2 + HEIGHT*2), float64(WIDTH*HEIGHT/4) -0.5)
			switch (*m)[x][y].state {
			case BLOCKED:
				dc.SetRGB(255, 255, 255)
			case START:
				dc.SetRGB(125, 0, 0)
			case END:
				dc.SetRGB(0, 125, 0)
			case PATH:
				dc.SetRGB(0, 0, 125)
			default:
				dc.SetRGB(50, 50, 50)
			}
			dc.Fill()
		}
	}
	dc.SavePNG(fileName)
	log.Println("Drawing finished and saved to file - ", fileName)
}