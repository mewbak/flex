package flex

import "testing"

func TestAlign_items_stretch(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetHeight(10)
	root.InsertChild(rootChild0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

}

func TestAlign_items_center(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(10)
	rootChild0.StyleSetHeight(10)
	root.InsertChild(rootChild0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 45, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 45, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

}

func TestAlign_items_flex_start(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetAlignItems(AlignFlexStart)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(10)
	rootChild0.StyleSetHeight(10)
	root.InsertChild(rootChild0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 90, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

}

func TestAlign_items_flex_end(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetAlignItems(AlignFlexEnd)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(10)
	rootChild0.StyleSetHeight(10)
	root.InsertChild(rootChild0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 90, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild0.LayoutGetHeight())

}

func TestAlign_baseline(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignBaseline)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(20)
	root.InsertChild(rootChild1, 1)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 30, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

}

func TestAlign_baseline_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignBaseline)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(20)
	root.InsertChild(rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	rootChild1child0.StyleSetWidth(50)
	rootChild1child0.StyleSetHeight(10)
	rootChild1.InsertChild(rootChild1child0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

}

func TestAlign_baseline_child_multiline(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignBaseline)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(60)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexDirection(FlexDirectionRow)
	rootChild1.StyleSetFlexWrap(WrapWrap)
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(25)
	root.InsertChild(rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	rootChild1child0.StyleSetWidth(25)
	rootChild1child0.StyleSetHeight(20)
	rootChild1.InsertChild(rootChild1child0, 0)

	rootChild1_child1 := NewNodeWithConfig(config)
	rootChild1_child1.StyleSetWidth(25)
	rootChild1_child1.StyleSetHeight(10)
	rootChild1.InsertChild(rootChild1_child1, 1)

	rootChild1_child2 := NewNodeWithConfig(config)
	rootChild1_child2.StyleSetWidth(25)
	rootChild1_child2.StyleSetHeight(20)
	rootChild1.InsertChild(rootChild1_child2, 2)

	rootChild1_child3 := NewNodeWithConfig(config)
	rootChild1_child3.StyleSetWidth(25)
	rootChild1_child3.StyleSetHeight(10)
	rootChild1.InsertChild(rootChild1_child3, 3)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 60, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1_child1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1_child2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child2.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1_child3.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child3.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child3.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1_child3.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 60, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1_child1.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1_child2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child2.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1_child3.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child3.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child3.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1_child3.LayoutGetHeight())

}

func TestAlign_baseline_child_multiline_override(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignBaseline)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(60)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexDirection(FlexDirectionRow)
	rootChild1.StyleSetFlexWrap(WrapWrap)
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(25)
	root.InsertChild(rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	rootChild1child0.StyleSetWidth(25)
	rootChild1child0.StyleSetHeight(20)
	rootChild1.InsertChild(rootChild1child0, 0)

	rootChild1_child1 := NewNodeWithConfig(config)
	rootChild1_child1.StyleSetAlignSelf(AlignBaseline)
	rootChild1_child1.StyleSetWidth(25)
	rootChild1_child1.StyleSetHeight(10)
	rootChild1.InsertChild(rootChild1_child1, 1)

	rootChild1_child2 := NewNodeWithConfig(config)
	rootChild1_child2.StyleSetWidth(25)
	rootChild1_child2.StyleSetHeight(20)
	rootChild1.InsertChild(rootChild1_child2, 2)

	rootChild1_child3 := NewNodeWithConfig(config)
	rootChild1_child3.StyleSetAlignSelf(AlignBaseline)
	rootChild1_child3.StyleSetWidth(25)
	rootChild1_child3.StyleSetHeight(10)
	rootChild1.InsertChild(rootChild1_child3, 3)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 60, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1_child1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1_child2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child2.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1_child3.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child3.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child3.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1_child3.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 60, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1_child1.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1_child2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child2.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1_child3.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child3.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child3.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1_child3.LayoutGetHeight())

}

func TestAlign_baseline_child_multiline_no_override_on_secondline(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignBaseline)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(60)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetFlexDirection(FlexDirectionRow)
	rootChild1.StyleSetFlexWrap(WrapWrap)
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(25)
	root.InsertChild(rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	rootChild1child0.StyleSetWidth(25)
	rootChild1child0.StyleSetHeight(20)
	rootChild1.InsertChild(rootChild1child0, 0)

	rootChild1_child1 := NewNodeWithConfig(config)
	rootChild1_child1.StyleSetWidth(25)
	rootChild1_child1.StyleSetHeight(10)
	rootChild1.InsertChild(rootChild1_child1, 1)

	rootChild1_child2 := NewNodeWithConfig(config)
	rootChild1_child2.StyleSetWidth(25)
	rootChild1_child2.StyleSetHeight(20)
	rootChild1.InsertChild(rootChild1_child2, 2)

	rootChild1_child3 := NewNodeWithConfig(config)
	rootChild1_child3.StyleSetAlignSelf(AlignBaseline)
	rootChild1_child3.StyleSetWidth(25)
	rootChild1_child3.StyleSetHeight(10)
	rootChild1.InsertChild(rootChild1_child3, 3)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 60, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1_child1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1_child2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child2.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1_child3.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child3.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child3.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1_child3.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 60, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 25, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1_child1.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child1.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1_child1.LayoutGetHeight())

	assertFloatEqual(t, 25, rootChild1_child2.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child2.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1_child2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1_child3.LayoutGetLeft())
	assertFloatEqual(t, 20, rootChild1_child3.LayoutGetTop())
	assertFloatEqual(t, 25, rootChild1_child3.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1_child3.LayoutGetHeight())

}

func TestAlign_baseline_child_top(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignBaseline)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetPosition(EdgeTop, 10)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(20)
	root.InsertChild(rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	rootChild1child0.StyleSetWidth(50)
	rootChild1child0.StyleSetHeight(10)
	rootChild1.InsertChild(rootChild1child0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

}

func TestAlign_baseline_child_top2(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignBaseline)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetPosition(EdgeTop, 5)
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(20)
	root.InsertChild(rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	rootChild1child0.StyleSetWidth(50)
	rootChild1child0.StyleSetHeight(10)
	rootChild1.InsertChild(rootChild1child0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 45, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 45, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

}

func TestAlign_baseline_double_nested_child(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignBaseline)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	root.InsertChild(rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetWidth(50)
	rootChild0Child0.StyleSetHeight(20)
	rootChild0.InsertChild(rootChild0Child0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(20)
	root.InsertChild(rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	rootChild1child0.StyleSetWidth(50)
	rootChild1child0.StyleSetHeight(15)
	rootChild1.InsertChild(rootChild1child0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 15, rootChild1child0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 15, rootChild1child0.LayoutGetHeight())

}

func TestAlign_baseline_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetAlignItems(AlignBaseline)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(20)
	root.InsertChild(rootChild1, 1)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

}

func TestAlign_baseline_child_margin(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignBaseline)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetMargin(EdgeLeft, 5)
	rootChild0.StyleSetMargin(EdgeTop, 5)
	rootChild0.StyleSetMargin(EdgeRight, 5)
	rootChild0.StyleSetMargin(EdgeBottom, 5)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(20)
	root.InsertChild(rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	rootChild1child0.StyleSetMargin(EdgeLeft, 1)
	rootChild1child0.StyleSetMargin(EdgeTop, 1)
	rootChild1child0.StyleSetMargin(EdgeRight, 1)
	rootChild1child0.StyleSetMargin(EdgeBottom, 1)
	rootChild1child0.StyleSetWidth(50)
	rootChild1child0.StyleSetHeight(10)
	rootChild1.InsertChild(rootChild1child0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 5, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 60, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 44, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 1, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 1, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 45, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, -10, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 44, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, -1, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 1, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

}

func TestAlign_baseline_child_padding(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignBaseline)
	root.StyleSetPadding(EdgeLeft, 5)
	root.StyleSetPadding(EdgeTop, 5)
	root.StyleSetPadding(EdgeRight, 5)
	root.StyleSetPadding(EdgeBottom, 5)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetPadding(EdgeLeft, 5)
	rootChild1.StyleSetPadding(EdgeTop, 5)
	rootChild1.StyleSetPadding(EdgeRight, 5)
	rootChild1.StyleSetPadding(EdgeBottom, 5)
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(20)
	root.InsertChild(rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	rootChild1child0.StyleSetWidth(50)
	rootChild1child0.StyleSetHeight(10)
	rootChild1.InsertChild(rootChild1child0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 5, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 55, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 5, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 45, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, -5, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, -5, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 5, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

}

func TestAlign_baseline_multiline(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignBaseline)
	root.StyleSetFlexWrap(WrapWrap)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(20)
	root.InsertChild(rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	rootChild1child0.StyleSetWidth(50)
	rootChild1child0.StyleSetHeight(10)
	rootChild1.InsertChild(rootChild1child0, 0)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(50)
	rootChild2.StyleSetHeight(20)
	root.InsertChild(rootChild2, 2)

	rootChild2_child0 := NewNodeWithConfig(config)
	rootChild2_child0.StyleSetWidth(50)
	rootChild2_child0.StyleSetHeight(10)
	rootChild2.InsertChild(rootChild2_child0, 0)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetWidth(50)
	rootChild3.StyleSetHeight(50)
	root.InsertChild(rootChild3, 3)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild2_child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2_child0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild3.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild2_child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2_child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 60, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild3.LayoutGetHeight())

}

func TestAlign_baseline_multiline_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetAlignItems(AlignBaseline)
	root.StyleSetFlexWrap(WrapWrap)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(30)
	rootChild1.StyleSetHeight(50)
	root.InsertChild(rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	rootChild1child0.StyleSetWidth(20)
	rootChild1child0.StyleSetHeight(20)
	rootChild1.InsertChild(rootChild1child0, 0)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(40)
	rootChild2.StyleSetHeight(70)
	root.InsertChild(rootChild2, 2)

	rootChild2_child0 := NewNodeWithConfig(config)
	rootChild2_child0.StyleSetWidth(10)
	rootChild2_child0.StyleSetHeight(10)
	rootChild2.InsertChild(rootChild2_child0, 0)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetWidth(50)
	rootChild3.StyleSetHeight(20)
	root.InsertChild(rootChild3, 3)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 70, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2_child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2_child0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild3.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 70, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2_child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2_child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild3.LayoutGetHeight())

}

func TestAlign_baseline_multiline_column2(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetAlignItems(AlignBaseline)
	root.StyleSetFlexWrap(WrapWrap)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(30)
	rootChild1.StyleSetHeight(50)
	root.InsertChild(rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	rootChild1child0.StyleSetWidth(20)
	rootChild1child0.StyleSetHeight(20)
	rootChild1.InsertChild(rootChild1child0, 0)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(40)
	rootChild2.StyleSetHeight(70)
	root.InsertChild(rootChild2, 2)

	rootChild2_child0 := NewNodeWithConfig(config)
	rootChild2_child0.StyleSetWidth(10)
	rootChild2_child0.StyleSetHeight(10)
	rootChild2.InsertChild(rootChild2_child0, 0)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetWidth(50)
	rootChild3.StyleSetHeight(20)
	root.InsertChild(rootChild3, 3)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 70, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2_child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2_child0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild3.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 70, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 50, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 30, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 40, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 70, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 30, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 10, rootChild2_child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2_child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 70, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild3.LayoutGetHeight())

}

func TestAlign_baseline_multiline_row_and_column(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetFlexDirection(FlexDirectionRow)
	root.StyleSetAlignItems(AlignBaseline)
	root.StyleSetFlexWrap(WrapWrap)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetWidth(50)
	rootChild0.StyleSetHeight(50)
	root.InsertChild(rootChild0, 0)

	rootChild1 := NewNodeWithConfig(config)
	rootChild1.StyleSetWidth(50)
	rootChild1.StyleSetHeight(50)
	root.InsertChild(rootChild1, 1)

	rootChild1child0 := NewNodeWithConfig(config)
	rootChild1child0.StyleSetWidth(50)
	rootChild1child0.StyleSetHeight(10)
	rootChild1.InsertChild(rootChild1child0, 0)

	rootChild2 := NewNodeWithConfig(config)
	rootChild2.StyleSetWidth(50)
	rootChild2.StyleSetHeight(20)
	root.InsertChild(rootChild2, 2)

	rootChild2_child0 := NewNodeWithConfig(config)
	rootChild2_child0.StyleSetWidth(50)
	rootChild2_child0.StyleSetHeight(10)
	rootChild2.InsertChild(rootChild2_child0, 0)

	rootChild3 := NewNodeWithConfig(config)
	rootChild3.StyleSetWidth(50)
	rootChild3.StyleSetHeight(20)
	root.InsertChild(rootChild3, 3)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild2_child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2_child0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 90, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild3.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1.LayoutGetLeft())
	assertFloatEqual(t, 40, rootChild1.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1.LayoutGetWidth())
	assertFloatEqual(t, 50, rootChild1.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild1child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild1child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild1child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild1child0.LayoutGetHeight())

	assertFloatEqual(t, 50, rootChild2.LayoutGetLeft())
	assertFloatEqual(t, 100, rootChild2.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild2.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild2.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild2_child0.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild2_child0.LayoutGetWidth())
	assertFloatEqual(t, 10, rootChild2_child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild3.LayoutGetLeft())
	assertFloatEqual(t, 90, rootChild3.LayoutGetTop())
	assertFloatEqual(t, 50, rootChild3.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild3.LayoutGetHeight())

}

func TestAlign_items_center_child_with_margin_bigger_than_parent(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetJustifyContent(JustifyCenter)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetWidth(52)
	root.StyleSetHeight(52)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetAlignItems(AlignCenter)
	root.InsertChild(rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetMargin(EdgeLeft, 10)
	rootChild0Child0.StyleSetMargin(EdgeRight, 10)
	rootChild0Child0.StyleSetWidth(52)
	rootChild0Child0.StyleSetHeight(52)
	rootChild0.InsertChild(rootChild0Child0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, root.LayoutGetWidth())
	assertFloatEqual(t, 52, root.LayoutGetHeight())

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 52, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 52, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 52, rootChild0Child0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, root.LayoutGetWidth())
	assertFloatEqual(t, 52, root.LayoutGetHeight())

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 52, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 52, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 52, rootChild0Child0.LayoutGetHeight())

}

func TestAlign_items_flex_end_child_with_margin_bigger_than_parent(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetJustifyContent(JustifyCenter)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetWidth(52)
	root.StyleSetHeight(52)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetAlignItems(AlignFlexEnd)
	root.InsertChild(rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetMargin(EdgeLeft, 10)
	rootChild0Child0.StyleSetMargin(EdgeRight, 10)
	rootChild0Child0.StyleSetWidth(52)
	rootChild0Child0.StyleSetHeight(52)
	rootChild0.InsertChild(rootChild0Child0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, root.LayoutGetWidth())
	assertFloatEqual(t, 52, root.LayoutGetHeight())

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 52, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 52, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 52, rootChild0Child0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, root.LayoutGetWidth())
	assertFloatEqual(t, 52, root.LayoutGetHeight())

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 52, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 10, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 52, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 52, rootChild0Child0.LayoutGetHeight())

}

func TestAlign_items_center_child_without_margin_bigger_than_parent(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetJustifyContent(JustifyCenter)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetWidth(52)
	root.StyleSetHeight(52)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetAlignItems(AlignCenter)
	root.InsertChild(rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetWidth(72)
	rootChild0Child0.StyleSetHeight(72)
	rootChild0.InsertChild(rootChild0Child0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, root.LayoutGetWidth())
	assertFloatEqual(t, 52, root.LayoutGetHeight())

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, -10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 72, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 72, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 72, rootChild0Child0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, root.LayoutGetWidth())
	assertFloatEqual(t, 52, root.LayoutGetHeight())

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, -10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 72, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 72, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 72, rootChild0Child0.LayoutGetHeight())

}

func TestAlign_items_flex_end_child_without_margin_bigger_than_parent(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetJustifyContent(JustifyCenter)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetWidth(52)
	root.StyleSetHeight(52)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetAlignItems(AlignFlexEnd)
	root.InsertChild(rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetWidth(72)
	rootChild0Child0.StyleSetHeight(72)
	rootChild0.InsertChild(rootChild0Child0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, root.LayoutGetWidth())
	assertFloatEqual(t, 52, root.LayoutGetHeight())

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, -10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 72, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 72, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 72, rootChild0Child0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 0, root.LayoutGetTop())
	assertFloatEqual(t, 52, root.LayoutGetWidth())
	assertFloatEqual(t, 52, root.LayoutGetHeight())

	assertFloatEqual(t, -10, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, -10, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 72, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 72, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 72, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 72, rootChild0Child0.LayoutGetHeight())

}

func TestAlign_center_should_size_based_on_content(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetAlignItems(AlignCenter)
	root.StyleSetMargin(EdgeTop, 20)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetJustifyContent(JustifyCenter)
	rootChild0.StyleSetFlexShrink(1)
	root.InsertChild(rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetFlexGrow(1)
	rootChild0Child0.StyleSetFlexShrink(1)
	rootChild0.InsertChild(rootChild0Child0, 0)

	rootChild0Child0Child0 := NewNodeWithConfig(config)
	rootChild0Child0Child0.StyleSetWidth(20)
	rootChild0Child0Child0.StyleSetHeight(20)
	rootChild0Child0.InsertChild(rootChild0Child0Child0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 20, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0Child0Child0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 20, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 40, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0Child0Child0.LayoutGetHeight())

}

func TestAlign_strech_should_size_based_on_parent(t *testing.T) {
	config := NewConfig()

	root := NewNodeWithConfig(config)
	root.StyleSetMargin(EdgeTop, 20)
	root.StyleSetWidth(100)
	root.StyleSetHeight(100)

	rootChild0 := NewNodeWithConfig(config)
	rootChild0.StyleSetJustifyContent(JustifyCenter)
	rootChild0.StyleSetFlexShrink(1)
	root.InsertChild(rootChild0, 0)

	rootChild0Child0 := NewNodeWithConfig(config)
	rootChild0Child0.StyleSetFlexGrow(1)
	rootChild0Child0.StyleSetFlexShrink(1)
	rootChild0.InsertChild(rootChild0Child0, 0)

	rootChild0Child0Child0 := NewNodeWithConfig(config)
	rootChild0Child0Child0.StyleSetWidth(20)
	rootChild0Child0Child0.StyleSetHeight(20)
	rootChild0Child0.InsertChild(rootChild0Child0Child0, 0)
	CalculateLayout(root, Undefined, Undefined, DirectionLTR)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 20, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0Child0Child0.LayoutGetHeight())

	CalculateLayout(root, Undefined, Undefined, DirectionRTL)

	assertFloatEqual(t, 0, root.LayoutGetLeft())
	assertFloatEqual(t, 20, root.LayoutGetTop())
	assertFloatEqual(t, 100, root.LayoutGetWidth())
	assertFloatEqual(t, 100, root.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0.LayoutGetHeight())

	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0.LayoutGetTop())
	assertFloatEqual(t, 100, rootChild0Child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0Child0.LayoutGetHeight())

	assertFloatEqual(t, 80, rootChild0Child0Child0.LayoutGetLeft())
	assertFloatEqual(t, 0, rootChild0Child0Child0.LayoutGetTop())
	assertFloatEqual(t, 20, rootChild0Child0Child0.LayoutGetWidth())
	assertFloatEqual(t, 20, rootChild0Child0Child0.LayoutGetHeight())

}
