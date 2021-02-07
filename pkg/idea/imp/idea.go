package imp

//go:generate mockgen -destination ../../mocks/mock_idea.go -package=Mocks github.com/mchirico/mock-playground/pkg/idea/imp IDEA
type IDEA interface {
	Start(string, bool, IDEA) bool
	Running(string, bool, IDEA) bool
	F(string) (map[string]string, error)
	IdeaCheck(IDEA, chan map[string]string, chan error)
	Calls(map[string]string, error)
}
