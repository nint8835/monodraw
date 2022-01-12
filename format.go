package monodraw

import "fmt"

type Header struct {
	Version int `json:"v"`
}

type File struct {
	Header  Header            `json:"header"`
	Objects map[string]Object `json:"objects"`
}

type ObjectType int

const (
	ObjectTypeDocument                   ObjectType = 0
	ObjectTypeCanvas                     ObjectType = 1
	ObjectTypeGroup                      ObjectType = 3
	ObjectTypeLine                       ObjectType = 5
	ObjectTypeTriangle                   ObjectType = 7
	ObjectTypeLineSegment                ObjectType = 9
	ObjectTypeLineSegmentEnd             ObjectType = 10
	ObjectTypeStaticPosition             ObjectType = 11
	ObjectTypeRectPosition               ObjectType = 12
	ObjectTypeLineSegmentPosition        ObjectType = 16
	ObjectTypeRectangularContentPosition ObjectType = 17
	ObjectTypeTixmap                     ObjectType = 18
	ObjectTypeLineDiagramPosition        ObjectType = 19
	ObjectTypeRectangle                  ObjectType = 20
	ObjectTypeDiamond                    ObjectType = 21
	ObjectTypeDiamondStandardPosition    ObjectType = 22
	ObjectTypeDiamondContentPosition     ObjectType = 23
	ObjectTypeLinePointIndexPosition     ObjectType = 24
	ObjectTypeImage                      ObjectType = 25
	ObjectTypeTable                      ObjectType = 26
	ObjectTypeText                       ObjectType = 27
	ObjectTypeEllipse                    ObjectType = 28
	ObjectTypeTableCell                  ObjectType = 29
)

func (t ObjectType) String() string {
	switch t {
	case ObjectTypeDocument:
		return "Document"
	case ObjectTypeCanvas:
		return "Canvas"
	case ObjectTypeGroup:
		return "Group"
	case ObjectTypeLine:
		return "Line"
	case ObjectTypeTriangle:
		return "Triangle"
	case ObjectTypeLineSegment:
		return "LineSegment"
	case ObjectTypeLineSegmentEnd:
		return "LineSegmentEnd"
	case ObjectTypeStaticPosition:
		return "StaticPosition"
	case ObjectTypeRectPosition:
		return "RectPosition"
	case ObjectTypeLineSegmentPosition:
		return "LineSegmentPosition"
	case ObjectTypeRectangularContentPosition:
		return "RectangularContentPosition"
	case ObjectTypeTixmap:
		return "Tixmap"
	case ObjectTypeLineDiagramPosition:
		return "LineDiagramPosition"
	case ObjectTypeRectangle:
		return "Rectangle"
	case ObjectTypeDiamond:
		return "Diamond"
	case ObjectTypeDiamondStandardPosition:
		return "DiamondStandardPosition"
	case ObjectTypeDiamondContentPosition:
		return "DiamondContentPosition"
	case ObjectTypeLinePointIndexPosition:
		return "LinePointIndexPosition"
	case ObjectTypeImage:
		return "Image"
	case ObjectTypeTable:
		return "Table"
	case ObjectTypeText:
		return "Text"
	case ObjectTypeEllipse:
		return "Ellipse"
	case ObjectTypeTableCell:
		return "TableCell"
	default:
		return fmt.Sprintf("Unknown type %d", t)
	}
}

type Object interface {
	Type() ObjectType
}
