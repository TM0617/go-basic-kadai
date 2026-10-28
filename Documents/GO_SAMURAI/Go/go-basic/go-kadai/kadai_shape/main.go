package main

import "fmt"

type ValidationError struct {
	Message string
	Field   string
	Value   float64
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("[バリデーションエラー] %s : %s (%.2f)", e.Field, e.Message, e.Value)
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return r.Width*2 + r.Height*2
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func validateShape(s Shape) error {
	switch val := s.(type) {
	case Rectangle:
		if val.Width <= 0 {
			//自分でエラーを作る。
			return ValidationError{
				Message: "幅は正の数を指定してください。",
				Field:   "Rectangle.Width",
				Value:   val.Width,
			}
		}
		if val.Height <= 0 {
			return ValidationError{
				Message: "高さは正の数を指定してください。",
				Field:   "Rectangle.Height",
				Value:   val.Height,
			}
		}
	case Circle:
		if val.Radius <= 0 {
			return ValidationError{
				Message: "半径は正の数を指定してください。",
				Field:   "Circle.Radius",
				Value:   val.Radius,
			}
		}
	default:
		return ValidationError{
			Message: "原因不明です。",
			Field:   "Type",
			Value:   0,
		}
	}
	return nil
}

func calc(s Shape) (float64, float64) {
	area := s.Area()
	perimeter := s.Perimeter()
	return area, perimeter
}

func main() {
	// 処理対象の図形スライスを作成
	shapes := []Shape{
		Rectangle{Width: 10.0, Height: 5.0},
		Circle{Radius: 3.0},
		Rectangle{Width: -2.0, Height: 5.0}, // 不正（幅が不正）
		Circle{Radius: -1.0},                // 不正（半径が不正）
		Rectangle{Width: 1.0, Height: 0.0},  // 不正（高さが不正）
	}
	var validShapes []Shape

	for index, shape := range shapes {
		shapeType := fmt.Sprintf("%T", shape)

		err := validateShape(shape)
		if err != nil {
			fmt.Printf("#%d:%s => NG: %v\n", index+1, shapeType, err)
			continue
		}

		fmt.Printf("#%d: %s => バリデーションOK\n", index+1, shapeType)
		validShapes = append(validShapes, shape)

	}
	for _, shape := range validShapes {
		shapeType := fmt.Sprintf("%T", shape)

		area, perimeter := calc(shape)

		fmt.Printf(
			"図形：%s｜面積：%.2f ｜周長：%.2f\n",
			shapeType, area, perimeter,
		)
	}
}
