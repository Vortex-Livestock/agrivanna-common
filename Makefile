# ------------------------------------------------------------------------------
# Install dependencies
# ------------------------------------------------------------------------------
.PHONY: install
install:
	@echo "Installing dependencies..."
	
	go mod tidy

	@echo "Dependencies installed."
