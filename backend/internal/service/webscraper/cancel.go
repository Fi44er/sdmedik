package webscraper

import "fmt"

func (s *service) CancelScraper() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.cancelFunc != nil {
		s.cancelFunc() // Отменяем текущий процесс
		s.cancelFunc = nil
		return nil
	}

	return fmt.Errorf("нет активного парсинга")
}
