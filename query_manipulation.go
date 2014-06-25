package gorethink

import (
	p "github.com/dancannon/gorethink/ql2"
)

// Returns the currently visited document.
var Row = newRqlTerm("Doc", p.Term_IMPLICIT_VAR, []interface{}{}, map[string]interface{}{})

func Literal(args ...interface{}) RqlTerm {
	return newRqlTerm("Literal", p.Term_LITERAL, args, map[string]interface{}{})
}

// Get a single field from an object. If called on a sequence, gets that field
// from every object in the sequence, skipping objects that lack it.
func (t RqlTerm) Field(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Field", p.Term_GET_FIELD, args, map[string]interface{}{})
}

// Test if an object has all of the specified fields. An object has a field if
// it has the specified key and that key maps to a non-null value. For instance,
//  the object `{'a':1,'b':2,'c':null}` has the fields `a` and `b`.
func (t RqlTerm) HasFields(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "HasFields", p.Term_HAS_FIELDS, args, map[string]interface{}{})
}

// Plucks out one or more attributes from either an object or a sequence of
// objects (projection).
func (t RqlTerm) Pluck(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Pluck", p.Term_PLUCK, args, map[string]interface{}{})
}

// The opposite of pluck; takes an object or a sequence of objects, and returns
// them with the specified paths removed.
func (t RqlTerm) Without(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Without", p.Term_WITHOUT, args, map[string]interface{}{})
}

// Merge two objects together to construct a new object with properties from both.
// Gives preference to attributes from other when there is a conflict.
func (t RqlTerm) Merge(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Merge", p.Term_MERGE, funcWrapArgs(args), map[string]interface{}{})
}

// Append a value to an array.
func (t RqlTerm) Append(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Append", p.Term_APPEND, args, map[string]interface{}{})
}

// Prepend a value to an array.
func (t RqlTerm) Prepend(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Prepend", p.Term_PREPEND, args, map[string]interface{}{})
}

// Remove the elements of one array from another array.
func (t RqlTerm) Difference(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Difference", p.Term_DIFFERENCE, args, map[string]interface{}{})
}

// Add a value to an array and return it as a set (an array with distinct values).
func (t RqlTerm) SetInsert(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "SetInsert", p.Term_SET_INSERT, args, map[string]interface{}{})
}

// Add a several values to an array and return it as a set (an array with
// distinct values).
func (t RqlTerm) SetUnion(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "SetUnion", p.Term_SET_UNION, args, map[string]interface{}{})
}

// Intersect two arrays returning values that occur in both of them as a set (an
// array with distinct values).
func (t RqlTerm) SetIntersection(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "SetIntersection", p.Term_SET_INTERSECTION, args, map[string]interface{}{})
}

// Remove the elements of one array from another and return them as a set (an
// array with distinct values).
func (t RqlTerm) SetDifference(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "SetDifference", p.Term_SET_DIFFERENCE, args, map[string]interface{}{})
}

// Insert a value in to an array at a given index. Returns the modified array.
func (t RqlTerm) InsertAt(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "InsertAt", p.Term_INSERT_AT, args, map[string]interface{}{})
}

// Insert several values in to an array at a given index. Returns the modified array.
func (t RqlTerm) SpliceAt(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "SpliceAt", p.Term_SPLICE_AT, args, map[string]interface{}{})
}

// Remove an element from an array at a given index. Returns the modified array.
func (t RqlTerm) DeleteAt(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "DeleteAt", p.Term_DELETE_AT, args, map[string]interface{}{})
}

// Change a value in an array at a given index. Returns the modified array.
func (t RqlTerm) ChangeAt(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "ChangeAt", p.Term_CHANGE_AT, args, map[string]interface{}{})
}

// Return an array containing all of the object's keys.
func (t RqlTerm) Keys(args ...interface{}) RqlTerm {
	return newRqlTermFromPrevVal(t, "Keys", p.Term_KEYS, args, map[string]interface{}{})
}

// Creates an object from a list of key-value pairs, where the keys must be strings.
func Object(args ...interface{}) RqlTerm {
	return newRqlTerm("Object", p.Term_OBJECT, args, map[string]interface{}{})
}
