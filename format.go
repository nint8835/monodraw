package monodraw

import (
	"encoding/json"
	"fmt"
)

type Header struct {
	Version int `json:"v"`
}

type File struct {
	Header     Header                     `json:"header"`
	Objects    map[string]interface{}     `json:"-"`
	RawObjects map[string]json.RawMessage `json:"objects"`
}

func (f *File) UnmarshalJSON(b []byte) error {
	type file File

	err := json.Unmarshal(b, (*file)(f))
	if err != nil {
		return err
	}

	f.Objects = map[string]interface{}{}

	for key, raw := range f.RawObjects {
		var baseObject BaseObject

		err = json.Unmarshal(raw, &baseObject)
		if err != nil {
			return err
		}

		var object interface{}

		switch baseObject.Type {
		case ObjectTypeDocument:
			object = &Document{}
		case ObjectTypeCanvas:
			object = &Canvas{}
		case ObjectTypeGroup:
			object = &Group{}
		default:
			fmt.Println("unknown object type")
			delete(f.RawObjects, key)
			continue
			//return fmt.Errorf("unknown object type %d", baseObject.Type)
		}

		err = json.Unmarshal(raw, &object)
		if err != nil {
			return err
		}

		f.Objects[key] = object
		delete(f.RawObjects, key)
	}

	return nil
}

func (f *File) MarshalJSON() ([]byte, error) {
	type file File

	if f.Objects != nil {
		for key, val := range f.Objects {
			b, err := json.Marshal(val)
			if err != nil {
				return nil, err
			}
			f.RawObjects[key] = b
		}
	}

	return json.Marshal((*file)(f))
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

type BaseObject struct {
	Type ObjectType `json:"type_id"`
}

type ShadowIntensity int

const (
	ShadowIntensityLow    ShadowIntensity = 1
	ShadowIntensityMedium ShadowIntensity = 2
	ShadowIntensityHigh   ShadowIntensity = 3
	ShadowIntensityFull   ShadowIntensity = 4
)

func (s ShadowIntensity) String() string {
	switch s {
	case ShadowIntensityLow:
		return "Low"
	case ShadowIntensityMedium:
		return "Medium"
	case ShadowIntensityHigh:
		return "High"
	case ShadowIntensityFull:
		return "Full"
	default:
		return fmt.Sprintf("Unknown shadow intensity %d", s)
	}
}

type Shape struct {
	ShapeType         int             `json:"shape_type"` // TODO: make enum
	Clickthrough      bool            `json:"cthrough"`
	Name              string          `json:"name"`
	NameEdited        bool            `json:"name_edited"`
	Origin            string          `json:"origin"` // TODO: make custom type?
	Tag               int             `json:"tag"`
	Hidden            bool            `json:"hidden"`
	Locked            bool            `json:"locked"`
	SurfacePosEnabled bool            `json:"surface_pos_enabled"`
	Attachable        bool            `json:"attachable"`
	ContentBlended    bool            `json:"content_blended"`
	ContentColor      int             `json:"content_color"` // TODO: make a proper colour type?
	ShadowEnabled     bool            `json:"shadow_enabled"`
	ShadowOffset      string          `json:"shadow_offset"`
	ShadowTixel       string          `json:"shadow_tixel"`
	ShadowIntensity   ShadowIntensity `json:"shadow_intensity"`
	ShadowColor       int             `json:"shadow_color"`
	SubshapeRefs      []string        `json:"subshape_refs"`
	PositionRefs      []string        `json:"position_refs"`
	DocRef            string          `json:"doc_ref"`
}

type Alphabet int

const (
	AlphabetUnicode Alphabet = 1
	AlphabetAscii   Alphabet = 2
)

func (a Alphabet) String() string {
	switch a {
	case AlphabetUnicode:
		return "Unicode"
	case AlphabetAscii:
		return "ASCII"
	default:
		return fmt.Sprintf("Unknown alphabet %d", a)
	}
}

type Document struct {
	BaseObject
	Alphabet        Alphabet       `json:"alphabet"`
	ShapeSeqNumbers map[string]int `json:"shape_seq_numbers"`
	CanvasRefs      []string       `json:"canvas_refs"`
}

type Canvas struct {
	BaseObject
	Shape
	DocRef string `json:"doc_ref"`
}

type Group struct {
	BaseObject
	Shape
	SubshapeSelection bool `json:"subshape_selection"`
}
