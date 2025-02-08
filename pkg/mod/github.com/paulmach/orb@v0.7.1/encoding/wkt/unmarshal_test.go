package wkt

import (
	"testing"

	"github.com/paulmach/orb"
)

func TestTrimSpaceBrackets(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "single point",
			s:        "(1 2)",
			expected: "1 2",
		},
		{
			name:     "double brackets",
			s:        "((1 2),(0.5 1.5))",
			expected: "(1 2),(0.5 1.5)",
		},
		{
			name:     "multiple values",
			s:        "(1 2,0.5 1.5)",
			expected: "1 2,0.5 1.5",
		},
		{
			name:     "multiple points",
			s:        "((1 2,3 4),(5 6,7 8))",
			expected: "(1 2,3 4),(5 6,7 8)",
		},
		{
			name:     "triple brackets",
			s:        "(((1 2,3 4)),((5 6,7 8),(1 2,5 4)))",
			expected: "((1 2,3 4)),((5 6,7 8),(1 2,5 4))",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if trimSpaceBrackets(tc.s) != tc.expected {
				t.Log(trimSpaceBrackets(tc.s))
				t.Log(tc.expected)
				t.Errorf("trim space and brackets error")
			}
		})
	}
}

func TestUnmarshalPoint(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		expected orb.Point
	}{
		{
			name:     "int",
			s:        "POINT(1 2)",
			expected: orb.Point{1, 2},
		},
		{
			name:     "float64",
			s:        "POINT(1.34 2.35)",
			expected: orb.Point{1.34, 2.35},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			geom, err := UnmarshalPoint(tc.s)
			if err != nil {
				t.Fatal(err)
			}

			if !geom.Equal(tc.expected) {
				t.Log(geom)
				t.Log(tc.expected)
				t.Errorf("incorrect wkt unmarshalling")
			}
		})
	}
}

func TestUnmarshalMultiPoint(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		expected orb.MultiPoint
	}{
		{
			name:     "empty",
			s:        "MULTIPOINT EMPTY",
			expected: orb.MultiPoint{},
		},
		{
			name:     "1 point",
			s:        "MULTIPOINT((1 2))",
			expected: orb.MultiPoint{{1, 2}},
		},
		{
			name:     "2 points",
			s:        "MULTIPOINT((1 2),(0.5 1.5))",
			expected: orb.MultiPoint{{1, 2}, {0.5, 1.5}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			geom, err := UnmarshalMultiPoint(tc.s)
			if err != nil {
				t.Fatal(err)
			}

			if !geom.Equal(tc.expected) {
				t.Log(geom)
				t.Log(tc.expected)
				t.Errorf("incorrect wkt unmarshalling")
			}
		})
	}
}

func TestUnmarshalLineString(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		expected orb.LineString
	}{
		{
			name:     "empty",
			s:        "LINESTRING EMPTY",
			expected: orb.LineString{},
		},
		{
			name:     "2 points",
			s:        "LINESTRING(1 2,0.5 1.5)",
			expected: orb.LineString{{1, 2}, {0.5, 1.5}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			geom, err := UnmarshalLineString(tc.s)
			if err != nil {
				t.Fatal(err)
			}

			if !geom.Equal(tc.expected) {
				t.Log(geom)
				t.Log(tc.expected)
				t.Errorf("incorrect wkt unmarshalling")
			}
		})
	}
}

func TestUnmarshalMultiLineString(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		expected orb.MultiLineString
	}{
		{
			name:     "empty",
			s:        "MULTILINESTRING EMPTY",
			expected: orb.MultiLineString{},
		},
		{
			name:     "2 lines",
			s:        "MULTILINESTRING((1 2,3 4),(5 6,7 8))",
			expected: orb.MultiLineString{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			geom, err := UnmarshalMultiLineString(tc.s)
			if err != nil {
				t.Fatal(err)
			}

			if !geom.Equal(tc.expected) {
				t.Log(geom)
				t.Log(tc.expected)
				t.Errorf("incorrect wkt unmarshalling")
			}
		})
	}
}

func TestUnmarshalPolygon(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		expected orb.Polygon
	}{
		{
			name:     "empty",
			s:        "POLYGON EMPTY",
			expected: orb.Polygon{},
		},
		{
			name:     "one ring",
			s:        "POLYGON((0 0,1 0,1 1,0 0))",
			expected: orb.Polygon{{{0, 0}, {1, 0}, {1, 1}, {0, 0}}},
		},
		{
			name:     "two rings",
			s:        "POLYGON((1 2,3 4),(5 6,7 8))",
			expected: orb.Polygon{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			geom, err := UnmarshalPolygon(tc.s)
			if err != nil {
				t.Fatal(err)
			}

			if !geom.Equal(tc.expected) {
				t.Log(geom)
				t.Log(tc.expected)
				t.Errorf("incorrect wkt unmarshalling")
			}
		})
	}
}

func TestUnmarshalMutilPolygon(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		expected orb.MultiPolygon
	}{
		{
			name:     "empty",
			s:        "MULTIPOLYGON EMPTY",
			expected: orb.MultiPolygon{},
		},
		{
			name:     "mulit-polygon",
			s:        "MULTIPOLYGON(((1 2,3 4)),((5 6,7 8),(1 2,5 4)))",
			expected: orb.MultiPolygon{{{{1, 2}, {3, 4}}}, {{{5, 6}, {7, 8}}, {{1, 2}, {5, 4}}}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			geom, err := UnmarshalMultiPolygon(tc.s)
			if err != nil {
				t.Fatal(err)
			}
			if !geom.Equal(tc.expected) {
				t.Log(geom)
				t.Log(tc.expected)
				t.Errorf("incorrect wkt unmarshalling")
			}
		})
	}
}

func TestUnmarshalCollection(t *testing.T) {
	cases := []struct {
		name     string
		s        string
		expected orb.Collection
	}{
		{
			name:     "empty",
			s:        "GEOMETRYCOLLECTION EMPTY",
			expected: orb.Collection{},
		},
		{
			name:     "point and line",
			s:        "GEOMETRYCOLLECTION(POINT(1 2),LINESTRING(3 4,5 6))",
			expected: orb.Collection{orb.Point{1, 2}, orb.LineString{{3, 4}, {5, 6}}},
		},
		{
			name: "lots of things",
			s:    "GEOMETRYCOLLECTION(POINT(1 2),LINESTRING(3 4,5 6),MULTILINESTRING((1 2,3 4),(5 6,7 8)),POLYGON((0 0,1 0,1 1,0 0)),POLYGON((1 2,3 4),(5 6,7 8)),MULTIPOLYGON(((1 2,3 4)),((5 6,7 8),(1 2,5 4)))",
			expected: orb.Collection{
				orb.Point{1, 2},
				orb.LineString{{3, 4}, {5, 6}},
				orb.MultiLineString{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
				orb.Polygon{{{0, 0}, {1, 0}, {1, 1}, {0, 0}}},
				orb.Polygon{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}},
				orb.MultiPolygon{{{{1, 2}, {3, 4}}}, {{{5, 6}, {7, 8}}, {{1, 2}, {5, 4}}}},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			geom, err := UnmarshalCollection(tc.s)
			if err != nil {
				t.Fatal(err)
			}
			if !geom.Equal(tc.expected) {
				t.Log(geom)
				t.Log(tc.expected)
				t.Errorf("incorrect wkt unmarshalling")
			}
		})
	}
}
