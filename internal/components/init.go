package components

import "fmt"

func Init() error {
	if err := InitLog(); err != nil {
		return fmt.Errorf("init log: %w", err)
	}
	if err := InitDB(); err != nil {
		return fmt.Errorf("init db: %w", err)
	}
	if err := InitRedis(); err != nil {
		return fmt.Errorf("init redis: %w", err)
	}
	return nil
}
