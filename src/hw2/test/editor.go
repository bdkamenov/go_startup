package main

type Editor interface {
	// Insert text starting from given position.
	Insert(position int, text string) Editor

	// Delete bufferLength items from offset.
	Delete(offset, length int) Editor

	// Undo reverts latest change.
	Undo() Editor

	// Redo re-applies latest undone change.
	Redo() Editor

	// String returns complete representation of what a file looks
	// like after all manipulations.
	String() string
}

type Piece struct {
	origin bool
	offset int
	length int
}

type EditorImpl struct {
	originString string
	buffer       string
	bufferLength int
	pieces       []Piece
	undoStack    [][]Piece
	redoStack    [][]Piece
}

func NewEditor(value string) Editor {

	e := Piece{true, 0, len(value)}
	return EditorImpl{originString: value, buffer: "", bufferLength: 0, pieces: []Piece{e},
		undoStack: make([][]Piece, 0, 10), redoStack: make([][]Piece, 0, 10)}
}

func (editor *EditorImpl) filterPieces() {

	filer := editor.pieces[:0]

	for _, v := range editor.pieces {

		if v.length > 0 {
			filer = append(filer, v)
		}
	}

	editor.pieces = filer
}

func (editor *EditorImpl) pieceIndexAndOffset(offset int) (Piece, int, int) {

	remainingOffset := offset

	for i, v := range editor.pieces {

		if remainingOffset <= v.length {

			return v, i, remainingOffset
		}

		remainingOffset -= v.length
	}

	return editor.pieces[len(editor.pieces)-1], len(editor.pieces) - 1, editor.pieces[len(editor.pieces)-1].length
}

func (editor *EditorImpl) appendUndoStack() {
	buff := make([]Piece, len(editor.pieces))
	copy(buff, editor.pieces)
	editor.undoStack = append(editor.undoStack, buff)
}

func (editor EditorImpl) Insert(position int, text string) Editor {

	if len(text) == 0 || position < 0 {
		return editor
	}

	addBufferOffset := editor.bufferLength
	editor.buffer += text
	editor.bufferLength += len(text)

	piece, index, offsetInPiece := editor.pieceIndexAndOffset(position)

	editor.appendUndoStack()
	editor.redoStack = nil

	// If we are inserting at the end of the Piece and at the end of the buffer, we can just increase its bufferLength
	if !piece.origin && offsetInPiece == piece.length && (piece.offset+piece.length) == addBufferOffset {
		editor.pieces[index].length += len(text)
		return editor
	}

	editor.pieces =
		append(editor.pieces[:index],
			append([]Piece{
				{origin: piece.origin, offset: piece.offset, length: offsetInPiece},
				{origin: false, offset: addBufferOffset, length: len(text)},
				{origin: piece.origin, offset: piece.offset + offsetInPiece, length: piece.length - offsetInPiece}}, editor.pieces[index+1:]...)...)

	editor.filterPieces()
	return editor
}

func (editor EditorImpl) Delete(offset, length int) Editor {

	if offset < 0 || length <= 0 {
		return editor
	}

	if offset > editor.bufferLength+len(editor.originString) {
		return editor
	}

	stringLen := len(editor.String())
	if length > stringLen {
		length = stringLen - offset
	}

	firstPiece, firstPieceIndex, offsetInFirstPiece := editor.pieceIndexAndOffset(offset)
	lastPiece, lastPieceIndex, offsetInLastPiece := editor.pieceIndexAndOffset(offset + length)

	editor.appendUndoStack()

	// If the delete spans only one piece and is at the very start or end of the piece, we can just modify it
	if firstPieceIndex == lastPieceIndex {

		if offsetInFirstPiece == 0 {
			editor.pieces[firstPieceIndex].offset += length
			editor.pieces[firstPieceIndex].length -= length
			return editor

		} else if offsetInLastPiece == lastPiece.length {
			editor.pieces[firstPieceIndex].length -= length
			return editor
		}
	}

	editor.pieces =
		append(editor.pieces[:firstPieceIndex],
			append([]Piece{
				{origin: firstPiece.origin, offset: firstPiece.offset, length: offsetInFirstPiece},
				{origin: lastPiece.origin, offset: lastPiece.offset + offsetInLastPiece, length: lastPiece.length - offsetInLastPiece}},
				editor.pieces[lastPieceIndex+1:]...)...)

	editor.filterPieces()

	return editor
}

func (editor EditorImpl) Undo() Editor {

	l := len(editor.undoStack)

	if l == 0 {

		return editor
	}

	editor.redoStack = append(editor.redoStack, editor.pieces)
	editor.pieces = editor.undoStack[l-1]
	editor.undoStack = editor.undoStack[:l-1]

	return editor
}

func (editor EditorImpl) Redo() Editor {

	l := len(editor.redoStack)

	if l == 0 {

		return editor
	}

	editor.appendUndoStack()

	editor.pieces = editor.redoStack[l-1]
	editor.redoStack = editor.redoStack[:l-1]

	return editor
}

func (editor EditorImpl) String() string {

	var result string
	for _, v := range editor.pieces {
		if v.origin {

			result += editor.originString[v.offset : v.offset+v.length]

		} else {

			result += editor.buffer[v.offset : v.offset+v.length]
		}
	}

	return result
}
