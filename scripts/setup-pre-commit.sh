#!/bin/bash

# Setup script for pre-commit hooks
# This helps developers install the necessary tools for secret detection

set -e

echo "🔧 Setting up pre-commit hooks for secret detection..."

# Check if pre-commit is installed
if ! command -v pre-commit &> /dev/null; then
    echo "📦 Installing pre-commit..."
    if command -v pip &> /dev/null; then
        pip install pre-commit
    elif command -v brew &> /dev/null; then
        brew install pre-commit
    else
        echo "❌ Please install pre-commit manually: https://pre-commit.com/#installation"
        exit 1
    fi
fi

# Install the pre-commit hooks
echo "🎣 Installing pre-commit hooks..."
pre-commit install

# Check if gitleaks is installed
if ! command -v gitleaks &> /dev/null; then
    echo "📦 Installing gitleaks..."
    if command -v brew &> /dev/null; then
        brew install gitleaks
    elif command -v go &> /dev/null; then
        go install github.com/gitleaks/gitleaks/v8@latest
    else
        echo "⚠️  Gitleaks not found. Please install it manually: https://github.com/gitleaks/gitleaks#installation"
    fi
fi

# Check if trufflehog is installed
if ! command -v trufflehog &> /dev/null; then
    echo "📦 Installing trufflehog..."
    if command -v brew &> /dev/null; then
        brew install trufflehog
    elif command -v go &> /dev/null; then
        go install github.com/trufflesecurity/trufflehog/v3@latest
    else
        echo "⚠️  TruffleHog not found. Please install it manually: https://github.com/trufflesecurity/trufflehog#installation"
    fi
fi

# Create initial secrets baseline
echo "🔍 Creating initial secrets baseline..."
if command -v detect-secrets &> /dev/null; then
    detect-secrets scan --baseline .secrets.baseline
else
    echo "⚠️  detect-secrets not found. Installing..."
    pip install detect-secrets
    detect-secrets scan --baseline .secrets.baseline
fi

# Run pre-commit on all files to check current state
echo "🧪 Running pre-commit on all files..."
pre-commit run --all-files || true

echo "✅ Pre-commit hooks setup complete!"
echo ""
echo "📋 Next steps:"
echo "   1. Review any findings from the secret scan above"
echo "   2. Fix any detected issues"
echo "   3. The hooks will now run automatically on each commit"
echo ""
echo "🔧 Manual commands:"
echo "   - Run all hooks: pre-commit run --all-files"
echo "   - Run gitleaks only: gitleaks detect"
echo "   - Run trufflehog only: trufflehog git file://. --only-verified"
