package ports
import (
    "github.com/nigelpage/HBCMembers/internal/domain"
)
type MemberRepository interface {
    Save(member *domain.Member) error
    FindByID(id int) (*domain.Member, error)
    FindAll() ([]*domain.Member, error)
}
