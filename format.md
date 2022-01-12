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

| Attribute             | Description                                                                                                                   | Type            |
|-----------------------|-------------------------------------------------------------------------------------------------------------------------------|-----------------|
| `locked`              | Locked state of this canvas.                                                                                                  | `bool`          |
| `hidden`              | Hidden state of this canvas.                                                                                                  | `bool`          |
| `shadow_enabled`      |                                                                                                                               | `bool`          |
| `name_edited`         | Denotes whether or not the name has been changed from it's default value by the user. Likely not triggerable by normal usage. | `bool`          |
| `shadow_color`        |                                                                                                                               | `int`           |
| `content_blended`     |                                                                                                                               | `bool`          |
| `shape_type`          | Type of shape. Will be `6` for canvas.                                                                                        | `int`           |
| `origin`              |                                                                                                                               | `str`           |
| `tag`                 |                                                                                                                               | `int`           |
| `name`                | Name of this canvas.                                                                                                          | `str`           |
| `position_refs`       |                                                                                                                               | `List[Unknown]` |
| `doc_ref`             | Object ID of the [Document](#0-document) containing this canvas.                                                              | `str`           |
| `surface_pos_enabled` |                                                                                                                               | `bool`          |
| `content_color`       |                                                                                                                               | `int`           |
| `subshape_refs`       | Object IDs of other objects contained within this canvas.                                                                     | `List[str]`     |
| `shadow_offset`       |                                                                                                                               | `str`           |
| `shadow_intensity`    |                                                                                                                               | `int`           |
| `shadow_tixel`        |                                                                                                                               | `str`           |
| `cthrough`            |                                                                                                                               | `bool`          |
| `attachable`          |                                                                                                                               | `bool`          |

### `3`: Group
Class name: `HTModelGroupShape`

A group of objects.

| Attribute             | Description                                                                                                           | Type            |
|-----------------------|-----------------------------------------------------------------------------------------------------------------------|-----------------|
| `locked`              | Locked state of this group.                                                                                           | `bool`          |
| `hidden`              | Hidden state of this group.                                                                                           | `bool`          |
| `shadow_enabled`      |                                                                                                                       | `bool`          |
| `name_edited`         | Denotes whether or not the name has been changed from it's default value by the user.                                 | `bool`          |
| `shadow_color`        |                                                                                                                       | `int`           |
| `content_blended`     |                                                                                                                       | `bool`          |
| `shape_type`          | Type of shape. Will be `8` for group.                                                                                 | `int`           |
| `origin`              |                                                                                                                       | `str`           |
| `tag`                 |                                                                                                                       | `int`           |
| `name`                | Name of this group.                                                                                                   | `str`           |
| `position_refs`       |                                                                                                                       | `List[Unknown]` |
| `doc_ref`             | Object ID of the [Document](#0-document) containing this group.                                                       | `str`           |
| `surface_pos_enabled` |                                                                                                                       | `bool`          |
| `content_color`       |                                                                                                                       | `int`           |
| `subshape_refs`       | Object IDs of other objects contained within this group.                                                              | `List[str]`     |
| `shadow_offset`       |                                                                                                                       | `str`           |
| `shadow_intensity`    |                                                                                                                       | `int`           |
| `shadow_tixel`        |                                                                                                                       | `str`           |
| `cthrough`            |                                                                                                                       | `bool`          |
| `attachable`          |                                                                                                                       | `bool`          |
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
