#!/bin/bash
set -e

# Script to regenerate AI golden files using real Ollama + qwen2.5-coder:0.5b
# This ensures all golden files match the current LLM output for E2E testing

echo "🤖 Regenerating AI golden files with Ollama..."
echo ""

# Array of test cases: "function_name:source_file:golden_file"
GOLDENS=(
    # Original test cases
    "CalculateDiscount:testdata/business_logic.go:testdata/goldens/business_logic_calculate_discount_ai.go"
    "Clamp:testdata/math_ops.go:testdata/goldens/math_ops_clamp_ai.go"
    "FilterPositive:testdata/data_processing.go:testdata/goldens/data_processing_filter_positive_ai.go"
    "HashPassword:testdata/user_service.go:testdata/goldens/user_service_hash_password_ai.go"

    # New test cases for improved coverage
    "Multiply:testdata/calculator.go:testdata/goldens/calculator_multiply_ai.go"
    "Divide:testdata/calculator.go:testdata/goldens/calculator_divide_ai.go"
    "Reverse:testdata/string_utils.go:testdata/goldens/string_utils_reverse_ai.go"
    "ParseKeyValue:testdata/string_utils.go:testdata/goldens/string_utils_parse_key_value_ai.go"
    "ContainsAny:testdata/string_utils.go:testdata/goldens/string_utils_contains_any_ai.go"

    # Previously had bad test names
    "FormatCurrency:testdata/business_logic.go:testdata/goldens/business_logic_format_currency_ai.go"
    "Factorial:testdata/math_ops.go:testdata/goldens/math_ops_factorial_ai.go"
)

for entry in "${GOLDENS[@]}"; do
    IFS=':' read -r func_name source_file golden_file <<< "$entry"

    echo "Generating $golden_file..."

    # Run gotests with AI, skip CLI output lines, save to golden file
    go run ./gotests -ai -only "^${func_name}$" "$source_file" 2>/dev/null | \
        grep -v "^⚠️" | \
        grep -v "^Generated" | \
        sed '/^$/d;1s/^//' > "$golden_file"

    echo "✓ $golden_file"
done

echo ""
echo "✅ All golden files regenerated successfully!"
echo ""
echo "Run E2E tests to validate:"
echo "  GOTESTS_E2E=true go test -tags=e2e -v ./internal/ai"
