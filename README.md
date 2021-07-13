[![GoDoc Reference](https://pkg.go.dev/badge/github.com/sf9v/packme)](https://pkg.go.dev/github.com/sf9v/packme)
![Github Actions](https://github.com/sf9v/packme/workflows/test/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/sf9v/packme/badge.svg?branch=main)](https://coveralls.io/github/sf9v/packme?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/sf9v/packme)](https://goreportcard.com/report/github.com/sf9v/packme)

# PackMe

A 3D bin packing library in Go.

## Packing implementations

All packing implementations comply to [Packer](packer.go) interface.

- [Best-fit packer](best_fit_packer.go) - A 3D bin packing implementation based on a paper by Dube, E., & Kanavathy L. (2006) titled *Optimizing Three-Dimensional Bin Packing Through Simulation*

## Example

```go
// Define box specifications
boxSpecs := []packme.BoxSpec{
    packme.NewBoxSpec("Box A", 1, packme.NewDimensions(30, 30, 30)),
    packme.NewBoxSpec("Box B", 1, packme.NewDimensions(5, 5, 40)),
    packme.NewBoxSpec("Box C", 1, packme.NewDimensions(20, 20, 30)),
}

// Define item specifications
itemSpecs := []packme.ItemSpec{
    packme.NewItemSpec("Item A1", 17, packme.NewDimensions(10, 10, 30)),
    packme.NewItemSpec("Item A2", 1, packme.NewDimensions(10, 10, 30)),
    packme.NewItemSpec("Tall Item", 1, packme.NewDimensions(5, 39.5, 5)),
    packme.NewItemSpec("Large Item", 1, packme.NewDimensions(50, 50, 100)),
}

// Create new instance of best-fit packer
packer := packme.NewBestFitPacker()

// Start packing!
packingScheme := packer.Pack(boxes, items)

// Print packing scheme
fmt.Println(packingScheme)

// Output:
// Box B items(1) dims(5x5x40) vol(1000)
// Tall Item dims(5x5x39.5) vol(987.5) pos(0,0,0) rot(h,l,w)
// 
// Box C items(4) dims(20x20x30) vol(12000)
// Item A1 dims(10x10x30) vol(3000) pos(0,0,0) rot(l,w,h)
// Item A1 dims(10x10x30) vol(3000) pos(10,0,0) rot(l,w,h)
// Item A1 dims(10x10x30) vol(3000) pos(0,10,0) rot(l,w,h)
// Item A1 dims(10x10x30) vol(3000) pos(10,10,0) rot(l,w,h)
//
// Box A items(9) dims(30x30x30) vol(27000)
// Item A1 dims(10x10x30) vol(3000) pos(0,0,0) rot(l,w,h)
// Item A1 dims(10x10x30) vol(3000) pos(10,0,0) rot(l,w,h)
// Item A1 dims(10x10x30) vol(3000) pos(20,0,0) rot(l,w,h)
// Item A1 dims(10x10x30) vol(3000) pos(0,10,0) rot(l,w,h)
// Item A1 dims(10x10x30) vol(3000) pos(10,10,0) rot(l,w,h)
// Item A1 dims(10x10x30) vol(3000) pos(20,10,0) rot(l,w,h)
// Item A1 dims(10x10x30) vol(3000) pos(0,20,0) rot(l,w,h)
// Item A1 dims(10x10x30) vol(3000) pos(10,20,0) rot(l,w,h)
// Item A1 dims(10x10x30) vol(3000) pos(20,20,0) rot(l,w,h)
//
// Not packed (6)
// Item A1 dims(10x30x10) vol(3000) pos(0,0,0) rot(l,h,w)
// Item A1 dims(10x30x10) vol(3000) pos(0,0,0) rot(l,h,w)
// Item A1 dims(10x30x10) vol(3000) pos(0,0,0) rot(l,h,w)
// Item A1 dims(10x30x10) vol(3000) pos(0,0,0) rot(l,h,w)
// Item A2 dims(10x30x10) vol(3000) pos(0,0,0) rot(l,h,w)
// Large Item dims(50x100x50) vol(250000) pos(0,0,0) rot(l,h,w)
```

## Credits

- [Dube, E., & Kanavathy L. (2006). *Optimizing Three-Dimensional Bin Packing Through Simulation.*](https://www.researchgate.net/publication/228974015_Optimizing_Three-Dimensional_Bin_Packing_Through_Simulation)
- [3D Bin Packing by Enzo Ruiz](https://github.com/enzoruiz/3dbinpacking)

## Licence

[MIT](LICENSE)