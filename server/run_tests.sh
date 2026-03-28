#!/bin/bash

# Calculator Server Test Coverage Script
# This script runs all tests and generates a coverage report

set -e

echo "================================"
echo "Running Calculator Server Tests"
echo "================================"

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Run tests with coverage
echo ""
echo "Running tests..."
go test -v ./... -coverprofile=coverage.out

# Check if tests passed
if [ $? -eq 0 ]; then
    echo -e "${GREEN}All tests passed!${NC}"
else
    echo -e "${RED}Some tests failed!${NC}"
    exit 1
fi

# Generate coverage report
echo ""
echo "Generating coverage report..."

# Create coverage directory
mkdir -p coverage

go tool cover -func=coverage.out > coverage/coverage_report.txt

# Display coverage summary
echo ""
echo "================================"
echo "Coverage Summary"
echo "================================"
cat coverage/coverage_report.txt | grep "total:"

# Generate HTML coverage report
go tool cover -html=coverage.out -o coverage/coverage.html

# Display per-package coverage
echo ""
echo "================================"
echo "Per-Package Coverage"
echo "================================"
cat coverage/coverage_report.txt | grep -E "^(calculator-server|ok|FAIL)"

# Check if coverage meets 95% threshold (simplified check without bc)
echo ""
echo "================================"
echo "Coverage Threshold Check"
echo "================================"

# Extract total coverage percentage
TOTAL_COVERAGE=$(cat coverage/coverage_report.txt | grep "total:" | awk '{print $3}' | sed 's/%//')
echo "Total coverage: $TOTAL_COVERAGE%"

# Simple check - just warn if below 95%
if [ -n "$TOTAL_COVERAGE" ]; then
    # Extract integer part for comparison
    COVERAGE_INT=$(echo "$TOTAL_COVERAGE" | cut -d. -f1)
    if [ "$COVERAGE_INT" -ge 95 ]; then
        echo -e "${GREEN}Coverage meets the 95% threshold!${NC}"
    else
        echo -e "${YELLOW}Warning: Coverage is below the 95% threshold${NC}"
    fi
else
    echo -e "${YELLOW}Could not determine coverage percentage${NC}"
fi

echo ""
echo "Coverage reports saved to coverage/"
echo "  - coverage/coverage_report.txt (text report)"
echo "  - coverage/coverage.html (HTML report)"
