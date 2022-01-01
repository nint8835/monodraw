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

### `1`: Canvas

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
Other type IDs not mentioned above did not appear in the example file. Might be worth some potential future exploration to see if they can be found / ignored.

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