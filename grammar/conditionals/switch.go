package conditionals

import "fmt"

func Switch() {
	// --- 1. 기본적인 switch 문 ---

	// switch 문은 '표현식'의 값을 평가하고, 일치하는 'case' 블록을 실행합니다.
	// Go에서는 break 문을 명시적으로 작성하지 않아도 case 블록이 실행된 후 자동으로 switch 문을 빠져나갑니다.
	
	fruit := "apple"

	fmt.Println("--- 1. 과일 종류 판별 ---")
	switch fruit {
	case "apple":
		fmt.Println("사과입니다.")
	case "banana":
		fmt.Println("바나나입니다.")
	default: // 어떤 case에도 해당하지 않을 때 실행됩니다.
		fmt.Println("알 수 없는 과일입니다.")
	}

	// --- 2. 여러 case를 한 번에 처리하기 ---

	// 같은 로직을 공유하는 여러 case를 콤마(,)로 묶어 처리할 수 있습니다.
	day := "saturday"

	fmt.Println("\n--- 2. 주말/평일 구분하기 ---")
	switch day {
	case "saturday", "sunday":
		fmt.Println("주말입니다.")
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("평일입니다.")
	default:
		fmt.Println("잘못된 요일입니다.")
	}

	// --- 3. 조건식이 없는 switch 문 (if-else if 대용) ---
	
	// switch 문에 표현식을 생략하면 'switch true'와 같이 작동합니다.
	// 각 case는 불리언(boolean) 조건식이 됩니다. 이는 'if-else if-else' 문과 유사합니다.
	score := 85

	fmt.Println("\n--- 3. 점수에 따라 등급 매기기 ---")
	switch {
	case score >= 90:
		fmt.Println("A 학점입니다.")
	case score >= 80: // score가 90보다 작고 80보다 크거나 같을 때
		fmt.Println("B 학점입니다.")
	case score >= 70:
		fmt.Println("C 학점입니다.")
	default:
		fmt.Println("F 학점입니다.")
	}

	// --- 4. `fallthrough` 키워드 ---
	
	// Go의 switch는 기본적으로 하나의 case만 실행하고 종료됩니다.
	// 하지만 `fallthrough` 키워드를 사용하면, 일치하는 case 실행 후 **다음 case까지 실행**합니다.
	// (자주 사용되지는 않습니다. 로직을 복잡하게 만들 수 있으니 주의해야 합니다.)
	
	grade := "B"

	fmt.Println("\n--- 4. `fallthrough` 키워드 사용 ---")
	switch grade {
	case "A":
		fmt.Println("매우 잘했습니다.")
	case "B":
		fmt.Println("잘했습니다.")
		fallthrough // 다음 case로 실행 흐름을 넘깁니다.
	case "C":
		fmt.Println("보통입니다.")
		fallthrough // 다음 case로 실행 흐름을 넘깁니다.
	default:
		fmt.Println("노력하세요.")
	}

	// --- 5. switch 문에 변수 선언 포함하기 ---
	
	// if 문과 마찬가지로, switch 문도 조건식 앞에 변수를 선언할 수 있습니다.
	// 여기서 선언된 변수(mod)는 switch 블록 내부에서만 유효합니다.
	fmt.Println("\n--- 5. switch에 변수 선언 포함 ---")
	switch mod := 11 % 3; mod {
	case 0:
		fmt.Println("11은 3의 배수입니다.")
	case 1:
		fmt.Println("11을 3으로 나눈 나머지는 1입니다.")
	case 2:
		fmt.Println("11을 3으로 나눈 나머지는 2입니다.")
	}
}