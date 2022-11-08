package genenv

import "fmt"

func (suite *GenenvTestSuite) TestStrConvToPascal() {
	sc := NewStrConv()
	suite.Equal("FooBar", sc.ToPascal("FOO_BAR"))
	suite.Equal("FooBar", sc.ToPascal("FOO BAR"))
	suite.Equal("FooBar", sc.ToPascal("FOO _ BAR"))
	suite.Equal("FoOBAr", sc.ToPascal("FO  O_B __AR"))
	suite.Equal("FoOBAr", sc.ToPascal("FO .- O_-,B ,_AR"))
}

func ExampleStrConv_ToPascal_underscore() {
	sc := NewStrConv()
	s := sc.ToPascal("FOO_BAR")
	fmt.Println(s)
	// output: FooBar
}

func ExampleStrConv_ToPascal_space() {
	sc := NewStrConv()
	s := sc.ToPascal("FOO BAR")
	fmt.Println(s)
	// output: FooBar
}

func ExampleStrConv_ToPascal_mixed() {
	sc := NewStrConv()
	s := sc.ToPascal("FO_O _ -_ B_ --_ _  AR")
	fmt.Println(s)
	// output: FoOBAr
}
