package logging

// DefaultLoggerFactory .
type DefaultLoggerFactory struct {
}

// GetLogger .
func (f *DefaultLoggerFactory) GetLogger(name string) Logger {
	logger := &DefaultLogger{
		Name: name,
	}
	logger.AddHandler(DefaultConsoleHandler)
	return logger
}

// GetParent .
func (f *DefaultLoggerFactory) GetParent(name string) Logger {
	return nil
}

var globalFactory LoggerFactory = &DefaultLoggerFactory{}

// GetLogger .
func GetLogger(moduleName string) Logger {
	return globalFactory.GetLogger(moduleName)
}

// Shutdown .
// This will flush and close all handlers.
func Shutdown() {

}

// SetLoggerFactory .
func SetLoggerFactory(factory LoggerFactory) {
	globalFactory = factory
}

func init() {

}
