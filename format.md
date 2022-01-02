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

Root type all other object types are eventually contained within.

| Attribute           | Description                                                                | Type             |
|---------------------|----------------------------------------------------------------------------|------------------|
| `alphabet`          | Type of alphabet to use for this document. `1` for Unicode, `2` for ASCII. | `int`            |
| `shape_seq_numbers` | Unknown. Appears to be `{"6": 1}` in all files I've inspected so far.      | `Dict[str, int]` |
| `canvas_refs`       | Array of object IDs referring to the canvases within this document.        | `List[str]`      |

### `1`: Canvas

Type that all other objects will go placed in. Effectively the root of the tree displayed in the sidebar.

Based on the keys appearing, this is likely a variant of [Group](#3-group), but this will require more research to
validate.

| Attribute             | Description                                                      | Type            |
|-----------------------|------------------------------------------------------------------|-----------------|
| `locked`              |                                                                  | `bool`          |
| `hidden`              |                                                                  | `bool`          |
| `shadow_enabled`      |                                                                  | `bool`          |
| `name_edited`         |                                                                  | `bool`          |
| `shadow_color`        |                                                                  | `int`           |
| `content_blended`     |                                                                  | `bool`          |
| `shape_type`          |                                                                  | `int`           |
| `origin`              |                                                                  | `str`           |
| `tag`                 |                                                                  | `int`           |
| `name`                | Name of this canvas.                                             | `str`           |
| `position_refs`       |                                                                  | `List[Unknown]` |
| `doc_ref`             | Object ID of the [Document](#0-document) containing this canvas. | `str`           |
| `surface_pos_enabled` |                                                                  | `bool`          |
| `content_color`       |                                                                  | `int`           |
| `subshape_refs`       | Object IDs of other objects contained within this canvas.        | `List[str]`     |
| `shadow_offset`       |                                                                  | `str`           |
| `shadow_intensity`    |                                                                  | `int`           |
| `shadow_tixel`        |                                                                  | `str`           |
| `cthrough`            |                                                                  | `bool`          |
| `attachable`          |                                                                  | `bool`          |

### `3`: Group

### `5`: Line

### `9`: Line Segment

### `10`: Line Segment End

### `11`: Frame Anchor?

### `12`: Frame Position?

### `16`: Line Segment Position

### `17`: Line Segment End Position

### `18`: Something about `border tixels`?

### `19`: Something to do with diagonal lines?

### `20`: Rectangle

### `21`: Diamond

### `22`: Diamond Position?

### `24`: Something related to lines & rectangles, dealing with an anchor or a position

### `27`: Text

### `28`: Ellipse

### Other

Other type IDs not mentioned above did not appear in the example file. Might be worth some potential future exploration
to see if they can be found / ignored.

Missing values:

- `2`
- `4`
- `6`
- `7`
- `8`
- `13`
- `14`
- `15`
- `23`
- `25`
- `26`