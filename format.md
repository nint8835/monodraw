# Monodraw File Format

Extension: `.monopic`

## Content

- 9 byte header (`FF 4D 4F 4E 4F 50 49 43 01`)
- Gzip-compressed JSON

## Structure

JSON contains an object with two keys:

- `header`: Object containing file info. Contains a single key, `v`, referring to the file version
- `objects`: Object with keys representing object IDs, and the values being object data.

## Object Types

The objects contained in the `objects` object can be of a variety of types, identified by the key `type_id`.

### Shape
Class name: `HTModelShape`

This object type does not have a unique type ID, but is rather a base class of most other object types

| Attribute             | Description                                                                                                                       | Type        |
|-----------------------|-----------------------------------------------------------------------------------------------------------------------------------|-------------|
| `shape_type`          | Unique identifier of the type of shape.                                                                                           | `int`       |
| `cthrough`            | Controls clickthrough (whether click events pass through this shape to the shapes underneath)                                     | `bool`      |
| `name`                | Name of the shape that will appear in the sidebar.                                                                                | `str`       |
| `name_edited`         | Whether the user has edited the name from it's default value. Will cause the app to stop trying to automatically update the name. | `bool`      |
| `origin`              | Origin point of the shape. Used to control the position. Stored as a string of format `x,y`.                                      | `str`       |
| `tag`                 |                                                                                                                                   | `int`       |
| `hidden`              | Whether or not this shape is hidden.                                                                                              | `bool`      |
| `locked`              | Whether or not this shape is locked.                                                                                              | `bool`      |
| `surface_pos_enabled` |                                                                                                                                   | `bool`      |
| `attachable`          | Whether shapes should be permitted to be attached to this shape's attachment points.                                              | `bool`      |
| `content_blended`     | Whether the shape should be blended.                                                                                              | `bool`      |
| `content_color`       | The shape's content color.                                                                                                        | `int`       |
| `shadow_enabled`      | Whether the shape has a shadow enabled.                                                                                           | `bool`      |
| `shadow_offset`       | Offset of the shadow, if enabled. Stored as a string of format `x,y`                                                              | `str`       |
| `shadow_tixel`        |                                                                                                                                   |             |
| `shadow_intensity`    | Intensity of the shadow, if enabled. Low = `1`, Medium = `2`, High = `3`, Full = `4`                                              | `int`       |
| `shadow_color`        | Color of the shadow.                                                                                                              | `int`       |
| `subshape_refs`       | Object IDs of any shapes contained within / belonging to this shape.                                                              | `List[str]` |
| `position_refs`       | Object IDs of any positioning-related objects belonging to this shape.                                                            | `List[str]` |


### `0`: Document
Class name: `HTModelDocument`

Root type all other object types are eventually contained within.

| Attribute           | Description                                                                      | Type             |
|---------------------|----------------------------------------------------------------------------------|------------------|
| `alphabet`          | Type of alphabet to use for this document. `1` for Unicode, `2` for ASCII.       | `int`            |
| `shape_seq_numbers` | Unknown. Appears to be `{"6": 1}` in all files I've inspected so far.            | `Dict[str, int]` |
| `canvas_refs`       | Array of object IDs referring to the [Canvases](#1-canvas) within this document. | `List[str]`      |

### `1`: Canvas
Class name: `HTModelCanvasShape`

Type that all other objects will go placed in. Effectively the root of the tree displayed in the sidebar.

If a file does not contain a canvas, attempting to edit it will do nothing.

Inherits from [Shape](#shape).

| Attribute | Description                                                      | Type  |
|-----------|------------------------------------------------------------------|-------|
| `doc_ref` | Object ID of the [Document](#0-document) containing this canvas. | `str` |

### `3`: Group
Class name: `HTModelGroupShape`

A group of objects.

Inherits from [Shape](#shape).

| Attribute             | Description                                                                                                           | Type            |
|-----------------------|-----------------------------------------------------------------------------------------------------------------------|-----------------|
| `doc_ref`             | Object ID of the [Document](#0-document) containing this group.                                                       | `str`           |
| `subshape_selection`  | Value of the "Allow sub-shape selection" option on the group. Will not be present if the option hasn't been modified. | `bool`          |

### `5`: Line
Class name: `HTModelLineShape`

### `7`: Triangle
Class name: `HTModelTriangleShape`

### `9`: Line Segment
Class name: `HTModelLineSegment`

### `10`: Line Segment End
Class name: `HTModelLineSegmentEnd`

### `11`: Static Position
Class name: `HTModelStaticPosition`

### `12`: Rect Position
Class name: `HTModelRectPosition`

### `16`: Line Segment Position
Class name: `HTModelLineSegmentPosition`

### `17`: Rectangular Content Position
Class name: `HTModelRectangularContentPosition`

### `18`: Tixmap
Class name: `HTModelTixmap`

### `19`: Line Diagram Position
Class name: `HTModelLineDiagramPosition`

### `20`: Rectangle
Class name: `HTModelRectangularShape`

### `21`: Diamond
Class name: `HTModelDiamondShape`

### `22`: Diamond Standard Position
Class name: `HTModelDIamondStandardPosition`

### `23`: Diamond Content Position
Class name: `HTModelDiamondContentPosition`

### `24`: Line Point Index Position
Class name: `HTModelLinePointIndexPosition`

### `25`: Image
Class name: `HTModelImageShape`

### `26`: Table
Class name: `HTModelTableShape`

### `27`: Text
Class name: `HTModelText`

### `28`: Ellipse
Class name: `HTModelEllipseShape`

### `29`: Table Cell
Class name: `HTModelTableCell`
