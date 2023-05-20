package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iwpnd/tyler"
	"github.com/urfave/cli/v2"
)

var fromExtentCommand cli.Command
var extentFlag cli.Float64SliceFlag
var zoomFlag cli.IntFlag

func tilesFromExtent(ctx *cli.Context) error {
	e := ctx.Float64Slice("extent")
	z := ctx.Int("zoom")

	if z >= 18 || z < 0 {
		return fmt.Errorf("zoom must be between 0 - 20, got: %d", z)
	}

	b := tyler.Extent{MinLng: e[0], MinLat: e[1], MaxLng: e[2], MaxLat: e[3]}
	tiles, err := b.ToTiles(z)
	if err != nil {
		return err
	}

	fc, err := tyler.ToFeatureCollection(tiles)
	if err != nil {
		return err
	}

	fmt.Println(string(fc))

	return nil
}

func init() {
	extentFlag = cli.Float64SliceFlag{
		Name:     "extent",
		Required: true,
		Usage:    "extent as [west, south, east, north]",
	}

	zoomFlag = cli.IntFlag{
		Name:     "zoom",
		Usage:    "zoom level 0-20",
		Required: true,
	}

	fromExtentCommand = cli.Command{
		Name:   "from-extent",
		Usage:  "feature collection of tiles intersecting a giving extent",
		Action: tilesFromExtent,
		Flags: []cli.Flag{
			&extentFlag,
			&zoomFlag,
		},
	}

}

func main() {
	app := &cli.App{
		Name:  "tyler",
		Usage: "tyler",
		Commands: []*cli.Command{
			&fromExtentCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
